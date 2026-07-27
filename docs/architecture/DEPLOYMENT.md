# Deployment Architecture

## Decision

Use Docker Compose as the single-machine orchestrator and Traefik as the local
reverse proxy and load balancer. Cloudflared is the only public ingress path and
connects to Traefik over a private Docker network. Kerberos is not part of this
design: it is a network authentication protocol, not a container load balancer.

Traefik was selected over HAProxy, Caddy, and Nginx because its Docker provider
discovers labelled Compose containers and updates routes when replicas change.
HAProxy is an excellent data-plane proxy but needs an additional discovery
mechanism here. Caddy is simpler for certificates, which Cloudflare already
terminates. Nginx Open Source requires generated configuration and reloads for
dynamic Compose replicas.

## Topology

```text
friends
  |
Cloudflare edge
  |
named Cloudflare Tunnel (outbound connection only)
  |
cloudflared container
  |
private ingress network
  |
Traefik :8787
  |----------------------|
app-1                 app-N
  |----------------------|
private service network
  |
Valkey (expiring game and room state)
```

No application, Valkey, metrics, Docker API proxy, or Traefik dashboard port is
published on the host. An optional loopback-only development binding uses port
`8787`, because port `8080` is already occupied on the target machine.

## Container discovery security

Mounting `/var/run/docker.sock` directly into a public-facing proxy grants a
dangerously broad Docker API capability. Production uses a pinned Docker socket
proxy on a dedicated network. It permits only the read-only endpoints Traefik
needs for container and network discovery. Traefik has no direct socket mount.

Additional controls:

- `exposedByDefault=false`; only explicitly labelled services are routable.
- The dashboard is disabled in production, not merely hidden by a hostname.
- Containers run with a read-only root filesystem, dropped Linux capabilities,
  `no-new-privileges`, resource limits, and explicit health checks.
- Internal networks are marked internal where compatible with cloudflared
  egress requirements.
- Images use version and digest pins. No `latest` tags.
- Secrets arrive as Docker secrets or host-mounted files outside Git.

## Load-balancing semantics

Traefik balances new HTTP requests and WebSocket upgrades across ready app
replicas. A WebSocket remains attached to its selected replica until it closes.
Sticky cookies are deliberately unnecessary: every game command is checked
against shared versioned state, and reconnect may land on any healthy replica.

Before horizontal scaling is enabled, the game-state compare-and-swap contract
must pass concurrency tests against multiple app replicas. Scaling a
single-writer in-memory game actor behind a load balancer would be incorrect.

Health endpoints have distinct meanings:

- `/health/live`: process event loop is responsive.
- `/health/ready`: process can safely accept new HTTP and WebSocket work.
- `/health/startup`: schema/protocol and engine initialization completed.
- `/health/drain`: readiness becomes false, new upgrades stop, existing games
  receive a reconnect instruction, then the process exits after a deadline.

Traefik checks readiness and removes failing replicas. Compose restart policies
restart crashed containers. These mechanisms do not replace application-level
state versioning or reconnect logic.

## Scaling workflow

The operator uses a documented wrapper around:

```text
docker compose up -d --scale app=3
```

The wrapper validates the Compose configuration, pulls pinned images, runs
preflight checks, applies compatible changes, waits for readiness, and runs a
two-client smoke game. A failed check stops the rollout and leaves the prior
healthy containers serving.

Compose is suitable for replicas on one PC, but it is not a multi-host
orchestrator. Moving beyond one host is a new architecture decision: Nomad or
Kubernetes may then replace Compose without changing the HTTP/WebSocket
contract.

## Cloudflare behavior

A named tunnel supplies the stable hostname and requires no inbound port
forwarding. One cloudflared instance already maintains redundant edge
connections. A second replica improves connector-process availability but does
not perform application-aware local balancing; Traefik remains the local load
balancer.

The Cloudflare tunnel token is supplied by the operator and never committed.
Quick tunnels are development-only. Publishing is the last release task, after
security, chaos, load, recovery, and privacy checks succeed.

## Deployment tests

- Scale from one to three app replicas while idle and during active games.
- Kill each app replica and verify reconnect plus command idempotency.
- Mark one replica unready and prove Traefik sends it no new upgrades.
- Kill Traefik and verify its restart; document the brief ingress outage.
- Kill one cloudflared replica while another remains connected.
- Disconnect Valkey and prove moves fail explicitly rather than split-brain.
- Verify no unintended host ports and no public dashboard.
- Inspect container privileges, mounts, networks, secret paths, and image pins.
- Run the clean-machine install and rollback procedures from documentation.


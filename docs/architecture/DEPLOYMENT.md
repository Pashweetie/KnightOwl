# Deployment Architecture

Terminology used below: application programming interface (API), continuous
integration (CI), Hypertext Transfer Protocol (HTTP), personal computer (PC),
pull request (PR), Secure Shell (SSH), and software bill of materials (SBOM).
GitOps means that reviewed Git state declares the desired deployed state.

## Decision

Use Docker Compose for local development and self-hosted single-node K3s for
production. K3s is a lightweight Kubernetes distribution and includes Traefik
for ingress routing. Cloudflared is the only public ingress path.

Flux is a Kubernetes controller that reads the reviewed `release` branch and
applies its desired state. An authenticated GitHub push event requests immediate
reconciliation, while periodic polling recovers from a missed event. Argo
Rollouts is an established Kubernetes controller that performs the blue-green
application switch; KnightOwl does not implement deployment switching.

Kustomize, the configuration-composition feature built into the Kubernetes
command-line tool, manages KnightOwl resources using readable files and
small overlays. Helm is Kubernetes package management: it installs versioned
packages called charts. Helm is used only when a third-party component's
maintained upstream installation supports it; KnightOwl does not need its own
chart initially.

Kerberos is not part of production: it authenticates network identities and is
not a deployment system or container load balancer.

## Topology

```text
friends
  |
Cloudflare edge
  |
named Cloudflare Tunnel (outbound connection only)
  |
cloudflared Pod
  |
Kubernetes guest Service
  |
Traefik
  |----------------------|
active app          preview app
  |----------------------|
Kubernetes service network
  |
Valkey and Stockfish workers
```

No application, Valkey, metrics, Kubernetes API, or Traefik dashboard port is
published on the host. Docker Compose development binds the operator service to
loopback port `8787`, because port `8080` is already occupied.

## Environments and authority

KnightOwl has three explicitly different environments:

1. **Local development** runs with Docker Compose on the contributor's computer,
   may rebuild source, binds the operator surface only to loopback, and uses
   disposable development data. An opt-in invite-tunnel profile may expose only
   its public guest surface as specified below.
2. **Pull-request review** occurs in the GitHub repository. Ticket
   branches cannot deploy production.
3. **Production** runs reviewed immutable images in K3s on the operator's PC.
   Cloudflared may expose Traefik only after the publication gate passes.

There is no automatic deployment on file save, test completion, commit, or
ordinary `dev` commit. A successful test is evidence, not deployment authority.

## Docker Compose invite-only tunnel

The local Docker Compose installation has two independent network surfaces:

```text
operator browser
  -> 127.0.0.1:8787
  -> operator service
       -> create room, create/revoke invites, local administration

invited friend's browser
  -> Cloudflare Tunnel
  -> public guest service
       -> redeem existing invite, play, spectate, reconnect
```

The operator service is bound to `127.0.0.1:8787` and is absent from the Docker
network reachable by cloudflared. The public guest service has no host port and
is reachable only by cloudflared and the minimum internal dependencies. It does
not register room-creation, invite-creation, administration, diagnostics,
metrics, or review routes. This is a server routing and network boundary, not a
hidden button or client-side authorization check.

The profile is disabled by default and must be started deliberately. It may use
either a dedicated development hostname on a named Cloudflare Tunnel or an
explicitly temporary quick-tunnel hostname. The operator interface displays the
actual active public base URL before generating shareable links; it never
invents or assumes a hostname.

An invite URL uses the public guest origin and places the capability in the URL
fragment. The initial HTTP request therefore receives only the guest shell;
browser code removes the fragment from history and redeems the capability over
same-origin HTTPS. Visiting the public root without a valid capability shows a
generic invite-required page. It cannot list rooms, create a room, request an
invite, or reveal whether a guessed room exists.

Host, white, black, and spectator links may be generated locally. A host may
open the resulting host link through the public guest surface to play away from
the operator PC, but that capability controls only its existing room. It never
inherits localhost room-creation authority.

Required isolation tests:

- public requests to every operator and creation route return the same
  non-revealing not-found response;
- the public container image or run mode has no creation handler registered;
- cloudflared cannot resolve or connect to the operator service;
- the operator port listens only on loopback and is unreachable through
  Traefik and Cloudflare;
- a valid invite can be claimed, played, spectated, reconnected, and revoked;
- invalid, expired, reused, and role-swapped capabilities reveal no room data;
- capability fragments do not appear in Cloudflare origin requests, access
  logs, error output, metrics, referrers, or browser history after redemption;
- stopping the profile removes public reachability without stopping the
  loopback-only development environment;
- public security scanning confirms that no room-directory, creation,
  administration, diagnostics, metrics, or Docker interface is exposed.

The production K3s design should reuse the same operator/public application
boundary, but this Compose profile is independently usable and testable; it
does not deploy to or require Kubernetes.

## Git branches and production trigger

The GitHub repository has two long-lived protected branches:

- `test/<ticket-id>-<short-name>` is a temporary branch for one feature or fix.
  Work may be committed and pushed there while it is implemented and tested.
- `dev` is the integration branch. A ticket enters `dev` only through a GitHub
  PR after the designated reviewer approves it and all applicable CI checks
  pass.
- `release` is the production desired-state branch. Direct human pushes and
  force pushes are prohibited. A reviewed promotion from `dev` selects an
  already built, tested, scanned, immutable image digest and the exact rendered
  Kubernetes configuration.

Merging a promotion into `release` is explicit production authorization. That
branch event starts deployment; no separate application-code commit is created
by the deployment system.

The event and recovery path is:

```text
ticket branch
  -> approved pull request into dev
  -> CI builds/tests/scans image identified by commit
  -> approved promotion pull request selects immutable image digest
  -> merge into release emits authenticated GitHub push event
  -> Flux inside K3s immediately reads release
  -> Argo Rollouts starts new version behind private preview Service
  -> readiness and two-browser preview smoke checks
  -> active Service switches to new version, or rollout aborts
```

Flux has read-only repository credentials and runs inside K3s. The Kubernetes
API remains private; GitHub-hosted jobs receive no cluster credential. A Flux
webhook receiver validates the GitHub event secret and accepts only relevant
push events. Its unguessable path is exposed through a dedicated Cloudflare
route with request-size, method, and rate limits. Polling remains the recovery
path if an event is delayed or missed. Replay, invalid signature, wrong branch,
event flood, secret rotation, and receiver outage tests are mandatory.

The deployment does not select a mutable image tag such as `latest`, build
source on the production PC, or permit automatic registry scanning to choose an
unreviewed version. The `release` commit records the image digest that production
must run.

Only one production reconciliation and smoke test may be active at a time. A
newer release queues rather than cancelling a rollout already modifying
production. Documentation-only changes can reconcile safely but cannot silently
change an image or runtime object.

Argo Rollouts keeps normal traffic on the blue active version while the green
preview version starts. Pre-promotion analysis runs health checks and a real
two-browser smoke game against the preview Service. Failure aborts without
switching normal traffic. After success, the active Service switches to green
and blue remains available for a bounded rollback window. Persistent-data
compatibility is checked separately because switching application traffic
cannot undo stored data.

The `release` branch remains the desired-state record. Rollout status and a
revert PR reconcile any lasting rollback rather than leaving hidden production
drift.

Branch controls, GitHub workflow behavior, Flux permissions, and rollback are
verified as repository tests where possible. GitHub plan features are not
assumed: any unavailable private-repository protection is enforced by the
ticket-branch process and least-privileged automation, and the limitation is
documented before publication.

## Ticketed delivery lifecycle

Deployment capabilities are delivered as separate ticket branches and PRs in
this order:

1. **Compose foundation:** pinned development services, private networks,
   health checks, a loopback-only operator surface, an optional invite-only
   tunnel profile, and clean removal.
2. **Image supply chain:** reproducible multi-stage builds, non-root runtime,
   digest pins, dependency inventory, SBOM, vulnerability scan, and provenance.
3. **Preflight and release bundle:** validate configuration, available disk and
   memory, required secrets, port conflicts, compatible state format, and
   rollback availability without changing production.
4. **Branch promotion and GitOps reconciliation:** build on `dev`, explicitly
   select a reviewed image digest in a `release` promotion, authenticate the
   GitHub event, and let in-cluster Flux reconcile it.
5. **Blue-green deployment and rollback:** use Argo Rollouts to start a private
   preview, run readiness and two-browser smoke analysis, switch traffic only
   after success, and retain the previous application version for a bounded
   rollback window.
6. **Named tunnel publication:** accept an operator-supplied secret outside Git,
   verify the intended hostname, publish only Traefik, and run the public smoke
   and exposure checks.
7. **Remote administration, only if needed:** deploy through a private network
   and a restricted SSH identity. It must invoke the same local rollout command
   rather than maintain a second deployment path.

Each ticket has acceptance criteria, tests, rollback behavior, and a GitHub PR.
Commits are allowed on its `test/<ticket>` branch but cannot enter `dev` until
the designated reviewer approves the PR. No `dev` commit deploys merely because
it exists.
Merging the separately reviewed promotion PR into `release` is the explicit
production action and deployment begins immediately.

## Release bundle and rollout

A release identifies:

- the source commit;
- exact image names and digests;
- toolchain and dependency lockfiles;
- generated SBOM and license notices;
- configuration-format and game-state compatibility;
- database backup requirements, if any;
- the previous known-good release;
- checks and results needed before and after rollout.

The release tooling accepts an explicit release identifier and supports `plan`,
`promote`, `status`, and `rollback` operations. `plan` is read-only and shows
image, configuration, state, resource, and service changes. `promote` prepares
the reviewed change into `release`; merging that promotion is deliberate
confirmation and emits the event that authorizes Flux to apply it. Production
never builds from a mutable checkout.

Blue-green deployment avoids switching ordinary traffic until preview checks
pass. Existing WebSockets remain attached to the old application until their
games end or the documented drain deadline is reached; new connections use the
new active version after promotion. Protocol and state-format compatibility are
required across the overlap. Emergency termination is explicit rather than
pretending continuity.

## Optional Kerberos learning lab

Kerberos is neither required nor recommended for the initial production
deployment. An optional, separately approved learning lab may follow completion
of the ordinary local deployment and rollback path.

The lab uses disposable principals, an isolated test realm, test containers or
virtual machines, and no production tunnel token, host identity, or production
keytab. Its goal is to authenticate a restricted remote deployment identity
before invoking the normal deployment command. It does not add Kerberos to the
KnightOwl application, replace SSH, create a second rollout implementation,
or become a prerequisite for playing chess.

The lab must cover key rotation, expiry, clock drift, hostname mismatch,
unavailable authentication service, least privilege, credential cleanup, and a
documented complete teardown. Promotion into production would require a later
architecture decision demonstrating an existing managed Kerberos environment
or a concrete benefit that exceeds its operational cost.

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

Kubernetes checks readiness, removes unready Pods from Services, and replaces
crashed Pods. These mechanisms do not replace application-level state versioning
or reconnect logic.

## Scaling workflow

Only after multiple application replicas are justified, the reviewed Kustomize
configuration changes the replica count. K3s schedules the resulting Pods and
Traefik routes through the stable Service.

Moving from single-node to multi-node K3s is a separate architecture and storage
decision. Adding worker nodes without independent control-plane and state
recovery does not create high availability.

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
- Kill Traefik and verify Kubernetes replacement; document the brief ingress
  outage.
- Trigger a `release` event and prove Flux selects only that reviewed commit.
- Deploy a failing green preview and prove Argo Rollouts keeps blue active.
- Pass preview analysis, switch to green, and prove bounded rollback to blue.
- Kill one cloudflared replica while another remains connected.
- Disconnect Valkey and prove moves fail explicitly rather than split-brain.
- Verify no unintended host ports and no public dashboard.
- Inspect container privileges, mounts, networks, secret paths, and image pins.

## Operator inputs required before publication

Publication requires a controlled domain and Cloudflare account, a named-tunnel
credential supplied outside Git, the intended public hostname, tested room and
engine limits, room-retention choices, a tested Valkey persistence policy, and
the selected aggregate metric retention. An optional outage contact method is
operated outside KnightOwl unless message retention is separately approved.

Safe protocol, isolation, logging, readiness, and failure defaults are
engineering responsibilities rather than operator configuration choices.
Persistent accounts, public discovery, stored chat, durable archives, prizes,
and multi-host operation require separate product and architecture approval.
- Run the clean-machine install and rollback procedures from documentation.

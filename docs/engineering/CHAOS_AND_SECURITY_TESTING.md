# Chaos and Security Testing

## Tools

- Toxiproxy injects Valkey and internal-worker latency, resets, bandwidth
  limits, and outages.
- Docker Compose lifecycle commands kill, pause, restart, and scale containers.
- k6 drives HTTP/WebSocket capacity and soak tests.
- Playwright verifies independent browsers, reconnects, and accessibility.
- OWASP ZAP scans the running same-origin application and WebSocket endpoints.
- Nuclei runs a pinned, reviewed template set for exposed configuration errors.
- Trivy scans dependencies, filesystems, SBOMs, and release images.
- A secret scanner and backend-specific SAST run in CI.
- Built-in Rust/Go fuzzing targets untrusted protocol and chess inputs.

Pumba is optional only if it adds a failure that Compose plus Toxiproxy cannot
produce. No tool is included merely to make the list look comprehensive.

## Required chaos matrix

| Fault | Assertion |
|---|---|
| Kill one app replica | Sockets reconnect through Traefik; accepted chain remains identical |
| Drain one app replica | No new upgrades reach it; active clients receive reconnect instruction |
| Add Valkey latency/jitter | Backpressure is bounded; command order and clocks remain correct |
| Drop Valkey connection after request | Retry returns original result or one explicit uncertainty/resync path; never a duplicate |
| Stop Valkey | Game mutations fail closed; replicas never fork state |
| Restart Valkey | AOF boundary is measured; recovered clients resync or receive honest terminal loss |
| Fill Valkey memory | Admission stops before eviction can silently remove active games |
| Kill Stockfish worker | Human games unaffected; job fails/retries within declared policy |
| Saturate Stockfish | Fair queue and resource caps protect move traffic |
| Kill Traefik | Compose restart behavior and public reconnect are measured |
| Kill one cloudflared replica | Remaining connector serves if configured |
| Fill application/Valkey disk | Readiness and alerts fire before uncontrolled corruption |
| Scale app 1→3→1 | No sticky-session dependence, split brain, or lost accepted command |

Each scenario has a reproducible command, precondition, observable invariant,
time bound, cleanup, and captured report. Passing once is not enough for flaky
concurrency tests; release CI repeats the critical command races.

## Minimum penetration scope

- Capability entropy, role separation, constant-time comparison, revocation,
  replay, confused deputy, invite reuse, and capability leakage.
- Exact Origin validation, cross-site WebSocket hijacking, CSRF on HTTP
  mutations, CORS, method/content-type confusion, and cache poisoning.
- Message size/depth/rate limits, malformed binary/JSON frames, decompression
  abuse, reconnect storms, and idempotency-table exhaustion.
- Illegal moves, stale versions, forged clocks, duplicate commands, sequence
  gaps, hash-chain changes, and spectator privilege escalation.
- FEN/PGN parser fuzzing, path traversal, command injection, SSRF, unsafe URL
  fetching, and Stockfish process argument isolation.
- CSP, clickjacking, MIME sniffing, DOM XSS, dependency supply chain, source
  maps, secret leakage, and error detail.
- Traefik dashboard and metrics exposure, Docker socket/proxy permissions,
  unintended host ports, container capabilities, writable mounts, and
  cloudflared secret handling.
- Canary privacy fields absent from logs, metrics, traces, crash output, and
  browser error reports.

## Release evidence

The exact built image digests are scanned and tested. Reports record tool and
template versions, configuration, target commit, findings, suppressions with
expiry, and remediation. A scanner exit code alone is not evidence that the
right authenticated routes, WebSockets, or container topology were exercised.


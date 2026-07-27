# Repository Quality Contract

## Intended layout

```text
apps/
  web/             TypeScript browser application
  server/          selected Rust or Go API/WebSocket service
  engine-worker/   Stockfish process adapter and admission control
packages/
  protocol/        canonical schema and cross-language fixtures
  chess-fixtures/  perft, rules, clock, PGN, and fuzz corpora
deploy/
  compose/         pinned production and test profiles
  traefik/         static/dynamic proxy configuration
  cloudflared/     ingress template without secrets
tests/
  e2e/             real-browser multi-user tests
  chaos/           Toxiproxy and container-failure scenarios
  security/        ZAP/Nuclei policies and capability probes
docs/
```

Generated artifacts never become the source of truth. Protocol schemas produce
language bindings and canonical fixtures, and CI fails on an uncommitted
generation diff.

## Boundaries

- Chess rules and clocks are pure domain code with no WebSocket, Valkey,
  framework, system-clock, or Stockfish import.
- The game-state repository exposes semantic atomic operations, not generic
  key/value access.
- HTTP and WebSocket handlers translate transport types into domain commands;
  they do not implement rules.
- Only the engine adapter executes Stockfish.
- Browser-local persistence has no server synchronization backdoor.
- Deployment configuration cannot grant application containers Docker API
  access.

Dependency-boundary tooling is selected for the winning backend and
dependency-cruiser is used for TypeScript. Knip rejects unused frontend exports
and dependencies.

## Required checks

Every change runs the relevant subset; the protected main branch requires all:

1. deterministic formatting and linting with warnings as errors;
2. dependency-boundary and unused-code checks;
3. unit tests;
4. chess perft, rule fixtures, clock model tests, and PGN round trips;
5. property/model tests for generated command sequences;
6. real-Valkey integration tests, including functions and TTLs;
7. protocol compatibility and canonical-hash fixtures in browser and backend;
8. multi-browser Playwright tests;
9. fuzz targets for protocol, FEN, PGN, and command state machines;
10. coverage thresholds on changed code and mutation testing for chess/clock
    invariants;
11. dependency, license, secret, SBOM, and container scans;
12. Compose validation and image smoke tests;
13. documentation-link and architecture-consistency checks.

Rust additionally uses `cargo fmt`, Clippy, locked builds, `cargo deny`, nextest,
Miri where supported, and sanitizer/fuzz jobs. Go additionally uses `gofmt`,
`go vet`, staticcheck, `go test -race`, fuzzing, and vulnerability checks.

## Test truthfulness

- Tests never replace Valkey with an in-memory fake for atomic/TTL behavior.
- Engine fixtures record the exact Stockfish build and settings.
- End-to-end tests use two independent browser contexts and real WebSockets.
- Clocks use an injected deterministic time source in unit/model tests.
- Random and property failures print replayable seeds.
- No test seed appears in production or demo views.

## Review rules

- Changes cite the task and acceptance criterion they satisfy.
- Protocol/state changes include compatibility and rollback notes.
- Dependencies require a stated purpose, maintained upstream, acceptable
  license, and security review.
- Image and lockfile upgrades are isolated where practical.
- Merged commits remain buildable; temporary spike code is deleted after the
  backend decision.
- Exceptions to a gate are time-bounded, documented, and cannot waive chess,
  privacy, capability, or state-consistency tests.


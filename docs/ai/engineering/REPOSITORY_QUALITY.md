# Repository Quality Contract

Terminology used below: application programming interface (API), continuous
integration (CI), Forsyth–Edwards Notation (FEN), Hypertext Transfer Protocol
(HTTP), Portable Game Notation (PGN), software bill of materials (SBOM), and
time to live (TTL).

## Library-first gate

Every code change begins with a dependency survey. Before custom code is
approved, its design note records:

1. the capability being solved;
2. maintained libraries and standards evaluated;
3. license compatibility and transitive-license risk;
4. release activity, maintainer depth, security policy, and known advisories;
5. bundle or runtime cost and supported platforms;
6. whether the public programming interface can be isolated behind an adapter;
7. the exact unmet requirement, if custom code remains necessary.

Prefer a small adapter around a proven library. Prefer a complete, consistently
supported platform standard where one exists. Custom implementation requires
an Architecture Decision Record and tests derived from the relevant standard
or upstream conformance suite. Review rejects code that skips this gate.

Dependencies are pinned with lockfiles, checked for known vulnerabilities and
license changes, updated regularly, and kept replaceable through narrow
adapters. Releases include generated dependency inventories and notices.

## Intended layout

```text
apps/
  web/             TypeScript browser application
  server/          selected backend application and WebSocket service
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
  private/         locally ignored operator validation material
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

## Merge-blocking checks

Every pull request runs the applicable checks below. Required checks fail on
warnings, skipped required tests, threshold misses, scanner findings above the
documented severity limit, or missing evidence. A pull request cannot merge
until every applicable check succeeds:

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
11. dependency, license, and private release-review gates;
12. Docker Compose and rendered Kubernetes validation plus image smoke tests;
13. documentation-link and architecture-consistency checks.

Language-specific checks are selected with the backend. Go candidates use
`gofmt`, `go vet`, staticcheck, `go test -race`, and fuzzing. TypeScript
candidates use deterministic formatting, strict compiler
settings, linting, dependency-boundary checks, unused-code checks, and
property/fuzz tests.

## Test truthfulness

- Tests never replace Valkey with an in-memory fake for atomic/TTL behavior.
- Engine fixtures record the exact Stockfish build and settings.
- End-to-end tests use two independent browser contexts and real WebSockets.
- Clocks use an injected deterministic time source in unit/model tests.
- Random and property failures print replayable seeds.
- No test seed appears in production or demo views.

## Review rules

- Every feature, fix, dependency change, and unrelated cleanup has a ticket and
  follows `TICKET_BRANCH_WORKFLOW.md`.
- Work is committed only to its `test/<ticket>` branch before review. It enters
  `dev` only through an approved GitHub pull request.
- The pull request includes the complete diff, test evidence, dependency
  decisions, risks, and rollback notes; a prose summary is not a substitute for
  reviewing changed files.
- Changes cite the task and acceptance criterion they satisfy.
- Protocol/state changes include compatibility and rollback notes.
- Dependencies require a stated purpose, maintained upstream, acceptable
  license, and security review.
- Image and lockfile upgrades are isolated where practical.
- Merged commits remain buildable; temporary spike code is deleted after the
  backend decision.
- Exceptions require their own reviewed policy change before the affected pull
  request. Chess, privacy, capability, state-consistency, secret scanning, and
  branch-protection checks cannot be waived.

Advisory reports may provide additional context, but no tool designated as a
quality gate may use a warning-only or continue-on-error configuration.

# Architecture Decisions

This file contains active decisions only. Detailed contracts:

- [System design](docs/architecture/SYSTEM_DESIGN.md)
- [Minimal game-state protocol](docs/architecture/MINIMAL_GAME_STATE_PROTOCOL.md)
- [Data platform and scale path](docs/architecture/DATA_PLATFORM.md)
- [Self-hosted database options](docs/architecture/SELF_HOSTED_DATABASES.md)
- [Backend language selection](docs/architecture/BACKEND_LANGUAGE.md)
- [Docker, Traefik, and Cloudflare deployment](docs/architecture/DEPLOYMENT.md)
- [Privacy and room identity](docs/architecture/PRIVACY_AND_IDENTITY.md)
- [Low-data authentication research](docs/architecture/LOW_DATA_AUTHENTICATION.md)
- [Peer-to-peer second-phase plan](docs/architecture/PEER_TO_PEER.md)
- [Operator decisions](docs/architecture/OPERATOR_DECISIONS.md)
- [Telemetry and retention](docs/architecture/TELEMETRY_AND_RETENTION.md)
- [XMage identity research](docs/research/XMAGE_IDENTITY.md)

## ADR-001: Friends-only product boundary

The first deployment is an unlisted chess service for invited friends. It has
no public registration, public directory, permanent profile, follower graph,
direct messages, public forums, cheating review, or advertising/behavioral
tracking.

Free-text social features are excluded; room communication uses fixed
reactions. Host controls and capability revocation replace community-scale
moderation. Premium-equivalent chess, engine, review, learning, puzzle, and
private tournament capabilities remain in scope, with progress/history stored
locally unless the user later authorizes persistent identity.

## ADR-002: Capability roles, not accounts

Host, white, black, and spectator access use separate high-entropy,
room-scoped capabilities. Only keyed hashes are stored. A claimed player invite
produces a rotating reconnect capability. Display labels are cosmetic,
room-local, non-unique, and expire with the room.

No username, password, email, OAuth/OIDC subject, or global player ID is stored.
WebAuthn passkeys are the leading future option only if a later approved feature
genuinely needs continuity across rooms.

## ADR-003: Minimal server authority

The browser renders and independently verifies the game, but the server
arbitrates seats, move order, chess rules, clocks, results, and active
tournament state. Matching browser hashes do not decide legality or races.

Commands use expected versions, unique IDs, and canonical previous hashes.
Accepted events form a BLAKE3 chain. Duplicate commands return their original
result; stale clients resynchronize.

## ADR-004: Valkey is the only initial server database

Valkey stores expiring room partitions, projections, event chains,
capabilities, idempotency records, presence, rate buckets, private tournaments,
and bounded analysis cache. One atomic Valkey function validates and changes
one game partition.

The initial deployment contains no SQL database. Browser IndexedDB stores
opt-in local PGNs, analysis, puzzle schedules, lesson progress, and preferences.
Stockfish needs no database.

## ADR-005: Scale is designed but not installed prematurely

Game keys use Valkey Cluster hash tags and repository operations map to one
partition. Scaling proceeds from app replicas, to Valkey replication/Sentinel
on real failure domains, to Valkey Cluster only at measured capacity limits.

ScyllaDB is the leading future Dynamo-style store only if durable multi-host
archives or recovery become approved requirements. ClickHouse is allowed only
for privacy-reviewed aggregate operational analytics at proven volume. Neither
is in the initial Compose stack.

## ADR-006: Backend language is selected by evidence

The browser uses TypeScript. Rust and Go implement the same production-shaped
room/WebSocket/Valkey/reconnect spike. Correctness, resource measurements,
operational behavior, dependency surface, and maintenance cost select the
backend. Rust is the leading candidate, not a predetermined result.

## ADR-007: Server deployment precedes P2P

Authoritative WebSocket multiplayer is implemented, qualified, and deployed
first. WebRTC is a planned second transport that must reuse the same canonical
commands and events, preserve server fallback, and pass NAT/TURN, privacy,
authority-transfer, reconnect, and spectator tests.

P2P is never required to play because some networks require an external TURN
relay and direct ICE can expose peer addresses.

## ADR-008: Traefik balances local containers

Docker Compose runs the single-PC deployment. Cloudflared connects only to
Traefik, which discovers explicitly labelled app replicas and performs
readiness-aware HTTP/WebSocket balancing. A restricted Docker API proxy avoids
a direct socket mount. The dashboard and internal services are not publicly
exposed.

An optional development binding uses loopback port `8787`; port `8080` is not
used.

## ADR-009: Same-origin security

Web, HTTP API, and WebSocket share one HTTPS origin. HTTP mutations use
SameSite/HttpOnly/Secure room-session cookies plus CSRF protection. WebSocket
upgrades validate exact Origin, cookie, protocol version, message schema,
authorization, size, rate, sequence, and idempotency.

Secrets never enter Git, images, browser bundles, URLs after capability claim,
logs, metrics, traces, or task documents.

## ADR-010: Server clocks and explicit failure

Clients render estimated clocks only. The app calculates deadlines from
monotonic time and stores recovery anchors. Move-versus-timeout races are
serialized by the same atomic game transition.

When Valkey is unavailable, no app replica accepts a move. Acknowledged state
recovers within the tested AOF policy; losses outside that boundary are
reported honestly rather than reconstructed from conflicting clients.

## ADR-011: Stockfish truth and isolation

Pinned Stockfish/NNUE workers run outside the app with CPU, memory, time,
concurrency, and queue limits. Returned analysis includes engine/network
identity, settings, score, nodes/depth/time, and MultiPV lines.

An educational bounded minimax/alpha-beta visualization is separate and is
never presented as Stockfish's internal NNUE/PVS search. Coach prose must be
grounded in legal variations or deterministic chess features.

## ADR-012: Original or licensed content only

KnightShift copies no Chess.com brand, source, UI assets, lesson text, puzzle
collection, labels, rating data, or proprietary explanations. Lessons are
original. Puzzles/opening data have recorded redistribution provenance.
Stockfish license/source obligations are included with distributions.

## ADR-013: Vertical, test-gated delivery

Work follows `TASKS.md` in order. Each feature includes domain behavior,
transport, UI, authorization, accessibility, telemetry, failure behavior,
tests, and documentation. A decorative or backend-only implementation does not
complete a feature; unavailable functionality is absent from the UI.

## ADR-014: Versioned rolling compatibility

Wire and stored-state formats carry explicit versions. Rolling releases read
the current and immediately prior version. Stored-state changes use
expand/migrate/contract or wait for old room TTLs. App drain stops new upgrades,
allows bounded completion, and makes clients reconnect through Traefik.

## ADR-015: Purpose-limited telemetry, not a no-law claim

Anonymous aggregate operational metrics are enabled. Short-lived,
pseudonymous security/diagnostic events may be retained for a fixed period.
Gameplay behavior tracking is disabled by default and requires a separate
review. Advertising, sale/sharing, third-party trackers, fingerprinting,
contacts, sensitive data, and hidden cross-game profiles are prohibited.

The product makes no promise that this eliminates legal obligations. It
documents actual fields, purposes, processors, retention, access, and deletion,
and tests that runtime behavior matches those claims.

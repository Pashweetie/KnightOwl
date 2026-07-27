# Architecture Overview

Terminology used below: Architecture Decision Record (ADR), append-only file
(AOF), application programming interface (API), Concise Binary Object
Representation (CBOR), central processing unit (CPU), cross-site request
forgery (CSRF), Hypertext Transfer Protocol (HTTP), Hypertext Transfer Protocol
Secure (HTTPS), Interactive Connectivity Establishment (ICE), identifier (ID),
Network Address Translation (NAT), Efficiently Updatable Neural Network (NNUE),
OpenID Connect (OIDC), peer-to-peer (P2P), personal computer (PC), Principal
Variation Search (PVS), Traversal Using Relays around Network Address
Translation (TURN), Web Authentication (WebAuthn), Web Real-Time Communication
(WebRTC), and user interface (UI). BLAKE3 is a proper algorithm name.

This file contains active decisions only. Detailed contracts:

- [System design](SYSTEM_DESIGN.md)
- [Minimal game-state protocol](MINIMAL_GAME_STATE_PROTOCOL.md)
- [Data platform and scale path](DATA_PLATFORM.md)
- [Future database scale options](../future/DATABASE_SCALE_OPTIONS.md)
- [Docker Compose development, K3s production, and Cloudflare deployment](DEPLOYMENT.md)
- [Privacy and room identity](PRIVACY_AND_IDENTITY.md)
- [Authentication options research](../research/AUTHENTICATION_OPTIONS.md)
- [Future peer-to-peer transport](../future/PEER_TO_PEER_TRANSPORT.md)
- [Telemetry and retention](TELEMETRY_AND_RETENTION.md)
- [Existing chess test ecosystem](../research/CHESS_TEST_ECOSYSTEM.md)
- [Private tournament contract](../product/TOURNAMENTS.md)
- [Pseudonymous identity research](../research/PSEUDONYMOUS_IDENTITY_RESEARCH.md)

## ADR-001: Friends-only product boundary

The first deployment is an unlisted chess service for invited friends. It has
no public registration, public directory, permanent profile, follower graph,
direct messages, public forums, cheating review, or advertising/behavioral
tracking.

Free-text social features are excluded; room communication uses fixed
reactions. Host controls and capability revocation replace community-scale
moderation. Premium-equivalent chess, engine, review, puzzle, drill, and
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

The initial deployment contains no Structured Query Language database. Browser
Indexed Database storage holds opt-in local Portable Game Notation records,
analysis, puzzle schedules, drill progress, and preferences.
Stockfish needs no database.

## ADR-005: Scale is designed but not installed prematurely

Game keys use Valkey Cluster hash tags and repository operations map to one
partition. Scaling proceeds from app replicas, to Valkey replication/Sentinel
on real failure domains, to Valkey Cluster only at measured capacity limits.

ScyllaDB is the leading future Dynamo-style store only if durable multi-host
archives or recovery become approved requirements. ClickHouse is allowed only
for privacy-reviewed aggregate operational analytics at proven volume. Neither
is in the initial deployment.

## ADR-006: Go backend

The browser uses TypeScript and the backend uses Go. Readability, maintained
library coverage, testing, debugging, and single-computer deployment are
weighted above raw speed. Ten concurrent friends is the acceptance load; large
synthetic capacity is not a selection gate. The initial Go vertical slice is
developed on its own ticket branch and enters `dev` only through an approved
GitHub pull request. Repository rulesets must block direct and force pushes to
`dev` and `release` before implementation begins. Rust is deferred unless a future measured constraint
justifies its added review cost.

Maintained libraries provide chess rules/notation, WebSockets, Valkey access,
deterministic CBOR, hashing, metrics, and testing infrastructure. KnightOwl
writes domain policy and adapters, not substitutes for established libraries.

No production code is written until the author has searched for maintained
libraries that solve the problem. The decision record names the candidates,
licenses, maintenance and security signals, fit gaps, and the reason for
adopting a library or writing the smallest necessary custom adapter. “Small
enough to write ourselves” is not sufficient justification.

The browser follows the same rule. KnightOwl composes an established chess
board, user-interface primitives, charts, validation, and testing libraries. It
does not invent a display protocol, component framework, drag-and-drop system,
dialog, menu, tooltip, chart renderer, or virtual list.

## ADR-007: Server deployment precedes P2P

Authoritative WebSocket multiplayer is implemented, qualified, and deployed
first. WebRTC is a planned second transport that must reuse the same canonical
commands and events, preserve server fallback, and pass NAT/TURN, privacy,
authority-transfer, reconnect, and spectator tests.

P2P is never required to play because some networks require an external TURN
relay and direct ICE can expose peer addresses.

## ADR-008: Docker Compose development and K3s production

Docker Compose runs local development. Its optional cloudflared profile exposes
only the invite guest surface; room creation and administration remain on
loopback port `8787`. Port `8080` is not used.

Production runs on a self-hosted, single-node K3s Kubernetes distribution.
K3s supplies Traefik for ingress and service routing. Flux, a controller that
keeps Kubernetes synchronized with reviewed Git state, watches `release`.
GitHub sends an authenticated push event to Flux after a `release` merge, with
polling as recovery for missed events.

Argo Rollouts, an established Kubernetes progressive-delivery controller,
performs blue-green application deployment. It starts the new version behind a
private preview Service, runs readiness and smoke analysis, and switches the
active Service only after success. The old version remains available for the
bounded rollback window. A single PC is not claimed to be highly available.

KnightOwl's readable Kubernetes resources use Kustomize, the configuration
composition tool built into `kubectl`. Helm, Kubernetes's package manager, is
used only when an established third-party component officially distributes and
supports a Helm package. KnightOwl does not create a custom Helm chart without
a demonstrated packaging need.

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
identity, settings, score, nodes/depth/time, and multiple principal variation
(MultiPV) lines.

An educational bounded minimax/alpha-beta visualization is separate and is
never presented as Stockfish's internal NNUE/PVS search. Coach prose must be
grounded in legal variations or deterministic chess features.

## ADR-012: Original or licensed content only

KnightOwl uses original or appropriately licensed brand, source, interface
assets, puzzles, opening data, labels, rating data, explanations, and visual
assets with recorded redistribution provenance.
Stockfish license/source obligations are included with distributions.

## ADR-013: Vertical, test-gated delivery

Work follows `docs/ai/delivery/ROADMAP.md` in order. Each feature includes domain behavior,
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

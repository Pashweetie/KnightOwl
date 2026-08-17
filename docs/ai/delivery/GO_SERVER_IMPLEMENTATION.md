# Go Server Implementation Plan

Terminology used below: central processing unit (CPU), continuous integration
(CI), Concise Binary Object Representation (CBOR), Hypertext Transfer Protocol
(HTTP), personal computer (PC), pull request (PR), and Universal Chess
Interface (UCI). BLAKE3 is a proper algorithm name.

Go is the selected backend language. The server is implemented as small,
production-intended increments. There is no disposable proof-of-concept server
and no ticket may claim behavior supplied only by a mock or placeholder.

Each increment requires an authorized ticket, a
`test/<ticket-id>-<short-name>` branch from current `dev`, focused
library-first evidence, applicable automated tests, one squashed commit, and a
PR into `dev`. Work follows the roadmap phase containing the behavior; this
plan does not authorize skipping ahead.

`server-001` is the first executable application ticket in roadmap phase 2,
after the ticket-branch workflow prerequisite.

## Ordered increments

### `server-001`: Service foundation

Create the Go workspace and a runnable server with explicit package boundaries,
configuration validation, privacy-safe structured logging, liveness, startup,
readiness, and graceful drain behavior. Build a minimal non-root container and
document local debugging. The server exposes no room or chess behavior yet and
does not pretend that health endpoints complete a user feature.

Adopt the Go standard HTTP library unless a documented gap requires a focused
dependency. Verify formatting, vetting, static analysis, race detection,
configuration failures, signals, drain deadlines, container user, and
dependency boundaries.

This ticket excludes rooms, WebSockets, chess, Valkey, and Stockfish. It stores
no user data. Rollback removes the new service and container without a
migration.

### `server-002`: Loopback room creation

Add the first usable operator behavior: create an expiring unlisted room only
through the loopback surface. Generate random host, white, black, and spectator
capabilities and store only keyed digests in real Valkey. Return claim links
with capabilities in fragments and never log capabilities, room positions,
labels, or personal network identifiers.

Adopt the official Valkey Go client behind a semantic room repository. Verify
real expiry, capability separation, entropy and encoding, storage inspection,
failure closure, log redaction, and absence of the creation route from the
guest surface.

This ticket excludes invite claiming and gameplay. Room records expire and
contain no account or contact identity. Rollback removes the creation route and
repository; disposable development keys expire or are deleted by the operator.

### `server-003`: Invite claim and WebSocket session

Let one invited participant claim the intended role over a same-origin
validated WebSocket and receive a rotating reconnect capability. Register only
invite claim, play, spectate, and reconnect routes on the guest surface.
Invalid, expired, reused, and role-swapped capabilities return the same
non-revealing failure.

Adopt `coder/websocket` only after the required protocol conformance check.
Verify origin and protocol versions, frame limits, strict schemas, rates,
one-use claims, concurrent claims, malformed and fuzzed input, and route and
network isolation.

This ticket excludes chess moves and public tunnel publication. It retains only
room-scoped capability digests and bounded presence required for the active
session. Rollback disables the guest routes and revokes or expires affected
room capabilities.

### `server-004`: Authoritative chess event

Accept one legal standard-chess move and reject an illegal move through a
qualified `corentings/chess` version 2 adapter. Encode the accepted event as
deterministic CBOR, hash it with BLAKE3, and commit it through one real Valkey
atomic operation that compares expected version and prior hash. A duplicate
command identifier returns the byte-identical original event; a losing
concurrent command receives stale state.

Use `fxamacker/cbor` and `zeebo/blake3`; do not implement chess rules,
serialization, hashing, or the Valkey protocol. Verify published chess
positions, canonical cross-language fixtures, two simultaneous commands,
idempotency, expiry refresh, Valkey failure, property sequences, fuzzing, and
race detection.

This ticket excludes clocks and the complete room action set. Events contain
game state but no account or contact identity and expire with the game. The PR
must document event-format compatibility; rollback disables new move
acceptance while retaining readable events until their normal expiry.

### `server-005`: Reconnect and process handoff

Reconnect a participant or spectator through another server process and replay
accepted events or a canonical snapshot. During drain, reject new upgrades,
notify existing sessions to reconnect, and close them by the deadline. No
process treats local memory as authoritative.

Verify multi-process reconnect, trimmed-event fallback, hash disagreement,
process exit, drain, real Valkey outage and restart, and absence of sticky
session assumptions.

This ticket excludes multi-host Valkey and production rollout automation. It
adds no durable identity or retention. Rollback restores the prior server
version only when its event and snapshot formats remain compatible; otherwise
the ticket must provide explicit development-state cleanup.

### `server-006`: Stockfish process adapter

Send one cancellable request to a real, pinned Stockfish process through UCI.
Keep engine execution outside the application boundary and return truthful
engine identity and settings. Apply CPU, memory, time, concurrency, and queue
limits so engine work cannot starve live games.

Verify startup failure, cancellation, timeout, malformed engine output,
overload, process exit, resource limits, and license/source obligations. This
increment belongs to the Bot Mode phase and is not a prerequisite for private
human multiplayer.

This ticket excludes bot personalities, teaching explanations, and review
views. Engine requests and bounded cache entries contain positions but no
player identity and expire according to the retention contract. Rollback
removes the worker and adapter after in-flight work is cancelled.

### `server-007`: Ten-friend qualification

After the preceding multiplayer increments exist, qualify the integrated
server on the target PC. Ten concurrent friends must create, claim, play,
spectate, and reconnect without error. Run a bounded soak and record CPU,
memory, live-move latency, build time, and image size as sanity evidence rather
than a scale competition.

The PR records exact commits and image digests, dependency and license
inventory, package map, build/test/debug/container commands, replayable seeds,
representative failure traces, reconnect and drain observations, profiler,
race, fuzz, vulnerability, and privacy findings. A correctness, readability,
boundary, or ten-friend failure creates a focused follow-up ticket; it does not
trigger an unreviewed rewrite.

This ticket adds no product behavior or retention. Test participants and
results must be real and clearly identified as qualification evidence, never
seeded into product views. Rollback removes generated test resources and
artifacts according to their documented cleanup procedure.

## Requirement trace

| Former combined-proof requirement | Owning increment |
| --- | --- |
| Health, startup, drain, non-root container, debugging | `server-001` |
| Loopback room creation, expiring rooms, keyed capability digests | `server-002` |
| Guest isolation, same-origin WebSocket claim, malformed input | `server-003` |
| Legal moves, canonical events, atomicity, duplicates, stale commands | `server-004` |
| Cross-process reconnect and graceful handoff | `server-005` |
| Cancellable pinned Stockfish request | `server-006` |
| Ten friends, bounded soak, integrated evidence and privacy audit | `server-007` |

Every former acceptance target is retained by its owning increment. Relevant
unit, integration, conformance, property, fuzz, race, security, and container
checks run in the earliest ticket that introduces the behavior and remain
regression gates afterward.

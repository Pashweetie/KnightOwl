# Minimal Game-State Protocol

## Constraint

KnightShift is for invited friends and does not need accounts, public ratings,
cheat detection, or a permanent server-side game archive. The browser owns the
board presentation, legal-move hints, analysis UI, and downloaded PGN.

The server must still arbitrate the small amount of live state that two clients
cannot safely agree on: seat ownership, move order, legal position, clock
deadlines, draw/resign actions, and the current version. Two matching
client-produced hashes prove agreement between those clients; they do not prove
that a move was legal or that a late clock command arrived first.

## Room and capability model

- A room identifier is random and unlisted.
- White, black, spectator, and host privileges use separate high-entropy
  capability tokens.
- Only a keyed hash of each capability is stored.
- Invite links are one-use where practical and may be revoked by the host.
- Player labels exist only inside a room and expire with it.

## Canonical command chain

Each client command contains protocol version, game identifier, expected
sequence, previous event hash, unique command identifier, command payload, and
the client's position hash. The server validates the capability, expected
version, chess rule, and clock before atomically accepting it.

The accepted event contains the authoritative sequence, command identifier,
move/action, resulting position hash, clock values, termination state, and a
BLAKE3 hash over the canonical binary encoding plus the previous event hash.
Both clients independently replay and verify this chain.

Duplicate command identifiers return the original accepted result. Stale
versions return the missing events. Conflicting commands at one version have
exactly one winner. A hash mismatch forces snapshot resynchronization and is
never silently merged.

## Storage

Active rooms and their compact event chains live in Valkey with atomic scripts
or transactions and explicit TTLs. Finished games remain briefly available for
reconnect and PGN download, then disappear. There is no gameplay database in
the initial release and no gameplay-state backup.

Clients may save PGN and local statistics in IndexedDB or download a file.
Clearing browser storage deletes local history. Local data is not silently
uploaded.

## Required tests

- Conflicting simultaneous moves accept exactly one move.
- Timeout versus arriving move resolves from server receive time and clock.
- Retries cannot duplicate a move, draw, resignation, or rematch.
- Reconnect on a different app replica rebuilds the same position and clock.
- Deliberately altered client hashes trigger resynchronization.
- Valkey loss fails closed and never creates two accepted histories.
- Room, capability hash, labels, and finished state actually expire.


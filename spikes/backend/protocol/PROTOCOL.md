# Spike Protocol v1

Terminology used below: American Standard Code for Information Interchange
(ASCII), Concise Binary Object Representation (CBOR), identifier (ID),
JavaScript Object Notation (JSON), time to live (TTL), user interface (UI),
uniform resource locator (URL), and 8-bit Unicode Transformation Format
(UTF-8). GET and POST are Hypertext Transfer Protocol method names, and BLAKE3
is a proper algorithm name.

This is a deliberately narrow infrastructure protocol. It proves transport,
capabilities, atomicity, canonical hashing, reconnect, and operations before
the full chess core is written.

## Encoding

External messages use UTF-8 JSON and reject duplicate keys, unknown fields,
non-integer numbers, invalid Unicode, nesting beyond 8, and frames over 4096
bytes. IDs are lowercase unpadded base64url or lowercase hexadecimal where
specified.

Canonical hashed events use deterministic CBOR, not JSON. Maps are forbidden in
the hashed structure; the event is the following fixed-order array:

```text
[
  1,                  # protocol version, uint
  game_id,            # 16 raw bytes
  sequence,           # uint
  previous_hash,      # 32 raw bytes
  command_id,         # 16 raw bytes
  actor,              # 0 white, 1 black
  move_uci,           # lowercase ASCII text
  resulting_fen,      # normalized ASCII text
  white_ms,           # uint
  black_ms,           # uint
  termination         # 0 ongoing
]
```

The event hash is BLAKE3 over the deterministic CBOR bytes. The initial
previous hash is 32 zero bytes.

## Room creation

`POST /v1/rooms` accepts:

```json
{"base_ms":300000,"increment_ms":0,"ttl_seconds":900}
```

The response contains the game ID and four claim URLs. Tokens occur in URL
fragments. The endpoint stores only keyed BLAKE3 capability digests in Valkey.
The spike permits all four URLs in one response because the caller is the room
creator; production UI displays/copies them individually.

## WebSocket

`GET /v1/socket` upgrades only with the configured exact `Origin`. The first
message claims a role:

```json
{"type":"claim","game_id":"...","role":"white","token":"...","after_sequence":0}
```

On success the server returns a snapshot and rotates a player reconnect token.
Subsequent commands use:

```json
{
  "type":"move",
  "game_id":"...",
  "command_id":"...",
  "expected_sequence":0,
  "previous_hash":"0000...",
  "move":"e2e4"
}
```

The infrastructure spike delegates move parsing, legality, resulting position,
and game outcome to the selected maintained chess library. It tests several
published positions and complete short games, but full conformance is still the
next delivery phase and cannot be claimed from the spike.

## Atomic results

- `accepted`: includes canonical event fields and hash.
- `duplicate`: contains the byte-identical originally accepted event.
- `stale`: contains current sequence/hash and events after the request version.
- `illegal`: state is unchanged.
- `unauthorized`: no state detail is disclosed.
- `unavailable`: Valkey mutation did not complete; the client must not assume
  acceptance and resynchronizes with the same command ID.

No application replica accepts a transition from local memory.

## Valkey partition

All keys include `{game-id}` so they share a future cluster slot:

```text
ks:game:{game-id}:state
ks:game:{game-id}:events
ks:game:{game-id}:commands
ks:game:{game-id}:caps
```

One persisted Valkey function compares capability generation, sequence,
previous hash, and command ID; appends the event; updates state; and refreshes
the bounded TTL atomically. The function receives the already validated
transition and canonical bytes. Full product code additionally verifies that
the stored starting projection matches the transition input.

## Health and drain

- `GET /health/live`: process can schedule work.
- `GET /health/ready`: process accepts new work and can reach Valkey.
- `POST /admin/drain`: loopback/private-operator only; readiness becomes false,
  new upgrades receive 503, sockets receive `server_draining`, then close by
  deadline.

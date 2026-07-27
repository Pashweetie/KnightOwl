# System Design

Terminology used below: append-only file (AOF), central processing unit (CPU),
Forsyth–Edwards Notation (FEN), identifier (ID), Efficiently Updatable Neural
Network (NNUE), personal computer (PC), Portable Game Notation (PGN), and
Structured Query Language (SQL). BLAKE3 is a proper algorithm name.

## Design target

KnightOwl runs on one privately operated PC for invited friends, accepts
public traffic only through an outbound Cloudflare Tunnel, stores no accounts
or contact information, and remains horizontally evolvable without operating a
distributed database prematurely.

## Runtime components

```text
browser
  |
Cloudflare edge and named Tunnel
  |
cloudflared
  |
Traefik
  |-------------------------|
app replica 1 ... app replica N
  |                    |
Valkey              Stockfish workers
```

- The TypeScript browser implements presentation, local history, move previews,
  PGN import/export, and analysis views.
- The selected backend implements room capabilities, authoritative rules and
  clocks, WebSocket protocol, game-state transitions, and worker admission
  control.
- Valkey stores only expiring rooms, compact event chains, capabilities,
  idempotency records, presence, rate buckets, room tournaments, and short-lived
  analysis cache.
- Stockfish workers have CPU/memory/concurrency limits and no database.
- Traefik performs readiness-aware local balancing.
- Cloudflared is the only ingress path; no router port-forwarding is required.

## Trust boundaries

The browser is untrusted even when all players are friends. This is primarily
for correctness under stale tabs, retries, packet reordering, and clock races,
not cheating enforcement.

The app accepts a game command only after:

1. exact Origin and protocol version validation;
2. capability verification;
3. schema, size, and rate validation;
4. expected sequence and previous-hash validation;
5. legal chess transition and clock calculation;
6. one atomic Valkey compare-and-set update.

Both browsers replay the returned canonical event and verify its BLAKE3 chain.
A disagreement causes snapshot recovery, never client-side majority voting.

## State model

Each game is a self-contained partition keyed by a random game ID:

- immutable settings and initial position;
- hashed seat capabilities;
- current FEN-compatible position plus repetition state;
- authoritative monotonic-clock anchors and remaining times;
- sequence and event-chain head;
- bounded idempotency table;
- ordered compact events;
- result and expiry.

All keys for one game use the same Valkey Cluster hash tag so future clustered
atomic functions remain single-slot. Room tournaments follow the same rule:
one tournament partition owns entrants, pairings, results, and scores.

## Time and clocks

The process uses a monotonic clock while running. Stored events include
remaining duration and a wall-clock recovery anchor. On restart or ownership
change, the app calculates a conservative deadline under a documented maximum
drift. No client timestamp decides a timeout. A timeout and move racing at the
same version are serialized by the state function.

## Reconnect and replica changes

The client retains only capability, last accepted sequence, and chain hash.
After reconnecting to any app replica it requests events after that sequence.
If the gap was trimmed or hashes differ, it receives a canonical snapshot plus
the subsequent chain.

Application replicas do not own durable game actors and Traefik does not use
sticky sessions. A graceful drain stops new upgrades, tells clients to
reconnect, allows a short completion period, and exits.

## Bot and review path

The browser submits a position and bounded analysis request. The app validates
the position and reserves worker capacity. A Stockfish worker returns engine
version, NNUE identifier, depth/nodes/time, score, and MultiPV lines. The app
may cache this response briefly by position plus engine settings.

There is no SQL write and no permanent analysis job. The browser can retain the
analysis locally. Interactive bot moves use the same worker pool with stricter
latency budgets and fair scheduling.

## Failure behavior

| Failure | Required behavior |
|---|---|
| App replica exits | Socket reconnects through Traefik; state reloads from Valkey. |
| App becomes unready | Traefik stops new requests/upgrades to it. |
| Valkey unavailable | Moves and room mutations fail closed; no replica invents state. |
| Valkey restarts | AOF restores within the declared persistence boundary; clients resync or receive an explicit ended/lost state. |
| Stockfish saturated | Human games remain responsive; bot/analysis requests queue briefly or return a real capacity error. |
| Traefik exits | K3s replaces it; existing public connections reconnect after a brief ingress outage. |
| Cloudflared exits | A replica continues if configured; otherwise local service remains healthy but unreachable publicly. |
| Disk fills | Readiness fails before corrupting state; operator alert identifies the exact volume. |

## Scale path

1. One app, one Valkey, one Stockfish worker group.
2. Multiple app replicas behind Traefik; still one Valkey.
3. Valkey replica plus Sentinel for failover on separate failure domains.
4. Valkey Cluster only after memory/throughput measurements require sharding.
5. ScyllaDB only if the product adds durable archives, multi-host survival, or
   retention beyond Valkey's intended boundary.
6. ClickHouse only for high-volume aggregate observability after a privacy
   review; it never accepts gameplay commands.

Stable game-partition and repository interfaces make these later stages
possible. They do not pretend that a single PC provides independent failure
domains.

# Data Platform

Terminology used below: append-only file (AOF), Cassandra Query Language (CQL),
Internet Protocol (IP), personal computer (PC), Portable Game Notation (PGN),
Structured Query Language (SQL), time to live (TTL), and user interface (UI).

## Initial deployment: Valkey only

KnightOwl has no SQL service in its initial deployment. Valkey is the live
state database, not merely a cache. It supplies:

- atomic functions/transactions for game version changes;
- hashes and streams for compact projections and event chains;
- TTL deletion for rooms, capabilities, rate buckets, and cached analysis;
- AOF persistence for bounded restart recovery;
- replication, Sentinel, and Cluster as later measured scale steps.

The exact AOF policy is selected by fault testing on the target disk. The UI
states the honest recovery boundary: acknowledgement means accepted by Valkey
under the configured policy, not guaranteed survival of every conceivable host
or disk failure.

## Partition contract

Every correctness-critical update occurs within one logical partition:

- `game:{game-id}:state`
- `game:{game-id}:events`
- `game:{game-id}:commands`
- `game:{game-id}:capabilities`

The braces represent a Valkey Cluster hash tag. Atomic code may not touch a
second partition. Cross-game views are derived and never decide move legality.
This rule prevents a future cluster migration from requiring a game rewrite.

## Why not install ScyllaDB now?

ScyllaDB is the leading self-hosted Dynamo-style candidate if durable,
multi-node state becomes a real requirement. Its data model suits
partition-keyed game events and its lightweight transactions provide
linearizable conditional changes.

It is deferred because one PC does not provide independent failure domains and
ScyllaDB adds compaction, repair, capacity, consistency, and upgrade work. For
short-lived friends-only rooms, that operational cost buys no user-visible
capability beyond properly configured Valkey.

Before adoption, a production-shaped bake-off must prove:

- one conditional writer for a hot game partition;
- idempotent retry and reconnect across node loss;
- TTL/tombstone behavior under room churn;
- backup/restore and repair;
- p99 latency on target hardware;
- safe rolling upgrade and rollback;
- resource coexistence with Stockfish.

The app talks through a narrow game-state repository, not through DynamoDB- or
CQL-shaped types in domain logic.

## Analytics and observability

There is no business/user analytics pipeline. Prometheus-compatible aggregate
metrics cover request counts, latency, errors, connections, resource pressure,
and queue depth without labels containing player, room, IP, move, or capability
values.

If metric cardinality or retention eventually exceeds that system, ClickHouse
may receive privacy-reviewed aggregate operational events. It is a separate,
derived store and never participates in game acceptance. Its SQL interface is
appropriate there, but ClickHouse is not installed in advance.

## Data lifetime

| Data | Location | Default lifetime |
|---|---|---|
| Unused invitation | Valkey | Short room-configured TTL |
| Active room/game | Valkey | Active duration plus reconnect grace |
| Finished game | Valkey | PGN download grace, then expiry |
| Idempotency entry | Valkey | At least game lifetime |
| Presence/reactions | Valkey | Seconds/minutes |
| Analysis cache | Valkey | Short bounded TTL and memory budget |
| Local PGN/history | Browser IndexedDB | Until the person clears it |
| Aggregate metrics | Metrics system | Operator-configured short retention |

TTL behavior is integration-tested by observing deletion, not inferred from
configuration.

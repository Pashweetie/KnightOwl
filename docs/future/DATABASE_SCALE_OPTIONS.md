# Future Database Scale Options

Terminology used below: high availability (HA), Structured Query Language
(SQL), and time to live (TTL).

This is a future trigger matrix, not a list of services to install.

| Need | First response | Later candidate | Why |
|---|---|---|---|
| More app throughput | Add stateless app replicas behind Traefik | — | App processes are cheap and share versioned state |
| Live-state HA | Valkey replica plus Sentinel on independent hardware | Managed/multi-host Valkey | Same model and atomic functions |
| Live-state capacity | Valkey Cluster using existing hash tags | — | Shards complete games without cross-slot transactions |
| Durable partitioned archives | Do not add unless users opt in | ScyllaDB | Dynamo-style scale and conditional single-partition updates |
| Aggregate operational analytics | Prometheus-compatible metrics | ClickHouse | Columnar analytics stays off the command path |
| Relational future feature | Re-evaluate that feature independently | PostgreSQL | SQL is allowed when relationships/transactions actually justify it |

ScyllaDB is not treated as magically faster for every request. Conditional
linearizable writes are more expensive than ordinary writes, TTL-heavy models
create tombstone/compaction work, and a useful cluster needs real failure
domains. Adoption follows benchmarks and operational drills, not forecasted
scale.

The repository boundary is semantic:

```text
create_room
claim_seat
apply_game_command(expected_version, command_id, transition)
read_events(after_sequence)
finish_and_expire
```

It is not a generic key-value abstraction and does not promise transparent
replacement of incompatible consistency models.

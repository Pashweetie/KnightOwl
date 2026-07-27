# Backend Language Selection

## Shortlist

| Choice | Strengths here | Costs here | Status |
|---|---|---|---|
| Rust with Tokio/Axum | Strong memory safety, controlled latency, compact containers, excellent explicit domain types | Highest learning and implementation cost | Leading candidate |
| Go | Simple operations, strong standard library, cheap concurrency, built-in race detector | Less expressive domain invariants; GC requires measurement | Required spike |
| Elixir/Phoenix | Excellent process supervision and socket model | Smaller chess/engine integration ecosystem; different operational stack | Reserve candidate |
| TypeScript/Node | Shared language with browser, fast iteration | Runtime validation burden and weaker compile-time domain guarantees | Frontend only unless spike evidence reverses decision |

## Decision process

Rust and Go must implement the same thin vertical spike before the backend is
selected:

1. capability-protected room creation;
2. two WebSocket players and spectators;
3. one legal move through the canonical hash chain;
4. atomic Valkey version comparison;
5. reconnect to another process;
6. graceful drain and health endpoints;
7. fuzzed command decoder and a clock race test.

Measure p50/p95/p99 command latency, resident memory per 1,000 sockets, CPU,
binary/image size, reconnect behavior, profiling quality, dependency surface,
and implementation complexity on the actual host. The winning implementation
becomes production code; the other spike is removed.

Current preference is Rust because correctness and predictable resource use
matter more than sharing a language with the web client. Go wins if it produces
equivalent correctness with substantially lower maintenance cost on the target
team and hardware.


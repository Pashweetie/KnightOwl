# Architecture Decisions

## ADR-001: Modular monolith first

Use one TypeScript API deployment with strongly separated domain modules and
independent workers. This keeps transactions and operations manageable on one
PC without preventing later extraction of matchmaking or analysis services.

## ADR-002: PostgreSQL is authoritative

PostgreSQL owns all durable business state. Redis is disposable coordination
state. A Redis flush may empty queues and presence but may not corrupt or erase
accepted product actions.

## ADR-003: Event-backed live games

Each accepted game command appends an ordered event in the same transaction as
its current-state projection. Sequence numbers and idempotency keys make replay
and retries safe. Periodic FEN snapshots accelerate recovery; PGN is derived.

## ADR-004: Clocks use server monotonic time

Clients render estimates only. The server calculates elapsed time from a
monotonic clock, records authoritative samples with moves, and reconstructs a
safe deadline after restart from persisted wall-clock anchors.

## ADR-005: Named tunnel only for production

Quick tunnels are permitted only for isolated development checks. Production
uses a named Cloudflare Tunnel, a controlled hostname, explicit ingress rules,
and a token stored outside Git.

## ADR-006: No premature microservices

Authentication, chess, ratings, learning, tournaments, and moderation remain
modules in the API until measured scaling or isolation requirements justify a
service boundary. Stockfish is isolated immediately because it executes
resource-intensive native code.

## ADR-007: Features ship vertically

Each phase includes schema, domain logic, API, UI, authorization, telemetry,
tests, and documentation. A backend-only or decorative UI implementation does
not complete a feature.


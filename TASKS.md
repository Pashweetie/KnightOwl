# KnightShift Delivery Plan

This checklist is the release gate. A box may be checked only when its automated
tests pass and the behavior has been exercised through the real application.
There will be no placeholder controls, invented statistics, simulated users, or
panels that merely display a toast.

## 0. Product contract

- [x] Reject and remove the prototype.
- [x] Remove its public tunnel, containers, image, and repository.
- [x] Define the supported product and explicit non-goals.
- [x] Define the architecture and persistence model.
- [x] Define acceptance criteria for every subsystem.

## 1. Foundation and authentication

- [ ] Create the TypeScript monorepo and locked dependency graph.
- [ ] Create PostgreSQL migrations and a disposable test database.
- [ ] Implement registration with normalized, unique email and username.
- [ ] Hash passwords with Argon2id and a server-side pepper.
- [ ] Implement login, logout, rotating sessions, and session revocation.
- [ ] Implement verified email changes and password reset tokens.
- [ ] Add CSRF protection, secure cookies, rate limits, and audit events.
- [ ] Add integration tests for success, failure, expiry, replay, and lockout.
- [ ] Pass the authentication threat-model checklist.

## 2. Live chess

- [ ] Implement authoritative chess rules and immutable game events.
- [ ] Implement seek creation for bullet, blitz, rapid, and classical pools.
- [ ] Implement atomic matchmaking with rating and latency expansion.
- [ ] Implement WebSocket authentication and origin validation.
- [ ] Implement authoritative monotonic clocks with increment and delay.
- [ ] Implement move, resign, draw offer/accept/decline, and abort.
- [ ] Implement disconnect grace, reconnect, multi-tab ownership, and recovery.
- [ ] Persist PGN, FEN snapshots, result, termination, and clock history.
- [ ] Add tests for every legal ending and illegal client action.
- [ ] Pass two-browser, restart-recovery, and concurrency tests.

## 3. Ratings, profiles, and history

- [ ] Implement separate Glicko-2 pools per time control and variant.
- [ ] Make rating updates transactional and idempotent.
- [ ] Implement public profiles with real statistics only.
- [ ] Implement private account and privacy preferences.
- [ ] Implement searchable, paginated game history and PGN export.
- [ ] Add rating-calculation fixtures and transaction-race tests.

## 4. Analysis and review

- [ ] Run a pinned Stockfish build in isolated workers with resource limits.
- [ ] Queue analysis jobs and persist progress and engine provenance.
- [ ] Produce evaluations, best lines, accuracy, and move classifications.
- [ ] Provide an interactive analysis board with variations and annotations.
- [ ] Support user PGN import/export.
- [ ] Add deterministic engine fixtures, cancellation, and overload tests.

## 5. Puzzles and learning

- [ ] Store licensed/original puzzle positions with verified solutions.
- [ ] Implement server-validated puzzle attempts and spaced repetition.
- [ ] Implement puzzle rating and user puzzle rating updates.
- [ ] Author original lessons with interactive checkpoints.
- [ ] Persist lesson progress, review queue, streaks, and mastery.
- [ ] Add solution-validation and progress-integrity tests.

## 6. Competition and community

- [ ] Implement arena and Swiss tournament state machines.
- [ ] Implement registration, pairing, scoring, tie-breaks, and withdrawals.
- [ ] Implement follows, blocks, challenges, clubs, and direct messages.
- [ ] Implement reports, moderation queues, sanctions, and appeals.
- [ ] Add permission, abuse-rate-limit, and tournament-invariant tests.

## 7. Operations and release

- [ ] Produce rootless Docker images with pinned base digests.
- [ ] Add PostgreSQL, Redis, worker, backup, and Cloudflare services.
- [ ] Implement health, readiness, metrics, structured logs, and alerts.
- [ ] Implement encrypted backups and successfully restore one.
- [ ] Run dependency, secret, container, and dynamic security scans.
- [ ] Run load tests for matchmaking, games, WebSockets, and analysis queues.
- [ ] Verify accessibility, responsive behavior, and browser compatibility.
- [ ] Complete data export, account deletion, privacy, and terms workflows.
- [ ] Perform a clean-machine installation from documented instructions.
- [ ] Publish through a named Cloudflare Tunnel and verify the public system.

## Release rule

KnightShift is not called complete and is not publicly published until every
box above is checked. Development previews remain local or access-restricted.


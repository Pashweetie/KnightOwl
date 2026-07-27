# KnightShift Product Contract

## Objective

KnightShift is a self-hostable, production-quality chess platform for real
people. It provides persistent accounts, rated and casual live chess, game
history, computer-assisted post-game review, puzzles, original lessons,
tournaments, social features, and moderation.

“Feature parity” means comparable user capabilities. It does not mean copying
Chess.com branding, source code, lesson text, puzzle collections, visual
assets, engine labels, rating data, or other proprietary material.

## Non-negotiable product rules

1. Every displayed user, game, rating, statistic, message, tournament, puzzle
   result, and progress value comes from persisted application state.
2. Every visible control performs its labeled action. Unimplemented features
   are absent rather than disabled, faked, or represented by notifications.
3. The server is authoritative for identity, permissions, chess rules, clocks,
   matchmaking, results, ratings, puzzle solutions, and tournament state.
4. Client input is untrusted. WebSocket messages and HTTP requests use the same
   authorization and validation standards.
5. Public release requires durable storage, backups, recovery testing,
   observability, abuse controls, and a permanent named tunnel.
6. Premium-equivalent features are available without an artificial paid tier.
   Operating costs and optional donations may be documented separately.

## Initial supported chess scope

- Standard chess, with the data model ready for variants but no advertised
  variant until it has its own rules and test suite.
- Casual and rated games.
- Bullet, blitz, rapid, and classical time-control pools.
- Fischer increment and Bronstein delay.
- Draw by agreement, stalemate, insufficient material, fivefold repetition,
  seventy-five-move automatic draw, claimable threefold repetition, and
  claimable fifty-move draw, according to the supported rules library.
- Resignation, abandonment, timeout, and authorized administrative termination.

## Architecture

The deployable system consists of:

- A React/TypeScript web client.
- A TypeScript HTTP and WebSocket API.
- PostgreSQL as the durable system of record.
- Redis for ephemeral matchmaking, presence, rate limits, and job coordination.
- Background workers for mail, analysis, rating settlement, and maintenance.
- Stockfish workers isolated from the API process.
- Cloudflare Tunnel as outbound-only ingress.
- Docker Compose for a single-machine installation, with clean seams for later
  multi-host deployment.

The API remains stateless except for active socket connections. Durable game
events are written before acknowledgement. Redis loss may disrupt presence or
queues but must not lose accounts, accepted moves, games, ratings, or progress.

## Core data domains

- Identity: users, credentials, sessions, verification tokens, roles, bans,
  privacy settings, and audit events.
- Chess: seeks, games, participants, moves, clock samples, results, and PGNs.
- Skill: rating pools, rating periods, rating transactions, and leaderboards.
- Analysis: jobs, engine versions, evaluations, lines, annotations, and quotas.
- Learning: puzzles, solutions, attempts, schedules, lessons, checkpoints, and
  progress.
- Competition: tournaments, entrants, rounds, pairings, scores, and tie-breaks.
- Social/moderation: relationships, challenges, clubs, messages, reports,
  evidence, actions, and appeals.

## Quality gates

### Correctness

- Chess state is reconstructible from persisted events.
- All balance-like mutations—ratings, tournament scores, quotas—are
  transactional and idempotent.
- Retried requests cannot duplicate moves, results, rating changes, or rewards.

### Security

- OWASP ASVS Level 2 is the baseline.
- Passwords use Argon2id; sessions are opaque, rotating, revocable, and stored
  hashed; cookies are Secure, HttpOnly, and SameSite.
- State-changing HTTP uses CSRF defenses. WebSockets validate session, Origin,
  payload schema, authorization, rate, and sequence.
- Secrets never enter images, Git, browser bundles, logs, or task documents.

### Reliability

- A process restart during a game preserves all acknowledged moves and restores
  authoritative clocks within the documented tolerance.
- PostgreSQL backups are encrypted, retained, and restored in a recorded test.
- Readiness fails when required dependencies cannot safely serve traffic.

### Performance

- On the target PC, 95% of accepted live moves are broadcast to both players
  within 250 ms excluding client network latency.
- Matchmaking does not pair the same account with itself and does not assign one
  seek to multiple games under concurrency.
- Engine work cannot starve live-game traffic.

### Accessibility and compatibility

- Keyboard play, visible focus, semantic labels, reduced motion, sufficient
  contrast, and screen-reader announcements are tested.
- Current stable Chrome, Firefox, Safari, and Edge are supported.
- Phone, tablet, laptop, and desktop layouts are verified with real controls.

## Publication gate

The public Cloudflare hostname is configured only after:

1. all task-list release checks pass;
2. the operator supplies a named-tunnel token and controlled domain;
3. backup/restore and security tests pass;
4. the operator reviews privacy, moderation, and acceptable-use settings.


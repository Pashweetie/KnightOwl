# KnightShift Product Contract

## Objective

KnightShift is a self-hostable, production-quality chess room for invited
friends. It provides capability-link multiplayer, casual live chess, local game
history, computer-assisted post-game review, a coaching bot, optional original
learning content, and ephemeral room tournaments without collecting account or
contact identity.

“Feature parity” means comparable user capabilities. It does not mean copying
Chess.com branding, source code, lesson text, puzzle collections, visual
assets, engine labels, rating data, or other proprietary material.

## Non-negotiable product rules

1. Every displayed player, game, statistic, tournament, puzzle result, and
   progress value comes from verified live state or clearly labelled local
   browser state.
2. Every visible control performs its labeled action. Unimplemented features
   are absent rather than disabled, faked, or represented by notifications.
3. The server is authoritative for room capabilities, chess rules, clocks,
   results, puzzle solutions, and active tournament state.
4. Client input is untrusted. WebSocket messages and HTTP requests use the same
   authorization and validation standards.
5. Public ingress requires tested expiry, recovery behavior, observability,
   resource controls, and a permanent named tunnel.
6. Premium-equivalent features are available without an artificial paid tier.
   Operating costs and optional donations may be documented separately.
7. KnightShift stores no login credentials, contact identity, public profile,
   advertising identifier, or hidden durable player profile. Purpose-limited
   operational/security telemetry follows the documented retention policy.

## Initial supported chess scope

- Standard chess, with the data model ready for variants but no advertised
  variant until it has its own rules and test suite.
- Casual invite-room games.
- Bullet, blitz, rapid, and classical time controls.
- Fischer increment and Bronstein delay.
- Draw by agreement, stalemate, insufficient material, fivefold repetition,
  seventy-five-move automatic draw, claimable threefold repetition, and
  claimable fifty-move draw, according to the supported rules library.
- Resignation, abandonment, timeout, and authorized administrative termination.

## Architecture

The deployable system consists of:

- A React/TypeScript web client.
- A Rust or Go HTTP and WebSocket API selected through the documented spike.
- Valkey for expiring room state, presence, rate limits, and job coordination.
- Background workers for analysis and maintenance.
- Stockfish workers isolated from the API process.
- Traefik for local health-aware load balancing across app containers.
- Cloudflare Tunnel as outbound-only ingress.
- Docker Compose for a single-machine installation, with clean seams for later
  multi-host deployment.

The API remains stateless except for active socket connections. Accepted game
events are atomically appended to an expiring Valkey chain before
acknowledgement. Valkey loss fails active games explicitly; there is no hidden
claim of durable recovery for data the product intentionally does not retain.

## Core data domains

- Rooms: expiring capabilities, room-local labels, seats, settings, and
  readiness.
- Chess: seeks, games, participants, moves, clock samples, results, and PGNs.
- Analysis: engine versions, evaluations, lines, and annotations; results are
  returned to the requesting room or retained locally by its browser.
- Learning: original/licensed puzzles and lessons with browser-local progress.
- Competition: expiring room tournaments, entrants, pairings, and scores.

## Quality gates

### Correctness

- Chess state is reconstructible from persisted events.
- Tournament scores and quotas are atomic and idempotent.
- Retried requests cannot duplicate moves, results, scores, or room actions.

### Security

- OWASP ASVS Level 2 is the baseline.
- Room capabilities are high entropy, revocable, stored only as keyed hashes,
  and never logged.
- State-changing HTTP uses CSRF defenses. WebSockets validate session, Origin,
  payload schema, authorization, rate, and sequence.
- Secrets never enter images, Git, browser bundles, logs, or task documents.

### Reliability

- An app-process restart preserves acknowledged moves while Valkey is healthy
  and restores authoritative clocks within the documented tolerance.
- Valkey persistence policy and its explicit data-loss boundary are tested.
- Readiness fails when required dependencies cannot safely serve traffic.

### Performance

- On the target PC, 95% of accepted live moves are broadcast to both players
  within 250 ms excluding client network latency.
- A room seat cannot be claimed twice under concurrency.
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
3. security, chaos, expiry, and recovery-behavior tests pass;
4. the operator reviews privacy and friends-only room controls.

# Product Overview

Terminology used below: application programming interface (API), Hypertext
Transfer Protocol (HTTP), and personal computer (PC).

## Objective

KnightOwl is a self-hostable, production-quality chess room for invited
friends. It provides capability-link multiplayer, casual live chess, local game
history, computer-assisted post-game review, Bot Mode with built-in teaching
assistance, optional original puzzles and drills, and ephemeral room
tournaments without collecting account or contact identity.

“Feature parity” means comparable user capabilities implemented with
KnightOwl-owned or appropriately licensed software, content, data, and assets.

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
7. KnightOwl stores no login credentials, contact identity, public profile,
   advertising identifier, or hidden durable player profile. Purpose-limited
   operational/security telemetry follows the documented retention policy.

## Initial supported chess scope

- Standard chess, with the data model ready for variants but no advertised
  variant until it has its own rules and test suite.
- Casual invite-room games.
- Bullet, blitz, rapid, and classical time controls.
- Fischer increment and Bronstein delay.

## Friends-only interaction

There is no public directory, public profile, follower system, direct message,
forum, club, public chat, moderation queue, cheating review, or identity-linked
sanction system.

The host can create and revoke invitations, choose room settings, assign
colors, remove a participant before play, disable spectators or reactions, end
the room, and export the game. Players can use room-local labels, ready up,
play, resign, offer or answer a draw, request a rematch, mute reactions,
reconnect, and export the game.

Communication uses a fixed set of reaction events. Each participant may
customize an event's local animation and sound, but cannot transmit custom
text, images, or media through KnightOwl.

The operator has no public room list. Room state, labels, reaction events, and
operational buckets expire. Runbooks cover service outages, resource pressure,
upgrade, rollback, and credential loss.
- Draw by agreement, stalemate, insufficient material, fivefold repetition,
  seventy-five-move automatic draw, claimable threefold repetition, and
  claimable fifty-move draw, according to the supported rules library.
- Resignation, abandonment, timeout, and authorized administrative termination.
- A live game and its clocks start only when White's first legal move is
  accepted; waiting in a joined room consumes no clock.

## Architecture

The deployable system consists of:

- A React/TypeScript web client.
- A Go server-authoritative Hypertext Transfer Protocol and WebSocket service.
- Valkey for expiring room state, presence, rate limits, and job coordination.
- Background workers for analysis and maintenance.
- Stockfish workers isolated from the API process.
- K3s-provided Traefik for production ingress and service routing.
- Cloudflare Tunnel as outbound-only ingress.
- Docker Compose for local development and single-node K3s for production, with
  clean seams for later multi-host deployment.

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
- Learning: original or licensed puzzles and drills with browser-local
  progress.
- Competition: expiring room tournaments, entrants, pairings, and scores.

## Quality gates

### Correctness

- Chess state is reconstructible from persisted events.
- Tournament scores and quotas are atomic and idempotent.
- Retried requests cannot duplicate moves, results, scores, or room actions.

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

- The private friends-only release makes no claim of formal Americans with
  Disabilities Act or Web Content Accessibility Guidelines certification.
- Keyboard play, visible focus, semantic controls, reduced motion, sufficient
  contrast, and non-color status cues remain baseline quality requirements.
- Full screen-reader chess-board narration is optional unless an intended
  player needs it; generic controls inherit assistive-technology behavior from
  the selected component library.
- Current stable Chrome, Firefox, Safari, and Edge are supported.
- Phone, tablet, laptop, and desktop layouts are verified with real controls.

## Publication gate

The public Cloudflare hostname is configured only after:

1. all task-list release checks pass;
2. the operator supplies a named-tunnel token and controlled domain;
3. the private release review, expiry, and recovery-behavior tests pass;
4. the operator reviews privacy and friends-only room controls.

# Private Tournament Contract

Terminology used below: International Chess Federation (FIDE), Tournament
Report Format (TRF), and time to live (TTL).

## Library boundary

KnightOwl never implements a Swiss pairing algorithm. It invokes a pinned
bbpPairings executable and exchanges the International Chess Federation
Tournament Report Format through a narrow adapter. Golden files from the
pairing program and the federation verification checklist qualify that adapter.

Arena tournaments are outside the initial release. An arena is a continuous
join-and-play event in which players complete as many games as possible during
a fixed time window. It differs from Swiss, which has scheduled rounds and one
pairing per entrant per round. Adding arena later requires its own product and
library review; KnightOwl will not improvise a pairing algorithm.

## Scope

Swiss tournaments are required release features for invited friends.
Tournament records expire with the room after an export grace period. There is
no global rating, public discovery, prize handling, or FIDE-rating claim.

## Shared lifecycle

```text
draft -> registration -> running -> completed -> expiring
                  \-> cancelled
```

The host selects format, time control, variant, start rule, capacity, spectator
policy, reaction policy, scoring/tie-break display, and expiry. Settings become
immutable at start except documented host emergency controls.

For every tournament board, the pairing can be published and both participants
can connect while the game remains `ready`. The game starts only when White's
first move is accepted. A tournament-specific no-show deadline may award a
forfeit, but it never starts White's clock before White moves.

Each participant has one tournament-scoped capability and one active entry.
Commands are versioned and idempotent. A game result can settle a pairing
exactly once. Standings are derived from accepted pairing results and can be
reconstructed within the tournament lifetime.

## Swiss

- The number of rounds and late-entry policy are fixed before start.
- `bbpPairings` performs Dutch-system pairings from a TRF representation.
- The exact engine version and per-round TRF input/output are retained until
  tournament expiry for explanation/replay.
- No player meets the same opponent twice.
- Bye eligibility, score grouping, color allocation, withdrawal, forfeit, and
  unplayed-game treatment follow the selected documented 2026 rule profile.
- KnightOwl validates the returned pairing set before publishing it.
- If the pairing engine fails or returns an invalid set, the round does not
  partially start; the host sees a real diagnostic and may retry or cancel.
- Pairings are deterministic for identical state and engine version.

Because participants are anonymous and unrated, initial ordering uses a host
seed, randomized order committed before round one, or opt-in room-local seed
numbers—not invented global ratings.

## Arena

Arena is an online product format, not a FIDE pairing system. Its rules must be
published in the room before registration:

- start/end time and whether games already started may finish;
- win/draw/loss points;
- streak/bonus definition, if enabled;
- rematch avoidance window;
- color-balancing priority;
- pairing wait/availability;
- withdrawal and reconnect;
- tie-break order;
- whether optional berserk/time-odds exists.

The matcher pairs only available entrants, never pairs a participant with
themself, avoids immediate repeats where possible, and atomically reserves both
entrants before creating one game. Scoring consumes one immutable game result
once.

## Host operations

Before start, the host may remove an entrant or cancel. During play, the host
may pause new pairings, disable reactions/spectators, withdraw a disconnected
entry, or cancel future rounds. The host cannot edit a finished board result
into a different chess result. Any exceptional forfeit is a separate,
visible tournament ruling.

## Interface requirements

- Registration clearly states format, time control, rounds/duration, scoring,
  tie-breaks, start condition, and expiry.
- Lobby shows readiness and actual entrants—never simulated names.
- Running view shows current board/opponent, round or time remaining, live
  standings, pairing status, withdrawal, and reconnect state.
- Standings expose how each displayed score/tie-break was calculated.
- Completed view exports tournament summary plus all available PGNs before
  expiry.
- Pairing, round, result, and standing updates use semantic text and keyboard
  navigation without relying on color or animation; specialized screen-reader
  narration remains optional.

## Required tests

- Upstream bbpPairings examples and current FIDE conformance expectations.
- Randomized Swiss tournaments across even/odd fields, withdrawals, byes,
  forfeits, late entries, score ties, and color histories.
- Arena double-reservation, repeat avoidance, result retry, simultaneous finish,
  tournament end while games run, and reconnect.
- Duplicate/stale host and participant commands.
- Kill/restart app replicas during registration, pairing, settlement, and
  completion.
- TTL expiry and export-before-expiry.
- Two-browser and multi-browser accessibility workflows.

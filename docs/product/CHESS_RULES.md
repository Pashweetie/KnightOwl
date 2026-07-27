# KnightOwl Chess Product Rules

Terminology used below: Forsyth–Edwards Notation (FEN), identifier (ID),
Portable Game Notation (PGN), Standard Algebraic Notation (SAN), Universal
Chess Interface (UCI), user interface (UI), and Coordinated Universal Time
(UTC).

This document fixes the standard-chess behavior before game implementation.
Variant rules receive separate documents and test suites.

## 1. Position rules

KnightOwl standard chess implements:

- the normal initial position;
- legal piece movement and capture;
- check and prohibition on leaving one's king in check;
- kingside and queenside castling, including attacked transit-square rules;
- en passant only on the immediately following move;
- promotion to queen, rook, bishop, or knight;
- checkmate and stalemate;
- repetition and move-count draw rules;
- dead-position detection.

The chosen rules library must be checked against independent perft positions and
targeted regression fixtures. UI legal-move highlighting is not evidence that a
move is legal; the server validates from authoritative state.

## 2. Game modes

### Live

Clock runs continuously during the player's turn. Supported controls:

- Bullet: estimated total per player below 3 minutes.
- Blitz: at least 3 and below 10 minutes.
- Rapid: at least 10 and below 60 minutes.
- Classical: at least 60 minutes.

The estimate is base time plus 40 increments/delays for display classification.

Fischer increment is added after an accepted move. Bronstein delay refunds up
to the configured delay but never increases remaining time above turn-start
time. Exact millisecond ordering is tested.

### Correspondence

Correspondence chess is deferred because days-long games require durable
participant continuity and retained state, both currently disallowed by the
identity/privacy decision. Its eventual design must explicitly define
deadlines, conditional moves, notification reliability, retention, and
recovery. Live-room TTLs must not be stretched into an accidental incomplete
implementation.

### Casual only

All initial human games are casual and unrated. KnightOwl enforces rules and
clock correctness but has no cheating policy. Coach/engine tools are visibly
separate from active human game controls to prevent accidental use, not to
police friends.

## 3. Start and abort

A live game remains in `ready` state after both players connect. Neither clock
runs. White's first server-accepted move atomically changes the state to
`started`; Black's clock begins from that accepted event. White is not charged
for time spent considering the first move.

Before White's first accepted move, either player or the host may abort without
a game result, unless private tournament rules define a start/no-show deadline.
After that move, abort is unavailable and players use resign/draw actions.
Repeated room creation and pre-start abort may hit operational rate limits.

## 4. Commands

Valid participant commands:

- move;
- resign;
- offer draw;
- accept or decline draw;
- claim eligible draw;
- request abort when eligible;
- request rematch after completion;
- enqueue or cancel premoves locally, with server validation on execution.

Every command has a unique idempotency ID and expected game version. The server
returns accepted, rejected with reason, or stale with resynchronization state.

## 5. Draws

Automatic:

- stalemate;
- dead position / insufficient mating possibility;
- fivefold repetition;
- seventy-five moves by each side without pawn move or capture, unless the last
  move checkmates.

Claimable by the player entitled under the position:

- threefold repetition, including a valid claim by intended move;
- fifty moves by each side without pawn move or capture, including a valid
  claim by intended move.

Agreement:

- one active offer at a time;
- an offer is implicitly declined by the opponent making a move;
- offer spam is rate-limited;
- accepting is idempotent and ends the game transactionally.

## 6. Timeout

When a player's authoritative clock reaches zero, that player loses unless the
opponent cannot possibly checkmate by any legal sequence from the resulting
position; in that case the result is a draw.

Timeout is decided by the server. A move and timeout racing at the deadline are
ordered under the locked game state using the server's authoritative receipt
time and clock calculation.

## 7. Disconnect and abandonment

Disconnect does not stop a live clock. Reconnection restores the game from the
last committed sequence.

An abandonment adjudication may end a game before clock expiration only under a
published policy based on game phase and remaining time. It must not alter the
board result semantics or fabricate a resignation. The termination reason is
stored separately as `abandoned`.

## 8. Premoves

- A user may queue zero to five premoves.
- Premoves are enabled by default and can be disabled from a plainly labelled
  board setting.
- While it is the opponent's turn, selecting the user's own piece changes the
  interaction banner to `Premove`; normal live-move and premove modes are never
  distinguished only by color.
- Every queued move is drawn with a distinct accessible style and numbered in
  execution order. The move list also has a `Queued premoves` region.
- The queue is cancellable one entry at a time and with `Clear all`; right-click
  or equivalent may be a shortcut but is never the only method.
- Escape cancels the newest queued move when board focus is active. Touch and
  screen-reader controls expose the same actions.
- After each opponent move, only the first queued entry is submitted.
- The server tests it against the new authoritative position.
- If legal, it is processed as a normal move with a fixed 100 ms clock charge.
- If illegal, it and all dependent later premoves are discarded.
- Promotion follows the premove promotion choice or the user's auto-queen
  setting recorded with the command.
- A premove is never accepted before the opponent move is committed.
- A reconnect/stale snapshot clears unconfirmed local premoves and announces
  that they were not played; it never guesses whether to resubmit them.
- Settings explain the fixed clock charge, queue limit, execution order, and
  invalidation behavior in plain language beside the control.

## 9. Move interaction

Supported inputs:

- click source then destination;
- drag source to destination;
- full keyboard square entry and promotion choice;
- screen-reader-friendly piece and legal-destination selection.

Settings:

- show legal destinations;
- show last move;
- auto-queen;
- premoves;
- board orientation;
- coordinates;
- animation level;
- sounds and low-time warning;
- focus mode.

Settings affect presentation only, except auto-queen and premove preferences,
which are included in explicit commands.

## 10. Notation and records

Each active game stores expiring ordered events and accepted move records.
Derived artifacts:

- SAN move list;
- UCI moves;
- PGN with time-control, result, termination, UTC, variant, opening, and clock
  annotations where enabled;
- current and periodic FEN;
- result and termination within the room lifetime.

PGN export never implies that imported comments or engine lines were generated
by KnightOwl. The browser can retain an exported/local PGN after server state
expires.

## 11. Spectators and active-game separation

The host controls whether spectators are allowed and may choose live or delayed
delivery. Spectator capabilities cannot issue player commands. Analysis,
opening explorer, tablebase, evaluation, and coach controls are absent from an
active human-game view; this prevents mistakes without claiming anti-cheating
enforcement. Fixed reactions are a separate muteable permission.

## 12. Required test families

- standard perft suites through multiple depths;
- castling rights gained/lost edge cases;
- en passant exposing or resolving check;
- underpromotion;
- every checkmate/stalemate/dead-position class;
- threefold/fivefold and 50/75-move boundaries;
- timeout versus possible/impossible mating material;
- increment/delay millisecond boundaries;
- move versus timeout races;
- duplicate, stale, reordered, and unauthorized commands;
- disconnect before/after commit and process restart;
- premove legality and fixed charge;
- premove first-use guidance, visible mode/queue/order, keyboard/touch/screen
  reader cancellation, stale reconnect clearing, and invalid dependent queues;
- PGN round-trip and event replay equivalence.

# Existing Chess Test Ecosystem

Terminology used below: Berkeley Software Distribution two-clause license
(BSD-2), continuous integration (CI), distance to zeroing move (DTZ),
Forsyth–Edwards Notation (FEN), International Chess Federation (FIDE), GNU
General Public License version 3 (GPL-3), Portable Game Notation (PGN),
Standard Algebraic Notation (SAN), Tournament Report Format (TRF), time to live
(TTL), Universal Chess Interface (UCI), user interface (UI), uniform resource
locator (URL), user experience (UX), and win/draw/loss (WDL). Codes such as
C0401 and E012023 are federation document identifiers.

## Principle

KnightOwl will integrate maintained chess libraries and import or run
established conformance data under compatible licenses. It will not create a
new move generator, PGN grammar, Swiss pairing algorithm, engine protocol, or
tablebase format.

KnightOwl-specific tests remain necessary for clocks, network ordering,
capabilities, premove UX, room lifecycle, tournament orchestration, and the
adapter boundaries. Those are product semantics that external chess libraries
cannot test.

## Rules and move generation

### Stockfish perft suite

Stockfish's official repository includes `tests/perft.sh` with standard and
Chess960 positions and exact node totals. It cites the established
ChessProgramming perft results and exercises castling, en passant, promotions,
checks, and complex move generation.

Use:

- run the selected backend chess adapter over the same FEN/depth cases;
- compare node totals to the pinned Stockfish script;
- run Stockfish `go perft` as a differential oracle for additional generated
  legal positions;
- record the exact Stockfish commit and suite license.

Source:
<https://github.com/official-stockfish/Stockfish/blob/master/tests/perft.sh>

### jchess-perft-dataset and chess-test-utils

`fathzer-games/jchess-perft-dataset` provides thousands of standard positions
and 960 Chess960 positions; `chess-test-utils` is a reusable framework around
that data. It is Java-oriented, so Go need not adopt its framework, but can
consume the documented dataset format if its license/provenance review passes.

Sources:

- <https://github.com/fathzer-games/jchess-perft-dataset>
- <https://github.com/fathzer-games/chess-test-utils>

### Mature-library differential testing

Compare randomized legal games at every ply across:

- the selected Go library;
- Stockfish position/perft;
- python-chess as an independent feature-rich reference;
- chessops in the browser.

Compare legal UCI move sets, normalized FEN, side to move, castling/en-passant
state, check, terminal outcome, and undo/redo round trips. Disagreement becomes
a minimized regression case; no majority vote silently chooses a result.

Reference suites:

- python-chess `test.py`:
  <https://github.com/niklasf/python-chess/blob/master/test.py>
- chessops test/fuzz sources:
  <https://github.com/niklasf/chessops>
- chess.js tests:
  <https://github.com/jhlywa/chess.js/tree/master/__tests__>

## Browser rules library

`chessops` is the leading browser candidate because it supports standard,
Chess960, Crazyhouse, King of the Hill, Three-check, SAN, PGN, streaming
parsing, and lichess/chessground compatibility. Its repository already includes
tests and fuzzing. It is GPL-3.0+, so the distribution/license decision must be
explicit.

`chess.js` is a BSD-2-Clause alternative for standard chess with an extensive
Vitest suite, benchmarks, and wide usage. It lacks chessops' integrated variant
coverage. We should not load both in production merely for confidence; the
unselected library may remain a CI oracle.

## Laws and adjudication

FIDE Laws of Chess are the normative source for standard over-the-board rules.
Online product policy must explicitly map differences such as premoves,
automatic versus claimable draws, disconnects, and server clock receipt.

Source: <https://handbook.fide.com/chapter/E012023>

Test cases must cite the exact article for:

- checkmate/stalemate;
- dead position;
- castling and en passant;
- promotion;
- threefold/fivefold repetition;
- 50/75-move rules;
- timeout with possible mating sequence.

## Notation and interchange

Use the PGN Specification and Implementation Guide for import/export grammar
and semantics rather than inventing a format. Test against the selected
library's corpus plus python-chess, chessops, and chess.js parser cases.

Source:
<https://www.saremba.de/chessgml/standards/pgn/pgn-complete.htm>

The parser security layer still needs KnightOwl limits for byte size, nesting,
variation depth, comments, tags, and processing time.

## Engine protocol and analysis

Use Stockfish's UCI implementation and documented commands. Test engine startup,
`uci`/`isready`, position/move transfer, MultiPV, cancellation, time/node/depth
limits, option validation, and worker death against a pinned official binary.
Do not implement a chess engine protocol or evaluation search.

Fishtest is for statistically validating Stockfish engine-strength changes; it
is not needed because KnightOwl does not modify Stockfish.

## Endgame tablebases

Use existing Syzygy probing through a maintained library or Stockfish. Validate
known WDL/DTZ positions against python-chess/Syzygy and the pinned table files.
Do not generate a proprietary tablebase or alter table results.

## Swiss tournaments

Do not implement the FIDE Dutch pairing algorithm. Use `bbpPairings`, an
open-source FIDE-endorsed engine, through its TRF input/output boundary after
license review. Pin its version/container and preserve the input/output artifact
for each generated round during the tournament TTL.

Verification includes:

- current FIDE Swiss rules effective 1 February 2026;
- bbpPairings own tests/examples;
- FIDE software endorsement/checker expectations;
- randomized complete tournaments checked for no repeated opponents, valid
  byes, score brackets, color constraints, deterministic reruns, withdrawal,
  and late-entry policy.

Sources:

- FIDE Swiss rules:
  <https://handbook.fide.com/chapter/C0401>
- FIDE Dutch system:
  <https://handbook.fide.com/chapter/C0403>
- bbpPairings:
  <https://github.com/BieremaBoyzProgramming/bbpPairings>
- FIDE software verification:
  <https://spp.fide.com/verification-checklist/>

KnightOwl must not describe itself as FIDE-endorsed. The external engine's
endorsement does not endorse KnightOwl's orchestration.

## Premoves

Premoves are an online-product interaction, not a FIDE rule or general chess
library responsibility. Reuse the chess library for legality after the
opponent's accepted move, and reuse established UI patterns from major chess
interfaces as research references. The queue, visualization, cancellation,
clock charge, stale-state behavior, and accessibility require small
KnightOwl-specific state-machine and end-to-end tests.

This is justified custom code because no rule library can know the product's
network/clock/UI contract. It must remain an adapter around legal-move
validation, not another move generator.

## Test adoption checklist

Before copying or vendoring any test corpus:

1. record upstream URL, commit/tag, license, and provenance;
2. prefer running upstream or consuming its released data over copying cases;
3. keep upstream data unchanged and add a thin adapter;
4. separate upstream fixtures from KnightOwl regressions;
5. schedule dependency/corpus updates and inspect diffs;
6. never claim that passing perft alone proves complete chess correctness.

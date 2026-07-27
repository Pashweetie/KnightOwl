# Bot Opponent and Coach Specification

## 1. Product modes

### Play a bot

An unrated game against a named engine profile. The profile is clearly labeled
as a computer and defines:

- target strength range;
- time budget and deliberate response delay;
- opening repertoire;
- engine skill and error distribution;
- optional style biases that are measured rather than described fictionally;
- whether hints are enabled.

### Coach game

An unrated bot game with teaching assistance:

- before a move: optional hint ladder;
- after the user's move: immediate or delayed feedback setting;
- after the bot move: explanation of the bot's plan;
- at any time: candidate-move comparison and shallow educational search tree;
- after the game: full review and linked learning recommendations.

Coach assistance is never available in rated or active human games.

## 2. What “min-max breakdown” means

Modern Stockfish is not a plain textbook minimax implementation. It uses
alpha-beta Principal Variation Search, extensive pruning/reductions, a
transposition table, quiescence search, and NNUE position evaluation. Therefore
KnightShift will not display a fabricated complete tree or pretend every
discarded branch was fully evaluated.

The coach provides two related views:

### Engine decision view

For the actual Stockfish decision:

- exact engine and NNUE digests;
- requested and reached depth/seldepth;
- nodes, nodes per second, hash usage, and elapsed time;
- MultiPV top candidate moves;
- score and win/draw/loss estimate for each candidate;
- principal variation for each candidate;
- mate/tablebase result when applicable;
- evaluation stability across completed iterative-deepening depths;
- explicit label that unreported/pruned branches are not known to have the
  displayed score.

### Educational minimax view

For explanation, KnightShift may run a separate bounded search:

- user-selectable two to four plies;
- top configurable legal moves per node;
- maximize/minimize layers visibly distinguished;
- leaf position score and explanation features;
- propagated min/max values;
- alpha/beta window and cutoff markers;
- transposition references rather than duplicated subtrees;
- warning that this small tree explains the concept and is not the full
  Stockfish search.

This view must be generated from real legal positions and real returned values.
It cannot contain illustrative fake branches in a live analysis result.

## 3. Human-readable explanation

Because current Stockfish uses NNUE, its final score does not provide a clean
handcrafted “+0.3 king safety, +0.2 space” decomposition. KnightShift keeps
engine truth separate from explanation.

The explanation pipeline analyzes the current position and candidate lines for
verifiable concepts:

- material and exchanges;
- checks, captures, and threats;
- hanging or overloaded pieces;
- pins, forks, skewers, discoveries, and mating nets;
- king exposure and forcing access;
- pawn structure: isolated, doubled, passed, backward, majorities;
- development, center control, mobility, space, open files/diagonals;
- weak squares, outposts, bad/good pieces;
- promotion races and tablebase facts;
- opening-book identity and departure;
- the concrete refutation of inferior candidates.

Every sentence links to evidence: highlighted squares/pieces or a candidate
line. If the system cannot support a causal statement, it uses bounded language
such as “the engine prefers” and shows the line instead of inventing a reason.

## 4. Hint ladder

Hints reveal progressively:

1. goal category: defend, improve, force, simplify, or calculate;
2. relevant region/piece;
3. tactical/strategic motif;
4. first move candidates without ranking;
5. strongest move;
6. principal variation.

Each level is recorded so training statistics distinguish unaided success from
assisted play.

## 5. Bot strength and mistakes

Stockfish `UCI_LimitStrength`/`UCI_Elo` and skill settings are a starting point,
not a promise of exact human rating. Profiles are calibrated by automated
matches at the production hardware/time budget. Published strength is a range
with calibration version and confidence.

Weaker bots choose among MultiPV candidates using a controlled distribution
based on evaluation loss, phase, and profile—not random illegal moves or
arbitrary blunders. A profile's claimed style must be supported by measured
features across a calibration match set.

## 6. Data and API

Persist:

- bot profile/configuration version;
- Stockfish/NNUE digests;
- hardware class and analysis budget;
- all game commands and clocks;
- hint requests and levels;
- candidate sets and engine provenance used for feedback;
- user feedback on explanation usefulness;
- final review.

Engine outputs are cached by position, engine digest, options, and budget.
Cache entries are never reused across mismatched provenance.

## 7. Acceptance gates

- bot games obey the same authoritative rules and clocks as human games;
- no coach endpoint serves an active human-game participant;
- candidate moves and PVs reproduce with the recorded engine build/options
  within documented nondeterminism;
- every displayed educational-tree branch is legal and its propagated value is
  mechanically verified;
- explanations never reference absent pieces, illegal lines, or unsupported
  tactical motifs;
- hint levels reveal only their promised information;
- skill profiles pass calibration bounds;
- engine timeout/crash cannot affect live human games;
- keyboard and screen-reader users can navigate candidate lines and tree levels.

## 8. Primary engine references

- Stockfish advanced topics and NNUE:
  https://official-stockfish.github.io/docs/stockfish-wiki/Advanced-topics.html
- Stockfish terminology (PVS, move ordering, quiescence):
  https://official-stockfish.github.io/docs/stockfish-wiki/Terminology.html
- Stockfish FAQ (evaluation, MultiPV, skill):
  https://official-stockfish.github.io/docs/stockfish-wiki/Stockfish-FAQ.html


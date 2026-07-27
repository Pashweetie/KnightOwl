# Capability Matrix and Difficulty Disclosure

Terminology used below: Computer Aggregated Precision Score (CAPS), central
processing unit (CPU), direct message (DM), Forsyth–Edwards Notation (FEN),
OpenID Connect (OIDC), peer-to-peer (P2P), Portable Game Notation (PGN),
Traversal Using Relays around Network Address Translation (TURN), television
(TV), and user interface (UI).

## Purpose

This is the scope map for a friends-only, self-hosted play-and-improve chess
service. Chess.com provides market-comparison research; KnightOwl requirements
use original or appropriately licensed implementations and assets.

The baseline was researched from Chess.com's public help center on 2026-07-27.
Product constraints added afterward are authoritative: no first-party login or
contact identity, no public community, no cheating enforcement, no hidden
durable player profiles, server multiplayer first, and puzzles are nice to have.

Status:

- **Release**: required for KnightOwl's complete release.
- **Local release**: complete capability stored on the person's browser.
- **Nice**: implemented after release requirements.
- **Deferred by privacy decision**: cannot honestly work across devices without
  durable identity/data; requires explicit approval, not a disguised shortcut.
- **Excluded**: explicitly outside the friends-only product.

## Play

| Capability | Status | KnightOwl behavior | Difficulty |
|---|---|---|---|
| Standard live chess | Release | Private invite games; bullet through classical; custom base/increment/delay | High correctness/concurrency |
| Complete board controls | Release | Click/drag/touch/keyboard, configurable legal cues, premove queue, promotion and auto-queen override | High accessibility/browser coverage |
| Game actions | Release | Draw offer/answer, resign, abort rules, rematch, reconnect, PGN | Medium |
| Spectating | Release | Host-controlled live spectators; no anti-cheat delay required | Medium fan-out |
| Correspondence chess | Excluded | Conflicts with expiring anonymous rooms and is not a premium learning service | Requires durable continuity |
| Chess variants | Post-release candidate | No variant is advertised until an established rules library and upstream conformance tests cover it | Library decision required |
| Bot Mode | Release | Pinned Stockfish profiles, hints, questions, grounded explanations, MultiPV, educational minimax tree, Undo, and Redo in one unified mode | Very high explanation quality and calibration |
| Odds games | Release | Time, material, and starting-position odds for casual rooms | Medium |

## Review and analysis

| Capability | Status | KnightOwl behavior | Difficulty |
|---|---|---|---|
| Full Game Review | Release | Evaluation graph, turning points, lines, opening/phase, engine provenance | High CPU/UI |
| Accuracy | Release | Original published/versioned formula, never represented as Chess.com Accuracy/CAPS | Very high calibration |
| Move classifications | Release | Original documented categories/thresholds grounded in engine and position facts | Very high calibration |
| Coach summaries | Release | Original deterministic explanations tied to legal variations/features | Very high; quality risk |
| Retry mistakes | Release | Replay key positions and compare legal alternatives | Medium |
| Self analysis | Release | Variations, annotations, arrows, engine settings, evaluation | High |
| Deeper analysis | Release | Bounded queued Stockfish analysis; result returned or retained locally | High resource admission |
| PGN/FEN import/export | Release | Strict parser with comments/variations and local files | High parser correctness |
| Opening explorer | Release | Permissively licensed corpus plus imported/local games | Very high data licensing/indexing |
| Endgame tablebases | Release | Local Syzygy for an operator-selected piece count with clear storage requirement | High storage/distribution |

## Puzzles

| Capability | Status | KnightOwl behavior | Difficulty |
|---|---|---|---|
| Puzzle trainer | Nice | Licensed/original corpus, complete validated lines, themes, retry, local schedule | High content/licensing |
| Custom filters | Nice | Theme, range, color, failed/unseen, opening, time | Medium after corpus |
| Rush/survival | Nice | Timed and survival sessions with real local records | Medium |
| Puzzle Battle | Nice | Private synchronized friend room; no global pairing/rating | High synchronization |
| Daily puzzle | Nice | Scheduled licensed/original puzzle and explanation; no comments | Editorial work |
| Game-derived recommendations | Nice | Local reviewed mistakes mapped to puzzle themes | Very high integration |
| Global puzzle rating/leaderboard | Deferred by privacy decision | Requires durable player and attempt records | Identity/privacy conflict |

## Training

| Capability | Status | KnightOwl behavior | Difficulty |
|---|---|---|---|
| Drills/endgames | Release | Curated positions, goals, hints, local repetition schedule | High content |
| Vision trainer | Release | Coordinates, colors, paths, blindfold, local scores | Medium |
| Recommendations | Local release | Explainable suggestions from local games, puzzles, and drill history | High browser analytics |
| Progress and streaks | Local release | IndexedDB with export/clear; no server sync | Medium |

## Statistics

| Capability | Status | KnightOwl behavior | Difficulty |
|---|---|---|---|
| Personal result/opening/time statistics | Local release | Calculated from locally retained/imported PGNs | High local query/UI |
| Insights | Local release | Accuracy, phase, openings, tactics, pieces, castling, time usage | Very high definitions |
| Advanced skill measures | Local release | Versioned confidence-bearing measures; no unsupported peer percentile | Very high validation |
| Sacrifice/highlight collection | Local release | Based on published KnightOwl classifier | Medium after review |
| Cross-device history | Deferred by privacy decision | Needs durable identity and server archive | Identity/privacy conflict |
| Global ratings/leaderboards/percentiles | Deferred by privacy decision | Honest global values require durable stable participants and games | Identity/privacy conflict |
| Achievements | Local release | Rule-defined, computed from local evidence, exportable | Medium |

## Private competition and interaction

| Capability | Status | KnightOwl behavior | Difficulty |
|---|---|---|---|
| Direct challenges | Release | Expiring role-specific invite links | Medium |
| Arena tournaments | Deferred | Continuous join-and-play format is not needed for the initial private tournament requirement | Separate future product and library decision |
| Swiss tournaments | Release | Invoke pinned bbpPairings through the International Chess Federation Tournament Report Format; KnightOwl does not implement the pairing algorithm | Integration and lifecycle only |
| Fixed reactions | Release | Muteable, host-disableable, no custom content | Low |
| Friends/follows/blocks | Excluded | Invite sharing happens outside KnightOwl | Explicit friends-only decision |
| Chat/DM/forums/clubs | Excluded | No user-generated text or public community | Explicit moderation/data decision |
| Public tournaments/classrooms | Excluded | Require identity, discovery, moderation, and child-safety programs | Separate product |
| News/blogs/broadcast TV | Excluded | Publishing/broadcast business, not a chess play-and-improve feature | Separate product |

## Identity, safety, and operations

| Capability | Status | KnightOwl behavior | Difficulty |
|---|---|---|---|
| Accountless join | Release | Invite claim, room label, rotating reconnect capability | High protocol complexity |
| Passkeys/OIDC | Deferred by privacy decision | Researched future options, not speculative dependencies | Medium/high |
| Themes/settings | Release | Independent board, sound, cursor, and full piece skins: Classic, Rounded, Diamond Cut, Magma Forge, Neon Arcade, Pixel Quest, Plush Court, and High Contrast | High original asset quality |
| Meme reaction effects | Release | Original pack selected by default plus browser-local user-imported packs; each event has an independent animation selector, with dedicated GIF/video/animation customization for checkmate | Medium import validation |
| Notifications | Release | In-app while connected; optional browser-local reminders | Medium/browser restrictions |
| Cheat detection/reports/appeals | Excluded | Friends-only decision; server still enforces rules and protocol correctness | Explicit decision |
| Data disclosure | Release | Runtime-generated data/retention/provider view | Medium |
| Local export/deletion | Release | PGN/preferences/progress export and one-action clear | Medium |
| Accessibility/mobile web | Release | Touch, keyboard, focus, reduced motion, contrast, and semantic library components; full screen-reader board narration optional | Medium ongoing quality assurance |
| Docker deployment | Release | Hardened Compose, Traefik replicas, Valkey, workers, cloudflared | High operations |
| P2P transport | Post-release planned | WebRTC direct/TURN with server fallback after dedicated spike | Very high networking |

## Hardest items, without understatement

1. A high-quality coach explanation and accuracy/classification system requires
   extensive chess-domain validation; Stockfish output alone is insufficient.
2. Opening, puzzle, and tablebase distribution requires licenses, provenance, storage,
   and update operations.
3. Global ratings, leaderboards, cross-device history, and correspondence are
   not honestly deliverable under the current no-durable-identity/no-archive
   decision. They are disclosed as blocked by that product tradeoff rather than
   faked or silently omitted.

Tournament pairing algorithms are not on this custom-work list. Swiss pairing
is delegated to bbpPairings. Arena remains conditional until a maintained,
license-compatible implementation passes the dependency proof; KnightOwl
will not recreate that algorithm merely to retain a matrix row.

## Roadmap traceability audit

Every included capability discussed across the documentation maps to the
following roadmap phase:

| Roadmap phase | Covered capabilities |
| --- | --- |
| 2 | Repository workflow; Docker Compose; K3s; Traefik; Valkey; Stockfish workers; Cloudflare Tunnel; Flux; Argo Rollouts; diagnostics; installation; upgrade; rollback |
| 3 | Standard rules and endings; bullet/blitz/rapid/classical clocks; increment; delay; timeout; odds games; PGN/FEN parsing; canonical events |
| 4 | Direct challenges; accountless room labels; readiness; White-first start; moves; draw; resign; abort; rematch; premoves; spectating; reconnect; fixed reactions |
| 5 | Mouse/touch/keyboard board; legal cues; promotion/auto-queen; orientation; clocks; notation; captured material; local history; notifications; data disclosure; export/deletion; responsive/accessibility behavior; board/sound/cursor/piece/interface themes; imported check/checkmate animations |
| 6 | Bot profiles and calibration; hints; questions; feedback timing; plan explanations; local speech; MultiPV; candidate lines; educational minimax; Undo/Redo |
| 7 | Game Review; accuracy; classifications; summaries; Retry Mistakes; self analysis; deeper analysis; opening explorer; tablebases; drills/endgames; vision training; recommendations; progress/streaks; statistics; Insights; skill measures; highlights; achievements |
| 8 | Optional puzzle trainer; filters; Rush; survival; Puzzle Battle; daily puzzle; game-derived recommendations |
| 9 | Private scheduled-round Swiss tournaments, pairing, standings, tie-breaks, withdrawal, forfeit, reconnect, spectators, and export |
| 10 | Private release review; chaos, load, recovery, retention, clean-install, tunnel, and two-player publication checks |

Excluded and privacy-blocked rows intentionally have no implementation phase.
Post-release candidates remain in `docs/future/` until separately approved.

## Baseline sources

- Premium comparison: <https://support.chess.com/en/articles/8562418>
- Live moves/premoves: <https://support.chess.com/en/articles/8708726>
  and <https://support.chess.com/en/articles/8562432>
- Daily chess: <https://support.chess.com/en/articles/8588171>
- Game Review: <https://support.chess.com/en/articles/8584089>
- Puzzles: <https://support.chess.com/en/articles/8608686>
- Insights: <https://support.chess.com/en/articles/8708925>
- Variants: <https://support.chess.com/en/articles/8583983>
- Arena tournaments: <https://support.chess.com/en/articles/8562889>
- Opening explorer: <https://support.chess.com/en/articles/8615183>

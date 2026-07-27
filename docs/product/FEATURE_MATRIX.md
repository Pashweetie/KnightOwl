# Capability Matrix and Difficulty Disclosure

## Purpose

This is the scope map for a friends-only, self-hosted alternative to Chess.com's
play-and-improve capabilities. It does not copy Chess.com code, branding, text,
assets, puzzle sets, formulas, data, or labels.

The baseline was researched from Chess.com's public help center on 2026-07-27.
Product constraints added afterward are authoritative: no first-party login or
contact identity, no public community, no cheating enforcement, no hidden
durable player profiles, server multiplayer first, and puzzles are nice to have.

Status:

- **Release**: required for KnightShift's complete release.
- **Local release**: complete capability stored on the person's browser.
- **Nice**: implemented after release requirements.
- **Deferred by privacy decision**: cannot honestly work across devices without
  durable identity/data; requires explicit approval, not a disguised shortcut.
- **Excluded**: explicitly outside the friends-only product.

## Play

| Capability | Status | KnightShift behavior | Difficulty |
|---|---|---|---|
| Standard live chess | Release | Private invite games; bullet through classical; custom base/increment/delay | High correctness/concurrency |
| Complete board controls | Release | Click/drag/touch/keyboard, configurable legal cues, premove queue, promotion and auto-queen override | High accessibility/browser coverage |
| Game actions | Release | Draw offer/answer, resign, abort rules, rematch, reconnect, PGN | Medium |
| Spectating | Release | Host-controlled live spectators; no anti-cheat delay required | Medium fan-out |
| Simultaneous exhibition | Release | Private host simul with measured participant limit and auto-switch | High multi-board UX |
| Correspondence chess | Deferred by privacy decision | Needs state lasting days/weeks and reliable participant continuity/notifications | High; incompatible with expiring anonymous rooms |
| Chess960 | Release | Correct starting positions, castling, notation, analysis | High rules/testing |
| 3-Check, King of the Hill, Crazyhouse | Release | Independent rules, termination, notation, engine/UI support | High |
| Bughouse/four-player | Release | Two boards, teams, transferred pieces, synchronized result | Very high protocol/UI |
| Bot opponents | Release | Pinned Stockfish profiles with calibrated behavior and no fictional humans | High calibration |
| Coach-assisted bot play | Release | Hints, questions, grounded explanations, MultiPV, educational minimax tree | Very high explanation quality |
| Odds games | Release | Time, material, and starting-position odds for casual rooms | Medium |

## Review and analysis

| Capability | Status | KnightShift behavior | Difficulty |
|---|---|---|---|
| Full Game Review | Release | Evaluation graph, turning points, lines, opening/phase, engine provenance | High CPU/UI |
| Accuracy | Release | Original published/versioned formula, never represented as Chess.com Accuracy/CAPS | Very high calibration |
| Move classifications | Release | Original documented categories/thresholds grounded in engine and position facts | Very high calibration |
| Coach summaries | Release | Original deterministic explanations tied to legal variations/features | Very high; quality risk |
| Retry mistakes | Release | Replay key positions and compare legal alternatives | Medium |
| Self analysis | Release | Variations, annotations, arrows, engine settings, evaluation | High |
| Deeper analysis | Release | Bounded queued Stockfish analysis; result returned or retained locally | High resource admission |
| PGN/FEN import/export | Release | Strict secure parser with comments/variations and local files | High parser/security |
| Opening explorer | Release | Permissively licensed corpus plus imported/local games | Very high data licensing/indexing |
| Endgame tablebases | Release | Local Syzygy for an operator-selected piece count with clear storage requirement | High storage/distribution |

## Puzzles

| Capability | Status | KnightShift behavior | Difficulty |
|---|---|---|---|
| Puzzle trainer | Nice | Licensed/original corpus, complete validated lines, themes, retry, local schedule | High content/licensing |
| Custom filters | Nice | Theme, range, color, failed/unseen, opening, time | Medium after corpus |
| Rush/survival | Nice | Timed and survival sessions with real local records | Medium |
| Puzzle Battle | Nice | Private synchronized friend room; no global pairing/rating | High synchronization |
| Daily puzzle | Nice | Scheduled licensed/original puzzle and explanation; no comments | Editorial work |
| Game-derived recommendations | Nice | Local reviewed mistakes mapped to puzzle themes | Very high integration |
| Global puzzle rating/leaderboard | Deferred by privacy decision | Requires durable player and attempt records | Identity/privacy conflict |

## Lessons and training

| Capability | Status | KnightShift behavior | Difficulty |
|---|---|---|---|
| Structured learning paths | Release | Original beginner-to-advanced curriculum with prerequisites | Extremely high authoring effort |
| Interactive lessons | Release | Board checkpoints and original correct/incorrect explanations | Extremely high content QA |
| Searchable library | Release | Level, phase, theme, opening, duration | Medium |
| Video lessons | Release | Original/licensed, captions, transcript, chapters | Extremely high production/licensing |
| Drills/endgames | Release | Curated positions, goals, hints, local repetition schedule | High content |
| Vision trainer | Release | Coordinates, colors, paths, blindfold, local scores | Medium |
| Recommendations | Local release | Explainable suggestions from local games/puzzles/lesson mastery | High browser analytics |
| Progress and streaks | Local release | IndexedDB with export/clear; no server sync | Medium |

## Statistics

| Capability | Status | KnightShift behavior | Difficulty |
|---|---|---|---|
| Personal result/opening/time statistics | Local release | Calculated from locally retained/imported PGNs | High local query/UI |
| Insights | Local release | Accuracy, phase, openings, tactics, pieces, castling, time usage | Very high definitions |
| Advanced skill measures | Local release | Versioned confidence-bearing measures; no unsupported peer percentile | Very high validation |
| Sacrifice/highlight collection | Local release | Based on published KnightShift classifier | Medium after review |
| Cross-device history | Deferred by privacy decision | Needs durable identity and server archive | Identity/privacy conflict |
| Global ratings/leaderboards/percentiles | Deferred by privacy decision | Honest global values require durable stable participants and games | Identity/privacy conflict |
| Achievements | Local release | Rule-defined, computed from local evidence, exportable | Medium |

## Private competition and interaction

| Capability | Status | KnightShift behavior | Difficulty |
|---|---|---|---|
| Direct challenges | Release | Expiring role-specific invite links | Medium |
| Arena tournaments | Release | Private room, ephemeral pairing/scoring/tie-breaks | High |
| Swiss tournaments | Release | Private rounds, score groups, colors, repeats, byes, tie-breaks | Very high validation |
| Vote chess | Release | Private room teams and move-vote windows | High |
| Fixed reactions | Release | Muteable, host-disableable, no custom content | Low |
| Friends/follows/blocks | Excluded | Invite sharing happens outside KnightShift | Explicit friends-only decision |
| Chat/DM/forums/clubs | Excluded | No user-generated text or public community | Explicit moderation/data decision |
| Public tournaments/classrooms | Excluded | Require identity, discovery, moderation, and child-safety programs | Separate product |
| News/blogs/broadcast TV | Excluded | Publishing/broadcast business, not a chess play-and-improve feature | Separate product |

## Identity, safety, and operations

| Capability | Status | KnightShift behavior | Difficulty |
|---|---|---|---|
| Accountless join | Release | Invite claim, room label, rotating reconnect capability | High security |
| Passkeys/OIDC | Deferred by privacy decision | Researched future options, not speculative dependencies | Medium/high |
| Themes/settings | Release | Accessible bundled themes and local preferences | Medium/design |
| Notifications | Release | In-app while connected; optional browser-local reminders | Medium/browser restrictions |
| Cheat detection/reports/appeals | Excluded | Friends-only decision; server still enforces rules and protocol correctness | Explicit decision |
| Data disclosure | Release | Runtime-generated data/retention/provider view | Medium |
| Local export/deletion | Release | PGN/preferences/progress export and one-action clear | Medium |
| Accessibility/mobile web | Release | Keyboard/screen reader/touch/reduced motion/contrast across target browsers | High ongoing QA |
| Docker deployment | Release | Hardened Compose, Traefik replicas, Valkey, workers, cloudflared | High operations |
| P2P transport | Post-release planned | WebRTC direct/TURN with server fallback after dedicated spike | Very high networking |

## Hardest items, without understatement

1. A large original lesson/video curriculum is primarily an editorial and
   production project, not a coding task.
2. A high-quality coach explanation and accuracy/classification system requires
   extensive chess-domain validation; Stockfish output alone is insufficient.
3. Bughouse, simuls, Swiss pairing, correspondence, and P2P host migration each
   need distinct state-machine and concurrency work.
4. Opening/puzzle/tablebase distribution requires licenses, provenance, storage,
   and update operations.
5. Global ratings, leaderboards, cross-device history, and correspondence are
   not honestly deliverable under the current no-durable-identity/no-archive
   decision. They are disclosed as blocked by that product tradeoff rather than
   faked or silently omitted.

## Baseline sources

- Premium comparison: <https://support.chess.com/en/articles/8562418>
- Live moves/premoves: <https://support.chess.com/en/articles/8708726>
  and <https://support.chess.com/en/articles/8562432>
- Daily chess: <https://support.chess.com/en/articles/8588171>
- Game Review: <https://support.chess.com/en/articles/8584089>
- Puzzles: <https://support.chess.com/en/articles/8608686>
- Lessons: <https://support.chess.com/en/articles/8609703>
- Insights: <https://support.chess.com/en/articles/8708925>
- Variants: <https://support.chess.com/en/articles/8583983>
- Arena tournaments: <https://support.chess.com/en/articles/8562889>
- Opening explorer: <https://support.chess.com/en/articles/8615183>

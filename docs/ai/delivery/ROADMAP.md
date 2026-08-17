# Delivery Roadmap

Terminology used below: application programming interface (API), continuous
integration (CI), International Chess Federation (FIDE), high availability
(HA), Network Address Translation (NAT), Efficiently Updatable Neural Network
(NNUE), OpenID Connect (OIDC), personal computer (PC), Portable Game Notation
(PGN), Structured Query Language (SQL), Tournament Report Format (TRF), time to
live (TTL), Traversal Using Relays around Network Address Translation (TURN),
and user interface (UI). BLAKE3 is a proper algorithm name; ZIP is the proper
name of the archive format rather than an abbreviation.

This is an ordered release gate. Work proceeds one section at a time. A box is
checked only after its real behavior and automated tests pass. There are no
placeholder controls, invented users/statistics, or fake service integrations.

## 1. Resolve the product and architecture

- [x] Define friends-only scope and explicit exclusions.
- [x] Define standard-chess rules and clock semantics.
- [x] Define White's first accepted move as the live-game start boundary.
- [x] Research established chess conformance suites, reference libraries,
  notation standards, pairing engines, and differential oracles.
- [x] Research the competitor capability baseline and define independently
  licensed implementation requirements.
- [x] Define unified Bot Mode with built-in teaching assistance and truthful
  minimax/Stockfish presentation.
- [x] Define anonymous capability rooms and the minimal authoritative state
  protocol.
- [x] Research anonymous/low-data authentication and define capability,
  passkey, and pairwise-OIDC tradeoffs.
- [x] Evaluate peer-to-peer WebRTC transport, NAT/TURN requirements, privacy,
  authority, reconnect, and multi-party limitations.
- [x] Research pseudonymous identity systems and define KnightOwl's independent
  low-data capability model.
- [x] Select Valkey for expiring live state and exclude SQL from the initial
  release.
- [x] Define the scale path: Valkey HA/Cluster first, ScyllaDB only if durable
  multi-host state becomes a measured requirement, and ClickHouse only for
  privacy-safe aggregate observability at proven volume.
- [x] Select Traefik for local Docker load balancing behind cloudflared.
- [x] Define friends-only room controls and eliminate public social moderation.
- [x] Define purpose-limited aggregate/diagnostic telemetry, forbidden tracking,
  retention, runtime disclosure, and the legal-risk boundary.
- [x] Define repository quality, security, chaos, and operational gates.
- [x] Select Go for the backend based on readability, library coverage,
  maintenance, and deployment rather than synthetic scale.
- [x] Select GPL-3-or-later licensing.
- [x] Select scheduled-round private Swiss as the required tournament format;
  exclude continuous arena tournaments from the initial release.
- [x] Approve the ordered Go server implementation tickets in
  `docs/ai/delivery/GO_SERVER_IMPLEMENTATION.md`; do not create a disposable
  proof-of-concept server.
- [ ] Review the final architecture documents for contradictions.
- [ ] Obtain explicit user sign-off on product scope, license, dependencies,
  architecture, difficult-feature disclosures, and implementation sequence.

## 2. Establish the repository and delivery system

- [ ] Configure the ticket-branch and pull-request workflow specified in
  `docs/ai/engineering/TICKET_BRANCH_WORKFLOW.md`, including protected `dev` and
  `release` branches and required continuous-integration checks.
- [ ] Deliver `server-001`, the Go service foundation, as the first executable
  application ticket after the branch-workflow prerequisite.
- [ ] Deliver the deployment system as separate reviewed tickets for Docker
  Compose development, K3s production, image supply chain, read-only preflight,
  digest-pinned promotion, authenticated Flux reconciliation, Argo Rollouts
  blue-green deployment, failure handling, rollback, and tunnel publication.
- [ ] Add an opt-in Docker Compose invite-tunnel profile: keep room creation and
  administration on loopback port `8787`, expose only the guest invite claim,
  play, spectate, and reconnect surface through cloudflared, and prove the
  network and route separation with private release validation.
- [ ] Prove that only an approved promotion merged into `release` can alter
  production; ordinary `dev` commits, pull requests, registry tags, missed or
  replayed events, and documentation-only changes cannot deploy an unreviewed
  image.
- [ ] Prove a clean-machine install, deliberate maintenance-window update,
  failed rollout, automatic restoration of the prior image digest, secret
  rotation, and complete removal before remote deployment is considered.
- [ ] Optionally create an isolated Kerberos learning-lab ticket only after the
  normal deployment path works; it must use disposable credentials and must not
  become an application or production dependency.
- [ ] Create the selected backend, TypeScript web client, protocol package,
  Stockfish worker, and end-to-end test workspace.
- [ ] Pin all toolchains, dependencies, container bases, and image digests.
- [ ] Add format, lint, dependency-boundary, unused-code, unit, integration,
  fuzz/property, accessibility, and end-to-end CI gates.
- [ ] Build hardened rootless images with SBOMs and provenance.
- [ ] Add Valkey, app, Stockfish worker, loopback operator, invite-only guest,
  and cloudflared Docker Compose development profiles.
- [ ] Add Kustomize production resources for K3s, K3s-packaged Traefik routes,
  Flux reconciliation, Argo Rollouts blue-green Services and analysis, network
  policies, persistent storage, and secrets references.
- [ ] Implement liveness, readiness, startup, drain, metrics, and private
  operator diagnostics.
- [ ] Verify that only loopback port `8787` is optional locally and that no
  internal service or dashboard is exposed.
- [ ] Document install, upgrade, scale, rollback, clean removal, and secret
  rotation.

## 3. Implement the chess core

- [ ] Qualify and integrate the selected maintained chess library behind a
  narrow adapter; do not create a move generator, legality engine, notation
  parser, repetition tracker, or terminal-position detector.
- [ ] Map every required rule and ending to verified library behavior or an
  upstream contribution; require a separately approved adapter only for a
  demonstrated gap.
- [ ] Implement Fischer increment, Bronstein delay, pause, reconnect, and
  server-authoritative timeout resolution.
- [ ] Implement casual time, material, and validated starting-position odds.
- [ ] Implement strict PGN and FEN import/export with comments, variations,
  result validation, and library-backed round trips.
- [ ] Implement canonical protocol encoding and BLAKE3 event chains.
- [ ] Add perft fixtures, published rule positions, property tests, differential
  tests, fuzzing, and deterministic clock tests.

## 4. Implement private multiplayer rooms

- [ ] Create unlisted rooms and separate hashed host/player/spectator
  capabilities.
- [ ] Implement one-use invites, color assignment, readiness, settings, and
  capability revocation.
- [ ] Implement move, resign, draw, abort, rematch, spectate, reconnect, and PGN
  export.
- [ ] Implement disconnect grace, abandonment, authorized room termination, and
  every automatic and claimable draw boundary.
- [ ] Keep games Ready until White's first accepted legal move starts both
  clocks in one authoritative transition.
- [ ] Execute visible queued premoves in order with cancellation, reconciliation,
  legality checks, and the documented clock charge.
- [ ] Implement atomic Valkey version checks, idempotency, snapshots, event TTL,
  and explicit failure behavior.
- [ ] Implement fixed reactions, local mute, and host disable controls.
- [ ] Add two-browser, simultaneous-command, multi-replica, restart, expiry, and
  tampered-client tests.

## 5. Implement the complete chess interface

- [ ] Complete and approve the frontend dependency proof in
  `docs/research/FRONTEND_LIBRARY_STRATEGY.md`; record license and maintenance
  decisions before writing interface code.
- [ ] Build responsive mouse, touch, and keyboard board interaction.
- [ ] Add legal-move cues, promotion, orientation, clocks, move list, captured
  material, status, settings, reconnect UI, numbered premoves, and configurable
  auto-queen with promotion override.
- [ ] Add coordinates, last-move and check indicators, analysis arrows, focus
  mode, sound controls, and low-time warnings.
- [ ] Add licensed cute cursor themes using the browser cursor standard,
  correct hotspots and fallbacks, local preferences, previews, and a System
  Default reset.
- [ ] Add independent full chess-piece skins: licensed Classic/Rounded plus
  original Diamond Cut, Magma Forge, Neon Arcade, Pixel Quest, Plush Court, and
  High Contrast sets, with twelve-piece previews, promotion/capture states, local
  preference, provenance inventory, fallback, performance budgets, and
  cross-browser visual tests.
- [ ] Create one original library-driven Reddit-style meme reaction pack and a
  browser-local pack importer/exporter using standard ZIP archives, validated
  manifests, non-executable bounded media, zero uploads or third-party
  requests, and Off/Subtle/Full controls; select the original pack by default
  on new browsers. Let each event, especially checkmate, use a separately
  selected imported GIF, video, or animation with preview, sound, duration,
  replay, and reset controls.
- [ ] Add an assisted Tenor workflow that opens Tenor externally and imports a
  user-downloaded local media file with optional source attribution; do not
  depend on the decommissioned Tenor API or scrape its pages.
- [ ] Add visible focus, reduced motion, contrast, non-color cues, and semantic
  controls through the selected component library.
- [ ] Treat full screen-reader chess-board navigation and announcements as an
  optional feature unless an intended player needs it.
- [ ] Implement IndexedDB game history, import/export, and clear-local-data.
- [ ] Add independent board, sound, cursor, piece, animation, and interface
  preferences with previews and resets.
- [ ] Add in-app game/connection notifications and optional browser-local
  reminders without a server notification profile.
- [ ] Add a runtime data disclosure plus one-action local export and deletion.
- [ ] Pass real Chrome, Firefox, Safari, and Edge desktop/mobile tests.

## 6. Implement Bot Mode

- [ ] Package a pinned Stockfish/NNUE build in a resource-isolated worker.
- [ ] Implement calibrated bot strength profiles without false rating claims.
- [ ] Implement MultiPV evaluations, principal variations, engine provenance,
  and resource limits.
- [ ] Implement a separate bounded educational minimax/alpha-beta tree labelled
  as an explanation rather than Stockfish internals.
- [ ] Ground every explanatory claim in a legal line, evaluation, motif detector, or
  endgame fact.
- [ ] Implement the full hint ladder, questions, immediate/delayed feedback,
  bot-plan explanations, candidate comparison, and local speech controls.
- [ ] Add visible Undo and Redo buttons for Bot Mode; undo one human
  turn plus its bot reply, cancel stale engine work, restore every derived
  view, and clear redo on a new branch.
- [ ] Add deterministic fixtures, calibration, overload, cancellation, and
  explanation-consistency tests.

## 7. Implement review, analysis, training, and local statistics

- [ ] Implement full Game Review with an evaluation graph, turning points,
  lines, opening/phase identification, engine provenance, grounded summaries,
  and locally reopenable results.
- [ ] Implement self analysis with variations, annotations, arrows, engine
  settings, evaluation, and position import.
- [ ] Implement documented accuracy and move-classification formulas plus a
  Retry Mistakes flow for replaying key positions.
- [ ] Implement bounded queued deeper analysis with visible queue state,
  resource budget, reached depth, nodes, engine version, and cancellation.
- [ ] Build an opening explorer from licensed games plus imported local games.
- [ ] Integrate operator-selected local Syzygy tablebases with explicit piece
  count and storage requirements.
- [ ] Add curated drills/endgames with goals, hints, verified solutions, and
  local repetition.
- [ ] Add coordinate, color, path, and blindfold vision trainers with local
  scores.
- [ ] Add explainable local recommendations from games, drills, and optional
  puzzle history.
- [ ] Add local result, opening, time, phase, tactical, piece, castling, and
  time-usage statistics and Insights.
- [ ] Add versioned skill measures, highlights/sacrifices, achievements,
  progress, and streaks from local evidence.
- [ ] Add engine, content-provenance, formula-calibration, and local-persistence
  tests.

## 8. Implement optional puzzle features

- [ ] Acquire or create a distributable puzzle corpus with provenance, complete
  verified lines, themes, and a documented rating method.
- [ ] Implement training, retry, local scheduling, and filters for theme, range,
  color, failed/unseen state, opening, and time.
- [ ] Implement optional Rush and survival sessions with browser-local records.
- [ ] Implement optional private Puzzle Battle with synchronized participants
  and authoritative results.
- [ ] Implement an optional daily puzzle with an explanation and no comments.
- [ ] Map locally reviewed mistakes to puzzle themes without uploading history.
- [ ] Add solution, synchronization, content, expiry, and local-progress tests.

Puzzle work is post-release “Nice” scope and does not block the core release.

## 9. Implement ephemeral room competition

- [ ] Integrate pinned bbpPairings through TRF for Swiss pairings; do not
  reimplement the FIDE Dutch algorithm.
- [ ] Implement private Swiss room orchestration; arena tournaments are outside
  the initial release.
- [ ] Implement entrants, pairings, scores, tie-breaks, withdrawals, reconnect,
  and PGN export without durable player identity.
- [ ] Implement registration close, late-entry policy, byes, no-shows,
  forfeits, host cancellation, spectator policy, and immutable started-event
  settings.
- [ ] Add invariant, concurrency, expiry, and recovery tests.

## 10. Qualify and publish

- [ ] Pass the ignored private security plan against the exact release images.
- [ ] Run Toxiproxy and container-kill chaos scenarios.
- [ ] Load-test room creation, WebSockets, moves, spectators, reconnects, proxy
  balancing, and Stockfish saturation on the target PC.
- [ ] Scale app replicas up/down during games and verify health-aware routing and
  graceful drain.
- [ ] Verify TTL deletion, log/metric privacy, disk exhaustion behavior, Valkey
  persistence boundaries, and clean-machine recovery.
- [ ] Perform a clean-machine installation from the operator documentation.
- [ ] Configure the operator-supplied named Cloudflare Tunnel and hostname.
- [ ] Run the public two-player smoke suite and publish only if all gates pass.

## Release rule

KnightOwl is not called complete or publicly published until every required
box in sections 1–7 and 9–10 is checked. Section 8 remains optional
post-release scope. Development previews remain local or access-restricted.

Excluded features receive no implementation task. Arena tournaments, variants,
passkeys/OIDC, peer-to-peer transport, and multi-host databases require their
future-feature approval before receiving tickets.

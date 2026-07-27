# KnightShift Delivery Plan

This is an ordered release gate. Work proceeds one section at a time. A box is
checked only after its real behavior and automated tests pass. There are no
placeholder controls, invented users/statistics, or fake service integrations.

## 0. Remove the rejected prototype

- [x] Remove its tunnel, containers, image, and remote repository.
- [x] Create a clean `knightshift` folder and repository.

## 1. Resolve the product and architecture

- [x] Define friends-only scope and explicit exclusions.
- [x] Define standard-chess rules and clock semantics.
- [x] Research the competitor capability baseline without copying proprietary
  content or branding.
- [x] Define the coaching bot and truthful minimax/Stockfish presentation.
- [x] Define anonymous capability rooms and the minimal authoritative state
  protocol.
- [x] Research anonymous/low-data authentication and define capability,
  passkey, and pairwise-OIDC tradeoffs.
- [x] Evaluate peer-to-peer WebRTC transport, NAT/TURN requirements, privacy,
  authority, reconnect, and multi-party limitations.
- [x] Inspect XMage's source-level anonymous/authenticated user models and
  document which low-friction ideas are safe to adopt.
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
- [ ] Complete Rust and Go backend spikes and record the measured selection.
- [ ] Review the final architecture documents for contradictions.

## 2. Establish the repository and delivery system

- [ ] Create the selected backend, TypeScript web client, protocol package,
  Stockfish worker, and end-to-end test workspace.
- [ ] Pin all toolchains, dependencies, container bases, and image digests.
- [ ] Add format, lint, dependency-boundary, unused-code, unit, integration,
  fuzz/property, accessibility, and end-to-end CI gates.
- [ ] Build hardened rootless images with SBOMs and provenance.
- [ ] Add Valkey, app replicas, Traefik, Docker API proxy, Stockfish worker, and
  cloudflared Compose profiles.
- [ ] Implement liveness, readiness, startup, drain, metrics, and private
  operator diagnostics.
- [ ] Verify that only loopback port `8787` is optional locally and that no
  internal service or dashboard is exposed.
- [ ] Document install, upgrade, scale, rollback, clean removal, and secret
  rotation.

## 3. Implement the chess core

- [ ] Implement immutable standard-chess position and move types.
- [ ] Implement every legal move and ending in the rules contract.
- [ ] Implement Fischer increment, Bronstein delay, pause, reconnect, and
  server-authoritative timeout resolution.
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
- [ ] Implement atomic Valkey version checks, idempotency, snapshots, event TTL,
  and explicit failure behavior.
- [ ] Implement fixed reactions, local mute, and host disable controls.
- [ ] Add two-browser, simultaneous-command, multi-replica, restart, expiry, and
  tampered-client tests.

## 5. Implement the complete chess interface

- [ ] Build responsive mouse, touch, and keyboard board interaction.
- [ ] Add legal-move cues, promotion, orientation, clocks, move list, captured
  material, status, settings, and reconnect UI.
- [ ] Add screen-reader board navigation and announcements, visible focus,
  reduced motion, and contrast compliance.
- [ ] Implement IndexedDB game history, import/export, and clear-local-data.
- [ ] Pass real Chrome, Firefox, Safari, and Edge desktop/mobile tests.

## 6. Implement bot play and coaching

- [ ] Package a pinned Stockfish/NNUE build in a resource-isolated worker.
- [ ] Implement calibrated bot strength profiles without false rating claims.
- [ ] Implement MultiPV evaluations, principal variations, engine provenance,
  and resource limits.
- [ ] Implement a separate bounded educational minimax/alpha-beta tree labelled
  as an explanation rather than Stockfish internals.
- [ ] Ground every coach claim in a legal line, evaluation, motif detector, or
  endgame fact.
- [ ] Add deterministic fixtures, calibration, overload, cancellation, and
  explanation-consistency tests.

## 7. Implement review and learning

- [ ] Implement local PGN analysis and interactive variations/annotations.
- [ ] Implement evidence-based move classifications and accuracy with documented
  formulas.
- [ ] Author original drills and lessons with local-only progress.
- [ ] Add licensed/original puzzles, verified solutions, local spaced
  repetition, and optional Rush/survival modes.
- [ ] Add engine, content-provenance, solution, and local-persistence tests.

## 8. Implement ephemeral room competition

- [ ] Implement private arena and Swiss room state machines.
- [ ] Implement entrants, pairings, scores, tie-breaks, withdrawals, reconnect,
  and PGN export without durable player identity.
- [ ] Add invariant, concurrency, expiry, and recovery tests.

## 9. Qualify and publish

- [ ] Run static analysis, dependency audit, secret scan, SBOM scan, Trivy,
  Nuclei, and authenticated OWASP ZAP against the release images.
- [ ] Run protocol fuzzing and authorization/capability penetration tests.
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

KnightShift is not called complete or publicly published until every box above
is checked. Development previews remain local or access-restricted.

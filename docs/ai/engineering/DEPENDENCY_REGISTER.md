# Dependency Decision Register

Terminology used below: application programming interface (API), Apache License
version 2.0 (Apache-2.0), Berkeley Software Distribution licenses (BSD),
Concise Binary Object Representation (CBOR), Cascading Style Sheets (CSS),
GNU General Public License (GPL), Hypertext Transfer Protocol (HTTP),
International Chess Federation (FIDE), International Software Consortium
License (ISC), Massachusetts Institute of Technology License (MIT), Portable
Game Notation (PGN), Structured Query Language (SQL), Tournament Report Format
(TRF), Universal Chess Interface (UCI), and user interface (UI). BLAKE3,
Stockfish, Valkey, Traefik, and Playwright are proper project names.

Status: selected dependency direction. Implementation tickets pin exact
versions and verify their transitive dependencies.

## Decision rule

Every implementation task selects a capability from this register before
writing code. If the selected dependency cannot satisfy the executable proof,
the task returns to research. Replacing it or writing custom code requires an
Architecture Decision Record that records the failed proof and alternatives.

Use a browser or language standard library when it completely solves the
problem. Use one primary library per responsibility. Do not add convenience
wrappers around wrappers.

## Browser runtime

| Responsibility | Selected dependency | License | Decision |
| --- | --- | --- | --- |
| Application view framework | React | MIT | Accepted by existing architecture |
| Build and development server | Vite | MIT | Recommended; production emits static files and does not run Vite |
| Page routing | React Router | MIT | Recommended; no custom history or route matcher |
| Generic interactive controls | React Aria Components | Apache-2.0 | Selected for its component coverage, adaptive interaction, internationalization, keyboard behavior, and maintained accessibility tests |
| Chess board | Chessground | GPL version 3 or later | Selected under KnightOwl's GPL-compatible license |
| Standard browser chess rules | chess.js | BSD two-clause | Recommended for standard chess previews, imports, local analysis, and differential checking; never authoritative for a live server game |
| Variant browser chess rules | chessops | GPL version 3 | Deferred with variants; do not ship two production rules libraries without need |
| Local structured storage | Dexie.js | Apache-2.0 | Recommended over raw Indexed Database calls because it supplies transactions, migrations, indexes, and browser workarounds |
| Request caching | TanStack Query | MIT | Use for ordinary HTTP resources only; live game events stay in the game adapter |
| Forms | React Hook Form | MIT | Recommended for settings, room creation, and import forms |
| Runtime schema validation | Zod | MIT | Recommended for form and local-file boundaries; network schema generation must have one canonical source |
| Coach charts and minimax tree | Apache ECharts | Apache-2.0 | Recommended; no custom chart renderer |
| Long lists | TanStack Virtual | MIT | Conditional on the measured rendering threshold |
| Interface animation | Motion | MIT | Recommended; no custom animation scheduler |
| Particle effects | canvas-confetti | ISC | Recommended with reduced-motion suppression always enabled |
| Rich reaction playback | dotLottie Web | MIT | Conditional on archive validation and bundle proof |
| ZIP pack handling | fflate | MIT | Recommended subject to decompression-limit proof; never extract paths directly to a filesystem |
| Icons | Lucide | ISC | Recommended with an icon inventory and recorded provenance |
| Cursor rendering | Browser CSS cursor standard | Platform | Accepted; a JavaScript cursor framework would be redundant |
| Audio playback and speech | Web Audio and Web Speech browser facilities | Platform | Accepted where browser support passes; no remote speech service |
| Live socket client | Browser WebSocket | Platform | Accepted; reconnect, sequence, and domain behavior remain KnightOwl policy |

## Go server

| Responsibility | Selected dependency | License | Decision |
| --- | --- | --- | --- |
| HTTP routing and middleware base | Go `net/http` | BSD-style Go license | Accepted; add a framework only after a demonstrated gap |
| WebSocket transport | coder/websocket | ISC | Recommended after Autobahn protocol-suite verification |
| Standard chess rules and notation | corentings/chess version 2 | MIT | Selected behind an adapter and verified by independent conformance fixtures |
| Valkey client | valkey-io/valkey-go | Apache-2.0 | Recommended official client; compare GLIDE only if cluster or telemetry proof reveals a gap |
| Metrics | Prometheus Go client | Apache-2.0 | Recommended for allowlisted aggregate metrics |
| Tracing and structured context | OpenTelemetry Go | Apache-2.0 | Conditional; disabled by default and cannot export identifying payloads |
| Logging | Go `log/slog` | BSD-style Go license | Accepted with allowlisted structured fields |
| Binary state encoding | fxamacker/cbor | MIT | Selected for deterministic CBOR behind the canonical protocol adapter |
| Event-chain hashing | zeebo/blake3 | public-domain dedication or MIT | Recommended implementation of BLAKE3; no local cryptographic implementation |
| Identifier generation | Go `crypto/rand` plus fixed encoding | BSD-style Go license | Accepted; no UUID dependency is necessary for random opaque capabilities |
| Testing containers | testcontainers-go | MIT | Recommended for Valkey, proxy, and worker integration tests |

## External executables and services

| Responsibility | Selected component | License | Boundary |
| --- | --- | --- | --- |
| Chess engine | Stockfish | GPL version 3 | Pinned isolated process through UCI; publish exact source and notices required by its license |
| Swiss pairing | bbpPairings | GPL | Pinned isolated process through TRF; KnightOwl never implements the Dutch pairing algorithm or claims FIDE endorsement |
| Live-state database | Valkey | BSD three-clause | Only initial server database; expiring game and tournament state |
| Lightweight production Kubernetes | K3s | Apache-2.0 | Self-hosted single-node production orchestrator; no claim of machine-level high availability |
| Production ingress and service routing | K3s-packaged Traefik | MIT | Routes only explicitly declared guest ingress and active application Services |
| Git-to-Kubernetes reconciliation | Flux | Apache-2.0 | Read-only access to the private `release` branch; authenticated event plus polling recovery |
| Blue-green application deployment | Argo Rollouts | Apache-2.0 | Maintains active and preview Services, runs pre-promotion analysis, and switches traffic after success |
| KnightOwl Kubernetes configuration | Kustomize through `kubectl` | Apache-2.0 | Readable base resources and small overlays; no custom template engine |
| Third-party Kubernetes packages | Helm, conditional | Apache-2.0 | Used only when a selected upstream component officially distributes a maintained chart |
| Outbound ingress | cloudflared | Apache-2.0 | Named Cloudflare Tunnel only after publication gates |
| Network fault injection | Toxiproxy | MIT | Test profile only |
| Load and soak testing | Grafana k6 | GNU Affero General Public License version 3 | Test executable only; scripts remain in the repository |
| Browser automation | Playwright | Apache-2.0 | Real multi-context interaction, compatibility, and deterministic visual testing |
| Component workshop | Storybook | MIT | Development and component-test build only |

## Content and asset dependencies

| Content | Source decision |
| --- | --- |
| Opening games, puzzles, and evaluations | Lichess database exports are a leading Creative Commons Zero source; pin snapshot, checksum, schema, and import provenance |
| Endgame truth | Syzygy tablebases at an operator-selected piece count; report exact installed coverage |
| Classic chess pieces | Colin M. L. Burnett source, GPL version 2 or later |
| Rounded chess pieces | Chessnut source, Apache-2.0 |
| Cursor choices | Kenney cursor pack, Creative Commons Zero |
| Default meme reactions | Original KnightOwl assets produced during approved visual work |
| User meme reactions | Browser-local imports; not a server or repository dependency |

## Explicit non-selections

- No general Go backend framework is adopted unless the standard HTTP library
  and selected focused dependencies demonstrate a reviewed gap.
- No Object-Relational Mapper because the initial server has no SQL database.
- No raw Redis-protocol reimplementation: use the official Valkey client.
- No custom chess rules, notation parser, Swiss pairing engine, board renderer,
  generic component kit, chart engine, animation scheduler, ZIP codec,
  cryptographic primitive, WebSocket framing, or load balancer.
- No simultaneous React Aria and Radix component stacks without a proven missing
  control.
- No Giphy or Tenor runtime integration. Tenor's API was decommissioned; local
  user downloads remain importable.
- No analytics, advertisement, session-replay, fingerprinting, or marketing
  software development kit.

## Primary project references

- React Aria Components: https://github.com/adobe/react-spectrum
- Chessground: https://github.com/lichess-org/chessground
- react-chessboard: https://github.com/Clariity/react-chessboard
- chess.js: https://github.com/jhlywa/chess.js
- Dexie.js: https://github.com/dexie/Dexie.js
- Lichess open databases: https://database.lichess.org/

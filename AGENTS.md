# KnightOwl Repository Instructions

Read `docs/ai/README.md` before changing anything. Follow its reading order,
then read the human product and architecture specifications linked by the
selected ticket.

## Authority

- Do not infer approval for a ticket, public tunnel publication, a pull-request
  merge, or a production promotion.
- Work on exactly one approved ticket at a time.
- Use a `test/<ticket-id>-<short-name>` branch created from current `dev`.
- Commits and pushes are permitted only on that ticket branch.
- Open a GitHub pull request into `dev`; never push directly to `dev` or
  `release`.
- Present exactly one squashed commit relative to `dev`.
- After every correction to a ticket branch, create and push a new revision
  before reporting the correction back to the repository owner.
- Do not report ticket work as complete until its branch is pushed and its
  GitHub pull request is open. A comparison URL is not a pull request.
- Do not merge until the designated reviewer approves that pull request.
- A separately reviewed pull request from `dev` into `release` authorizes
  immediate production deployment.
- Never invent or configure an agent, ChatGPT, OpenAI, or vendor Git identity.
  Use only the repository owner's configured identity.

## Engineering rules

- Before writing custom code, document the maintained libraries and platform
  standards evaluated. Prefer a focused adapter around an established library.
- Do not create chess move generation, board rendering, Swiss pairing,
  WebSocket framing, Valkey protocol, cryptography, charting, generic interface
  components, container orchestration, or blue-green deployment when the
  selected maintained dependency covers it.
- No placeholder, fake integration, fabricated result, seed data presented as
  real, or incomplete control counts as a finished feature.
- Keep the initial acceptance target to approximately ten friends. Record
  performance as a sanity check; do not optimize for hypothetical scale.
- Preserve the accountless, low-data, invite-only privacy boundary.
- Explain every acronym before its first use in each document.
- Keep unrelated cleanup out of the ticket.
- Run the ticket's required tests and report failures honestly.

## Accepted technical direction

- Go backend; TypeScript React browser.
- GNU General Public License version 3 or later.
- Maintained chess library behind a qualified adapter; Chessground is the
  leading board candidate.
- Valkey is the only initial server database.
- Stockfish is an isolated engine process.
- Private scheduled-round Swiss tournaments use bbpPairings.
- Docker Compose is local development; its optional Cloudflare Tunnel exposes
  existing game invites only while room creation stays on loopback port `8787`.
- Production is self-hosted single-node K3s with Traefik.
- Flux reconciles reviewed `release` state after an authenticated GitHub event,
  with polling recovery.
- Argo Rollouts provides blue-green deployment and preview analysis.
- Kustomize manages KnightOwl Kubernetes resources; Helm is conditional for
  upstream third-party packages only.

If any current code, task, or document contradicts these instructions, stop the
ticket and reconcile the contradiction in its own reviewed change.

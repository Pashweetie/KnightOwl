# KnightOwl

KnightOwl is a planned self-hosted chess service for small groups of friends.
It combines private real-time multiplayer, teaching-oriented Bot Mode, Swiss
tournaments, game review, optional puzzles, and browser-local customization.

> **Current status:** architecture review. This repository contains
> specifications and implementation-ticket definitions; it does not yet contain a runnable
> KnightOwl application.

## Developer starting point

1. Read the [human documentation index](docs/README.md).
2. Select an authorized ticket from the
   [implementation catalog](docs/ai/tickets/README.md).
3. Follow the [change workflow](docs/ai/engineering/TICKET_BRANCH_WORKFLOW.md).

The Go service foundation is the first executable application ticket. Run it
locally with `go run ./apps/server/cmd/knightowl`; it reads
`config/development.json` and listens on
`127.0.0.1:8787` by default. Its liveness, startup, readiness, and drain
endpoints are `/health/live`, `/health/startup`, `/health/ready`, and
`/admin/drain`. Pass `-config path/to/file.json` to use a different JSON
configuration file.

The server is a program that waits for requests. “Draining” is the standard
infrastructure term for a graceful shutdown: the server first tells the network
“do not send me anything new,” then gives requests that already arrived a few
seconds to finish, and finally exits. This matters during a restart or update
because it avoids cutting somebody off mid-action.

## Selected technical direction

- Go backend and TypeScript React browser
- Valkey for expiring authoritative game and room state
- Stockfish for Bot Mode and post-game analysis
- Docker Compose for local development
- K3s, Flux, and Argo Rollouts for self-hosted production deployment
- Cloudflare Tunnel for invite-only internet access
- GNU General Public License version 3 or later

The initial operating target is approximately ten friends. The design retains
clean service boundaries for later scaling without optimizing prematurely.

## Contribution workflow

All work starts from a ticket and uses a branch named
`test/<ticket-id>-<short-name>`.

1. Read the ticket and linked specifications.
2. Evaluate maintained libraries before proposing custom code.
3. Implement and test one reviewable vertical change.
4. Open a pull request from the ticket branch into `dev`.
5. Squash the pull request to one commit.
6. Merge only after required checks pass and the designated reviewer approves.
7. Promote `dev` to `release` through a separate reviewed pull request.

Direct and force pushes to `dev` and `release` are prohibited. Implementation
remains blocked until the live GitHub rulesets enforce that policy. After the
deployment system exists, merging into `release` initiates production
deployment.

See [Ticket Branch and Pull Request Workflow](docs/ai/engineering/TICKET_BRANCH_WORKFLOW.md)
for the complete change-review process.

## Repository map

| Path | Audience and purpose |
| --- | --- |
| `docs/product/` | Human-readable product behavior and scope |
| `docs/architecture/` | Human-readable active technical decisions |
| `docs/research/` | Human-readable evidence and comparisons |
| `docs/future/` | Evaluated features outside the initial release |
| `docs/ai/` | Prescriptive agent workflow, quality gates, roadmap, and tickets |
| `spikes/` | Temporary executable research created only when an approved decision requires it |

`AGENTS.md` contains repository instructions for automated coding agents. Human
contributors do not need it as their primary onboarding document.

## First planned change

The architecture baseline is approved. Before implementation begins, the
repository and continuous-integration foundations are delivered through their
own reviewed tickets. Go server behavior then arrives as the ordered, usable
increments in the
[Go server implementation plan](docs/ai/delivery/GO_SERVER_IMPLEMENTATION.md),
not as a disposable proof of concept.

# Implementation Guide

This is the entry point for implementation. It ties the product, architecture,
research, delivery plan, and verification documents together. No production
implementation begins until the architecture baseline has recorded approval.

Executable research and production work both require an authorized ticket.
Authorization is recorded on the applicable issue and pull request.

## Source-of-truth order

When documents disagree, resolve the conflict before writing code. The order is:

1. `docs/product/OVERVIEW.md` defines the promised product, exclusions, and release gates.
2. `docs/architecture/OVERVIEW.md` contains accepted Architecture Decision Records and
   system-wide constraints.
3. `docs/product/` defines observable feature behavior and acceptance cases.
4. `docs/architecture/` defines boundaries, protocols, data, deployment, and
   failure behavior.
5. `docs/ai/engineering/` defines repository, documentation, security, chaos, and
   testing quality.
6. `docs/research/` supplies evidence and candidates; it is not an accepted
   technology decision by itself.
7. `docs/ai/delivery/ROADMAP.md` is the ordered execution and release checklist.
8. `docs/ai/tickets/` defines the current reviewable work units; a ticket cannot
   weaken a higher-order requirement.

The task list never overrides a specification. A checked task means the
specified behavior and its automated tests are complete, not that a placeholder
exists.

## Required reading by work area

| Work area | Read first | Supporting evidence |
| --- | --- | --- |
| Any change | `docs/architecture/OVERVIEW.md`, `docs/ai/engineering/REPOSITORY_QUALITY.md`, `docs/ai/engineering/DOCUMENTATION_STYLE.md` | Relevant research document |
| Ticket, branch, review, or merge | `docs/ai/engineering/TICKET_BRANCH_WORKFLOW.md` | The ticket and GitHub pull request |
| Executable research | `docs/ai/tickets/README.md` and the active GitHub issue | The research-specific specification |
| Dependency or custom-code decision | `docs/ai/engineering/DEPENDENCY_REGISTER.md`, `docs/ai/engineering/REPOSITORY_QUALITY.md` | Relevant upstream source and focused conformance evidence |
| Chess gameplay, clocks, premoves | `docs/product/CHESS_RULES.md`, `docs/architecture/MINIMAL_GAME_STATE_PROTOCOL.md` | `docs/research/CHESS_TEST_ECOSYSTEM.md` |
| Browser interface | `docs/product/FEATURE_MATRIX.md`, `docs/product/CHESS_RULES.md` | `docs/research/FRONTEND_LIBRARY_STRATEGY.md` |
| Product scope changes | `docs/product/OVERVIEW.md`, `docs/product/FEATURE_MATRIX.md` | Current official product sources listed in the matrix |
| Rooms and identity | `docs/architecture/PRIVACY_AND_IDENTITY.md` | `docs/research/AUTHENTICATION_OPTIONS.md`, `docs/research/PSEUDONYMOUS_IDENTITY_RESEARCH.md` |
| Bot Mode | `docs/product/BOT_MODE.md`, `docs/architecture/SYSTEM_DESIGN.md` | Stockfish references in the bot specification |
| Tournaments | `docs/product/TOURNAMENTS.md` | Pairing and notation references in `docs/research/CHESS_TEST_ECOSYSTEM.md` |
| Data and retention | `docs/architecture/DATA_PLATFORM.md`, `docs/architecture/TELEMETRY_AND_RETENTION.md` | `docs/future/DATABASE_SCALE_OPTIONS.md` |
| Deployment | `docs/architecture/DEPLOYMENT.md`, `docs/architecture/SYSTEM_DESIGN.md` | Operator-input section in the deployment document |
| Security and resilience | Applicable requirements in `docs/ai/engineering/REPOSITORY_QUALITY.md` | Private operator test plan and threat/failure cases in architecture documents |
| Future peer-to-peer work | `docs/future/PEER_TO_PEER_TRANSPORT.md` | The canonical game protocol |

## Change flow

Every implementation unit follows this sequence:

1. Select the next unchecked item in `docs/ai/delivery/ROADMAP.md` and create one reviewable ticket;
   do not work across phases merely because adjacent code is convenient.
2. Link the exact product behavior, acceptance cases, and architecture decision
   in that ticket.
3. Perform the library-first survey required by the repository quality
   contract. Record adoption or custom-code justification before coding.
4. Resolve unexplained terms, contradictions, privacy impact, failure behavior,
   and operator choices in documentation.
5. Obtain recorded approval when the phase or decision has not been signed off.
6. Add or select upstream conformance fixtures and write the failing acceptance
   test.
7. Implement the smallest complete vertical behavior behind narrow adapters.
8. Run unit, integration, browser, security, chaos, and load checks applicable
   to the change.
9. Commit and push coherent work only to the ticket's `test/<ticket>` branch as
   described in `docs/ai/engineering/TICKET_BRANCH_WORKFLOW.md`.
10. Open a GitHub pull request into `dev` and update it in response to review.
11. Merge only after the designated reviewer approves that pull request and all
    required checks pass.
12. Update `docs/ai/delivery/ROADMAP.md` only when the reviewed real behavior is merged.

## Approval gates before first production code

- The release feature set and exclusions have recorded approval.
- The server architecture and low-data identity model have recorded approval.
- The Go server implementation sequence has recorded approval.
- Frontend and backend dependency decisions include license review.
- The Cloudflare Tunnel hostname remains unpublished until the publication gate
  in `docs/product/OVERVIEW.md` passes.

## Definition of done

A feature is done only when it uses real services and real data paths, has no
placeholder or fabricated result, handles documented errors and reconnects,
passes relevant automated acceptance tests, preserves the privacy and retention
contract, and is understandable from these documents without private context.

# Ticket Catalog

Terminology used below: continuous integration (CI), pull request (PR), and
Uniform Resource Locator (URL).

This directory defines approved ticket shapes before they are created as GitHub
issues. `docs/ai/delivery/ROADMAP.md` remains the ordered release checklist; these files define
reviewable units of work.

## Ticket execution rule

1. Work only on the first authorized Ready ticket.
2. Create the GitHub issue from the matching definition without silently
   changing scope or acceptance criteria.
3. Create `test/<ticket-id>-<short-name>` from current `dev`.
4. Put library research in the issue or PR before custom implementation.
5. Commit and push only to the ticket branch.
6. Open a PR into `dev`, link the issue, and wait for designated review.
7. Never begin the next ticket merely because the current branch is waiting for
   review.

Statuses:

- **Defined:** scope is written but work is not authorized.
- **Ready:** prerequisites are met and the ticket has recorded authorization.
- **Active:** one ticket branch is being worked.
- **In review:** PR is awaiting designated review.
- **Done:** approved PR is merged into `dev`.
- **Blocked:** a named prerequisite prevents meaningful work.

Create tickets from the roadmap in dependency order. Each issue uses
`TEMPLATE.md`, links its roadmap item, and remains small enough for one
pull-request review. A ticket cannot replace, weaken, or omit its roadmap
requirement.

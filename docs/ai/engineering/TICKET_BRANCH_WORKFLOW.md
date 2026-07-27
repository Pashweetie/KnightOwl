# Ticket Branch and Pull Request Workflow

Terminology used below: continuous integration (CI) and pull request (PR).

Every feature, fix, dependency update, infrastructure change, and unrelated
cleanup starts with a ticket and its own Git branch. GitHub's PR diff is the
review surface.

## Branch contract

- `test/<ticket-id>-<short-name>` is a temporary working branch for exactly one
  ticket. Commits and pushes are permitted while the work is developed and
  tested. Before review, its history is squashed to exactly one commit relative
  to `dev`.
- `dev` is the protected integration branch. Changes enter it only through a PR
  from the matching ticket branch after all required CI checks pass and the
  designated reviewer approves the GitHub diff.
- `release` is the protected production branch. Changes enter it only through a
  separately reviewed promotion PR from `dev`. Merging it begins production
  deployment immediately.
- GitHub rulesets must block direct pushes, force pushes, branch deletion, and
  bypasses on `dev` and `release`, including administrator bypasses. Until the
  live rules are verified, implementation tickets remain blocked.
- No automated identity may approve or merge its own feature PR.

Approval of one PR does not approve another ticket, a later revision, or a
production promotion. A PR approval authorizes only the merge it names.

## One ticket, one work package

Before implementation work begins, create a ticket that records:

- a short identifier and title;
- the user-visible outcome;
- behavior that is deliberately out of scope;
- linked product requirements and architecture decisions;
- acceptance cases and applicable quality checks;
- the library-first survey or the exact survey work required;
- privacy, security, data-retention, migration, and rollback effects;
- expected files and dependencies, where known.

A ticket must describe one reviewable vertical change. Split it when independent
parts could reasonably be accepted, rejected, or rolled back separately.
Unrelated cleanup receives its own ticket and branch.

## Working and review procedure

1. Create or select the ticket and its acceptance criteria.
2. Branch from current `dev` using the required `test/<ticket-id>-<short-name>`
   form.
3. Perform the library-first survey before custom implementation.
4. Commit coherent checkpoints to the ticket branch using the repository's
   configured human-controlled Git identity; never invent an agent or vendor
   author identity.
5. Push only the ticket branch and open a draft PR into `dev`. Pushing a branch
   or providing a comparison URL does not complete this step.
6. Keep the PR description current with scope, exclusions, library decisions,
   privacy/security effects, tests, risks, migrations, and rollback.
7. Run required CI checks. Failed or skipped checks are visible and explained.
8. Before marking the PR ready, squash the branch to one commit relative to
   `dev`, force-push with lease, and verify that the tree diff is unchanged.
9. Mark the PR ready only when its complete behavior and documentation are
   reviewable without placeholders.
10. The designated reviewer reviews GitHub's changed-files view and either
   requests changes or approves the PR.
11. Any material commit after approval dismisses approval and requires another
    review.
12. Merge only after designated-reviewer approval and all required checks pass.
13. Delete the ticket branch after the merge unless it is needed for a
    documented investigation.

Ticket work is not reported as complete until the pull request exists on
GitHub. If repository authentication cannot create it, report the work as
blocked rather than complete.

Silence, general encouragement, ticket approval, architecture approval, or
approval of another PR is not merge approval.

## Pull request requirements

Every PR must:

- link exactly one primary ticket;
- contain exactly one commit relative to its target branch;
- contain no unrelated file or generated noise;
- state user-visible behavior and deliberate exclusions;
- identify adopted libraries and rejected alternatives;
- list exact test commands and results;
- identify any skipped test with a reason;
- include screenshots or recordings for visual interaction changes;
- include rendered Kubernetes changes for deployment modifications;
- include compatibility, migration, privacy, security, and rollback notes where
  applicable;
- contain no secrets, room capabilities, personal data, fabricated results, or
  agent/vendor branding.

Large features use a sequence of independently usable vertical tickets rather
than one long-lived branch with an unreadable final diff.

## Production promotion

A feature PR into `dev` cannot deploy production. A promotion PR from `dev` to
`release` shows the complete accumulated source and desired Kubernetes change,
identifies immutable image digests, links successful qualification evidence,
and explains active-game and rollback behavior.

The designated reviewer approves that promotion separately. Merging it is the
event that authorizes immediate blue-green deployment.

## Required branch rules

Repository rulesets for both protected branches must require a pull request, one
approval, dismissal of stale approvals, resolution of review conversations,
successful required status checks, linear history, and signed commits. The
rules must block branch deletion, force pushes, and direct updates. `release` accepts
promotion pull requests from `dev` only; a required workflow rejects every
other source branch.

Rulesets use no bypass actor. Emergency changes use the normal pull-request
path. A repository administrator may edit a ruleset in GitHub, so exported
ruleset definitions are versioned and a scheduled audit fails if the live
configuration drifts. No implementation ticket becomes Ready until these live
settings are verified.

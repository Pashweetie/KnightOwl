# Go Backend Vertical-Slice Proof

Terminology used below: central processing unit (CPU), continuous integration
(CI), identifier (ID), personal computer (PC), and pull request (PR). BLAKE3 is
a proper algorithm name; `RESULTS.md` is a filename.

Go is the selected backend language. This proof verifies that the proposed
libraries, package boundaries, tests, debugging, and container deployment
produce code maintainers can understand and safely change. It is not a
decorative mock server or a large-scale benchmark.

The proof receives its own ticket and `test/<ticket>` branch. Commits are
allowed on that branch. It enters `dev` only through a GitHub PR approved by
the designated reviewer.

## Required behavior

The proof must:

1. expose liveness, readiness, startup, and drain endpoints;
2. create an expiring room only through the loopback operator surface;
3. create random host, white, black, and spectator capabilities while storing
   only keyed capability digests;
4. expose an invite-only guest surface with no room-creation handler;
5. claim a seat over a same-origin-validated WebSocket;
6. use the selected maintained chess library to accept legal standard-chess
   moves and reject illegal ones;
7. atomically compare expected version and prior hash in real Valkey;
8. return the original event for a duplicate command identifier;
9. return stale state for a losing concurrent command;
10. reconnect through a different process and read the accepted event;
11. generate byte-identical canonical events and BLAKE3 hashes;
12. send one cancellable request to a real pinned Stockfish process;
13. stop accepting new upgrades while draining;
14. reject malformed and fuzzed transport input without panic or process crash;
15. build a minimal non-root production container with a documented debugging
    procedure.

## Acceptance targets

On the target PC:

- ten concurrent friends can create, claim, play, spectate, and reconnect
  without error;
- two simultaneous commands preserve one authoritative event order;
- invalid, expired, reused, and role-swapped capabilities fail closed without
  leaking room existence;
- the public guest surface cannot reach room creation, administration,
  diagnostics, metrics, or operator routes;
- readiness, drain, reconnect, and Valkey-failure behavior match the
  architecture;
- race detector, fuzz tests, static analysis, unit tests, real-Valkey
  integration tests, and a bounded soak pass;
- logs and test artifacts contain no raw capabilities, room positions, labels,
  or personal network identifiers;
- code remains within the documented domain, transport, storage, and engine
  boundaries;
- the PR explains every direct dependency and proves an established library was
  evaluated before custom code.

CPU, memory, latency, build time, and image size are recorded only as sanity and
operational context. There is no 10,000-socket target, throughput competition,
or Rust comparison. A ten-user correctness failure, unreadable structure, or
unsuitable library boundary reopens the design.

## Review evidence

`RESULTS.md` records:

- ticket and PR;
- exact source commit and image digest;
- direct and transitive dependency inventory with licenses;
- package map and concepts a contributor must learn;
- exact build, test, debug, and container commands;
- test results and replayable failure seeds;
- errors and stack traces from representative failures;
- CPU and memory at idle and at the ten-user scenario;
- reconnect and graceful-drain observations;
- profiler, race, fuzz, vulnerability, and license findings;
- any custom code with its rejected-library evidence;
- the accept, revise, or reopen decision.

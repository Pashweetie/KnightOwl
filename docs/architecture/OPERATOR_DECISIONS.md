# Operator Decisions

Implementation can proceed locally without these answers. Public publication
cannot.

## Required before named-tunnel publication

- Controlled domain and Cloudflare account.
- Named-tunnel token supplied outside Git.
- Desired public hostname.
- Maximum concurrent rooms, spectators, and Stockfish jobs based on load tests.
- Finished-game download grace and maximum active-room duration.
- Whether to run one cloudflared connector or two local connector processes.
- Whether Valkey AOF uses the tested lower-latency or lower-loss policy.
- Short aggregate metric retention.
- A contact method, if any, for friends to report an outage; it must not be
  built into KnightShift unless the operator accepts retaining messages.

## Explicitly not delegated to the operator

The implementation must choose safe defaults for capability size, secure
headers, protocol limits, container privileges, dashboard exposure, logging
redaction, resource isolation, readiness, and failure behavior. These are
engineering duties, not configuration questions.

## Decisions required only if scope expands

Accounts, cross-device synchronization, durable history, ratings,
leaderboards, public rooms, free-text chat, clubs, public tournaments, and
multi-host deployment each require a separate product, privacy, abuse, data,
and operations review before implementation.


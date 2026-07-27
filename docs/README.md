# KnightOwl Documentation

This is the human documentation index. Read:

1. [Product Overview](product/OVERVIEW.md) defines the supported experience,
   release scope, and exclusions.
2. [Architecture Overview](architecture/OVERVIEW.md) records system-wide
   technical decisions and links to detailed designs.
3. [Research](research/) for evidence behind a decision.
4. [Future Features](future/README.md) for evaluated work outside the release.

Implementation agents use the separate [automated implementation
guide](ai/README.md).

## Document authority

The product overview controls release scope. The architecture overview controls
system-wide technical decisions. Detailed product and architecture documents
control their named feature or subsystem. Research records evidence but cannot
override a decision. Agent instructions cannot change human product or
architecture requirements.

## Product

Product documents define observable behavior. They do not select deployment or
storage mechanisms.

- [Overview](product/OVERVIEW.md)
- [Capability matrix](product/FEATURE_MATRIX.md)
- [Chess rules and clocks](product/CHESS_RULES.md)
- [Bot Mode](product/BOT_MODE.md)
- [Tournaments](product/TOURNAMENTS.md)

## Architecture

Architecture documents define system boundaries, protocols, data ownership,
privacy, and deployment.

- [Overview and decision records](architecture/OVERVIEW.md)
- [System design](architecture/SYSTEM_DESIGN.md)
- [Game-state protocol](architecture/MINIMAL_GAME_STATE_PROTOCOL.md)
- [Identity and privacy](architecture/PRIVACY_AND_IDENTITY.md)
- [Data platform](architecture/DATA_PLATFORM.md)
- [Telemetry and retention](architecture/TELEMETRY_AND_RETENTION.md)
- [Deployment](architecture/DEPLOYMENT.md)

## Research

Research documents preserve evidence and evaluated alternatives. They are not
architecture decisions unless adopted by an architecture document.

- [Chess test ecosystem](research/CHESS_TEST_ECOSYSTEM.md)
- [Frontend library strategy](research/FRONTEND_LIBRARY_STRATEGY.md)
- [Authentication options](research/AUTHENTICATION_OPTIONS.md)
- [Pseudonymous identity research](research/PSEUDONYMOUS_IDENTITY_RESEARCH.md)

## Future features

Future documents preserve designs that are explicitly outside the initial
release.

- [Future-feature policy](future/README.md)
- [Peer-to-peer transport](future/PEER_TO_PEER_TRANSPORT.md)
- [Database scale options](future/DATABASE_SCALE_OPTIONS.md)

## Automated implementation material

The roadmap, ticket definitions, dependency register, quality gates, and
change workflow are grouped under [`docs/ai/`](ai/README.md). They are
prescriptive implementation inputs rather than the primary human product
narrative.

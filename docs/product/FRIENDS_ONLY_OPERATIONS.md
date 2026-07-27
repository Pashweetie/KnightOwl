# Friends-Only Social and Operations Model

## Deliberately absent

The initial product has no registration, public profiles, searchable directory,
followers, direct messages, public forums, clubs, public chat, moderation
queue, cheating review, or identity-linked sanctions.

## Room interactions

The host can create and revoke invitations, choose room settings, assign
colors, remove someone before play, disable spectators or reactions, end the
room, and export PGN. Players can choose a temporary room label, ready up,
play, resign, offer or answer a draw, request a rematch, mute reactions,
reconnect, and export PGN.

Communication is a small fixed set of non-customizable reactions. There is no
free text or image upload, which removes the need to retain user-generated
messages or run community moderation. A participant can mute reactions locally
and the host can disable them for the room.

## Operator controls

- No public room list; knowing a valid capability is required.
- Conservative room/player/socket creation limits use short-lived keyed
  buckets without durable source addresses.
- Room state, labels, reaction events, and rate buckets expire.
- Metrics are aggregate counters and latency histograms, never room IDs,
  capability values, labels, IP addresses, or moves.
- Runbooks cover tunnel outage, proxy outage, app crash, state-store outage,
  engine exhaustion, disk pressure, upgrade, rollback, and suspected secret
  exposure.

Persistent identities, cross-device history, ratings, leaderboards, custom
chat, or public discovery require a new explicit product and privacy decision.


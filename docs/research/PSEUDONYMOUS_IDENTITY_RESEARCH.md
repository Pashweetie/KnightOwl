# Pseudonymous Identity Research

Terminology used below: identifier (ID), Internet Protocol (IP), Secure Hash
Algorithm 256-bit (SHA-256), time to live (TTL), and universally unique
identifier (UUID). H2 is the proper name of the database product.

## Research basis

Inspected the official `magefree/mage` repository at commit
`d199ed2732b556435c7a5a2819e622fa9e37411b` (2026-07-26). Relevant sources:

- `Mage.Server/config/config.xml`
- `Mage.Server/src/main/java/mage/server/Session.java`
- `Mage.Server/src/main/java/mage/server/User.java`
- `Mage.Server/src/main/java/mage/server/UserManagerImpl.java`
- `Mage.Server/src/main/java/mage/server/AuthorizedUser.java`
- `Mage.Server/src/main/java/mage/server/AuthorizedUserRepository.java`
- `Mage.Server/src/main/java/mage/server/record/UserStatsRepository.java`

Official repository: <https://github.com/magefree/mage>

## XMage findings

XMage has two server-configured modes rather than one pseudonymous account
system.

### Registration-disabled mode

The default source configuration has `authenticationActivated="false"` and
states that users need not register. `Session.REGISTRATION_DISABLED_MESSAGE`
tells clients they may use any name and an empty password.

On connection:

- the name must satisfy configured length/character rules;
- the server creates a random in-memory UUID for the active user;
- the connection has a session ID and a restore-session ID;
- only one active user with a given name is normally allowed;
- anonymous duplicate/reconnect handling compares the connection host;
- active user/session state is held in concurrent in-memory maps;
- disconnected sessions receive a short reconnection window and are later
  removed.

This is frictionless pseudonymity, not cryptographic authentication. A name is
effectively first-come-first-served while active. A changed network address can
prevent reclaiming it unless the restore-session ID works.

### Registration-enabled mode

When authentication is enabled, XMage registers:

- unique username;
- unique email;
- salted password verifier;
- hash algorithm and iteration count;
- active/locked/chat-locked state;
- last-connection timestamp.

The current source uses Apache Shiro SHA-256 hashing with a random salt and
1,024 iterations, persisted through ORMLite in an H2 database. Registration
generates a password and sends it to the supplied email. This is a conventional
first-party credential/contact database and does not meet KnightOwl's
requirements.

### Persistent statistics

Independent of the login impression, XMage's result repositories build durable
statistics keyed by player name:

- match and tournament counts;
- quit/idle/timeout history;
- general, constructed, and limited Glicko ratings;
- table/match/tournament result records.

`UserStatsRepository` stores these in SQLite. Therefore registration-disabled
XMage should not be described as “no user data”; a recurring chosen name can
become a durable pseudonymous activity record.

### Operational exposure

The active `User` contains host, name, UUID, session IDs, activity timestamps,
client version, tables/games, and optional user data. Administrative user views
include host and session information, and connection logs include the chosen
name. That may be reasonable for a public game server, but it is more
operational identity data than KnightOwl intends to retain.

## General design findings

- Joining should require no registration ceremony.
- A short room-local display name is enough for human recognition.
- A reconnect credential should be separate from the display name.
- Active sessions need explicit connected/disconnected/offline transitions and
  finite grace periods.
- Names need normalization, safe characters, bounded length, and collision
  handling.
- Server configuration should make the identity/retention mode obvious.

## Rejected patterns

- Do not use source host/IP as anonymous identity or reconnect authority.
- Do not let a display name itself grant an existing seat.
- Do not store passwords or email.
- Do not silently turn repeated names into durable ratings/history.
- Do not expose session IDs, addresses, or internal user IDs in operator views
  unless strictly required and redacted.
- Do not log room labels as routine connection identifiers.

## Independent KnightOwl model

KnightOwl uses an independently specified capability model:

1. An invite capability grants permission to claim a room role.
2. Claiming creates a random internal participant ID and rotating reconnect
   capability.
3. A display label is cosmetic, room-local, non-unique, and never authorizes
   anything.
4. The app does not persist host/IP, label, participant identity, game record,
   or rating after the room TTL.
5. The browser may keep PGNs and local statistics after clear disclosure.

XMage is one research input, not a protocol dependency or implementation
template. The resulting model avoids name squatting as an authentication
mechanism, survives ordinary IP changes, and does not create a hidden global
profile.

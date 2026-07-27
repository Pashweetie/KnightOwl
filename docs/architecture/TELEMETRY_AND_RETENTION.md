# Telemetry and Retention

Terminology used below: continuous integration (CI), central processing unit
(CPU), Forsyth–Edwards Notation (FEN), Hypertext Transfer Protocol (HTTP),
identifier (ID), Internet Protocol (IP), Portable Game Notation (PGN), and
uniform resource locator (URL). “US” means United States.

## Goal

Collect enough operational evidence to run, secure, debug, and scale the
friends server without building advertising, identity, or behavioral profiles.
This reduces privacy and security exposure; it does not guarantee that no law
applies. Deployment jurisdiction, audience, operator status, and actual data
flows still matter and require appropriate review.

## Three telemetry classes

### Class A: anonymous aggregate metrics

Enabled by default:

- current HTTP/WebSocket connections;
- room/game counts by coarse state and time-control category;
- request/command latency histograms;
- result and error counts by bounded machine code;
- reconnect and transport-success counts;
- Valkey/worker/proxy CPU, memory, disk, queue, and saturation;
- Stockfish request duration grouped by bounded configuration class;
- build/protocol version adoption counts.

These metrics contain no room ID, capability, participant ID, display label,
source address, user-agent, move, position, PGN, URL, or arbitrary error text.
Low-volume dimensions are suppressed or coarsened to reduce singling out.

### Class B: short-lived security/diagnostic events

Enabled with strict retention:

- rotating keyed source-network bucket for rate enforcement;
- random request/connection correlation ID;
- endpoint and bounded outcome code;
- coarse timestamp and duration;
- app replica/build/protocol version;
- resource-limit or validation category.

The rotating key is destroyed on schedule so buckets cannot be joined across
periods. Events omit URL query/fragment, cookies, headers, message bodies,
display labels, room IDs, positions, moves, and raw exceptions. Default
retention is 24 hours, configurable downward. Access is local-operator only.

### Class C: gameplay/product research

Disabled by default and not required for operation:

- move/position sequences;
- persistent browser/device pseudonyms;
- cross-game cohorts;
- drill/puzzle behavior;
- detailed feature funnels;
- long-lived address or fingerprint data.

Adding any Class C field requires an explicit purpose, necessity analysis,
user-facing disclosure/choice where appropriate, retention/deletion plan,
access controls, threat model, and jurisdiction-specific review. Advertising,
sale/sharing, third-party trackers, fingerprinting, precise location, contacts,
and sensitive-category inference are prohibited.

## Data inventory

Every collected field has a machine-readable catalog entry:

- name and type;
- class;
- exact purpose and consumer;
- source and destinations;
- identifying/linkability analysis;
- retention and deletion mechanism;
- access role;
- security classification;
- configuration default;
- test proving redaction and expiry.

CI fails when telemetry code emits an unregistered field or label. Production
startup fails if configured retention exceeds the compiled maximum without an
explicit operator override.

## Request logs

Routine reverse-proxy access logs are off. Diagnostic access logging uses a
fixed allowlist and bounded codes, never an unchecked standard format.
Application logging is structured and schema-validated. Arbitrary values are
kept in local debug memory only when deliberately enabled for a short incident,
then purged.

## User-facing truth

The site has a concise “Data on this server” view generated from active
configuration. It states what the server currently collects, why, retention,
whether optional telemetry is enabled, what remains only in the browser, how to
clear local data, and which infrastructure providers process traffic.

It never claims “anonymous” for keyed or linkable pseudonymous records.

## Verification

- Submit canary IP, user-agent, label, room, capability, FEN, move, PGN, cookie,
  query, and error values and scan every log/metric/trace/crash sink.
- Assert Class A labels remain within a cardinality budget.
- Advance time and prove Class B records and rotating keys disappear.
- Attempt correlation across key rotations.
- Inspect metrics/diagnostic endpoints for accidental raw values.
- Verify the public data view matches runtime configuration.

## Legal-risk boundary

The design follows data minimization and finite retention. Official guidance
still treats IP addresses, cookie IDs, and re-identifiable pseudonyms as
personal data in some regimes. Pseudonymisation reduces risk but is not the
same as irreversible anonymisation.

Primary guidance:

- European Commission:
  <https://commission.europa.eu/law/law-topic/data-protection/information-business-and-organisations/application-regulation/application-gdpr_en>
- US Federal Trade Commission:
  <https://www.ftc.gov/business-guidance/resources/protecting-personal-information-guide-business>
- California Attorney General:
  <https://oag.ca.gov/privacy/ccpa>

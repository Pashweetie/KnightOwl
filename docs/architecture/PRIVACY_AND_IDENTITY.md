# Privacy and Room Identity

## No account system

KnightShift stores no username, password, password hash, email address, phone
number, OAuth/OIDC subject, recovery secret, public profile, durable player ID,
friend graph, or server-side game history.

Third-party authentication is not integrated because the current product has
no cross-device identity feature that needs it. Adding authentication later
would not remove the operator's privacy responsibilities and requires a new
explicit decision.

## Capabilities

A random URL capability grants only one room role: host, white, black, or
spectator. Tokens contain at least 128 bits of cryptographic randomness. The
server stores only keyed hashes, compares in constant time, never logs tokens,
and expires them with the room.

Capabilities are bearer secrets. The UI warns participants not to repost them.
The host can revoke unused invites and spectator access. Referrer policy,
same-origin routing, CSP, and redacted error handling prevent common accidental
leaks.

Temporary player labels are optional, room-local, length/character restricted,
and deleted with the room. They never become searchable identities.

## Network metadata

Network infrastructure necessarily processes source addresses to deliver
traffic. The application may retain only the documented short-lived rotating
security bucket and bounded diagnostic events.

- Proxy access logs are disabled by default or use the registered allowlisted
  diagnostic schema.
- Short-lived abuse buckets use a rotating keyed digest and expire quickly.
- Application logs omit raw address, user-agent, room, capability, label,
  FEN, PGN, and move payloads.
- Metrics use bounded aggregate operational labels only.
- Crash reports and traces are local and scrubbed before retention.

Cloudflare remains an external processor of traffic; using a tunnel does not
make legal/privacy obligations disappear. Operator documentation must describe
the actual deployment rather than promise legal immunity.

## Browser-local data

PGNs, analysis, lesson progress, puzzle schedules, and preferences may be stored
in IndexedDB only after an explicit local-storage explanation. They remain on
that browser, can be exported, and have a one-action clear function. The app
does not upload them for synchronization.

## Social-data minimization

Communication is limited to fixed reactions. No free text, images, voice,
direct messages, forums, profiles, or public discovery are accepted. Muting is
local; the host can disable reactions. This is the primary moderation design.

## Verification

- Automated log tests submit unique canary values for every forbidden field and
  assert none occur in logs, metrics, traces, or error pages.
- TTL tests prove room records and capability hashes disappear.
- Browser tests prove local export and deletion.
- Security tests try capability guessing, leakage through referrers/errors,
  privilege swapping, replay, and revoked-link use.
- The deployment audit proves internal dashboards and data ports are not
  externally reachable.

# Authentication Options Research

Terminology used below: Content Security Policy (CSP),
cryptographically secure pseudorandom number generator (CSPRNG),
hash-based message authentication code (HMAC), Hypertext Transfer Protocol
(HTTP), Hypertext Transfer Protocol Secure (HTTPS), identifier (ID), OpenID
Connect (OIDC), operating system (OS), one-time password (OTP), personally
identifiable information (PII), Proof Key for Code Exchange (PKCE), time to live
(TTL), uniform resource locator (URL), user experience (UX), and World Wide Web
Consortium (W3C).

## Authentication is not always required

KnightOwl has three separate questions:

1. **May this browser enter a room?** An invitation capability answers this.
2. **May this connection act as host/white/black/spectator?** A role capability
   answers this.
3. **Is this the same person who visited last month?** Only persistent identity
   authentication answers this.

The current friends-only product needs the first two and deliberately does not
need the third. Calling a room capability “anonymous authentication” obscures
the design: it is bearer authorization scoped to one expiring room.

## Comparison

| Method | Server-retained data | Advantages | Costs and privacy limits | Suitable use |
|---|---|---|---|---|
| Expiring room capability | Keyed token hash, role, expiry | No person identity; simplest; revocable per room | Link theft grants its role; no cross-device identity or recovery | Required default |
| Device key pair | Random device public key and local private key | No email/provider; stronger than a bearer cookie | Stable server-correlatable device; loss means loss; browser storage can be cleared | Optional remembered device |
| WebAuthn/passkey | Random user handle, credential ID, public key, counters/metadata | Phishing-resistant proof; server never receives private key/biometric; synced passkeys may work across devices | Creates a durable pseudonymous account; provider/platform may sync it; recovery and deletion must be designed | Best future persistent option |
| Pairwise-subject OIDC | Issuer plus pairwise subject or a keyed derivative, session records | Outsources login/recovery; provider can support many devices | Provider knows login activity; stable KnightOwl identity; tokens/claims must be processed; pairwise support varies | Optional future identity |
| Email magic link/OTP | Email plus challenge and delivery metadata | Familiar recovery | Direct contact data and email-delivery operations; explicitly conflicts with current constraint | Rejected |
| Social login with public subject | Stable provider subject and usually profile claims | Easy onboarding | Correlation, provider dependency, claim minimization mistakes | Rejected |
| Crypto wallet signature | Public address and challenge records | No email or password | Public ledger correlation, wallet UX, phishing/loss, financial-identity implications | Rejected |
| Privacy Pass token | Issuer/key version, redeemed-token protection as required | Proves an anonymous property/rate entitlement without identity | Not a durable account, room membership, or recovery mechanism | Possible abuse-control research |
| Cloudflare Access | Identity assertion/headers and Cloudflare policy state | Excellent for an operator allowlist | Common policies use email/IdP identity; places the entire site behind third-party identity | Optional private deployment gate, not in-app identity |

## Default capability protocol

### Token construction

- Generate at least 128 random bits with the operating-system CSPRNG.
- Create distinct host, white, black, and spectator capabilities.
- Put the opaque token in the URL fragment for the landing page so it is not
  sent in the initial HTTP request or Referer.
- Browser JavaScript exchanges it once over same-origin HTTPS for an HttpOnly,
  Secure, SameSite session cookie scoped to the room.
- Store only `HMAC(server-secret, token)` in Valkey, with room role, generation,
  use policy, and TTL.
- Remove the fragment from browser history immediately.

The fragment approach is not magic secrecy: browser extensions, screenshots,
clipboard history, or malicious same-origin code can still expose it. CSP,
dependency control, no third-party scripts, redacted telemetry, and revocation
remain necessary.

### Lifecycle

- Room and capability creation occurs only through the loopback operator
  surface; the public guest surface can claim, reconnect, and exercise a role
  but cannot create rooms or capabilities.
- On the loopback operator surface, the host capability creates and revokes
  invites or ends the room. On the public guest surface, it permits only the
  approved in-room host controls and cannot mint another capability.
- Player invites default to one successful claim, which returns a separate
  reconnect session; forwarding the original link then does not steal a seat.
- Reconnect sessions rotate after use and expire with the game.
- Spectator invites can be multi-use with a configured cap.
- No capability works outside its room or role.
- A server-secret rotation invalidates all capabilities unless a short
  overlapping key generation is deliberately configured.

## Future passkey profile

If cross-device local history, stable ratings, or durable tournament identity
is later approved, passkeys are the lowest-data first-party option.

KnightOwl would store:

- a random, non-PII user handle;
- one or more credential IDs and public keys;
- signature counter where meaningful;
- creation/last-use security timestamps;
- explicitly opted-in product data.

It would not request attestation identifying authenticator models, and it would
use discoverable credentials so a username lookup is unnecessary. User
verification is required. Display names remain room-local and are not account
identifiers.

Tradeoffs that cannot be designed away:

- The credential ID and public key are stable identifiers at this relying
  party, even though they are not useful across relying parties.
- A passkey may be synchronized by an OS/vendor account; KnightOwl does not
  receive that vendor identity, but the vendor participates in availability.
- Without email, provider identity, recovery codes, or a second passkey, losing
  all authenticators means irreversible account loss.
- Recovery codes are themselves credentials and must be stored hashed.
- Account deletion and inactive retention still require explicit behavior.

The WebAuthn specification requires the user handle to avoid identifying
information and recommends a random value. It also notes that credential IDs
and public keys are necessarily visible to the relying party but scoped to it.

## Pairwise OIDC fallback

If users prefer provider recovery, require authorization code flow with PKCE,
exact redirect URIs, issuer/audience/nonce/state checks, and a provider that
supports pairwise subject identifiers. Request only `openid`; do not request
email, profile, contacts, or offline access. Discard ID/access tokens after
validation and derive the internal key with a keyed domain-separated function.

Pairwise subjects reduce correlation between different relying parties. They
do not make the user anonymous to the identity provider, and they still create
a durable pseudonymous record at KnightOwl.

## Adopted result

Ship capability rooms only. Design the protocol so a future authenticated
principal can own a capability, but do not build passkeys or OIDC until a
requested feature genuinely requires continuity across rooms. If that happens,
offer passkeys first and pairwise OIDC as an explicit recovery/convenience
tradeoff.

## Primary references

- W3C Web Authentication Level 3:
  <https://www.w3.org/TR/webauthn-3/>
- OpenID Connect Core, pairwise identifiers:
  <https://openid.net/specs/openid-connect-core-1_0.html#PairwiseAlg>
- Privacy Pass HTTP Authentication Scheme:
  <https://www.rfc-editor.org/rfc/rfc9577.html>

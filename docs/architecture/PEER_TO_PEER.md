# Peer-to-Peer Game Transport Evaluation

## Status

WebRTC is planned as a second transport, not adopted as the sole multiplayer
transport. The server-authoritative Valkey protocol is implemented, qualified,
and released first. A later vertical spike adds WebRTC only if it preserves the
same command, hash-chain, clock, reconnect, and fallback semantics.

## Feasible hybrid

```text
browser A ---- signaling ---- browser B
    \                            /
     +---- WebRTC data channel -+
                 |
          TURN when direct ICE fails
```

The existing app can provide expiring offer/answer/ICE signaling without
retaining gameplay. An ordered reliable WebRTC data channel transports chess
commands. DTLS protects the channel. Both peers retain the canonical event
chain and PGN locally.

## What becomes easier

- Direct-path move latency and server bandwidth.
- Gameplay can continue through a short signaling outage after connection.
- The server need not retain active move history for a successful direct game.
- Browser Stockfish/WASM can make bot games and analysis entirely local on
  capable hardware.

## What becomes harder

### Connection establishment

WebRTC does not eliminate infrastructure. It needs an out-of-band signaling
exchange plus ICE. Some NAT/firewall pairs cannot establish a direct route and
require TURN. Hosting TURN on the same home network generally needs public UDP
reachability, so avoiding port forwarding means using an external TURN service
or accepting connection failures.

### Privacy

Direct ICE can expose public/network addresses to the other peer. Relay-only
ICE hides that path but sends traffic through TURN and consumes relay bandwidth.
The UI must present this as a real choice rather than claiming P2P is inherently
more private.

### Authority and clocks

For casual friends, one peer can be elected host-authority, but then:

- a malicious or broken host can decide clocks/results;
- host sleep/background throttling affects deadlines;
- host departure requires authority transfer;
- simultaneous commands need a deterministic term/version protocol.

A symmetric two-peer protocol cannot guarantee progress when peers disagree.
Consensus between two members has no majority after a partition. A small
server witness restores arbitration, at which point the architecture is hybrid.

### Reconnect and spectators

A closed WebRTC transport cannot switch underneath an existing SCTP
association. Reconnect performs ICE/signaling again and restores from the
shared chain. Spectators either create a host-centered fan-out, a peer mesh, or
use the server; all scale worse operationally than one server broadcast.

### Tournaments and premium-equivalent services

Pairings, synchronized puzzle battles, standings, room membership, and
multi-party events still require a coordinator. P2P does not materially simplify
these services. Durable ratings/history remain incompatible with the current
no-identity/no-retention scope regardless of transport.

## Adoption spike

A P2P spike is worthwhile only after server multiplayer works:

1. use the exact production command/event encoding over `RTCDataChannel`;
2. connect across representative home, mobile, corporate, and CGNAT networks;
3. measure direct versus TURN success and setup time;
4. test address exposure and relay-only mode;
5. background/sleep both browsers during live clocks;
6. disconnect and migrate the elected host;
7. add two, ten, and fifty spectators;
8. fall back to server transport without changing game identity or duplicating
   a command.

Adopt it when direct success is high, fallback is seamless, privacy is clearly
communicated, and its maintenance cost is justified by measured server
bandwidth/latency. Do not require it for a user to play.

## Primary references

- W3C WebRTC Recommendation:
  <https://www.w3.org/TR/webrtc/>
- WebRTC data channels:
  <https://www.rfc-editor.org/rfc/rfc8831.html>
- TURN:
  <https://www.rfc-editor.org/rfc/rfc8656.html>
- WebRTC security architecture:
  <https://www.rfc-editor.org/rfc/rfc8827.html>

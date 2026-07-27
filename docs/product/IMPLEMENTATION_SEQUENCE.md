# Implementation Sequence

The ordered source of truth is `TASKS.md`. The dependency chain is:

1. **Architecture closure**
   - Remove contradictory assumptions and finish the Rust/Go spike.
2. **Delivery foundation**
   - Reproducible repository, CI, hardened images, Valkey, Traefik, and private
     local diagnostics.
3. **Chess core**
   - Rules, clocks, canonical encoding, event hashing, and exhaustive tests.
4. **Private multiplayer**
   - Capability rooms, atomic state, reconnect, spectating, and PGN.
5. **Complete interface**
   - Responsive accessible board, local history, and cross-browser verification.
6. **Bot coach and analysis**
   - Isolated Stockfish, truthful MultiPV, educational minimax, and grounded
     explanations.
7. **Learning**
   - Original lessons, licensed/original puzzles, and browser-local progress.
8. **Private competition**
   - Ephemeral arena and Swiss tournaments.
9. **Qualification and publication**
   - Security, chaos, load, privacy, accessibility, clean-install, and named
     Cloudflare Tunnel gates.

No later phase may compensate for a failing earlier invariant. Vertical slices
may be used during a phase, but unfinished controls remain absent from the UI.


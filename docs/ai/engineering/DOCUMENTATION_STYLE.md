# Documentation Style

## Abbreviations

Every document explains an abbreviation at its first use in that document:

```text
Portable Game Notation (PGN)
```

Later uses in the same document may use `PGN`. A link to a glossary does not
replace this requirement because documents must remain understandable alone.

Names that are not abbreviations, such as Valkey, Stockfish, Traefik, Go, Rust,
Docker, BLAKE3, and WebSocket, do not need artificial expansions. Initialisms,
protocol abbreviations, organizations, file formats, test-tool names derived
from abbreviations, and operational shorthand do.

The documentation check maintains an allowlist only for ordinary units and
universally written product names. New unexplained all-capital tokens fail the
check until expanded or deliberately reviewed.

## Plain language

- Prefer a concrete term over internal shorthand.
- Define a specialized chess or infrastructure term before relying on it.
- State whether a requirement is implemented, planned, deferred, or excluded.
- Do not use “anonymous” for merely pseudonymous or short-lived data.
- Do not claim standards compliance, endorsement, security, privacy, or
  completion without the evidence that proves it.

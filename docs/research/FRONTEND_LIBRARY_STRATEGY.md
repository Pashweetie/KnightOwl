# Frontend Library Strategy

Terminology used below: Cascading Style Sheets (CSS). GNU is the proper name of
the GNU Project, and ZIP is the proper name of the archive format rather than
an abbreviation. Moving Picture Experts Group 4 (MPEG-4) is the media
container standard used below. “US” appears only inside a documentation uniform
resource locator.

Status: research shortlist; no implementation is approved.

## Rule

KnightOwl does not invent frontend display protocols or generic interface
components. A small executable proof must compare shortlisted libraries against
real chess workflows before implementation. Only product-specific composition,
themes, server integration, and documented library gaps may be custom.

## Responsibility map

| Responsibility | Library-first candidates | Required decision |
| --- | --- | --- |
| Chess board interaction | Chessground or react-chessboard | Choose the license policy first, then verify keyboard behavior, numbered premoves, promotion, touch, and stale-state updates |
| Browser chess rules | chess.js or chessops | Select after backend differential and license tests |
| Dialogs, menus, popovers, tabs, tooltips, switches, and focus | React Aria Components or Radix Primitives | Select one primary family after a representative proof |
| Forms and validation | React Hook Form plus the selected schema library | Avoid duplicate validation models |
| Request and response state | TanStack Query where caching is useful | Keep live moves in the game protocol adapter |
| Local interface state | Framework state first; Zustand only if measured complexity warrants it | Do not add a global store by default |
| Coach evaluation and minimax charts | Apache ECharts | Verify tree layout, responsive rendering, structured fallback, and bundle size |
| Large lists | TanStack Virtual | Adopt only at a measured rendering threshold |
| Component workshop | Storybook | Represent meaningful states as executable stories |
| Browser and visual regression tests | Playwright | Use deterministic local screenshot baselines |
| Custom cursors | Browser Cascading Style Sheets `cursor` property plus reviewed Creative Commons Zero assets | Use the platform standard; do not add a cursor-rendering framework |
| Meme and interface animation | Motion plus canvas-confetti; dotLottie Web for licensed character sequences | Use event-driven adapters and do not create a rendering engine |
| Chess-piece themes | Selected board library's piece-renderer interface plus reviewed existing sets | Bundle several proven sets; do not build a second board renderer |

## Board comparison

| Candidate | License | Existing capability | Material gap or risk |
| --- | --- | --- | --- |
| Chessground | GNU General Public License version 3 or later | Click, drag, touch, premoves, arrows, animation, orientation, drawing, responsive board, variants, and a small dependency-free runtime | Requires the distributed combined browser work to use a compatible license; framework adapter and KnightOwl keyboard proof still required |
| react-chessboard | Massachusetts Institute of Technology License | Current React and TypeScript component with drag, touch, responsive sizing, custom pieces, animation, event hooks, and stated accessibility support | Does not document a premove queue as a built-in feature; upstream roadmap still calls for accessibility and test-suite improvements |
| chessboard-element | Massachusetts Institute of Technology License | Framework-independent Web Component, drag, responsive layout, custom rendering, and small bundle | Release cadence and community are weak, and it does not provide premoves |
| chessboard.js | Massachusetts Institute of Technology License | Mature simple board, drag, positions, spare pieces, and documented examples | Old architecture and jQuery dependency; no touch or premove claim strong enough for this product |

Chessground has the closest capability fit and avoids recreating board
interaction behavior, especially premoves. Its license is the only unresolved
product decision. If KnightOwl is approved under a compatible free-software
license, Chessground is the recommended board. If a permissive KnightOwl
license is required, react-chessboard becomes the preferred base, and the
premove queue is a documented KnightOwl-specific adapter because the
shortlisted permissive boards do not supply it.

No candidate contains chess legality. The selected browser chess-rules library
and the authoritative server remain responsible for legal moves.

## Chess-piece skins

Piece appearance is independent from board colors, cursor theme, and meme pack.
The settings panel shows all twelve pieces on both light and dark sample
squares, previews a real middlegame position, and applies the selected set
immediately without changing game state.

Traditional piece sources remain available, but the feature is primarily a
skin system. Each skin supplies distinct artwork for all six white and six
black pieces at every shipped resolution. It may also supply documented
selection, legal-target, capture, promotion, check, and victory accents. A
global color filter over one traditional set does not qualify as a separate
skin.

The release skin catalog is:

| Choice | Source | License and status |
| --- | --- | --- |
| Classic | Colin M. L. Burnett vector pieces | GNU General Public License version 2 or later; confirmed and compatible if KnightOwl adopts Chessground's required license |
| Rounded | Chessnut piece graphics | Apache License version 2.0; confirmed |
| Diamond Cut | Original faceted crystal pieces with distinct light and dark gem palettes, readable silhouettes, and a bounded sparkle accent | Original KnightOwl asset work; required |
| Magma Forge | Original volcanic-stone pieces with glowing fissures, distinct silhouettes, and restrained ember effects | Original KnightOwl asset work; required |
| Neon Arcade | Original emissive arcade pieces with distinct geometry and restrained glow | Original KnightOwl asset work; required |
| Pixel Quest | Existing licensed pack if one passes provenance review, otherwise original pixel artwork | Required; source decision remains open |
| Plush Court | Original rounded, cute fabric-like pieces | Required |
| High Contrast | Strong silhouettes and outlines based on a confirmed source set | Required functional skin |

No unfinished skin appears in the selector. Each set records author, source
page, license text, file digests, original vector or bitmap sources, generated
sizes, and the reproducible generation command.

Diamond Cut is not a single gem icon reused six times. The king has a crown
profile, queen a tiara profile, bishop a cut-mitre profile, knight a faceted
horse profile, rook a crystal-tower profile, and pawn a small jewel profile.
White and black remain identifiable without relying solely on hue. Sparkles use
Motion and stop when the browser requests reduced motion; the pieces themselves
remain fully visible without animation.

Magma Forge uses sculpted volcanic forms rather than applying an orange filter
to Classic. Every role has a distinct stone silhouette and fissure pattern.
Ember and heat-shimmer accents are bounded, stop under reduced motion, and
cannot obscure squares, legal targets, check state, or piece color.

The board-library adapter supplies pieces through its documented custom-piece
or styling interface. It does not replace drag, click, animation, orientation,
or square rendering. Failed or missing assets atomically fall back to Classic
instead of leaving invisible pieces. Themes are cached locally, but the
selected name and provenance remain visible in settings.

Tests cover all pieces, both colors, capture animation, promotion choices,
orientation, high pixel density, zoom, phone layout, every bundled board color,
missing assets, theme switching during a game, and screenshot regressions in
supported browsers.

The skin preview reports approximate download and decoded-memory size. Heavy
skins are still bundled locally and lazy-loaded; they never fetch from a theme
marketplace during a game. A skin performance budget prevents sparkle, glow,
or high-resolution textures from delaying pointer movement or live move
rendering.

React Aria Components and Radix Primitives provide established interaction and
focus behavior for generic controls. KnightOwl will theme and compose the
selected family, not copy its implementation into locally owned components.

Apache ECharts provides established scalable vector graphics and canvas chart
rendering, including tree charts. The coach also exposes the underlying values
in a structured view even when full screen-reader certification is out of
scope.

## Proof before selection

The board proof must demonstrate real click, drag, touch, keyboard alternative,
promotion, orientation, legal targets, last move, check, arrows, resize, visible
numbered premoves, cancellation, server rejection, stale sequence, reconnect,
and opponent updates. It must not duplicate move-legality logic.

The component proof covers dialog, menu, select, tooltip, tabs, toast, and
confirmation flows with mouse, touch, and keyboard. The chart proof renders
real engine candidates and the bounded minimax tree without fabricated data.
All proofs run in current Chrome, Firefox, Safari, and Edge.

## Permitted custom code

Custom frontend code is limited to:

- adapting versioned KnightOwl game events to selected libraries;
- chess-specific behavior that the chosen board cannot express;
- KnightOwl styling and responsive page composition;
- a documented compatibility or security correction pending an upstream fix.

Each exception records the missing capability or upstream issue and has focused
tests. Generic widgets and rendering engines are not exceptions.

## Cursor themes

Cursor customization uses the browser's standard `cursor` property, image
fallback list, explicit hotspot, and mandatory system keyword fallback. This is
a mature platform facility, so adding a JavaScript cursor renderer would
violate the library-first rule rather than satisfy it.

The initial asset candidate is Kenney's Creative Commons Zero cursor pack,
which provides bitmap and vector sources. Selection still requires an asset
inventory and provenance record. Themes should include cute chess-oriented
choices, a preview grid, separate normal and actionable pointer states, and a
one-click System Default choice. Preferences remain in browser-local storage.

Cursor images target 32 by 32 pixels because Chromium and Firefox may reject
oversized images. The feature must retain meaningful system shapes for text,
resize, forbidden, wait, grab, and grabbing states; visual personality cannot
hide interaction meaning. Custom cursors are desktop-only where necessary
because touch devices and iPad operating system pointer behavior differ.

## Meme reaction effects

The animation layer subscribes to confirmed domain events; it never infers
check, checkmate, result, or promotion from board pixels. Motion handles
interface transitions, canvas-confetti handles bounded particles, and the
official dotLottie Web player handles licensed or original reaction sequences.

The desired tone is deliberately silly, recognizable Reddit-style reaction
humor. KnightOwl ships one custom pack made from original characters, audio,
and animation. It may also include public-domain or explicitly licensed assets
whose provenance is recorded.

The original KnightOwl meme pack is selected by default on a new browser,
with Full intensity unless the browser requests reduced motion. Reduced motion
selects Off on first use. The player can change pack, intensity, event toggles,
and sound independently at any time; importing a pack does not silently make it
active.

The supplied reference for the default checkmate reaction is the Terry Crews
Old Spice “mind blown” reaction posted to Reddit in 2012:
https://www.reddit.com/r/gifs/comments/zca0r/im_tired_of_that_same_old_mind_blown_gif_from_tim/

The bundled original pack adapts its comic timing into a chess sequence: a
triumphant king reacts, then its crown and head burst into obviously cartoon
chess pieces and confetti while the persistent Checkmate result stays visible.
The exact advertising clip may be selected from a person's browser-local pack
but is not copied into the repository. The Reddit post's original image host is
no longer reliably fetchable, so the post cannot be a runtime media dependency.

Each person can import additional animation packs into their own browser.
Imports never upload to KnightOwl, synchronize through a room, enter server
logs, or become visible to another participant automatically. A user may
therefore privately map their own local media to events, while the repository
and release images distribute only the original pack.

The reaction settings expose an explicit event-to-animation mapping. Checkmate
has its own selector and may use any compatible bundled or imported GIF, WebP,
video, or dotLottie animation independently from check, win, loss, draw,
promotion, puzzle, bot, and tournament events. The selector previews the exact
media, sound, crop mode, duration, loop count, and reduced-motion fallback
before activation. A reset restores the bundled default without deleting the
imported pack.

KnightOwl does not call, scrape, proxy, or embed Giphy, Tenor, or another
reaction service at runtime. This avoids undocumented integrations, changing
service terms, tracking exposure, and uncertain rights.

The pack screen does provide an assisted Tenor workflow: Open Tenor Search in a
new browser tab, download the chosen standard-definition, high-definition, or
MPEG-4 file through Tenor's own interface, return to KnightOwl, and import
that local file. The importer can record the pasted Tenor share page as source
attribution, but it does not fetch or parse that page. This keeps Tenor content
accessible even though Google fully decommissioned the Tenor application
programming interface on 30 June 2026.

The checkmate timing reference is:
https://tenor.com/view/mind-blown-explosion-overwhelmed-woah-wow-gif-4820143

Tenor describes it as an 18.3-second, 480 by 360 mind-blown reaction. The
original Graphics Interchange Format download is about 29.4 megabytes, so the
MPEG-4 download is preferred for local playback. Import limits must admit this
reference after decoding checks while still rejecting decompression bombs.

### Pack format and isolation

The import container is a standard ZIP archive read with an established ZIP
library. It contains a versioned JavaScript Object Notation manifest plus local
media. The manifest records pack name, author, optional license/source,
content digests, media dimensions, duration, event mapping, sound loudness, and
fallback behavior. It does not contain executable code.

Allowed media are bounded WebP, Graphics Interchange Format, animated Portable
Network Graphics, MPEG-4, WebM, and Ogg or Waveform Audio File Format audio.
dotLottie is allowed only after its archive and external-reference validator
passes the security proof. Scalable Vector Graphics, Hypertext Markup Language,
scripts, remote uniform resource locators, nested archives, absolute paths, and
parent-directory paths are rejected. Archive entry count, compressed size,
expanded size, compression ratio, decoded dimensions, frame duration, and
audio duration are bounded before storage.

The importer uses a maintained ZIP reader and schema validator, verifies every
digest, decodes media in an isolated worker where supported, and stores the
approved pack in browser Indexed Database storage. Playback uses local object
URLs with a restrictive Content Security Policy. Removing a pack revokes its
object URLs and deletes its browser records. Export reproduces the same
validated standard ZIP format.

Settings expose Off, Subtle, and Full intensity, pack selection, preview, mute,
and per-event toggles. A reduced-motion browser preference defaults effects to
Off, and the confetti adapter always enables reduced-motion suppression. An
effect is a non-interactive overlay: it cannot capture input, obscure the result
or controls, alter clocks, or delay an acknowledgement. Duration, decoded
dimensions, memory, audio level, and particle count have strict bounds. A new
state, reconnect, navigation, Undo, or Redo cancels obsolete media.

Tests drive real confirmed events and verify mapping, cancellation, archive
traversal defense, decompression-bomb limits, malformed media, external
reference rejection, digest failures, storage clearing, zero network requests,
input pass-through, mute, intensity, and reduced-motion behavior. They wait on
library completion events, never arbitrary delays.

## Primary references

- Chessground: https://github.com/lichess-org/chessground
- react-chessboard: https://github.com/Clariity/react-chessboard
- chessboard-element: https://github.com/justinfagnani/chessboard-element
- chessboard.js: https://chessboardjs.com/
- React Aria Components: https://react-aria.adobe.com/
- Radix Primitives: https://www.radix-ui.com/primitives
- Apache ECharts: https://echarts.apache.org/en/
- Storybook testing: https://storybook.js.org/docs/writing-tests
- Playwright visual comparisons: https://playwright.dev/docs/test-snapshots
- Browser cursor standard guidance:
  https://developer.mozilla.org/en-US/docs/Web/CSS/Reference/Properties/cursor
- Kenney Creative Commons Zero cursor pack:
  https://opengameart.org/content/cursor-pack-1
- Motion: https://motion.dev/
- canvas-confetti: https://github.com/catdad/canvas-confetti
- dotLottie Web: https://github.com/LottieFiles/dotlottie-web

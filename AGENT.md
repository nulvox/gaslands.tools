# Agent Brief: Gaslands.tools Development

You are building a browser-based team roster builder for the Gaslands Refuelled tabletop game. Read `CLAUDE.md` thoroughly before starting — it contains the full spec, architecture, data model, and mandatory development rules.

## Your Development Process

You MUST follow this exact loop for every task in `TODO.md`:

```
1. Read the current task from TODO.md
2. Write a failing test for the task's acceptance criteria
3. Run the test — confirm it fails for the expected reason
4. Write the minimal production code to pass the test
5. Run ALL tests — confirm everything passes
6. Lint ALL changed files
7. Verify the feature works (for UI tasks: describe what to check in browser)
8. Mark the task [x] in TODO.md
9. Commit with a message referencing the task
10. Move to the next task
```

**Never write production code without a failing test first.**
**Never mark a task done if any test fails or any lint error exists.**
**Never skip tasks or work on multiple tasks simultaneously.**

## Special Phase: PDF Extraction (Phase 1)

Phase 1 uses Python, not Go. For these tasks:

- Use `uv` for all Python package management (never `pip`)
- The extractor lives in `tools/extract-pdf/`
- Input: `Gaslands Refuelled Post-Apocalyptic Vehicular Mayhem.pdf` (in project root)
- Output: Structured markdown files in `data/` directory
- After automated extraction, manually review and correct the data against the rulebook
- The corrected data feeds into Go `internal/gamedata/` package in Phase 2

```bash
# Run the extractor
cd tools/extract-pdf && uv run python extract.py

# Review output
ls ../../data/
```

## Linting Commands

After every code change, run the relevant linters:

```bash
# Go
golangci-lint run ./...

# HTML validation
# Ensure: no unclosed tags, no duplicate IDs, valid attributes

# CSS (manual check: no syntax errors, print.css scoped to @media print)

# JS (ensure no console.error in browser, strict mode)
```

## Build Commands

```bash
# Full build
make build

# Compile Go to WASM only
GOOS=js GOARCH=wasm go build -o docs/main.wasm ./cmd/wasm/

# Copy WASM exec helper
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" docs/

# Serve locally
make serve
# or
python3 -m http.server 8080 --directory docs
```

## Testing

```bash
# Pure Go tests (core logic + gamedata)
go test ./internal/...

# With verbose output
go test -v ./internal/...

# Specific package
go test ./internal/team/...
go test ./internal/gamedata/...
```

Design `internal/team/` and `internal/gamedata/` packages as pure Go with NO `syscall/js` imports. Only `cmd/wasm/main.go` should import `syscall/js`. This keeps all core logic testable with standard `go test`.

## Architecture Guidance

### Separation of Concerns

```
internal/gamedata/   — Pure Go. Built-in database of vehicles, weapons, upgrades, perks, sponsors.
                       Lookup functions: GetVehicleType(name), ListWeapons(), GetSponsor(name), etc.
                       All data is hardcoded from the extracted+reviewed rulebook data.

internal/team/       — Pure Go. Team/vehicle data model, constructors, validation, serialization.
  team.go            — Team, Vehicle, WeaponInstance, etc. structs + constructors
  defaults.go        — Default team config (50 Cans, empty vehicle list)
  validate.go        — Budget, build slots, sponsor perk class, variant rules
  serialization.go   — JSON marshal/unmarshal with version handling
  export.go          — HTML export via Go html/template (embedded template)

cmd/wasm/main.go     — Bridge layer. Thin wrapper that:
                       - Registers Go functions as JS globals
                       - Converts JS values <-> Go structs via JSON
                       - select{} to keep WASM alive

web/js/              — Frontend JS that calls WASM-exposed functions
  app.js             — Boot: load WASM, init UI, restore teams from localStorage
  team-ui.js         — Render team header, budget bar, notes, validation warnings
  vehicle-ui.js      — Render vehicle cards, weapon/upgrade/perk editors, database dropdowns
  tabs.js            — Tab bar: new team, switch, close, rename
  storage.js         — localStorage CRUD (gaslands_team_{id}, gaslands_manifest)
  import-export.js   — JSON file download/upload, HTML export download, print trigger
  share.js           — URL hash encoding/decoding for shareable links
  theme.js           — Dark mode toggle (gaslands_theme in localStorage)
```

### Data Flow

```
User adds weapon from dropdown
  -> JS captures weapon name selection
  -> JS calls addWeaponToVehicle(teamJSON, vehicleID, weaponName) in WASM
  -> Go: looks up weapon in gamedata, adds WeaponInstance to vehicle
  -> Go: recalculates costs, validates budget + slots
  -> Go: returns {team: updatedJSON, warnings: [...]}
  -> JS: re-renders vehicle card with new weapon
  -> JS: updates budget bar
  -> JS: saves to localStorage
```

### Key Design Decisions

1. **Go WASM as pure logic engine**: All team manipulation, validation, cost calculation, and game database lookups happen in Go. JS never directly mutates team data.

2. **JSON as the bridge format**: Data crosses the JS/Go boundary as JSON strings. Simple, debuggable, no complex `syscall/js` value marshaling.

3. **localStorage, not server**: Teams stored as JSON in localStorage. Key format: `gaslands_team_{id}`. Manifest key: `gaslands_manifest` (list of team IDs + metadata).

4. **HTML export in Go**: The styled roster page is generated by Go's `html/template` with an embedded template. The WASM bridge returns the HTML string; JS triggers the download.

5. **Print = HTML export view + browser print**: No separate print layout. The HTML export template IS the print-friendly layout. JS opens it in a new window/iframe and calls `window.print()`.

6. **GitHub Pages via `docs/`**: Build outputs to `docs/`. Repo configured for GitHub Pages from `docs/` on main branch.

7. **Database + Custom**: Every game item category has a built-in database AND a custom entry option. Users can mix database items and homebrew content freely.

## Reference Output

See `examples/the-warden-roster.html` for the target HTML export appearance. Study it carefully — the export template should produce output matching this style:

- Post-apocalyptic color scheme with CSS variables
- Google Fonts: Black Ops One, Share Tech Mono, Oswald
- Vehicle cards with stat grids, cost breakdowns, slot bars
- Tabbed layout (Roster, Sponsor, Strategy, Key Rules)
- Budget summary with pill badges
- Self-contained single HTML file (inline CSS, minimal JS)

## What "Done" Looks Like

The finished application:
1. Loads in a browser from static files (GitHub Pages compatible)
2. Creates new teams with configurable budget and sponsor selection
3. Adds vehicles from a built-in database of all Gaslands vehicle types
4. Equips vehicles with weapons, upgrades, and perks from database or custom entries
5. Tracks budget spending in real-time with soft validation warnings
6. Validates build slots, sponsor perk classes, and variant restrictions (soft warnings)
7. Multiple teams via tabs
8. Auto-saves to localStorage on every edit
9. Exports team as downloadable `.gaslands.json` file
10. Imports team from uploaded `.json` file
11. Generates standalone styled HTML roster page for download
12. Prints roster via browser print dialog (using HTML export view)
13. Shares teams via encoded URL hash
14. All Go code has test coverage, all linting passes

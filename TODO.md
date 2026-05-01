# TODO — Gaslands.tools Incremental Build

Work through these tasks in order. Each task is a small, testable increment.
Do NOT skip ahead. Mark `[x]` only after all tests pass and linting is clean.

---

## Phase 0: Project Skeleton

- [x] **0.1** Initialize Go module (`go mod init gaslands.tools`), verify directory structure (`cmd/wasm/`, `internal/gamedata/`, `internal/team/`, `web/css/`, `web/js/`), create minimal `cmd/wasm/main.go` that compiles to WASM and prints "WASM loaded" to the JS console.
- [x] **0.2** Verify `Makefile` targets work: `make build` produces `docs/main.wasm`, `make test` runs (even if no tests yet), `make lint` runs golangci-lint, `make serve` starts local server. Fix any issues.
- [x] **0.3** Create minimal `web/index.html` that loads `wasm_exec.js` and `main.wasm`, and displays a loading message until WASM is ready. Create `web/js/app.js` that initializes the WASM module. Verify in browser: console shows "WASM loaded".
- [x] **0.4** Verify `.golangci.yml` works. Run `golangci-lint run ./...` — must pass cleanly.

## Phase 1: PDF Data Extraction

- [x] **1.1** Set up the Python extractor project: `cd tools/extract-pdf && uv sync`. Verify `uv run python extract.py --help` runs without error. The script should accept a `--pdf` argument pointing to the rulebook PDF and an `--output-dir` argument for the data directory.
- [x] **1.2** Implement full text extraction from the PDF. The script should extract all text content page by page, preserving structure. Output a raw `data/full-text.md` for review.
- [x] **1.3** Implement structured extraction: parse the raw text to identify and separate vehicle types (with stats tables), weapons, upgrades, perks (with classes), and sponsors (with perk classes and sponsor perks). Output separate files: `data/vehicles.md`, `data/weapons.md`, `data/upgrades.md`, `data/perks.md`, `data/sponsors.md`.
- [x] **1.4** Run the extractor against the actual PDF: `cd tools/extract-pdf && uv run python extract.py --pdf "../../Gaslands Refuelled Post-Apocalyptic Vehicular Mayhem.pdf" --output-dir ../../data/`. Review the output for completeness and accuracy. Manually correct any extraction errors in the markdown files.
- [x] **1.5** Validate extracted data against the example roster in `examples/the-warden-roster.html`. Confirm: The Warden sponsor exists with Aggression+Badass perk classes and Prison Cars+Fireworks sponsor perks. Car/Truck vehicle types have correct stats. Weapons (Handgun, Machine Gun, Heavy Machine Gun, Minigun) have correct costs and dice. Perks (Battlehammer, Grinderman) have correct costs and classes. Upgrades (Ram) have correct cost and slots. Fix any discrepancies in the data files.

## Phase 2: Game Database (Pure Go, Fully Testable)

- [x] **2.1** Define `VehicleType` struct in `internal/gamedata/vehicles.go` with fields: Name, WeightClass, BaseCost, Hull, Handling, MaxGear, Crew, BuildSlots, SpecialRules. Populate with all vehicle types from `data/vehicles.md`. Test: lookup "Car" returns correct stats (Hull 8, Handling 3, MaxGear 5, Crew 2, Slots 2, Middleweight). Test: lookup "Truck" returns correct stats.
- [x] **2.2** Define `Weapon` struct in `internal/gamedata/weapons.go` with fields: Name, Cost, AttackDice, Range, Slots, SpecialRules. Populate with all weapons from `data/weapons.md`. Test: lookup "Heavy Machine Gun" returns 3 Cans, "3D6". Test: "Handgun" costs 0.
- [x] **2.3** Define `Upgrade` struct in `internal/gamedata/upgrades.go` with fields: Name, Cost, Slots, Description, Restrictions. Populate from `data/upgrades.md`. Test: lookup "Ram" returns correct cost and slot usage.
- [x] **2.4** Define `Perk` struct in `internal/gamedata/perks.go` with fields: Name, Cost, Class, Description. Populate from `data/perks.md`. Test: lookup "Battlehammer" returns Aggression class, 4 Cans. Test: list all Aggression perks.
- [x] **2.5** Define `Sponsor` struct in `internal/gamedata/sponsors.go` with fields: Name, FlavorText, PerkClasses []string, SponsorPerks []SponsorPerk. `SponsorPerk` has Name, Description. Populate from `data/sponsors.md`. Test: lookup "The Warden" returns PerkClasses=["Aggression","Badass"] and has Prison Cars + Fireworks sponsor perks.
- [x] **2.6** Create `GameDatabase` aggregate in `internal/gamedata/gamedata.go` with lookup and listing functions: `GetVehicleType(name)`, `ListVehicleTypes()`, `GetWeapon(name)`, `ListWeapons()`, `GetUpgrade(name)`, `ListUpgrades()`, `GetPerk(name)`, `ListPerks()`, `ListPerksByClass(class)`, `GetSponsor(name)`, `ListSponsors()`. Test: all lookups work, listing returns complete data, missing items return appropriate errors.

## Phase 3: Core Team/Vehicle Model (Pure Go, Fully Testable)

- [x] **3.1** Define `Team` struct in `internal/team/team.go` with fields: ID, Name, Budget. Implement `NewTeam()` constructor that generates a UUID, sets Name="", Budget=50. Test: new team has non-empty ID, empty name, budget 50.
- [x] **3.2** Add `SponsorSelection` struct (Name string, IsCustom bool, CustomPerkClasses []string) and `Sponsor` field to Team. Test: setting a sponsor by name.
- [x] **3.3** Define `Vehicle` struct with fields: ID, CustomName, Role, VehicleType string. Implement `NewVehicle(vehicleType string)` constructor. Test: new vehicle has UUID, empty custom name, correct vehicle type ref.
- [x] **3.4** Add `Vehicles []Vehicle` field to Team. Implement `AddVehicle(vehicleType)` and `RemoveVehicle(vehicleID)` methods. Test: add two vehicles, remove one by ID, verify list.
- [x] **3.5** Define `WeaponInstance`, `UpgradeInstance`, `PerkInstance` structs (as specified in CLAUDE.md). Add `Weapons`, `Upgrades`, `Perks` fields to Vehicle. Test: add/remove items from each list.
- [x] **3.6** Add Variant string field to Vehicle. Add Notes string field to both Team and Vehicle. Test: setting variant and notes.
- [x] **3.7** Implement `VehicleCost(v Vehicle, db GameDatabase)` — calculates total cost: base vehicle cost (from DB, adjusted by variant/sponsor) + sum of weapon costs + sum of upgrade costs + sum of perk costs. Test: vehicle with known weapons/upgrades/perks returns correct total. Test: custom items use their stored cost.
- [x] **3.8** Implement `TeamCost(t Team, db GameDatabase)` — sum of all vehicle costs. Implement `TeamHull(t Team, db GameDatabase)` — sum of all vehicle hulls. Test: team with multiple vehicles returns correct totals.
- [x] **3.9** Implement `SlotsUsed(v Vehicle)` — sum of weapon slots + upgrade slots. Implement `SlotsAvailable(v Vehicle, db GameDatabase)` — vehicle type's build slots minus slots used. Test: correct slot math.

## Phase 4: Serialization (JSON Import/Export)

- [x] **4.1** Implement `Team.ToJSON() ([]byte, error)` and `FromJSON([]byte) (*Team, error)` in `internal/team/serialization.go`. Test: round-trip a team with vehicles, weapons, upgrades, perks through JSON — all fields match.
- [x] **4.2** Add a `Version` field to JSON output (start at `"1.0"`). `FromJSON` checks version and returns error for unknown versions. Missing version defaults to "1.0". Test: version handling.
- [x] **4.3** Test edge cases: empty team (no vehicles), vehicle with no weapons/upgrades/perks, custom items (IsCustom=true), unicode characters in names/notes, team with many vehicles.

## Phase 5: Validation (Soft Warnings)

- [x] **5.1** Implement `ValidateBudget(team, db) []string` in `internal/team/validate.go` — returns warning if total cost exceeds budget, including the overage amount. Test: under budget = no warning; at budget = no warning; over budget = warning with "over by X Cans".
- [x] **5.2** Implement `ValidateBuildSlots(vehicle, db) []string` — returns warning if slots used exceed vehicle's available slots. Test: under/at/over slot limit.
- [x] **5.3** Implement `ValidateSponsorPerks(team, db) []string` — returns warning for each perk whose class is not in the sponsor's allowed classes. Skip validation for custom sponsors. Test: valid perk class = no warning; invalid = warning naming the perk and class.
- [x] **5.4** Implement `ValidateVariants(vehicle, db) []string` — returns warning if variant requirements not met (e.g., Prison Car requires Middleweight). Test: valid variant + weight = no warning; invalid = warning.
- [x] **5.5** Implement `ValidateTeam(team, db) []string` — combines all validations. Returns empty list for valid team. Also warns on empty team (no vehicles). Test: various combinations of issues.

## Phase 6: WASM Bridge

- [x] **6.1** In `cmd/wasm/main.go`, expose `createTeam()` to JS — returns a new team as JSON string. Verify from browser console: valid JSON with default fields.
- [x] **6.2** Expose `updateTeam(jsonStr)` — accepts modified team JSON, validates, returns `{team: json, warnings: []}`. Test from console.
- [x] **6.3** Expose `validateTeam(jsonStr)` — returns just the warnings array as JSON. Test from console.
- [x] **6.4** Expose game database query functions: `getVehicleTypes()`, `getWeapons()`, `getUpgrades()`, `getPerks()`, `getSponsors()` — each returns the full list as JSON. Test from console.
- [x] **6.5** Expose `importTeam(jsonStr)` and `exportTeamJSON(jsonStr)` — import parses and validates external JSON; export returns clean JSON for download. Test from console.
- [x] **6.6** Expose `exportTeamHTML(jsonStr)` — returns a styled standalone HTML string for the team roster. Uses Go `html/template` with an embedded template matching the style of `examples/the-warden-roster.html`. Test: output is valid HTML with correct team data.

## Phase 7: Basic UI Shell

- [x] **7.1** Create `web/css/style.css` with post-apocalyptic theme: CSS variables for rust red, amber, soot black color scheme. Layout with header bar, tab bar, budget summary bar, main content area. Dark mode via `[data-theme="dark"]`. Google Fonts: Black Ops One, Share Tech Mono, Oswald.
- [x] **7.2** Implement tab bar in `web/js/tabs.js`: "New Team" button creates a tab, clicking a tab activates it, tabs show team name (or "New Team"). No team rendering yet — just tab switching with placeholder content.
- [x] **7.3** Wire tab creation to `createTeam()` WASM call. When a new tab is created, call `createTeam()`, store result. Verify: click "New Team" → new tab → console shows team JSON.
- [x] **7.4** Implement `web/js/theme.js`: dark mode toggle button, stores preference in `gaslands_theme` localStorage key. Verify: toggle works, persists on reload.

## Phase 8: Team Roster UI

- [x] **8.1** In `web/js/team-ui.js`, implement `renderTeamHeader(team)` — renders editable Name field and Sponsor dropdown (populated from `getSponsors()` WASM call, plus "Custom..." option). On change, update via WASM. Verify: edit name → tab updates; select sponsor.
- [x] **8.2** Implement `renderBudgetBar(team)` — always-visible summary bar showing: Budget (editable input), Spent, Remaining (color-coded: green if under, amber if at, red if over), Vehicle Count, Total Hull. Updates on every team change. Verify: numbers update as vehicles are added/modified.
- [x] **8.3** Implement `renderBudgetInput(team)` — budget is an editable number input in the budget bar. Common presets as quick-select buttons (20, 50, 100, 150 Cans). On change, update via WASM. Verify: changing budget updates remaining and validation.
- [x] **8.4** Implement `renderVehicleList(team)` — container for vehicle cards. "Add Vehicle" button opens a vehicle type selector dropdown (populated from `getVehicleTypes()`). Selecting a type calls WASM to add the vehicle and re-renders. Verify: add a Car, add a Truck, both appear as cards.
- [x] **8.5** Implement `renderTeamNotes(team)` — freeform textarea for strategy notes, lore, anything. On change, update via WASM. Verify: edit notes, data persists.
- [x] **8.6** Implement `renderValidationWarnings(team)` — non-intrusive panel (collapsible) showing warnings from `validateTeam()`. Updates live on edit. Verify: go over budget → warning appears; fix → disappears.
- [x] **8.7** Compose all team-level renderers into `renderTeamRoster(team)` called on tab activation and data changes. Verify: full team view renders.

## Phase 9: Vehicle Card UI

- [x] **9.1** In `web/js/vehicle-ui.js`, implement `renderVehicleCard(vehicle, team)` — card layout showing: custom name (editable), vehicle type label, role (editable), stat grid (Hull, Handling, MaxGear, Crew, Slots), cost badge. Verify: card displays correct stats from database.
- [x] **9.2** Add variant selector to vehicle card — dropdown of available variants for the vehicle's weight class (plus "None"). On change, update via WASM. Verify: selecting Prison Car adjusts displayed hull and cost.
- [x] **9.3** Implement `renderWeaponEditor(vehicle)` — list of equipped weapons with remove buttons. "Add Weapon" dropdown populated from `getWeapons()` plus "Custom..." option. For custom: show name, cost, dice, range, slots inputs. On add/remove, update via WASM. Verify: add Machine Gun → appears in list with correct cost; remove → disappears.
- [x] **9.4** Implement `renderUpgradeEditor(vehicle)` — same pattern as weapons. Dropdown from `getUpgrades()` plus custom. On add/remove, update via WASM. Verify: add Ram → appears with correct cost and slot usage.
- [x] **9.5** Implement `renderPerkEditor(vehicle)` — same pattern. Dropdown from `getPerks()` filtered by sponsor's allowed classes (plus "Custom..."). On add/remove, update via WASM. Verify: add Battlehammer → appears with correct class and cost.
- [x] **9.6** Implement `renderCostBreakdown(vehicle)` — table showing: base cost, each weapon cost, each upgrade cost, each perk cost, variant adjustments, total. Verify: matches manual calculation.
- [x] **9.7** Implement `renderSlotBar(vehicle)` — visual bar showing slots used vs. available (like the example's percentage fill bars). Color-coded: green under limit, red over. Verify: adding slotted items updates the bar.
- [x] **9.8** Implement `renderVehicleNotes(vehicle)` — freeform textarea per vehicle. On change, update via WASM. Verify: edit notes, data persists.
- [x] **9.9** Add delete vehicle button to each card (with confirmation). On delete, remove vehicle and re-render. Verify: delete vehicle → card removed → budget/totals update.
- [x] **9.10** Integrate vehicle cards into the team roster view. Each vehicle in `team.Vehicles` renders as a card via `renderVehicleCard()`. Cards update individually on change. Verify: full team with multiple vehicles renders and edits correctly.

## Phase 10: Persistence (localStorage)

- [x] **10.1** Implement `web/js/storage.js`: `saveTeam(team)` writes to `localStorage` keyed by `gaslands_team_{id}`. `loadTeam(id)` reads and parses. `listTeams()` returns all stored team IDs and names. `deleteTeam(id)` removes. Manifest key: `gaslands_manifest`.
- [x] **10.2** Wire auto-save: every time `updateTeam()` returns successfully, save to localStorage. On app load, restore all teams from localStorage and recreate tabs. Verify: create team, add vehicle, reload page → team and vehicle persist.
- [x] **10.3** Handle storage errors gracefully: if localStorage is full or unavailable, show a non-blocking warning banner. Verify: simulate full storage → banner appears.

## Phase 11: Import/Export & Sharing

- [x] **11.1** Implement JSON export in `web/js/import-export.js`: "Export JSON" button triggers download of current team as `{name}.gaslands.json`. Verify: export, open file, valid JSON with all team data.
- [x] **11.2** Implement JSON import: "Import JSON" button opens file picker. On select, read contents, call `importTeam()` WASM, create new tab. Verify: export team, delete it, import file → team restored with all vehicles.
- [x] **11.3** Handle import errors: malformed JSON, wrong version, missing fields. Show clear error messages. Verify: import random file → friendly error.
- [x] **11.4** Implement HTML export: "Export HTML" button calls `exportTeamHTML()` WASM, receives styled HTML string, triggers download as `{name}-roster.html`. Verify: exported HTML opens in browser and matches style of `examples/the-warden-roster.html`.
- [x] **11.5** Implement URL sharing in `web/js/share.js`: "Share" button encodes team JSON (compressed) into URL hash, copies URL to clipboard. On app load, check hash for shared team and import it. Verify: share → paste URL in new tab → team loads.

## Phase 12: Print Layout

- [x] **12.1** Create `web/css/print.css` with `@media print` rules: hide app UI (tabs, buttons, editors), show only the styled roster view. Verify: print preview shows clean roster.
- [x] **12.2** Implement print flow: "Print" button calls `exportTeamHTML()` WASM to get styled HTML, opens in new window/iframe, triggers `window.print()`. The printed output uses the HTML export template (same as the downloadable HTML). Verify: print preview matches the HTML export style.

## Phase 13: Polish & Edge Cases

- [ ] **13.1** Tab management: close button on tabs (with confirmation if team has data), rename tab by double-clicking, tab label shows team name or "New Team" for unnamed.
- [ ] **13.2** Empty state: when no teams exist, show a welcome/getting-started message with prominent "Create Team" button and brief feature overview.
- [ ] **13.3** Keyboard accessibility: all interactive elements reachable via Tab key, Enter/Space to activate, Escape to cancel dropdowns. Vehicle cards navigable.
- [ ] **13.4** Responsive layout: ensure usable at 1024px+ width. Vehicle cards stack vertically on narrower viewports. Budget bar collapses gracefully.
- [ ] **13.5** Performance: test with 5+ teams, 10+ vehicles per team. Ensure tab switching is instant, no lag on edits, dropdown populations are fast.

## Phase 14: Build & Deploy

- [ ] **14.1** Verify `make build` produces a complete `docs/` directory: `index.html`, `main.wasm`, `wasm_exec.js`, CSS files, JS files. Verify: `make serve` → app works in browser.
- [ ] **14.2** Verify `.github/workflows/pages.yml` is correct for the `nulvox/gaslands.tools` repo. Update deployment URL references.
- [ ] **14.3** Final test pass: run all Go tests, lint all files. Create a team from scratch: set sponsor to The Warden, set budget to 50, add a Truck + 2 Cars, equip weapons/upgrades/perks, verify budget tracking. Export JSON, delete team, import JSON → team restored. Export HTML → opens as styled roster. Print → looks correct. Share URL → loads in new tab. All green.
- [ ] **14.4** Write a minimal `README.md`: what this is, how to build, how to use, link to live site (nulvox.github.io/gaslands.tools).

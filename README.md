# Gaslands.tools

A browser-based team roster builder and manager for [Gaslands Refuelled](https://gaslands.com/). Build teams, add vehicles, equip weapons/upgrades/perks, track budgets, and export styled roster sheets.

**Live site:** [nulvox.github.io/gaslands.tools](https://nulvox.github.io/gaslands.tools)

## Features

- Create and manage multiple team rosters with tabs
- Full game database: 16 vehicle types, 38 weapons, 10 upgrades, 72 perks, 13 sponsors
- Real-time budget tracking and build slot validation
- Sponsor perk class filtering
- Custom items support (weapons, upgrades, perks)
- JSON import/export
- Styled HTML roster export
- URL-based team sharing
- Print-friendly output
- localStorage persistence
- Dark/light theme toggle

## Tech Stack

- **Core logic:** Go 1.23+, compiled to WebAssembly
- **Frontend:** Vanilla HTML/CSS/JavaScript (no framework)
- **Storage:** localStorage
- **Hosting:** GitHub Pages (static files)

## Build

```sh
# Prerequisites: Go 1.23+, golangci-lint

# Build WASM and copy web files to docs/
make build

# Run tests
make test

# Lint
make lint

# Serve locally at http://localhost:8080
make serve
```

## License

See [LICENSE](LICENSE).

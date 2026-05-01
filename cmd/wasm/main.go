package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"gaslands.tools/internal/gamedata"
	"gaslands.tools/internal/team"
)

// createTeam returns a new team as a JSON string.
func createTeam(_ js.Value, _ []js.Value) any {
	t := team.NewTeam()
	data, err := t.ToJSON()
	if err != nil {
		return errorResult(err)
	}
	return string(data)
}

// updateTeam accepts team JSON, validates, and returns {team: json, warnings: []}.
func updateTeam(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return errorResult(fmt.Errorf("updateTeam requires 1 argument"))
	}
	jsonStr := args[0].String()

	t, err := team.FromJSON([]byte(jsonStr))
	if err != nil {
		return errorResult(err)
	}

	warnings := team.ValidateTeam(t)
	data, err := t.ToJSON()
	if err != nil {
		return errorResult(err)
	}

	return marshalResult(map[string]any{
		"team":     json.RawMessage(data),
		"warnings": warnings,
	})
}

// validateTeam returns just the warnings array as a JSON string.
func validateTeam(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return errorResult(fmt.Errorf("validateTeam requires 1 argument"))
	}

	t, err := team.FromJSON([]byte(args[0].String()))
	if err != nil {
		return errorResult(err)
	}

	warnings := team.ValidateTeam(t)
	return marshalResult(warnings)
}

// getVehicleTypes returns all vehicle types as JSON.
func getVehicleTypes(_ js.Value, _ []js.Value) any {
	return marshalResult(gamedata.ListVehicleTypes())
}

// getWeapons returns all weapons as JSON.
func getWeapons(_ js.Value, _ []js.Value) any {
	return marshalResult(gamedata.ListWeapons())
}

// getUpgrades returns all upgrades as JSON.
func getUpgrades(_ js.Value, _ []js.Value) any {
	return marshalResult(gamedata.ListUpgrades())
}

// getPerks returns all perks as JSON.
func getPerks(_ js.Value, _ []js.Value) any {
	return marshalResult(gamedata.ListPerks())
}

// getSponsors returns all sponsors as JSON.
func getSponsors(_ js.Value, _ []js.Value) any {
	return marshalResult(gamedata.ListSponsors())
}

// importTeam parses and validates external JSON, returns {team, warnings}.
func importTeam(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return errorResult(fmt.Errorf("importTeam requires 1 argument"))
	}

	t, err := team.FromJSON([]byte(args[0].String()))
	if err != nil {
		return errorResult(err)
	}

	warnings := team.ValidateTeam(t)
	data, err := t.ToJSON()
	if err != nil {
		return errorResult(err)
	}

	return marshalResult(map[string]any{
		"team":     json.RawMessage(data),
		"warnings": warnings,
	})
}

// exportTeamJSON returns clean JSON for download.
func exportTeamJSON(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return errorResult(fmt.Errorf("exportTeamJSON requires 1 argument"))
	}

	t, err := team.FromJSON([]byte(args[0].String()))
	if err != nil {
		return errorResult(err)
	}

	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return errorResult(err)
	}
	return string(data)
}

// marshalResult serializes a value to a JSON string.
func marshalResult(v any) string {
	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf(`{"error":%q}`, err.Error())
	}
	return string(data)
}

// exportTeamHTML returns a styled standalone HTML string for the team roster.
func exportTeamHTML(_ js.Value, args []js.Value) any {
	if len(args) < 1 {
		return errorResult(fmt.Errorf("exportTeamHTML requires 1 argument"))
	}

	t, err := team.FromJSON([]byte(args[0].String()))
	if err != nil {
		return errorResult(err)
	}

	html, err := team.ExportHTML(t)
	if err != nil {
		return errorResult(err)
	}
	return html
}

// errorResult formats an error as a JSON error string.
func errorResult(err error) string {
	return fmt.Sprintf(`{"error":%q}`, err.Error())
}

func main() {
	js.Global().Set("createTeam", js.FuncOf(createTeam))
	js.Global().Set("updateTeam", js.FuncOf(updateTeam))
	js.Global().Set("validateTeam", js.FuncOf(validateTeam))
	js.Global().Set("getVehicleTypes", js.FuncOf(getVehicleTypes))
	js.Global().Set("getWeapons", js.FuncOf(getWeapons))
	js.Global().Set("getUpgrades", js.FuncOf(getUpgrades))
	js.Global().Set("getPerks", js.FuncOf(getPerks))
	js.Global().Set("getSponsors", js.FuncOf(getSponsors))
	js.Global().Set("importTeam", js.FuncOf(importTeam))
	js.Global().Set("exportTeamJSON", js.FuncOf(exportTeamJSON))
	js.Global().Set("exportTeamHTML", js.FuncOf(exportTeamHTML))

	js.Global().Get("console").Call("log", "WASM loaded — Gaslands.tools ready")
	fmt.Println("Gaslands.tools WASM module initialized")

	select {}
}

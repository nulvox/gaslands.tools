package team

import (
	"encoding/json"
	"testing"
)

// 4.1: Round-trip serialization
func TestToJSON_FromJSON_RoundTrip(t *testing.T) {
	tm := NewTeam()
	tm.Name = "Test Team"
	tm.Budget = 100
	tm.Sponsor = SponsorSelection{Name: "The Warden"}
	tm.Notes = "Strategy notes"

	v := tm.AddVehicle("Car")
	v.CustomName = "Iron Coffin"
	v.Role = "Brawler"
	v.Variant = "Prison Car"
	v.Notes = "Front line"
	v.Weapons = []WeaponInstance{
		{Name: "Machine Gun", Cost: 2, AttackDice: "2D6", Range: "Double", Slots: 1},
	}
	v.Upgrades = []UpgradeInstance{
		{Name: "Ram", Cost: 4, Slots: 1, Description: "Rams things"},
	}
	v.Perks = []PerkInstance{
		{Name: "Battlehammer", Cost: 4, Class: "Aggression", Description: "+1 die"},
	}
	tm.Vehicles[0] = *v

	data, err := tm.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON error: %v", err)
	}

	restored, err := FromJSON(data)
	if err != nil {
		t.Fatalf("FromJSON error: %v", err)
	}

	if restored.ID != tm.ID {
		t.Errorf("ID: got %q, want %q", restored.ID, tm.ID)
	}
	if restored.Name != tm.Name {
		t.Errorf("Name: got %q, want %q", restored.Name, tm.Name)
	}
	if restored.Budget != tm.Budget {
		t.Errorf("Budget: got %d, want %d", restored.Budget, tm.Budget)
	}
	if restored.Sponsor.Name != tm.Sponsor.Name {
		t.Errorf("Sponsor: got %q, want %q", restored.Sponsor.Name, tm.Sponsor.Name)
	}
	if restored.Notes != tm.Notes {
		t.Errorf("Notes: got %q, want %q", restored.Notes, tm.Notes)
	}
	if len(restored.Vehicles) != 1 {
		t.Fatalf("Vehicles: got %d, want 1", len(restored.Vehicles))
	}

	rv := restored.Vehicles[0]
	if rv.CustomName != "Iron Coffin" {
		t.Errorf("Vehicle CustomName: got %q", rv.CustomName)
	}
	if rv.Role != "Brawler" {
		t.Errorf("Vehicle Role: got %q", rv.Role)
	}
	if rv.Variant != "Prison Car" {
		t.Errorf("Vehicle Variant: got %q", rv.Variant)
	}
	if len(rv.Weapons) != 1 || rv.Weapons[0].Name != "Machine Gun" {
		t.Error("Weapon not preserved")
	}
	if len(rv.Upgrades) != 1 || rv.Upgrades[0].Name != "Ram" {
		t.Error("Upgrade not preserved")
	}
	if len(rv.Perks) != 1 || rv.Perks[0].Name != "Battlehammer" {
		t.Error("Perk not preserved")
	}
}

// 4.2: Version handling
func TestToJSON_IncludesVersion(t *testing.T) {
	tm := NewTeam()
	data, err := tm.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON error: %v", err)
	}

	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("unmarshal raw: %v", err)
	}

	version, ok := raw["version"]
	if !ok {
		t.Fatal("expected version field in JSON")
	}
	if version != "1.0" {
		t.Errorf("version: got %v, want %q", version, "1.0")
	}
}

func TestFromJSON_MissingVersionDefaultsTo1(t *testing.T) {
	jsonStr := `{"id":"abc","name":"Test","budget":50,"sponsor":{"name":""},"vehicles":[]}`
	tm, err := FromJSON([]byte(jsonStr))
	if err != nil {
		t.Fatalf("FromJSON error: %v", err)
	}
	if tm.Version != "1.0" {
		t.Errorf("version: got %q, want %q", tm.Version, "1.0")
	}
}

func TestFromJSON_UnknownVersionError(t *testing.T) {
	jsonStr := `{"id":"abc","name":"Test","budget":50,"version":"99.0","sponsor":{"name":""},"vehicles":[]}`
	_, err := FromJSON([]byte(jsonStr))
	if err == nil {
		t.Fatal("expected error for unknown version")
	}
}

// 4.3: Edge cases
func TestRoundTrip_EmptyTeam(t *testing.T) {
	tm := NewTeam()
	data, err := tm.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON error: %v", err)
	}
	restored, err := FromJSON(data)
	if err != nil {
		t.Fatalf("FromJSON error: %v", err)
	}
	if len(restored.Vehicles) != 0 {
		t.Errorf("expected 0 vehicles, got %d", len(restored.Vehicles))
	}
}

func TestRoundTrip_VehicleNoEquipment(t *testing.T) {
	tm := NewTeam()
	tm.AddVehicle("Buggy")
	data, err := tm.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON error: %v", err)
	}
	restored, err := FromJSON(data)
	if err != nil {
		t.Fatalf("FromJSON error: %v", err)
	}
	if len(restored.Vehicles) != 1 {
		t.Fatalf("expected 1 vehicle, got %d", len(restored.Vehicles))
	}
	rv := restored.Vehicles[0]
	if len(rv.Weapons) != 0 {
		t.Errorf("expected 0 weapons, got %d", len(rv.Weapons))
	}
	if len(rv.Upgrades) != 0 {
		t.Errorf("expected 0 upgrades, got %d", len(rv.Upgrades))
	}
	if len(rv.Perks) != 0 {
		t.Errorf("expected 0 perks, got %d", len(rv.Perks))
	}
}

func TestRoundTrip_CustomItems(t *testing.T) {
	tm := NewTeam()
	v := tm.AddVehicle("Car")
	v.Weapons = []WeaponInstance{
		{Name: "Custom Laser", Cost: 7, Slots: 2, IsCustom: true},
	}
	v.Upgrades = []UpgradeInstance{
		{Name: "Custom Armour", Cost: 5, Slots: 1, IsCustom: true},
	}
	v.Perks = []PerkInstance{
		{Name: "Custom Perk", Cost: 3, Class: "Custom", IsCustom: true},
	}
	tm.Vehicles[0] = *v

	data, err := tm.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON error: %v", err)
	}
	restored, err := FromJSON(data)
	if err != nil {
		t.Fatalf("FromJSON error: %v", err)
	}

	rv := restored.Vehicles[0]
	if !rv.Weapons[0].IsCustom {
		t.Error("weapon IsCustom not preserved")
	}
	if !rv.Upgrades[0].IsCustom {
		t.Error("upgrade IsCustom not preserved")
	}
	if !rv.Perks[0].IsCustom {
		t.Error("perk IsCustom not preserved")
	}
}

func TestRoundTrip_UnicodeNames(t *testing.T) {
	tm := NewTeam()
	tm.Name = "Les Flamb\u00e9s \u2014 \u706b\u7130\u6226\u58eb"
	tm.Notes = "\u2620 Danger \u2620"
	v := tm.AddVehicle("Car")
	v.CustomName = "El Diablo \U0001F525"
	tm.Vehicles[0] = *v

	data, err := tm.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON error: %v", err)
	}
	restored, err := FromJSON(data)
	if err != nil {
		t.Fatalf("FromJSON error: %v", err)
	}
	if restored.Name != tm.Name {
		t.Errorf("unicode name: got %q, want %q", restored.Name, tm.Name)
	}
	if restored.Notes != tm.Notes {
		t.Errorf("unicode notes: got %q, want %q", restored.Notes, tm.Notes)
	}
	if restored.Vehicles[0].CustomName != "El Diablo \U0001F525" {
		t.Errorf("unicode vehicle name: got %q", restored.Vehicles[0].CustomName)
	}
}

func TestRoundTrip_ManyVehicles(t *testing.T) {
	tm := NewTeam()
	for range 10 {
		tm.AddVehicle("Car")
	}

	data, err := tm.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON error: %v", err)
	}
	restored, err := FromJSON(data)
	if err != nil {
		t.Fatalf("FromJSON error: %v", err)
	}
	if len(restored.Vehicles) != 10 {
		t.Errorf("expected 10 vehicles, got %d", len(restored.Vehicles))
	}
}

func TestFromJSON_MalformedJSON(t *testing.T) {
	_, err := FromJSON([]byte("not json"))
	if err == nil {
		t.Fatal("expected error for malformed JSON")
	}
}

package team

import (
	"testing"

	"gaslands.tools/internal/gamedata"
)

// 3.1: NewTeam constructor
func TestNewTeam(t *testing.T) {
	tm := NewTeam()
	if tm.ID == "" {
		t.Error("expected non-empty ID")
	}
	if tm.Name != "" {
		t.Errorf("expected empty name, got %q", tm.Name)
	}
	if tm.Budget != 50 {
		t.Errorf("expected budget 50, got %d", tm.Budget)
	}
}

func TestNewTeam_UniqueIDs(t *testing.T) {
	t1 := NewTeam()
	t2 := NewTeam()
	if t1.ID == t2.ID {
		t.Error("expected unique IDs for different teams")
	}
}

// 3.2: SponsorSelection
func TestTeam_SetSponsor(t *testing.T) {
	tm := NewTeam()
	tm.Sponsor = SponsorSelection{Name: "The Warden"}
	if tm.Sponsor.Name != "The Warden" {
		t.Errorf("sponsor name: got %q, want %q", tm.Sponsor.Name, "The Warden")
	}
	if tm.Sponsor.IsCustom {
		t.Error("expected IsCustom=false for database sponsor")
	}
}

// 3.3: NewVehicle constructor
func TestNewVehicle(t *testing.T) {
	v := NewVehicle("Car")
	if v.ID == "" {
		t.Error("expected non-empty vehicle ID")
	}
	if v.VehicleType != "Car" {
		t.Errorf("vehicle type: got %q, want %q", v.VehicleType, "Car")
	}
	if v.CustomName != "" {
		t.Errorf("expected empty custom name, got %q", v.CustomName)
	}
}

// 3.4: AddVehicle / RemoveVehicle
func TestTeam_AddRemoveVehicle(t *testing.T) {
	tm := NewTeam()
	v1 := tm.AddVehicle("Car")
	v2 := tm.AddVehicle("Truck")

	if len(tm.Vehicles) != 2 {
		t.Fatalf("expected 2 vehicles, got %d", len(tm.Vehicles))
	}

	tm.RemoveVehicle(v1.ID)
	if len(tm.Vehicles) != 1 {
		t.Fatalf("expected 1 vehicle after remove, got %d", len(tm.Vehicles))
	}
	if tm.Vehicles[0].ID != v2.ID {
		t.Error("expected Truck to remain after removing Car")
	}
}

// 3.5: WeaponInstance, UpgradeInstance, PerkInstance
func TestVehicle_AddWeapon(t *testing.T) {
	v := NewVehicle("Car")
	w := WeaponInstance{Name: "Machine Gun", Cost: 2, AttackDice: "2D6", Range: "Double", Slots: 1}
	v.Weapons = append(v.Weapons, w)
	if len(v.Weapons) != 1 {
		t.Fatalf("expected 1 weapon, got %d", len(v.Weapons))
	}
	if v.Weapons[0].Name != "Machine Gun" {
		t.Errorf("weapon name: got %q, want %q", v.Weapons[0].Name, "Machine Gun")
	}
}

func TestVehicle_AddUpgrade(t *testing.T) {
	v := NewVehicle("Truck")
	u := UpgradeInstance{Name: "Ram", Cost: 4, Slots: 1}
	v.Upgrades = append(v.Upgrades, u)
	if len(v.Upgrades) != 1 {
		t.Fatalf("expected 1 upgrade, got %d", len(v.Upgrades))
	}
}

func TestVehicle_AddPerk(t *testing.T) {
	v := NewVehicle("Car")
	p := PerkInstance{Name: "Battlehammer", Cost: 4, Class: "Aggression"}
	v.Perks = append(v.Perks, p)
	if len(v.Perks) != 1 {
		t.Fatalf("expected 1 perk, got %d", len(v.Perks))
	}
}

// 3.6: Variant and Notes
func TestVehicle_VariantAndNotes(t *testing.T) {
	v := NewVehicle("Car")
	v.Variant = "Prison Car"
	v.Notes = "Front line fighter"
	if v.Variant != "Prison Car" {
		t.Errorf("variant: got %q, want %q", v.Variant, "Prison Car")
	}

	tm := NewTeam()
	tm.Notes = "Aggressive team"
	if tm.Notes != "Aggressive team" {
		t.Errorf("team notes: got %q", tm.Notes)
	}
}

// 3.7: VehicleCost
func TestVehicleCost(t *testing.T) {
	v := NewVehicle("Car")
	v.Weapons = []WeaponInstance{
		{Name: "Machine Gun", Cost: 2, Slots: 1},
		{Name: "Handgun", Cost: 0, Slots: 0},
	}
	v.Upgrades = []UpgradeInstance{
		{Name: "Ram", Cost: 4, Slots: 1},
	}
	v.Perks = []PerkInstance{
		{Name: "Battlehammer", Cost: 4},
	}

	cost := VehicleCost(v)
	// Car base (12) + MG (2) + Ram (4) + Battlehammer (4) = 22
	expected := 12 + 2 + 4 + 4
	if cost != expected {
		t.Errorf("vehicle cost: got %d, want %d", cost, expected)
	}
}

func TestVehicleCost_CustomItem(t *testing.T) {
	v := NewVehicle("Car")
	v.Weapons = []WeaponInstance{
		{Name: "Custom Laser", Cost: 7, Slots: 2, IsCustom: true},
	}

	cost := VehicleCost(v)
	// Car base (12) + Custom (7) = 19
	if cost != 19 {
		t.Errorf("vehicle cost with custom: got %d, want %d", cost, 19)
	}
}

// 3.8: TeamCost and TeamHull
func TestTeamCost(t *testing.T) {
	tm := NewTeam()
	v1 := tm.AddVehicle("Car")
	v1.Weapons = []WeaponInstance{{Name: "Machine Gun", Cost: 2, Slots: 1}}
	tm.Vehicles[0] = *v1

	v2 := tm.AddVehicle("Truck")
	v2.Upgrades = []UpgradeInstance{{Name: "Ram", Cost: 4, Slots: 1}}
	tm.Vehicles[1] = *v2

	cost := TeamCost(tm)
	// Car (12+2) + Truck (15+4) = 33
	if cost != 33 {
		t.Errorf("team cost: got %d, want %d", cost, 33)
	}
}

func TestTeamHull(t *testing.T) {
	tm := NewTeam()
	tm.AddVehicle("Car")   // Hull 10
	tm.AddVehicle("Truck") // Hull 12

	hull := TeamHull(tm)
	if hull != 22 {
		t.Errorf("team hull: got %d, want %d", hull, 22)
	}
}

// 3.9: SlotsUsed and SlotsAvailable
func TestSlotsUsed(t *testing.T) {
	v := NewVehicle("Car")
	v.Weapons = []WeaponInstance{
		{Name: "Machine Gun", Cost: 2, Slots: 1},
	}
	v.Upgrades = []UpgradeInstance{
		{Name: "Ram", Cost: 4, Slots: 1},
	}

	used := SlotsUsed(v)
	if used != 2 {
		t.Errorf("slots used: got %d, want %d", used, 2)
	}
}

func TestSlotsAvailable(t *testing.T) {
	v := NewVehicle("Car") // 2 build slots
	v.Weapons = []WeaponInstance{
		{Name: "Machine Gun", Cost: 2, Slots: 1},
	}

	available := SlotsAvailable(v)
	// Car has 2 slots, 1 used = 1 available
	if available != 1 {
		t.Errorf("slots available: got %d, want %d", available, 1)
	}
}

// Ensure gamedata import is used
var _ = gamedata.GetVehicleType

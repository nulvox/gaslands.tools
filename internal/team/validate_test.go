package team

import (
	"strings"
	"testing"
)

// 5.1: ValidateBudget
func TestValidateBudget_UnderBudget(t *testing.T) {
	tm := NewTeam()
	tm.Budget = 50
	tm.AddVehicle("Buggy") // base cost 8

	warnings := ValidateBudget(tm)
	if len(warnings) != 0 {
		t.Errorf("expected no warnings, got %v", warnings)
	}
}

func TestValidateBudget_AtBudget(t *testing.T) {
	tm := NewTeam()
	tm.Budget = 12
	tm.AddVehicle("Car") // base cost 12

	warnings := ValidateBudget(tm)
	if len(warnings) != 0 {
		t.Errorf("expected no warnings at exact budget, got %v", warnings)
	}
}

func TestValidateBudget_OverBudget(t *testing.T) {
	tm := NewTeam()
	tm.Budget = 10
	tm.AddVehicle("Car") // base cost 12

	warnings := ValidateBudget(tm)
	if len(warnings) != 1 {
		t.Fatalf("expected 1 warning, got %d: %v", len(warnings), warnings)
	}
	if !strings.Contains(warnings[0], "2") {
		t.Errorf("expected warning to mention overage amount, got %q", warnings[0])
	}
}

// 5.2: ValidateBuildSlots
func TestValidateBuildSlots_UnderLimit(t *testing.T) {
	v := NewVehicle("Car") // 2 build slots
	v.Weapons = []WeaponInstance{
		{Name: "Machine Gun", Cost: 2, Slots: 1},
	}

	warnings := ValidateBuildSlots(v)
	if len(warnings) != 0 {
		t.Errorf("expected no warnings, got %v", warnings)
	}
}

func TestValidateBuildSlots_AtLimit(t *testing.T) {
	v := NewVehicle("Car") // 2 build slots
	v.Weapons = []WeaponInstance{
		{Name: "Machine Gun", Cost: 2, Slots: 1},
		{Name: "Rockets", Cost: 3, Slots: 1},
	}

	warnings := ValidateBuildSlots(v)
	if len(warnings) != 0 {
		t.Errorf("expected no warnings at slot limit, got %v", warnings)
	}
}

func TestValidateBuildSlots_OverLimit(t *testing.T) {
	v := NewVehicle("Car") // 2 build slots
	v.Weapons = []WeaponInstance{
		{Name: "Machine Gun", Cost: 2, Slots: 1},
		{Name: "Rockets", Cost: 3, Slots: 1},
		{Name: "Minigun", Cost: 4, Slots: 1},
	}

	warnings := ValidateBuildSlots(v)
	if len(warnings) != 1 {
		t.Fatalf("expected 1 warning, got %d: %v", len(warnings), warnings)
	}
}

// 5.3: ValidateSponsorPerks
func TestValidateSponsorPerks_ValidClass(t *testing.T) {
	tm := NewTeam()
	tm.Sponsor = SponsorSelection{Name: "The Warden"} // Aggression, Badass
	v := tm.AddVehicle("Car")
	v.Perks = []PerkInstance{
		{Name: "Battlehammer", Cost: 4, Class: "Aggression"},
	}
	tm.Vehicles[0] = *v

	warnings := ValidateSponsorPerks(tm)
	if len(warnings) != 0 {
		t.Errorf("expected no warnings for valid perk class, got %v", warnings)
	}
}

func TestValidateSponsorPerks_InvalidClass(t *testing.T) {
	tm := NewTeam()
	tm.Sponsor = SponsorSelection{Name: "The Warden"} // Aggression, Badass
	v := tm.AddVehicle("Car")
	v.Perks = []PerkInstance{
		{Name: "Hell For Leather", Cost: 5, Class: "Speed"},
	}
	tm.Vehicles[0] = *v

	warnings := ValidateSponsorPerks(tm)
	if len(warnings) != 1 {
		t.Fatalf("expected 1 warning for invalid perk class, got %d: %v", len(warnings), warnings)
	}
	if !strings.Contains(warnings[0], "Hell For Leather") {
		t.Errorf("expected warning to name the perk, got %q", warnings[0])
	}
	if !strings.Contains(warnings[0], "Speed") {
		t.Errorf("expected warning to name the class, got %q", warnings[0])
	}
}

func TestValidateSponsorPerks_CustomSponsorSkipped(t *testing.T) {
	tm := NewTeam()
	tm.Sponsor = SponsorSelection{Name: "Custom", IsCustom: true}
	v := tm.AddVehicle("Car")
	v.Perks = []PerkInstance{
		{Name: "Anything", Cost: 5, Class: "Unknown"},
	}
	tm.Vehicles[0] = *v

	warnings := ValidateSponsorPerks(tm)
	if len(warnings) != 0 {
		t.Errorf("expected no warnings for custom sponsor, got %v", warnings)
	}
}

func TestValidateSponsorPerks_NoSponsor(t *testing.T) {
	tm := NewTeam()
	v := tm.AddVehicle("Car")
	v.Perks = []PerkInstance{
		{Name: "Anything", Cost: 5, Class: "Speed"},
	}
	tm.Vehicles[0] = *v

	warnings := ValidateSponsorPerks(tm)
	// No sponsor set = no sponsor validation
	if len(warnings) != 0 {
		t.Errorf("expected no warnings when no sponsor set, got %v", warnings)
	}
}

// 5.4: ValidateVariants
func TestValidateVariants_ValidPrisonCar(t *testing.T) {
	v := NewVehicle("Car") // Middleweight
	v.Variant = "Prison Car"

	warnings := ValidateVariants(v)
	if len(warnings) != 0 {
		t.Errorf("expected no warnings for valid variant, got %v", warnings)
	}
}

func TestValidateVariants_InvalidPrisonCarOnLightweight(t *testing.T) {
	v := NewVehicle("Buggy") // Lightweight
	v.Variant = "Prison Car"

	warnings := ValidateVariants(v)
	if len(warnings) != 1 {
		t.Fatalf("expected 1 warning, got %d: %v", len(warnings), warnings)
	}
	if !strings.Contains(warnings[0], "Middleweight") {
		t.Errorf("expected warning to mention Middleweight, got %q", warnings[0])
	}
}

func TestValidateVariants_NoVariant(t *testing.T) {
	v := NewVehicle("Car")

	warnings := ValidateVariants(v)
	if len(warnings) != 0 {
		t.Errorf("expected no warnings for no variant, got %v", warnings)
	}
}

// 5.5: ValidateTeam
func TestValidateTeam_ValidTeam(t *testing.T) {
	tm := NewTeam()
	tm.Budget = 50
	tm.Sponsor = SponsorSelection{Name: "The Warden"}
	v := tm.AddVehicle("Car")
	v.Perks = []PerkInstance{
		{Name: "Battlehammer", Cost: 4, Class: "Aggression"},
	}
	tm.Vehicles[0] = *v

	warnings := ValidateTeam(tm)
	if len(warnings) != 0 {
		t.Errorf("expected no warnings for valid team, got %v", warnings)
	}
}

func TestValidateTeam_EmptyTeam(t *testing.T) {
	tm := NewTeam()

	warnings := ValidateTeam(tm)
	if len(warnings) != 1 {
		t.Fatalf("expected 1 warning for empty team, got %d: %v", len(warnings), warnings)
	}
	if !strings.Contains(warnings[0], "no vehicles") {
		t.Errorf("expected warning about no vehicles, got %q", warnings[0])
	}
}

func TestValidateTeam_MultipleIssues(t *testing.T) {
	tm := NewTeam()
	tm.Budget = 10 // will be over budget
	tm.Sponsor = SponsorSelection{Name: "The Warden"}
	v := tm.AddVehicle("Car") // base 12, already over 10
	v.Perks = []PerkInstance{
		{Name: "Hell For Leather", Cost: 5, Class: "Speed"}, // wrong class
	}
	v.Weapons = []WeaponInstance{
		{Name: "W1", Cost: 1, Slots: 1},
		{Name: "W2", Cost: 1, Slots: 1},
		{Name: "W3", Cost: 1, Slots: 1}, // over 2 slots
	}
	tm.Vehicles[0] = *v

	warnings := ValidateTeam(tm)
	// Should have: budget warning, slot warning, perk class warning
	if len(warnings) < 3 {
		t.Errorf("expected at least 3 warnings, got %d: %v", len(warnings), warnings)
	}
}

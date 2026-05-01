package team

import (
	"strings"
	"testing"
)

func TestExportHTML_ValidHTML(t *testing.T) {
	tm := NewTeam()
	tm.Name = "Test Team"
	tm.Budget = 50
	tm.Sponsor = SponsorSelection{Name: "The Warden"}
	v := tm.AddVehicle("Car")
	v.CustomName = "Iron Coffin"
	v.Role = "Brawler"
	v.Weapons = []WeaponInstance{
		{Name: "Machine Gun", Cost: 2, AttackDice: "2D6", Range: "Double", Slots: 1},
	}
	v.Upgrades = []UpgradeInstance{
		{Name: "Ram", Cost: 4, Slots: 1, Description: "Front ram"},
	}
	v.Perks = []PerkInstance{
		{Name: "Battlehammer", Cost: 4, Class: "Aggression"},
	}
	tm.Vehicles[0] = *v

	html, err := ExportHTML(tm)
	if err != nil {
		t.Fatalf("ExportHTML error: %v", err)
	}

	// Basic structure checks
	if !strings.Contains(html, "<!DOCTYPE html>") {
		t.Error("missing DOCTYPE")
	}
	if !strings.Contains(html, "</html>") {
		t.Error("missing closing html tag")
	}
	if !strings.Contains(html, "Test Team") {
		t.Error("missing team name")
	}
	if !strings.Contains(html, "Iron Coffin") {
		t.Error("missing vehicle name")
	}
	if !strings.Contains(html, "Machine Gun") {
		t.Error("missing weapon name")
	}
	if !strings.Contains(html, "Ram") {
		t.Error("missing upgrade name")
	}
	if !strings.Contains(html, "Battlehammer") {
		t.Error("missing perk name")
	}
	if !strings.Contains(html, "The Warden") {
		t.Error("missing sponsor name")
	}
}

func TestExportHTML_BudgetSummary(t *testing.T) {
	tm := NewTeam()
	tm.Name = "Budget Test"
	tm.Budget = 50
	tm.AddVehicle("Car") // base 12

	html, err := ExportHTML(tm)
	if err != nil {
		t.Fatalf("ExportHTML error: %v", err)
	}

	if !strings.Contains(html, "50") {
		t.Error("missing budget value")
	}
	if !strings.Contains(html, "12") {
		t.Error("missing spent value")
	}
}

func TestExportHTML_EmptyTeam(t *testing.T) {
	tm := NewTeam()
	tm.Name = "Empty"

	html, err := ExportHTML(tm)
	if err != nil {
		t.Fatalf("ExportHTML error: %v", err)
	}

	if !strings.Contains(html, "Empty") {
		t.Error("missing team name")
	}
}

func TestExportHTML_SponsorInfo(t *testing.T) {
	tm := NewTeam()
	tm.Name = "Sponsor Test"
	tm.Sponsor = SponsorSelection{Name: "Miyazaki"}
	tm.AddVehicle("Car")

	html, err := ExportHTML(tm)
	if err != nil {
		t.Fatalf("ExportHTML error: %v", err)
	}

	if !strings.Contains(html, "Miyazaki") {
		t.Error("missing sponsor name")
	}
	if !strings.Contains(html, "DARING") {
		t.Error("missing sponsor perk class")
	}
	if !strings.Contains(html, "Virtuoso") {
		t.Error("missing sponsor perk")
	}
}

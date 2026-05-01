// Package team provides the team and vehicle data model for Gaslands Refuelled.
//
// This package contains no syscall/js dependencies and is fully testable
// with standard Go tests.
package team

import (
	"crypto/rand"
	"fmt"

	"gaslands.tools/internal/gamedata"
)

// SponsorSelection represents the team's chosen sponsor.
type SponsorSelection struct {
	Name     string `json:"name"`
	IsCustom bool   `json:"isCustom,omitempty"`
}

// WeaponInstance represents a weapon equipped on a vehicle.
type WeaponInstance struct {
	Name         string `json:"name"`
	Cost         int    `json:"cost"`
	AttackDice   string `json:"attackDice,omitempty"`
	Range        string `json:"range,omitempty"`
	Slots        int    `json:"slots"`
	SpecialRules string `json:"specialRules,omitempty"`
	IsCustom     bool   `json:"isCustom,omitempty"`
}

// UpgradeInstance represents an upgrade installed on a vehicle.
type UpgradeInstance struct {
	Name        string `json:"name"`
	Cost        int    `json:"cost"`
	Slots       int    `json:"slots"`
	Description string `json:"description,omitempty"`
	IsCustom    bool   `json:"isCustom,omitempty"`
}

// PerkInstance represents a perk assigned to a vehicle's driver/crew.
type PerkInstance struct {
	Name        string `json:"name"`
	Cost        int    `json:"cost"`
	Class       string `json:"class,omitempty"`
	Description string `json:"description,omitempty"`
	IsCustom    bool   `json:"isCustom,omitempty"`
}

// Vehicle represents a single vehicle in a team roster.
type Vehicle struct {
	ID          string            `json:"id"`
	CustomName  string            `json:"customName,omitempty"`
	Role        string            `json:"role,omitempty"`
	VehicleType string            `json:"vehicleType"`
	Variant     string            `json:"variant,omitempty"`
	Weapons     []WeaponInstance  `json:"weapons"`
	Upgrades    []UpgradeInstance `json:"upgrades"`
	Perks       []PerkInstance    `json:"perks"`
	Notes       string            `json:"notes,omitempty"`
}

// Team represents a team roster.
type Team struct {
	ID       string           `json:"id"`
	Name     string           `json:"name"`
	Sponsor  SponsorSelection `json:"sponsor"`
	Budget   int              `json:"budget"`
	Vehicles []Vehicle        `json:"vehicles"`
	Notes    string           `json:"notes,omitempty"`
	Version  string           `json:"version"`
}

// newID generates a random UUID-like identifier.
func newID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// NewTeam creates a new team with default values.
func NewTeam() *Team {
	return &Team{
		ID:      newID(),
		Budget:  50,
		Version: "1.0",
	}
}

// NewVehicle creates a new vehicle of the given type.
func NewVehicle(vehicleType string) *Vehicle {
	return &Vehicle{
		ID:          newID(),
		VehicleType: vehicleType,
	}
}

// AddVehicle creates a new vehicle of the given type and adds it to the team.
func (t *Team) AddVehicle(vehicleType string) *Vehicle {
	v := NewVehicle(vehicleType)
	t.Vehicles = append(t.Vehicles, *v)
	return &t.Vehicles[len(t.Vehicles)-1]
}

// RemoveVehicle removes the vehicle with the given ID from the team.
func (t *Team) RemoveVehicle(id string) {
	for i, v := range t.Vehicles {
		if v.ID == id {
			t.Vehicles = append(t.Vehicles[:i], t.Vehicles[i+1:]...)
			return
		}
	}
}

// VehicleCost calculates the total cost of a vehicle including base cost,
// weapons, upgrades, and perks.
func VehicleCost(v *Vehicle) int {
	vt, ok := gamedata.GetVehicleType(v.VehicleType)
	cost := 0
	if ok {
		cost = vt.BaseCost
	}
	for _, w := range v.Weapons {
		cost += w.Cost
	}
	for _, u := range v.Upgrades {
		cost += u.Cost
	}
	for _, p := range v.Perks {
		cost += p.Cost
	}
	return cost
}

// TeamCost calculates the total cost of all vehicles in the team.
func TeamCost(t *Team) int {
	total := 0
	for i := range t.Vehicles {
		total += VehicleCost(&t.Vehicles[i])
	}
	return total
}

// TeamHull calculates the total hull points of all vehicles in the team.
func TeamHull(t *Team) int {
	total := 0
	for _, v := range t.Vehicles {
		vt, ok := gamedata.GetVehicleType(v.VehicleType)
		if ok {
			total += vt.Hull
		}
	}
	return total
}

// SlotsUsed calculates the total build slots consumed by weapons and upgrades.
func SlotsUsed(v *Vehicle) int {
	used := 0
	for _, w := range v.Weapons {
		used += w.Slots
	}
	for _, u := range v.Upgrades {
		used += u.Slots
	}
	return used
}

// SlotsAvailable calculates the remaining build slots for a vehicle.
func SlotsAvailable(v *Vehicle) int {
	vt, ok := gamedata.GetVehicleType(v.VehicleType)
	if !ok {
		return 0
	}
	return vt.BuildSlots - SlotsUsed(v)
}

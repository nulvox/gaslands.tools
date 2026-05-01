package team

import (
	"fmt"

	"gaslands.tools/internal/gamedata"
)

// ValidateBudget returns a warning if the team's total cost exceeds the budget.
func ValidateBudget(t *Team) []string {
	cost := TeamCost(t)
	if cost > t.Budget {
		over := cost - t.Budget
		return []string{fmt.Sprintf("Team is over budget by %d Cans (%d/%d)", over, cost, t.Budget)}
	}
	return nil
}

// ValidateBuildSlots returns a warning if the vehicle's slots used exceed available slots.
func ValidateBuildSlots(v *Vehicle) []string {
	vt, ok := gamedata.GetVehicleType(v.VehicleType)
	if !ok {
		return nil
	}
	used := SlotsUsed(v)
	if used > vt.BuildSlots {
		over := used - vt.BuildSlots
		return []string{fmt.Sprintf("%s: %d build slots over limit (%d/%d)", v.VehicleType, over, used, vt.BuildSlots)}
	}
	return nil
}

// ValidateSponsorPerks returns warnings for perks whose class is not in the sponsor's allowed classes.
func ValidateSponsorPerks(t *Team) []string {
	if t.Sponsor.Name == "" || t.Sponsor.IsCustom {
		return nil
	}
	sponsor, ok := gamedata.GetSponsor(t.Sponsor.Name)
	if !ok {
		return nil
	}

	allowed := make(map[string]bool)
	for _, c := range sponsor.PerkClasses {
		allowed[c] = true
	}

	var warnings []string
	for _, v := range t.Vehicles {
		for _, p := range v.Perks {
			if p.IsCustom {
				continue
			}
			if !allowed[p.Class] {
				warnings = append(warnings, fmt.Sprintf("Perk %q (class %s) is not allowed by sponsor %s", p.Name, p.Class, sponsor.Name))
			}
		}
	}
	return warnings
}

// ValidateVariants returns warnings if variant requirements are not met.
func ValidateVariants(v *Vehicle) []string {
	if v.Variant == "" {
		return nil
	}
	vt, ok := gamedata.GetVehicleType(v.VehicleType)
	if !ok {
		return nil
	}

	var warnings []string
	if v.Variant == "Prison Car" && vt.WeightClass != "Middleweight" {
		warnings = append(warnings, fmt.Sprintf("Prison Car variant requires Middleweight vehicle, %s is %s", v.VehicleType, vt.WeightClass))
	}
	return warnings
}

// ValidateTeam combines all validation checks and returns all warnings.
func ValidateTeam(t *Team) []string {
	var warnings []string

	if len(t.Vehicles) == 0 {
		warnings = append(warnings, "Team has no vehicles")
		return warnings
	}

	warnings = append(warnings, ValidateBudget(t)...)
	warnings = append(warnings, ValidateSponsorPerks(t)...)
	for i := range t.Vehicles {
		warnings = append(warnings, ValidateBuildSlots(&t.Vehicles[i])...)
		warnings = append(warnings, ValidateVariants(&t.Vehicles[i])...)
	}
	return warnings
}

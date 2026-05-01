package gamedata

// Upgrade represents a vehicle upgrade from the Gaslands Refuelled rulebook.
type Upgrade struct {
	Name         string
	Cost         int
	Slots        int
	Description  string
	Restrictions string
}

var upgrades = []Upgrade{
	{Name: "Armour Plating", Cost: 4, Slots: 1, Description: "+2 Hull Points"},
	{Name: "Experimental Nuclear Engine", Cost: 5, Slots: 0, Description: "+2 Max Gear (max 6). Long Straight permitted in any Gear. Electrical", Restrictions: "Mishkin only. Not Lightweight"},
	{Name: "Experimental Teleporter", Cost: 7, Slots: 0, Description: "Electrical. See special rules", Restrictions: "Mishkin only"},
	{Name: "Exploding Ram", Cost: 3, Slots: 0, Description: "Ammo 1. +6 attack dice on first Collision on declared facing"},
	{Name: "Extra Crewmember", Cost: 4, Slots: 0, Description: "+1 Crew, up to max of 2x starting Crew"},
	{Name: "Improvised Sludge Thrower", Cost: 2, Slots: 1, Description: "See special rules"},
	{Name: "Nitro Booster", Cost: 6, Slots: 0, Description: "Ammo 1. Forced Long Straight move forward"},
	{Name: "Ram", Cost: 4, Slots: 1, Description: "+2 Smash Attack dice on declared facing. No Hazard Tokens from that Collision"},
	{Name: "Roll Cage", Cost: 4, Slots: 1, Description: "May choose to ignore 2 hits from Flip"},
	{Name: "Tank Tracks", Cost: 4, Slots: 1, Description: "-1 Max Gear. +1 Handling. All Terrain"},
}

// GetUpgrade returns the upgrade with the given name.
func GetUpgrade(name string) (Upgrade, bool) {
	for _, u := range upgrades {
		if u.Name == name {
			return u, true
		}
	}
	return Upgrade{}, false
}

// ListUpgrades returns all upgrades.
func ListUpgrades() []Upgrade {
	result := make([]Upgrade, len(upgrades))
	copy(result, upgrades)
	return result
}

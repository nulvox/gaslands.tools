package gamedata

// Weapon represents a weapon from the Gaslands Refuelled rulebook.
type Weapon struct {
	Name         string
	Cost         int
	AttackDice   string
	Range        string
	Slots        int
	SpecialRules string
}

var weapons = []Weapon{
	// Basic Weapons
	{Name: "Handgun", Cost: 0, AttackDice: "1D6", Range: "Medium", Slots: 0, SpecialRules: "Crew Fired"},
	{Name: "Machine Gun", Cost: 2, AttackDice: "2D6", Range: "Double", Slots: 1},
	{Name: "Heavy Machine Gun", Cost: 3, AttackDice: "3D6", Range: "Double", Slots: 1},
	{Name: "Minigun", Cost: 5, AttackDice: "4D6", Range: "Double", Slots: 1},

	// Advanced Weapons
	{Name: "125mm Cannon", Cost: 6, AttackDice: "8D6", Range: "Double", Slots: 3, SpecialRules: "Ammo 3. Blast"},
	{Name: "Arc Lightning Projector", Cost: 6, AttackDice: "6D6", Range: "Double", Slots: 2, SpecialRules: "Ammo 1. Electrical. Mishkin only"},
	{Name: "Bazooka", Cost: 4, AttackDice: "3D6", Range: "Double", Slots: 2, SpecialRules: "Ammo 3. Blast"},
	{Name: "BFG", Cost: 1, AttackDice: "10D6", Range: "Double", Slots: 3, SpecialRules: "Ammo 1"},
	{Name: "Combat Laser", Cost: 5, AttackDice: "3D6", Range: "Double", Slots: 1, SpecialRules: "Splash"},
	{Name: "Death Ray", Cost: 3, AttackDice: "3D6", Range: "Double", Slots: 1, SpecialRules: "Ammo 1. Electrical. Mishkin only"},
	{Name: "Flamethrower", Cost: 4, AttackDice: "6D6", Range: "Large Burst", Slots: 2, SpecialRules: "Ammo 3. Splash. Fire. Indirect"},
	{Name: "Grabber Arm", Cost: 6, AttackDice: "3D6", Range: "Short", Slots: 1},
	{Name: "Grav Gun", Cost: 2, AttackDice: "(3D6)", Range: "Double", Slots: 1, SpecialRules: "Ammo 1. Electrical. Mishkin only"},
	{Name: "Harpoon", Cost: 2, AttackDice: "(5D6)", Range: "Double", Slots: 1},
	{Name: "Kinetic Super Booster", Cost: 6, AttackDice: "(6D6)", Range: "Double", Slots: 2, SpecialRules: "Ammo 1. Electrical. Mishkin only"},
	{Name: "Magnetic Jammer", Cost: 2, AttackDice: "-", Range: "Double", Slots: 0, SpecialRules: "Electrical. Mishkin only"},
	{Name: "Mortar", Cost: 4, AttackDice: "4D6", Range: "Double", Slots: 1, SpecialRules: "Ammo 3. Indirect"},
	{Name: "Rockets", Cost: 5, AttackDice: "6D6", Range: "Double", Slots: 2, SpecialRules: "Ammo 3"},
	{Name: "Thumper", Cost: 4, AttackDice: "-", Range: "Medium", Slots: 2, SpecialRules: "Ammo 1. Electrical. Indirect. 360-degree. Mishkin only"},
	{Name: "Wall of Amplifiers", Cost: 4, AttackDice: "-", Range: "Medium", Slots: 3, SpecialRules: "360-degree arc of fire"},
	{Name: "Wreck Lobber", Cost: 4, AttackDice: "-", Range: "Double/Dropped", Slots: 4, SpecialRules: "Ammo 3"},
	{Name: "Wrecking Ball", Cost: 2, AttackDice: "*", Range: "Short", Slots: 3, SpecialRules: "See special rules"},

	// Crew Fired Weapons
	{Name: "Blunderbuss", Cost: 2, AttackDice: "2D6", Range: "Small Burst", Slots: 0, SpecialRules: "Crew Fired. Splash"},
	{Name: "Gas Grenades", Cost: 1, AttackDice: "(1D6)", Range: "Medium", Slots: 0, SpecialRules: "Ammo 5. Crew Fired. Indirect. Blitz"},
	{Name: "Grenades", Cost: 1, AttackDice: "1D6", Range: "Medium", Slots: 0, SpecialRules: "Ammo 5. Crew Fired. Blast. Indirect. Blitz"},
	{Name: "Magnum", Cost: 3, AttackDice: "1D6", Range: "Double", Slots: 0, SpecialRules: "Crew Fired. Blast"},
	{Name: "Molotov Cocktails", Cost: 1, AttackDice: "1D6", Range: "Medium", Slots: 0, SpecialRules: "Ammo 5. Crew Fired. Fire. Indirect. Blitz"},
	{Name: "Shotgun", Cost: 4, AttackDice: "*", Range: "Long", Slots: 0, SpecialRules: "Crew Fired"},
	{Name: "Steel Nets", Cost: 2, AttackDice: "(3D6)", Range: "Short", Slots: 0, SpecialRules: "Crew Fired. Blast"},
	{Name: "Submachine Gun", Cost: 5, AttackDice: "3D6", Range: "Medium", Slots: 0, SpecialRules: "Crew Fired"},

	// Dropped Weapons
	{Name: "Caltrop Dropper", Cost: 1, AttackDice: "2D6", Range: "Dropped", Slots: 1, SpecialRules: "Ammo 3. Small Burst"},
	{Name: "Glue Dropper", Cost: 1, AttackDice: "-", Range: "Dropped", Slots: 1, SpecialRules: "Ammo 1"},
	{Name: "Mine Dropper", Cost: 1, AttackDice: "4D6", Range: "Dropped", Slots: 1, SpecialRules: "Ammo 3. Small Burst. Blast"},
	{Name: "Napalm Dropper", Cost: 1, AttackDice: "4D6", Range: "Dropped", Slots: 1, SpecialRules: "Ammo 3. Small Burst. Fire"},
	{Name: "Oil Slick Dropper", Cost: 2, AttackDice: "-", Range: "Dropped", Slots: 0, SpecialRules: "Ammo 3"},
	{Name: "RC Car Bombs", Cost: 3, AttackDice: "4D6", Range: "Dropped", Slots: 0, SpecialRules: "Ammo 3"},
	{Name: "Sentry Gun", Cost: 3, AttackDice: "2D6", Range: "Dropped", Slots: 0, SpecialRules: "Ammo 3"},
	{Name: "Smoke Dropper", Cost: 1, AttackDice: "-", Range: "Dropped", Slots: 0, SpecialRules: "Ammo 3"},
}

// GetWeapon returns the weapon with the given name.
func GetWeapon(name string) (Weapon, bool) {
	for _, w := range weapons {
		if w.Name == name {
			return w, true
		}
	}
	return Weapon{}, false
}

// ListWeapons returns all weapons.
func ListWeapons() []Weapon {
	result := make([]Weapon, len(weapons))
	copy(result, weapons)
	return result
}

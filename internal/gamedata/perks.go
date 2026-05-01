package gamedata

// Perk represents a driver/crew perk from the Gaslands Refuelled rulebook.
type Perk struct {
	Name        string
	Cost        int
	Class       string
	Description string
}

var perks = []Perk{
	// Aggression
	{Name: "Double-Barreled", Cost: 2, Class: "Aggression", Description: "Up to three Crew Fired weapons gain +1 to hit"},
	{Name: "Boarding Party", Cost: 2, Class: "Aggression", Description: "Ignores Distracted rule"},
	{Name: "Battlehammer", Cost: 4, Class: "Aggression", Description: "+1 Smash Attack die per Hazard Token held"},
	{Name: "Terrifying Lunatic", Cost: 5, Class: "Aggression", Description: "Enemy ending Movement Step within Short range gains Hazard Token"},
	{Name: "Grinderman", Cost: 5, Class: "Aggression", Description: "Smash Attack hits add Hazard Tokens instead of removing Hull Points"},
	{Name: "Murder Tractor", Cost: 5, Class: "Aggression", Description: "May make piledriver attacks like a War Rig"},

	// Badass
	{Name: "Powder Keg", Cost: 1, Class: "Badass", Description: "+1 to explosion check. Counts as one weight-class heavier when exploding"},
	{Name: "Crowd Pleaser", Cost: 1, Class: "Badass", Description: "If this vehicle wipes out, gain 1 Audience Vote"},
	{Name: "Road Warrior", Cost: 2, Class: "Badass", Description: "Remove a Hazard Token after causing hits on enemy"},
	{Name: "Cover Me", Cost: 2, Class: "Badass", Description: "Transfer a Hazard Token to friendly vehicle within Double range"},
	{Name: "Madman", Cost: 3, Class: "Badass", Description: "At 4+ Hazard Tokens, transfer one to vehicle within Medium range"},
	{Name: "Bullet-Time", Cost: 3, Class: "Badass", Description: "After Slide, one weapon counts as Turret-mounted"},

	// Built
	{Name: "Dead Weight", Cost: 2, Class: "Built", Description: "Gain 2 Hazard Tokens to count as heavier weight class"},
	{Name: "Barrel Roll", Cost: 2, Class: "Built", Description: "Choose Flip template placement from side or rear edge"},
	{Name: "Bruiser", Cost: 4, Class: "Built", Description: "Enemy gains Hazard Token in Collision if you don't Evade"},
	{Name: "Splashback", Cost: 5, Class: "Built", Description: "When losing Hull Points, 1D6 attack vs vehicles within Medium range"},
	{Name: "Crusher", Cost: 7, Class: "Built", Description: "Gains Up and Over special rule"},
	{Name: "Feel No Pain", Cost: 8, Class: "Built", Description: "Cancel all hits if 2 or fewer uncancelled hits from attack"},

	// Daring
	{Name: "Chrome-Whisperer", Cost: 2, Class: "Daring", Description: "May Push It any number of times per Movement Step"},
	{Name: "Slippery", Cost: 3, Class: "Daring", Description: "-2 Smash Attack dice for attackers targeting this vehicle"},
	{Name: "Handbrake Artist", Cost: 3, Class: "Daring", Description: "Choose any direction when applying Spin result"},
	{Name: "Evasive", Cost: 5, Class: "Daring", Description: "Gain Hazard Tokens for +1 per Evade die"},
	{Name: "Powerslide", Cost: 5, Class: "Daring", Description: "Use any template except Long Straight for Slide"},
	{Name: "Stunt Driver", Cost: 7, Class: "Daring", Description: "Ignore obstructions during Movement Step, then gain 3 Hazard Tokens"},

	// Horror
	{Name: "Purifying Flames", Cost: 1, Class: "Horror", Description: "Suffer damage to repair Hull Points on friendly vehicle (Fire)"},
	{Name: "Ecstatic Visions", Cost: 1, Class: "Horror", Description: "Gain Hazard Tokens to discard from friendly vehicle"},
	{Name: "Sympathy For The Devil", Cost: 1, Class: "Horror", Description: "Add friendly vehicle's Gear to Evade check"},
	{Name: "Highway To Hell", Cost: 2, Class: "Horror", Description: "Suffer damage on straight to leave Napalm template"},
	{Name: "Violent Manifestation", Cost: 3, Class: "Horror", Description: "On respawn, explode against vehicles within Medium range"},
	{Name: "Angel Of Death", Cost: 4, Class: "Horror", Description: "Suffer damage to add attack dice to weapon"},

	// Military
	{Name: "Dead-Eye", Cost: 2, Class: "Military", Description: "+1 to hit at Double range but not Medium"},
	{Name: "Loader", Cost: 2, Class: "Military", Description: "Reduce Crew by 1 for +1 to hit with one weapon"},
	{Name: "Fully Loaded", Cost: 2, Class: "Military", Description: "+1 attack die if weapon has 3+ Ammo Tokens"},
	{Name: "Rapid Fire", Cost: 2, Class: "Military", Description: "Additional Attack Step with same weapon once per round"},
	{Name: "Headshot", Cost: 4, Class: "Military", Description: "Critical Hits inflict 3 hits instead of 2"},
	{Name: "Return Fire", Cost: 5, Class: "Military", Description: "Take 2 Hazard Tokens to attack when targeted by shooting"},

	// Precision
	{Name: "Mister Fahrenheit", Cost: 2, Class: "Precision", Description: "Max 2 Hazard Tokens from Collisions per activation"},
	{Name: "Moment Of Glory", Cost: 2, Class: "Precision", Description: "Once per game, change any number of Skid Dice results"},
	{Name: "Restraint", Cost: 2, Class: "Precision", Description: "Remove Hazard Token instead of gaining one when shifting down"},
	{Name: "Expertise", Cost: 3, Class: "Precision", Description: "+1 Handling Value"},
	{Name: "Trick Driving", Cost: 3, Class: "Precision", Description: "Select template as if Gear was one higher or lower"},
	{Name: "Easy Rider", Cost: 5, Class: "Precision", Description: "Discard one Skid Die result once per round"},

	// Pursuit
	{Name: "On Your Tail", Cost: 2, Class: "Pursuit", Description: "Enemy gains Hazard Token when Spin/Slide ends within Short range"},
	{Name: "Schadenfreude", Cost: 2, Class: "Pursuit", Description: "Remove all Hazard Tokens when vehicle wipes out within Short range"},
	{Name: "Taunt", Cost: 2, Class: "Pursuit", Description: "Place Skid Die result on target within Short range"},
	{Name: "Out Run", Cost: 2, Class: "Pursuit", Description: "Vehicles in lower Gear within Short range gain Hazard Token"},
	{Name: "Pit", Cost: 4, Class: "Pursuit", Description: "PIT reaction forces target to move using Hazardous template"},
	{Name: "Unnerving Eye Contact", Cost: 5, Class: "Pursuit", Description: "Enemies within Short range can't use Shift to remove Hazard Tokens"},

	// Reckless
	{Name: "Drive Angry", Cost: 1, Class: "Reckless", Description: "Gain 1 Hazard Token at start of activation"},
	{Name: "Hog Wild", Cost: 2, Class: "Reckless", Description: "+2 Smash Attack dice during Wipeout Step Collision"},
	{Name: "In For A Penny", Cost: 2, Class: "Reckless", Description: "Double Smash Attack dice at 6+ Hazard Tokens this activation"},
	{Name: "Don't Come Knocking", Cost: 4, Class: "Reckless", Description: "Gain 4 Hazard Tokens; can't gain or lose any until next activation"},
	{Name: "Bigger'n You", Cost: 4, Class: "Reckless", Description: "Double weight-class Smash Attack bonuses/penalties"},
	{Name: "Beerserker", Cost: 5, Class: "Reckless", Description: "Reduce damage outside activation by 1 (min 1)"},

	// Speed
	{Name: "Hot Start", Cost: 1, Class: "Speed", Description: "Start game in random Gear (D6)"},
	{Name: "Slipstream", Cost: 2, Class: "Speed", Description: "Slipstream reaction on tailgate Collision; no Hazard Tokens"},
	{Name: "Overload", Cost: 2, Class: "Speed", Description: "Roll extra Skid Die; must change up or gain Hazard Token"},
	{Name: "Downshift", Cost: 3, Class: "Speed", Description: "Forced Short Straight after changing down Gears"},
	{Name: "Time Extended!", Cost: 3, Class: "Speed", Description: "Remove any Hazard Tokens after passing race gate"},
	{Name: "Hell For Leather", Cost: 5, Class: "Speed", Description: "Long Straight permitted in any Gear"},

	// Technology
	{Name: "Rocket Thrusters", Cost: 1, Class: "Technology", Description: "Choose Flip template from Long Straight, Veer, or Gentle"},
	{Name: "Whizbang", Cost: 1, Class: "Technology", Description: "Gain random Speed perk each game"},
	{Name: "Gyroscope", Cost: 1, Class: "Technology", Description: "Gain random Daring perk each game"},
	{Name: "Satellite Navigation", Cost: 2, Class: "Technology", Description: "Set aside Shift results for team use later"},
	{Name: "Mobile Mechanic", Cost: 3, Class: "Technology", Description: "Reduce Crew by 1 to repair 1 Hull Point"},
	{Name: "Eureka!", Cost: 4, Class: "Technology", Description: "Once per game, declare any weapon for one attack"},

	// Tuning
	{Name: "Fenderkiss", Cost: 2, Class: "Tuning", Description: "-2 Smash Attack dice for both vehicles"},
	{Name: "Rear Drive", Cost: 2, Class: "Tuning", Description: "Pivot about front edge for Spin results"},
	{Name: "Delicate Touch", Cost: 3, Class: "Tuning", Description: "Ignore hazard icons on movement templates"},
	{Name: "Momentum", Cost: 3, Class: "Tuning", Description: "Set aside Slide/Spin to re-roll Skid Die"},
	{Name: "Purring", Cost: 6, Class: "Tuning", Description: "Max 1 Hazard Token from each of Spin, Slide, Hazard per activation"},
	{Name: "Skiing", Cost: 6, Class: "Tuning", Description: "Take 3 Hazard Tokens to be ignored by other vehicles"},
}

// GetPerk returns the perk with the given name.
func GetPerk(name string) (Perk, bool) {
	for _, p := range perks {
		if p.Name == name {
			return p, true
		}
	}
	return Perk{}, false
}

// ListPerks returns all perks.
func ListPerks() []Perk {
	result := make([]Perk, len(perks))
	copy(result, perks)
	return result
}

// ListPerksByClass returns all perks in the given class.
func ListPerksByClass(class string) []Perk {
	var result []Perk
	for _, p := range perks {
		if p.Class == class {
			result = append(result, p)
		}
	}
	return result
}

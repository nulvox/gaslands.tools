package gamedata

// SponsorPerk represents a unique perk granted by a sponsor.
type SponsorPerk struct {
	Name        string
	Description string
}

// Sponsor represents a team sponsor from the Gaslands Refuelled rulebook.
type Sponsor struct {
	Name         string
	FlavorText   string
	PerkClasses  []string
	SponsorPerks []SponsorPerk
}

var sponsors = []Sponsor{
	{
		Name:        "Rutherford",
		FlavorText:  "Grant Rutherford is the son of a militaristic, American, oil baron. His teams gain access to military surplus, missile launchers, tanks, helicopters, and as much ammo as they can carry.",
		PerkClasses: []string{"Badass", "Military"},
		SponsorPerks: []SponsorPerk{
			{Name: "Military Hardware", Description: "This team may purchase a single Tank and a single Helicopter."},
			{Name: "Well Stocked", Description: "Weapons with ammo 3 have ammo 4 instead."},
			{Name: "Might Is Right", Description: "This team may not purchase Lightweight vehicles."},
			{Name: "Televised Carnage", Description: "6+ hits in one Attack Step gains 1 Audience Vote."},
		},
	},
	{
		Name:        "Miyazaki",
		FlavorText:  "Yuri Miyazaki grew up in the rubble of Tokyo, fighting her way to the top of the speedway circuit. Her drivers are unsurpassed in their skill and finesse.",
		PerkClasses: []string{"Daring", "Precision"},
		SponsorPerks: []SponsorPerk{
			{Name: "Virtuoso", Description: "First Push It per activation costs no Hazard Token."},
			{Name: "Elegance", Description: "May not purchase vehicles with base Handling 2 or lower."},
			{Name: "Showing Off", Description: "Gain Showing Off token for resolving Spin, Slide, and Gear change without Wipeout. All tokens = Audience Votes."},
		},
	},
	{
		Name:        "Mishkin",
		FlavorText:  "Andre Mishkin proved that technology is just as solid an answer as skill or ferocity on the track. He sends designs for unusual weapons and hi-tech vehicles from Mars.",
		PerkClasses: []string{"Military", "Technology"},
		SponsorPerks: []SponsorPerk{
			{Name: "Thumpermonkey", Description: "This team may purchase electrical weapons and upgrades."},
			{Name: "Dynamo", Description: "After activating in Gear Phase 4+, add +1 Ammo Token to an electrical weapon or upgrade."},
			{Name: "All the Toys", Description: "First use of a named weapon in the game gains 1 Audience Vote."},
		},
	},
	{
		Name:        "Idris",
		FlavorText:  "Yandi Idris was an addict of speed. The Cult of Speed spread like wildfire after his meteoric rise. He crossed the finishing line in a ball of fire and was never found.",
		PerkClasses: []string{"Precision", "Speed"},
		SponsorPerks: []SponsorPerk{
			{Name: "N2O Addict", Description: "Purchase Nitro at half cost."},
			{Name: "Speed Demon", Description: "Nitro Hazard Tokens cap at 3 instead of 5."},
			{Name: "Cult of Speed", Description: "Long Straight in Gear Phase 1-3 gains 1 Audience Vote."},
			{Name: "Kiss My Asphalt", Description: "May not purchase Gyrocopters."},
		},
	},
	{
		Name:        "Slime",
		FlavorText:  "Slime rules a wild and feral city in the Australian wastes known as Anarchy. The wild-eyed gangs are led by Slime's henchwomen, the Chooks.",
		PerkClasses: []string{"Tuning", "Reckless"},
		SponsorPerks: []SponsorPerk{
			{Name: "Pinball", Description: "Side-edge Collision with Smash Attack grants another Movement Step."},
			{Name: "Spiked Fist", Description: "Ram upgrade requires zero build slots."},
			{Name: "Live Fast", Description: "More Hazard Tokens than Hull Points at Wipeout Step gains 1 Audience Vote."},
		},
	},
	{
		Name:        "The Warden",
		FlavorText:  "Warden Cadeila runs the Sao Paulo People's Penitentiary. Her prisoners are granted a shot at freedom — welded into solid steel \"coffin cars\" and sent out to race.",
		PerkClasses: []string{"Aggression", "Badass"},
		SponsorPerks: []SponsorPerk{
			{Name: "Prison Cars", Description: "Middleweight vehicles may reduce base cost by 4 Cans (min 5) and Hull by 2. Once per vehicle, 0 build slots."},
			{Name: "Fireworks", Description: "If this vehicle Explodes, gain 1 Audience Vote (Middleweight) or 2 (Heavyweight), plus discard all Ammo Tokens."},
		},
	},
	{
		Name:        "Scarlett Annie",
		FlavorText:  "A dashing and flamboyant buccaneer, Scarlett Annie's cult following comes from her association with the \"Death Valley Death Run\" documentary TV series.",
		PerkClasses: []string{"Tuning", "Aggression"},
		SponsorPerks: []SponsorPerk{
			{Name: "Crew Quarters", Description: "Purchase Extra Crewmember upgrade at half cost."},
			{Name: "Raiders", Description: "Reduce Crew to remove Hull Points from vehicles in base contact."},
			{Name: "Raise the Sails", Description: "Reduce Crew by 1 to add 1 Shift to Skid Dice."},
			{Name: "Press Gang or Keelhaul", Description: "When contacted vehicle is Wrecked, gain 1 crew or 2 Audience Votes."},
		},
	},
	{
		Name:        "Highway Patrol",
		FlavorText:  "The Highway Patrol are the last law in a world gone crazy. Their only power: a badge of bronze. Their only weapon: 600 horsepower of fuel-injected steel.",
		PerkClasses: []string{"Speed", "Pursuit"},
		SponsorPerks: []SponsorPerk{
			{Name: "Hot Pursuit", Description: "Nominate one enemy vehicle as the bogey at game start."},
			{Name: "Bogey at 12 o'Clock", Description: "If bogey is in front arc beyond Double range, resolve another Movement Step."},
			{Name: "Siren", Description: "If in bogey's rear arc, bogey must reduce Gear by 1 or gain 2 Hazard Tokens."},
			{Name: "Steel Justice", Description: "Bogey wipeout = 2 Audience Votes. Bogey Wrecked = 4 Audience Votes."},
		},
	},
	{
		Name:        "Verney",
		FlavorText:  "The newly-freed Verney specialises in building unique Frankenstein's monsters of vehicles for anyone who can afford his high-quality customs.",
		PerkClasses: []string{"Technology", "Built"},
		SponsorPerks: []SponsorPerk{
			{Name: "MicroPlate Armour", Description: "Purchase MicroPlate Armour upgrade: 6 Cans, +2 Hull, 0 build slots."},
			{Name: "Trunk of Junk", Description: "Attack with any number of dropped weapons in a single activation."},
			{Name: "Tombstone", Description: "+1 Evade if shot from rear. Gain 2 Hazard Tokens to make all Collisions Head-on."},
			{Name: "That's Entertainment", Description: "Gain 1 Audience Vote when a dropped weapon template is removed."},
		},
	},
	{
		Name:        "Maxxine",
		FlavorText:  "Maxxine is the face of The Black Swans. It's ballet, but the dancers weigh 4,000 pounds and are dripping in engine oil.",
		PerkClasses: []string{"Tuning", "Pursuit"},
		SponsorPerks: []SponsorPerk{
			{Name: "Dizzy", Description: "Resolve any number of Spin results separately, allowing >90 degree Spins."},
			{Name: "Maxximum Drift", Description: "2 Slides = Medium Straight template. 3+ Slides = Long Straight."},
			{Name: "Meshuggah", Description: "Slide/Spin ending near friendly vehicle without Collision gains 1 Audience Vote."},
		},
	},
	{
		Name:        "The Order of the Inferno",
		FlavorText:  "Yandi Idris is not dead. He rides on in the living flame. Only by knowing the flames can we know true freedom.",
		PerkClasses: []string{"Horror", "Speed"},
		SponsorPerks: []SponsorPerk{
			{Name: "Fire Walk With Me", Description: "Reduce Fire damage by up to 3 (min 1)."},
			{Name: "Burning Man", Description: "+1 to all Evade dice while On Fire."},
			{Name: "Cult of Flame", Description: "If more enemy vehicles On Fire than friendly, gain 1 Audience Vote per friendly On Fire."},
		},
	},
	{
		Name:        "Beverly",
		FlavorText:  "Beverly was a stupid story told to scare children. She wasn't real. The Devil on the Highway.",
		PerkClasses: []string{"Horror", "Built"},
		SponsorPerks: []SponsorPerk{
			{Name: "Graveyard Shift", Description: "All vehicles except one must gain Ghost Rider at game start."},
			{Name: "Ghost Rider", Description: "Ignored by and ignores other vehicles. No Collisions, no shooting."},
			{Name: "Soul Anchor", Description: "If all vehicles have Ghost Rider, remove all from play."},
			{Name: "At the Crossroads", Description: "Respawn for 1 Audience Vote, but vehicle gains Ghost Rider."},
			{Name: "Inexorable", Description: "May respawn even if other rules prevent it."},
			{Name: "Soul Harvest", Description: "Contact enemy = Soul Token. Contact friendly = trade for Audience Votes or Hull repairs."},
		},
	},
	{
		Name:        "Rusty's Bootleggers",
		FlavorText:  "Zeke Rusty and his boys run moonshine. Their stills are volatile, their vehicles ramshackle, but the liquor is fine.",
		PerkClasses: []string{"Reckless", "Built"},
		SponsorPerks: []SponsorPerk{
			{Name: "Party Hard", Description: "Gain Audience Votes for having more Hazard Tokens than nearby enemies combined."},
			{Name: "Dutch Courage", Description: "Wipeout at 8 Hazard Tokens instead of 6."},
			{Name: "As Straight as I'm Able", Description: "No Hazard Token from articulated rule on non-Straight templates."},
			{Name: "Over the Limit", Description: "Straight templates never permitted. Veer is permitted and Trivial in any Gear."},
			{Name: "Trailer Trash", Description: "May purchase Trailers. Must include trailer-equipped vehicle or War Rig."},
			{Name: "Haulage", Description: "Each trailer/War Rig gets one free trailer cargo upgrade."},
		},
	},
}

// GetSponsor returns the sponsor with the given name.
func GetSponsor(name string) (Sponsor, bool) {
	for _, s := range sponsors {
		if s.Name == name {
			return s, true
		}
	}
	return Sponsor{}, false
}

// ListSponsors returns all sponsors.
func ListSponsors() []Sponsor {
	result := make([]Sponsor, len(sponsors))
	copy(result, sponsors)
	return result
}

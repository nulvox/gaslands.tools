package gamedata

import "testing"

func TestGetVehicleType_Car(t *testing.T) {
	vt, ok := GetVehicleType("Car")
	if !ok {
		t.Fatal("expected to find Car vehicle type")
	}
	if vt.WeightClass != "Middleweight" {
		t.Errorf("Car weight class: got %q, want %q", vt.WeightClass, "Middleweight")
	}
	if vt.Hull != 10 {
		t.Errorf("Car hull: got %d, want %d", vt.Hull, 10)
	}
	if vt.Handling != 3 {
		t.Errorf("Car handling: got %d, want %d", vt.Handling, 3)
	}
	if vt.MaxGear != 5 {
		t.Errorf("Car max gear: got %d, want %d", vt.MaxGear, 5)
	}
	if vt.Crew != 2 {
		t.Errorf("Car crew: got %d, want %d", vt.Crew, 2)
	}
	if vt.BuildSlots != 2 {
		t.Errorf("Car build slots: got %d, want %d", vt.BuildSlots, 2)
	}
	if vt.BaseCost != 12 {
		t.Errorf("Car base cost: got %d, want %d", vt.BaseCost, 12)
	}
}

func TestGetVehicleType_Truck(t *testing.T) {
	vt, ok := GetVehicleType("Truck")
	if !ok {
		t.Fatal("expected to find Truck vehicle type")
	}
	if vt.WeightClass != "Middleweight" {
		t.Errorf("Truck weight class: got %q, want %q", vt.WeightClass, "Middleweight")
	}
	if vt.Hull != 12 {
		t.Errorf("Truck hull: got %d, want %d", vt.Hull, 12)
	}
	if vt.Handling != 2 {
		t.Errorf("Truck handling: got %d, want %d", vt.Handling, 2)
	}
	if vt.MaxGear != 4 {
		t.Errorf("Truck max gear: got %d, want %d", vt.MaxGear, 4)
	}
	if vt.Crew != 3 {
		t.Errorf("Truck crew: got %d, want %d", vt.Crew, 3)
	}
	if vt.BuildSlots != 3 {
		t.Errorf("Truck build slots: got %d, want %d", vt.BuildSlots, 3)
	}
	if vt.BaseCost != 15 {
		t.Errorf("Truck base cost: got %d, want %d", vt.BaseCost, 15)
	}
}

func TestGetVehicleType_NotFound(t *testing.T) {
	_, ok := GetVehicleType("Spaceship")
	if ok {
		t.Error("expected Spaceship to not be found")
	}
}

// Weapon tests

func TestGetWeapon_HeavyMachineGun(t *testing.T) {
	w, ok := GetWeapon("Heavy Machine Gun")
	if !ok {
		t.Fatal("expected to find Heavy Machine Gun")
	}
	if w.Cost != 3 {
		t.Errorf("HMG cost: got %d, want %d", w.Cost, 3)
	}
	if w.AttackDice != "3D6" {
		t.Errorf("HMG dice: got %q, want %q", w.AttackDice, "3D6")
	}
}

func TestGetWeapon_HandgunFree(t *testing.T) {
	w, ok := GetWeapon("Handgun")
	if !ok {
		t.Fatal("expected to find Handgun")
	}
	if w.Cost != 0 {
		t.Errorf("Handgun cost: got %d, want %d", w.Cost, 0)
	}
}

// Upgrade tests

func TestGetUpgrade_Ram(t *testing.T) {
	u, ok := GetUpgrade("Ram")
	if !ok {
		t.Fatal("expected to find Ram")
	}
	if u.Cost != 4 {
		t.Errorf("Ram cost: got %d, want %d", u.Cost, 4)
	}
	if u.Slots != 1 {
		t.Errorf("Ram slots: got %d, want %d", u.Slots, 1)
	}
}

// Perk tests

func TestGetPerk_Battlehammer(t *testing.T) {
	p, ok := GetPerk("Battlehammer")
	if !ok {
		t.Fatal("expected to find Battlehammer")
	}
	if p.Class != "Aggression" {
		t.Errorf("Battlehammer class: got %q, want %q", p.Class, "Aggression")
	}
	if p.Cost != 4 {
		t.Errorf("Battlehammer cost: got %d, want %d", p.Cost, 4)
	}
}

func TestListPerksByClass_Aggression(t *testing.T) {
	perks := ListPerksByClass("Aggression")
	if len(perks) == 0 {
		t.Fatal("expected Aggression perks, got none")
	}
	found := false
	for _, p := range perks {
		if p.Name == "Battlehammer" {
			found = true
		}
		if p.Class != "Aggression" {
			t.Errorf("perk %q has class %q, expected Aggression", p.Name, p.Class)
		}
	}
	if !found {
		t.Error("expected Battlehammer in Aggression perks")
	}
}

// Sponsor tests

func TestGetSponsor_TheWarden(t *testing.T) {
	s, ok := GetSponsor("The Warden")
	if !ok {
		t.Fatal("expected to find The Warden")
	}
	if len(s.PerkClasses) != 2 {
		t.Fatalf("Warden perk classes: got %d, want 2", len(s.PerkClasses))
	}
	classes := map[string]bool{}
	for _, c := range s.PerkClasses {
		classes[c] = true
	}
	if !classes["Aggression"] || !classes["Badass"] {
		t.Errorf("Warden perk classes: got %v, want [Aggression, Badass]", s.PerkClasses)
	}

	// Check sponsor perks
	if len(s.SponsorPerks) < 2 {
		t.Fatalf("Warden sponsor perks: got %d, want at least 2", len(s.SponsorPerks))
	}
	perkNames := map[string]bool{}
	for _, sp := range s.SponsorPerks {
		perkNames[sp.Name] = true
	}
	if !perkNames["Prison Cars"] {
		t.Error("expected Prison Cars sponsor perk")
	}
	if !perkNames["Fireworks"] {
		t.Error("expected Fireworks sponsor perk")
	}
}

// Aggregate tests

func TestListVehicleTypes(t *testing.T) {
	vts := ListVehicleTypes()
	if len(vts) < 10 {
		t.Errorf("expected at least 10 vehicle types, got %d", len(vts))
	}
}

func TestListWeapons(t *testing.T) {
	ws := ListWeapons()
	if len(ws) < 4 {
		t.Errorf("expected at least 4 weapons, got %d", len(ws))
	}
}

func TestListUpgrades(t *testing.T) {
	us := ListUpgrades()
	if len(us) < 5 {
		t.Errorf("expected at least 5 upgrades, got %d", len(us))
	}
}

func TestListPerks(t *testing.T) {
	ps := ListPerks()
	if len(ps) < 20 {
		t.Errorf("expected at least 20 perks, got %d", len(ps))
	}
}

func TestListSponsors(t *testing.T) {
	ss := ListSponsors()
	if len(ss) < 10 {
		t.Errorf("expected at least 10 sponsors, got %d", len(ss))
	}
}

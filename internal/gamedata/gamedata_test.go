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

package gamedata

// VehicleType represents a vehicle type from the Gaslands Refuelled rulebook.
type VehicleType struct {
	Name         string
	WeightClass  string
	BaseCost     int
	Hull         int
	Handling     int
	MaxGear      int
	Crew         int
	BuildSlots   int
	SpecialRules string
}

var vehicleTypes = []VehicleType{
	// Basic Vehicle Types
	{Name: "Buggy", WeightClass: "Lightweight", BaseCost: 6, Hull: 6, Handling: 4, MaxGear: 6, Crew: 2, BuildSlots: 2, SpecialRules: "Roll Cage"},
	{Name: "Car", WeightClass: "Middleweight", BaseCost: 12, Hull: 10, Handling: 3, MaxGear: 5, Crew: 2, BuildSlots: 2},
	{Name: "Performance Car", WeightClass: "Middleweight", BaseCost: 15, Hull: 8, Handling: 4, MaxGear: 6, Crew: 1, BuildSlots: 2, SpecialRules: "Slip Away"},
	{Name: "Truck", WeightClass: "Middleweight", BaseCost: 15, Hull: 12, Handling: 2, MaxGear: 4, Crew: 3, BuildSlots: 3},
	{Name: "Heavy Truck", WeightClass: "Heavyweight", BaseCost: 25, Hull: 14, Handling: 2, MaxGear: 3, Crew: 4, BuildSlots: 5},
	{Name: "Bus", WeightClass: "Heavyweight", BaseCost: 30, Hull: 16, Handling: 2, MaxGear: 3, Crew: 8, BuildSlots: 3},

	// Advanced Vehicle Types
	{Name: "Drag Racer", WeightClass: "Lightweight", BaseCost: 5, Hull: 4, Handling: 4, MaxGear: 6, Crew: 1, BuildSlots: 2, SpecialRules: "Jet Engine"},
	{Name: "Bike", WeightClass: "Lightweight", BaseCost: 5, Hull: 4, Handling: 5, MaxGear: 6, Crew: 1, BuildSlots: 1, SpecialRules: "Full Throttle. Pivot"},
	{Name: "Bike with Sidecar", WeightClass: "Lightweight", BaseCost: 8, Hull: 4, Handling: 5, MaxGear: 6, Crew: 2, BuildSlots: 2, SpecialRules: "Full Throttle. Pivot"},
	{Name: "Ice Cream Truck", WeightClass: "Middleweight", BaseCost: 8, Hull: 10, Handling: 2, MaxGear: 4, Crew: 2, BuildSlots: 2, SpecialRules: "Infuriating Jingle"},
	{Name: "Gyrocopter", WeightClass: "Middleweight", BaseCost: 10, Hull: 4, Handling: 4, MaxGear: 6, Crew: 1, BuildSlots: 0, SpecialRules: "Airwolf. Airborne"},
	{Name: "Ambulance", WeightClass: "Middleweight", BaseCost: 20, Hull: 12, Handling: 2, MaxGear: 5, Crew: 3, BuildSlots: 3, SpecialRules: "Uppers. Downers"},
	{Name: "Monster Truck", WeightClass: "Heavyweight", BaseCost: 25, Hull: 10, Handling: 3, MaxGear: 4, Crew: 2, BuildSlots: 2, SpecialRules: "All Terrain. Up and Over"},
	{Name: "Helicopter", WeightClass: "Heavyweight", BaseCost: 30, Hull: 8, Handling: 3, MaxGear: 4, Crew: 3, BuildSlots: 4, SpecialRules: "Airwolf. Airborne. Pivot. Up and Over. All Terrain. Restricted"},
	{Name: "Tank", WeightClass: "Heavyweight", BaseCost: 40, Hull: 20, Handling: 4, MaxGear: 3, Crew: 3, BuildSlots: 4, SpecialRules: "Turret. Restricted"},
	{Name: "War Rig", WeightClass: "Heavyweight", BaseCost: 40, Hull: 26, Handling: 2, MaxGear: 4, Crew: 5, BuildSlots: 5, SpecialRules: "Articulated. Ponderous. Piledriver"},
}

// GetVehicleType returns the vehicle type with the given name.
func GetVehicleType(name string) (VehicleType, bool) {
	for _, vt := range vehicleTypes {
		if vt.Name == name {
			return vt, true
		}
	}
	return VehicleType{}, false
}

// ListVehicleTypes returns all vehicle types.
func ListVehicleTypes() []VehicleType {
	result := make([]VehicleType, len(vehicleTypes))
	copy(result, vehicleTypes)
	return result
}

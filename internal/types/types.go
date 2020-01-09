package types

import (
	"fmt"
)

// Part stores the name and stats of a generic part
type Part struct {
	Name         string  `json:"name"`
	Speed        float64 `json:"speed"`
	Acceleration float64 `json:"acceleration"`
	Weight       float64 `json:"weight"`
	Handling     float64 `json:"handling"`
	Traction     float64 `json:"traction"`
	MiniTurbo    float64 `json:"mini-turbo"`
}

// Build stores a combination of character, vehicle, tire and glider
type Build struct {
	Character    *Part   `json:"character"`
	Vehicle      *Part   `json:"vehicle"`
	Tire         *Part   `json:"tire"`
	Glider       *Part   `json:"glider"`
	Speed        float64 `json:"speed"`
	Acceleration float64 `json:"acceleration"`
	Weight       float64 `json:"weight"`
	Handling     float64 `json:"handling"`
	Traction     float64 `json:"traction"`
	MiniTurbo    float64 `json:"mini-turbo"`
}

// NewBuild initializes a new Build object based on the given parts
func NewBuild(char Part, vehicle Part, tires Part, glider Part) Build {
	build := Build{
		Character: &char,
		Vehicle:   &vehicle,
		Tire:      &tires,
		Glider:    &glider,
	}
	build.Speed = char.Speed + vehicle.Speed + tires.Speed + glider.Speed
	build.Acceleration = char.Acceleration + vehicle.Acceleration + tires.Acceleration + glider.Acceleration
	build.Weight = char.Weight + vehicle.Weight + tires.Weight + glider.Weight
	build.Handling = char.Handling + vehicle.Handling + tires.Handling + glider.Handling
	build.Traction = char.Traction + vehicle.Traction + tires.Traction + glider.Traction
	build.MiniTurbo = char.MiniTurbo + vehicle.MiniTurbo + tires.MiniTurbo + glider.MiniTurbo
	return build
}

// PrintBuild pretty prints the desired build
func (b *Build) PrintBuild() {
	fmt.Printf("Character: %s\nVehicle: %s\nTires: %s\nGlider: %s\n", b.Character.Name, b.Vehicle.Name, b.Tire.Name, b.Glider.Name)
	fmt.Printf("Speed: %0.2f\tAcceleration: %0.2f\tWeight: %0.2f\nHandling: %0.2f\tTraction: %0.2f\t\tMiniturbo: %0.2f\n\n", b.Speed, b.Acceleration, b.Weight, b.Handling, b.Traction, b.MiniTurbo)
}

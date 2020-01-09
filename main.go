package main

import (
	"flag"
	"fmt"

	"github.com/AdrienCos/mk8dx_pareto/internal/generator"
	"github.com/AdrienCos/mk8dx_pareto/internal/loader"
	"github.com/AdrienCos/mk8dx_pareto/internal/pareto"
)

var dataPath string = "data/"
var charactersFilename string = "characters.csv"
var vehiclesFilename string = "vehicles.csv"
var tiresFilename string = "tires.csv"
var glidersFilename string = "gliders.csv"

func main() {
	// Get the correct paths to each CSV file
	characterPath := flag.String("c", dataPath+charactersFilename, "Path the the characters CSV")
	vehiclesPath := flag.String("v", dataPath+vehiclesFilename, "Path to the vehicles CSV")
	tiresPath := flag.String("t", dataPath+tiresFilename, "Path to the tires CSV")
	glidersPath := flag.String("g", dataPath+glidersFilename, "Path to the gliders CSV")
	flag.Parse()
	// Load all data in memory
	characters := loader.LoadParts(*characterPath)
	vehicles := loader.LoadParts(*vehiclesPath)
	tires := loader.LoadParts(*tiresPath)
	gliders := loader.LoadParts(*glidersPath)
	// Compute all types.Builds
	builds := generator.GenerateBuilds(characters, vehicles, tires, gliders)
	// Select the Pareto pair
	criteria1 := pareto.SortAcceleration
	criteria2 := pareto.SortSpeed
	// types.Build the frontier
	frontier := pareto.ExtractFrontier(builds, criteria1, criteria2)
	// Print the results
	fmt.Println(len(frontier))
	for _, b := range frontier {
		b.PrintBuild()
	}
	return
}

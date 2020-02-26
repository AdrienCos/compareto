package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/AdrienCos/compareto/internal/generator"
	"github.com/AdrienCos/compareto/internal/loader"
	"github.com/AdrienCos/compareto/internal/pareto"
)

var dataPath string = "/data/"
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
	criteria1String := flag.String("c1", "speed", "First criteria to use for the Pareto frontier")
	criteria2String := flag.String("c2", "acceleration", "Second criteria to use for the Pareto frontier")
	flag.Parse()
	// Check user input
	c1, err := pareto.GetCriteriaFromFlag(*criteria1String)
	if err != nil {
		log.Fatal(err)
		return
	}
	c2, err := pareto.GetCriteriaFromFlag(*criteria2String)
	if err != nil {
		log.Fatal(err)
		return
	}
	// Load all data in memory
	vehicles := loader.LoadParts(*vehiclesPath)
	characters := loader.LoadParts(*characterPath)
	tires := loader.LoadParts(*tiresPath)
	gliders := loader.LoadParts(*glidersPath)
	// Compute all types.Builds
	builds := generator.GenerateBuilds(characters, vehicles, tires, gliders)
	// types.Build the frontier
	frontier := pareto.ExtractFrontier(builds, c1, c2)
	// Print the results
	fmt.Println(len(frontier))
	for _, b := range frontier {
		b.PrintBuild()
	}
	return
}

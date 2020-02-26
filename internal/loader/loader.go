package loader

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/AdrienCos/compareto/internal/types"

	"github.com/markbates/pkger"
)

func getFloat(s string) float64 {
	out, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		log.Fatal("Could not read value from CSV", err)
	}
	return out
}

// LoadParts reads the CSV file at the given address, and turns it into a slice of Part
func LoadParts(path string) []types.Part {
	csvFile, err := pkger.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';'
	reader.Comment = '#'
	var parts []types.Part
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		newPart := types.Part{
			Name:         line[0],
			Speed:        getFloat(line[1]),
			Acceleration: getFloat(line[2]),
			Weight:       getFloat(line[3]),
			Handling:     getFloat(line[4]),
			Traction:     getFloat(line[5]),
			MiniTurbo:    getFloat(line[6]),
		}
		newPart.Total = newPart.Speed + newPart.Acceleration + newPart.Weight + newPart.Handling + newPart.Traction + newPart.MiniTurbo
		parts = append(parts, newPart)
	}
	return parts

}

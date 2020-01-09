package loader

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/AdrienCos/mk8dx_pareto/internal/types"
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
	csvFile, _ := os.Open(path)
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
		parts = append(parts, types.Part{
			Name:         line[0],
			Speed:        getFloat(line[1]),
			Acceleration: getFloat(line[2]),
			Weight:       getFloat(line[3]),
			Handling:     getFloat(line[4]),
			Traction:     getFloat(line[5]),
			MiniTurbo:    getFloat(line[6]),
		})
	}
	return parts

}

package pareto

import (
	"fmt"
	"sort"

	"github.com/AdrienCos/compareto/internal/types"
)

// Criteria is a 'less'-type function that compares two types.Builds based on a specific stat
type Criteria func(b1, b2 *types.Build) bool

// BuildsSorter is a sort type for sorting types.Builds according to the given function
type BuildsSorter struct {
	b  []types.Build
	by Criteria
}

func (a BuildsSorter) Len() int           { return len(a.b) }
func (a BuildsSorter) Swap(i, j int)      { a.b[i], a.b[j] = a.b[j], a.b[i] }
func (a BuildsSorter) Less(i, j int) bool { return a.by(&a.b[i], &a.b[j]) }

// SortSpeed is the speed criteria
func SortSpeed(b1, b2 *types.Build) bool {
	return b1.Speed > b2.Speed
}

// SortAcceleration is the acceleration criteria
func SortAcceleration(b1, b2 *types.Build) bool {
	return b1.Acceleration > b2.Acceleration
}

// SortWeight is the acceleration criteria
func SortWeight(b1, b2 *types.Build) bool {
	return b1.Weight > b2.Weight
}

// SortHandling is the weigth criteria
func SortHandling(b1, b2 *types.Build) bool {
	return b1.Handling > b2.Handling
}

// SortTraction is the traction criteria
func SortTraction(b1, b2 *types.Build) bool {
	return b1.Traction > b2.Traction
}

// SortMiniTurbo is the mini-turbo criteria
func SortMiniTurbo(b1, b2 *types.Build) bool {
	return b1.MiniTurbo > b2.MiniTurbo
}

func orderAndPick(b []types.Build, c1, c2 Criteria) []types.Build {
	// Sort the types.Builds
	buildsSorter := BuildsSorter{b, c1}
	sort.Sort(buildsSorter)
	// Construct the Pareto frontier
	var frontier []types.Build
	frontier = append(frontier, b[0])
	for _, build := range b {
		if c2(&build, &frontier[len(frontier)-1]) {
			frontier = append(frontier, build)
		}
	}
	return frontier
}

// ExtractFrontier types.Builds the Pareto frontier of the given types.Builds based on the given criterion
func ExtractFrontier(b []types.Build, c1, c2 Criteria) []types.Build {
	frontier := orderAndPick(b, c1, c2)
	frontier = orderAndPick(frontier, c2, c1)
	return frontier
}

// InvalidCriteria is an error sent when the user asks for a non-existent criteria
type InvalidCriteria struct {
	Flag string
}

func (e InvalidCriteria) Error() string {
	return fmt.Sprintf("%s is an invalid criteria.", e.Flag)
}

// GetCriteriaFromFlag returns the Criteria corresponding to the user-given flag.
func GetCriteriaFromFlag(s string) (Criteria, error) {
	switch s {
	case "speed":
		return SortSpeed, nil
	case "acceleration":
		return SortAcceleration, nil
	case "weight":
		return SortWeight, nil
	case "handling":
		return SortHandling, nil
	case "traction":
		return SortTraction, nil
	case "miniturbo":
		return SortMiniTurbo, nil
	default:
		return nil, InvalidCriteria{s}
	}
}

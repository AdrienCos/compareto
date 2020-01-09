package generator

import (
	"github.com/AdrienCos/mk8dx_pareto/internal/types"
)

// GenerateBuilds creates and returns all possible types.Builds based on the given characters, vehicles, tires, and gliders
func GenerateBuilds(c []types.Part, v []types.Part, t []types.Part, g []types.Part) []types.Build {
	var builds []types.Build
	for _, char := range c {
		for _, vehi := range v {
			for _, tire := range t {
				for _, glid := range g {
					builds = append(builds, types.NewBuild(char, vehi, tire, glid))
				}
			}
		}
	}
	return builds
}

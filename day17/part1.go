package day17

import (
	"fmt"
	"strings"
)

const numConwaySteps = 6

var conwayCubeStartState2 = `.#.
..#
###`

var conwayCubeStartState = `#.#.##.#
#.####.#
...##...
#####.##
#....###
##..##..
#..####.
#...#.#.`

func loadInitialConwayCubeState() [][]bool {
	lines := strings.Split(conwayCubeStartState, "\n")
	res := make([][]bool, len(lines)+2*numConwaySteps)
	singleInputLineLen := len(lines[0])
	// first 6 rows, all false
	for i := 0; i < numConwaySteps; i++ {
		row := make([]bool, singleInputLineLen+2*numConwaySteps)
		for j := 0; j < singleInputLineLen+2*numConwaySteps; j++ {
			row[j] = false
		}
		res[i] = row
	}
	// next len(lines) rows
	for i := numConwaySteps; i < numConwaySteps+len(lines); i++ {
		row := make([]bool, singleInputLineLen+2*numConwaySteps)
		for j := 0; j < numConwaySteps; j++ {
			row[j] = false
		}
		for j, c := range lines[i-numConwaySteps] {
			row[j+numConwaySteps] = string(c) == "#"
		}
		for j := singleInputLineLen + numConwaySteps; j < singleInputLineLen+2*numConwaySteps; j++ {
			row[j] = false
		}
		res[i] = row
	}
	// last 6 rows, all false
	for i := len(lines) + numConwaySteps; i < len(lines)+2*numConwaySteps; i++ {
		row := make([]bool, singleInputLineLen+2*numConwaySteps)
		for j := 0; j < singleInputLineLen+2*numConwaySteps; j++ {
			row[j] = false
		}
		res[i] = row
	}

	return res
}

func generateBlankPlane(maxY, maxX int) [][]bool {
	res := make([][]bool, maxY)
	for y := 0; y < maxY; y++ {
		row := make([]bool, maxX)
		res[y] = row
	}

	return res
}

func generateBlankSpace3D(maxZ, maxY, maxX int) [][][]bool {
	res := make([][][]bool, maxZ)
	for z := 0; z < maxZ; z++ {
		res[z] = generateBlankPlane(maxY, maxX)
	}
	return res
}

func countActiveConwayNeighbors(space [][][]bool, x, y, z int) int {
	directions := [][3]int{
		{0, 0, 1},
		{0, 0, -1},
		{0, 1, 0},
		{0, -1, 0},
		{0, 1, 1},
		{0, 1, -1},
		{0, -1, 1},
		{0, -1, -1},
		{1, 0, 0},
		{-1, 0, 0},
		{1, 0, 1},
		{1, 0, -1},
		{-1, 0, 1},
		{-1, 0, -1},
		{1, 1, 0},
		{1, -1, 0},
		{-1, 1, 0},
		{-1, -1, 0},
		{1, 1, 1},
		{1, 1, -1},
		{1, -1, 1},
		{-1, 1, 1},
		{1, -1, -1},
		{-1, 1, -1},
		{-1, -1, 1},
		{-1, -1, -1},
	}

	total := 0
	for _, direction := range directions {
		if z+direction[0] >= len(space) || z+direction[0] < 0 {
			continue
		}
		if y+direction[1] >= len(space[0]) || y+direction[1] < 0 {
			continue
		}
		if x+direction[2] >= len(space[0][0]) || x+direction[2] < 0 {
			continue
		}
		if space[z+direction[0]][y+direction[1]][x+direction[2]] {
			total++
		}
	}
	return total
}

func cloneConwaySpace(space [][][]bool) [][][]bool {
	res := make([][][]bool, len(space))
	for z := range space {
		newPlane := make([][]bool, len(space[0]))
		for y := range space[z] {
			newRow := make([]bool, len(space[0][0]))
			for x := range space[z][y] {
				newRow[x] = space[z][y][x]
			}
			newPlane[y] = newRow
		}
		res[z] = newPlane
	}
	return res
}

func stepConwaySpace(space [][][]bool) [][][]bool {
	cloned := cloneConwaySpace(space)
	//fmt.Printf("%#v\n", cloned)
	for z := range space {
		for y := range space[z] {
			for x := range space[z][y] {
				activeNeighborCount := countActiveConwayNeighbors(space, x, y, z)
				if space[z][y][x] && activeNeighborCount != 2 && activeNeighborCount != 3 {
					fmt.Println("#: found ", activeNeighborCount, " neighbors for ", z, y, x)
					cloned[z][y][x] = false
				} else if !space[z][y][x] && activeNeighborCount == 3 {
					fmt.Println(".: found 3 neighbors for ", z, y, x)
					cloned[z][y][x] = true
				}
			}
		}
	}
	return cloned
}

func printState(space [][][]bool) {
	for z := range space {
		if isPlaneBlank(space[z]) {
			continue
		}
		fmt.Println("Z=", z, " -----------------")
		for y := range space[z] {
			for x := range space[z][y] {
				if space[z][y][x] {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Printf("\n")
		}
	}
}

func countActiveConwayCubes(space [][][]bool) int {
	total := 0
	for z := range space {
		for y := range space[z] {
			for x := range space[z][y] {
				if space[z][y][x] {
					total++
				}
			}
		}
	}
	return total
}

func ActiveConwayCubes() {
	initialPlane := loadInitialConwayCubeState()
	space := append(generateBlankSpace3D(numConwaySteps, len(initialPlane), len(initialPlane[0])), initialPlane)
	space = append(space, generateBlankSpace3D(numConwaySteps, len(initialPlane), len(initialPlane[0]))...)
	printState(space)
	for i := 0; i < numConwaySteps; i++ {
		space = stepConwaySpace(space)
		fmt.Println("after step ", i, ":")
		printState(space)
		fmt.Println("total active cubes: ", countActiveConwayCubes(space))
	}
}

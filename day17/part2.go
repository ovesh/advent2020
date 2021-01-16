package day17

import "fmt"

var directions = prepare4DDirections()

func prepare4DDirections() [][4]int {
	res := [][4]int{}
	for w := -1; w <= 1; w++ {
		for z := -1; z <= 1; z++ {
			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					if w == 0 && z == 0 && y == 0 && x == 0 {
						continue
					}
					res = append(res, [4]int{w, z, y, x})
				}
			}
		}
	}
	return res
}

func countActiveConwayNeighbors4D(space [][][][]bool, x, y, z, w int) int {
	total := 0
	for _, direction := range directions {
		if w+direction[0] >= len(space) || w+direction[0] < 0 {
			continue
		}
		if z+direction[1] >= len(space[0]) || z+direction[1] < 0 {
			continue
		}
		if y+direction[2] >= len(space[0][0]) || y+direction[2] < 0 {
			continue
		}
		if x+direction[3] >= len(space[0][0][0]) || x+direction[3] < 0 {
			continue
		}
		if space[w+direction[0]][z+direction[1]][y+direction[2]][x+direction[3]] {
			total++
		}
	}
	return total
}

func cloneConwaySpace4D(space [][][][]bool) [][][][]bool {
	res := make([][][][]bool, len(space))
	for w := range space {
		new3DSpace := make([][][]bool, len(space[0]))
		for z := range space[w] {
			newPlane := make([][]bool, len(space[0][0]))
			for y := range space[w][z] {
				newRow := make([]bool, len(space[0][0][0]))
				for x := range space[w][z][y] {
					newRow[x] = space[w][z][y][x]
				}
				newPlane[y] = newRow
			}
			new3DSpace[z] = newPlane
		}
		res[w] = new3DSpace
	}
	return res
}

func stepConwaySpace4D(space [][][][]bool) [][][][]bool {
	cloned := cloneConwaySpace4D(space)
	//fmt.Printf("%#v\n", cloned)
	for w := range space {
		for z := range space[w] {
			for y := range space[w][z] {
				for x := range space[w][z][y] {
					activeNeighborCount := countActiveConwayNeighbors4D(space, x, y, z, w)
					if space[w][z][y][x] && activeNeighborCount != 2 && activeNeighborCount != 3 {
						//fmt.Println("#: found ", activeNeighborCount, " neighbors for ", w, z, y, x)
						cloned[w][z][y][x] = false
					} else if !space[w][z][y][x] && activeNeighborCount == 3 {
						//fmt.Println(".: found 3 neighbors for ", w, z, y, x)
						cloned[w][z][y][x] = true
					}
				}
			}
		}
	}
	return cloned
}

func countActiveConwayCubes4D(space [][][][]bool) int {
	total := 0
	for w := range space {
		for z := range space[w] {
			for y := range space[w][z] {
				for x := range space[w][z][y] {
					if space[w][z][y][x] {
						total++
					}
				}
			}
		}
	}
	return total
}

func generateBlankSpace4D(maxW, maxZ, maxY, maxX int) [][][][]bool {
	res := make([][][][]bool, maxW)
	for w := 0; w < maxW; w++ {
		res[w] = generateBlankSpace3D(maxZ, maxY, maxX)
	}
	return res
}

func isPlaneBlank(plane [][]bool) bool {
	for y := range plane {
		for x := range plane[y] {
			if plane[y][x] {
				return false
			}
		}
	}
	return true
}

func printState4D(space [][][][]bool) {
	for w := range space {
		for z := range space[w] {
			if isPlaneBlank(space[w][z]) {
				continue
			}
			fmt.Println("W=", w, "Z=", z, " -----------------")
			for y := range space[w][z] {
				for x := range space[w][z][y] {
					if space[w][z][y][x] {
						fmt.Printf("#")
					} else {
						fmt.Printf(".")
					}
				}
				fmt.Printf("\n")
			}
		}
	}

}

func ActiveConwayCubes4D() {
	initialPlane := loadInitialConwayCubeState()
	space := append(generateBlankSpace3D(numConwaySteps, len(initialPlane), len(initialPlane[0])), initialPlane)
	space = append(space, generateBlankSpace3D(numConwaySteps, len(initialPlane), len(initialPlane[0]))...)
	space4D := append(generateBlankSpace4D(numConwaySteps, 2*numConwaySteps+1, len(initialPlane), len(initialPlane[0])), space)
	space4D = append(space4D, generateBlankSpace4D(numConwaySteps, 2*numConwaySteps+1, len(initialPlane), len(initialPlane[0]))...)
	//printState4D(space4D)
	fmt.Println("total active cubes: ", countActiveConwayCubes4D(space4D))
	for i := 0; i < numConwaySteps; i++ {
		space4D = stepConwaySpace4D(space4D)
		//fmt.Println("after step ", i, ":")
		//printState4D(space4D)
		fmt.Println("total active cubes: ", countActiveConwayCubes4D(space4D))
	}
}

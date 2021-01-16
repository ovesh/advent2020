package day03

import "fmt"

func howManyTrees(right, down int) int {
	curCol := 0
	totalHits := 0
	rowLen := len(trees[0])
	for row := 0; row < len(trees); row += down {
		if string(trees[row][curCol%rowLen]) == "#" {
			totalHits++
		}
		curCol += right
	}

	fmt.Println("total hits: ", totalHits)
	return totalHits
}

func HowManyTrees2() {
	rightDowns := [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	product := 1
	for _, rightDown := range rightDowns {
		hits := howManyTrees(rightDown[0], rightDown[1])
		product *= hits
	}

	fmt.Println("product: ", product)
}

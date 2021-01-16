package day20

import (
	"fmt"
	"strconv"
	"strings"
)

type tile struct {
	id int
	// up, right, down, left
	borders        [4][10]bool
	flippedBorders [4][10]bool

	pixels [10][10]bool
}

func flipArray(a [10]bool) [10]bool {
	res := [10]bool{}
	for i := range a {
		res[len(a)-i-1] = a[i]
	}
	return res
}

func newTile(id int, rawInputLines []string) tile {
	borders := [4][10]bool{}
	flippedBorders := [4][10]bool{}
	upBorder := [10]bool{}
	for i, c := range rawInputLines[0] {
		sc := string(c)
		if sc == "#" {
			upBorder[i] = true
		}
	}
	borders[0] = upBorder
	flippedBorders[0] = flipArray(upBorder)

	downBorder := [10]bool{}
	for i, c := range rawInputLines[9] {
		sc := string(c)
		if sc == "#" {
			downBorder[i] = true
		}
	}
	borders[2] = downBorder
	flippedBorders[2] = flipArray(downBorder)

	rightBorder := [10]bool{}
	leftBorder := [10]bool{}
	for i, line := range rawInputLines {
		left := string(line[0])
		if left == "#" {
			leftBorder[i] = true
		}
		right := string(line[9])
		if right == "#" {
			rightBorder[i] = true
		}
	}
	borders[1] = rightBorder
	borders[3] = leftBorder
	flippedBorders[1] = flipArray(rightBorder)
	flippedBorders[3] = flipArray(leftBorder)

	pixels := [10][10]bool{}
	for i, line := range rawInputLines {
		for j, c := range line {
			sc := string(c)
			if sc == "#" {
				pixels[i][j] = true
			}
		}
	}

	return tile{id: id, borders: borders, flippedBorders: flippedBorders, pixels: pixels}
}

type match struct {
	// up, right, down, left
	aSide string
	bSide string
	// 0, 90, 180, 270
	aRotation int
	//bRotation          int
	aFlippedVertical   bool
	aFlippedHorizontal bool
	//bFlippedVertical   bool
	//bFlippedHorizontal bool
}

func bordersMatch(a, b [10]bool) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (a tile) matches(b tile) []match {
	res := []match{}
	if bordersMatch(a.borders[0], b.borders[0]) {
		res = append(res, match{
			aSide:              "up",
			bSide:              "up",
			aRotation:          180,
			aFlippedVertical:   false,
			aFlippedHorizontal: true,
		})
	}
	if bordersMatch(a.flippedBorders[0], b.borders[0]) {
		res = append(res, match{
			aSide:              "up",
			bSide:              "up",
			aRotation:          180,
			aFlippedVertical:   false,
			aFlippedHorizontal: false,
		})
	}
	if bordersMatch(a.borders[0], b.borders[1]) {
		res = append(res, match{
			aSide:              "up",
			bSide:              "right",
			aRotation:          270,
			aFlippedVertical:   false,
			aFlippedHorizontal: true,
		})
	}
	if bordersMatch(a.flippedBorders[0], b.borders[1]) {
		res = append(res, match{
			aSide:              "up",
			bSide:              "right",
			aRotation:          270,
			aFlippedVertical:   false,
			aFlippedHorizontal: false,
		})
	}
	if bordersMatch(a.borders[0], b.borders[2]) {
		res = append(res, match{
			aSide:              "up",
			bSide:              "down",
			aRotation:          0,
			aFlippedVertical:   false,
			aFlippedHorizontal: false,
		})
	}
	if bordersMatch(a.flippedBorders[0], b.borders[2]) {
		res = append(res, match{
			aSide:              "up",
			bSide:              "down",
			aRotation:          0,
			aFlippedVertical:   false,
			aFlippedHorizontal: true,
		})
	}
	if bordersMatch(a.borders[0], b.borders[3]) {
		res = append(res, match{
			aSide:              "up",
			bSide:              "left",
			aRotation:          90,
			aFlippedVertical:   false,
			aFlippedHorizontal: false,
		})
	}
	if bordersMatch(a.flippedBorders[0], b.borders[3]) {
		res = append(res, match{
			aSide:              "up",
			bSide:              "left",
			aRotation:          90,
			aFlippedVertical:   false,
			aFlippedHorizontal: true,
		})
	}
	if bordersMatch(a.borders[1], b.borders[0]) {
		res = append(res, match{
			aSide:              "right",
			bSide:              "up",
			aRotation:          90,
			aFlippedVertical:   true,
			aFlippedHorizontal: false,
		})
	}
	if bordersMatch(a.flippedBorders[1], b.borders[0]) {
		res = append(res, match{
			aSide:              "right",
			bSide:              "up",
			aRotation:          90,
			aFlippedVertical:   false,
			aFlippedHorizontal: false,
		})
	}
	if bordersMatch(a.borders[1], b.borders[1]) {
		res = append(res, match{
			aSide:              "right",
			bSide:              "right",
			aRotation:          0,
			aFlippedVertical:   false,
			aFlippedHorizontal: true,
		})
	}
	if bordersMatch(a.flippedBorders[1], b.borders[1]) {
		res = append(res, match{
			aSide:              "right",
			bSide:              "right",
			aRotation:          0,
			aFlippedVertical:   true,
			aFlippedHorizontal: true,
		})
	}
	if bordersMatch(a.borders[1], b.borders[2]) {
		res = append(res, match{
			aSide:              "right",
			bSide:              "down",
			aRotation:          270,
			aFlippedVertical:   false,
			aFlippedHorizontal: false,
		})
	}
	if bordersMatch(a.flippedBorders[1], b.borders[2]) {
		res = append(res, match{
			aSide:              "right",
			bSide:              "down",
			aRotation:          270,
			aFlippedVertical:   true,
			aFlippedHorizontal: false,
		})
	}
	if bordersMatch(a.borders[1], b.borders[3]) {
		res = append(res, match{
			aSide:              "right",
			bSide:              "left",
			aRotation:          0,
			aFlippedVertical:   false,
			aFlippedHorizontal: false,
		})
	}
	if bordersMatch(a.flippedBorders[1], b.borders[3]) {
		res = append(res, match{
			aSide:              "right",
			bSide:              "left",
			aRotation:          0,
			aFlippedVertical:   true,
			aFlippedHorizontal: false,
		})
	}
	if bordersMatch(a.borders[2], b.borders[0]) {
		res = append(res, match{
			aSide:              "down",
			bSide:              "up",
			aRotation:          0,
			aFlippedVertical:   false,
			aFlippedHorizontal: false,
		})
	}
	if bordersMatch(a.flippedBorders[2], b.borders[0]) {
		res = append(res, match{
			aSide:              "down",
			bSide:              "up",
			aRotation:          0,
			aFlippedVertical:   false,
			aFlippedHorizontal: true,
		})
	}
	if bordersMatch(a.borders[2], b.borders[1]) {
		res = append(res, match{
			aSide:              "down",
			bSide:              "right",
			aRotation:          90,
			aFlippedVertical:   false,
			aFlippedHorizontal: false,
		})
	}
	if bordersMatch(a.flippedBorders[2], b.borders[1]) {
		res = append(res, match{
			aSide:              "down",
			bSide:              "right",
			aRotation:          90,
			aFlippedVertical:   false,
			aFlippedHorizontal: true,
		})
	}
	if bordersMatch(a.borders[2], b.borders[2]) {
		res = append(res, match{
			aSide:              "down",
			bSide:              "down",
			aRotation:          0,
			aFlippedVertical:   true,
			aFlippedHorizontal: false,
		})
	}
	if bordersMatch(a.flippedBorders[2], b.borders[2]) {
		res = append(res, match{
			aSide:              "down",
			bSide:              "down",
			aRotation:          0,
			aFlippedVertical:   true,
			aFlippedHorizontal: true,
		})
	}
	if bordersMatch(a.borders[2], b.borders[3]) {
		res = append(res, match{
			aSide:              "down",
			bSide:              "left",
			aRotation:          90,
			aFlippedVertical:   true,
			aFlippedHorizontal: false,
		})
	}
	if bordersMatch(a.flippedBorders[2], b.borders[3]) {
		res = append(res, match{
			aSide:              "down",
			bSide:              "left",
			aRotation:          270,
			aFlippedVertical:   false,
			aFlippedHorizontal: false,
		})
	}
	if bordersMatch(a.borders[3], b.borders[0]) {
		res = append(res, match{
			aSide:              "left",
			bSide:              "up",
			aRotation:          270,
			aFlippedVertical:   false,
			aFlippedHorizontal: false,
		})
	}
	if bordersMatch(a.flippedBorders[3], b.borders[0]) {
		res = append(res, match{
			aSide:              "left",
			bSide:              "up",
			aRotation:          270,
			aFlippedVertical:   true,
			aFlippedHorizontal: false,
		})
	}
	if bordersMatch(a.borders[3], b.borders[1]) {
		res = append(res, match{
			aSide:              "left",
			bSide:              "right",
			aRotation:          0,
			aFlippedVertical:   false,
			aFlippedHorizontal: false,
		})
	}
	if bordersMatch(a.flippedBorders[3], b.borders[1]) {
		res = append(res, match{
			aSide:              "left",
			bSide:              "right",
			aRotation:          0,
			aFlippedVertical:   true,
			aFlippedHorizontal: false,
		})
	}
	if bordersMatch(a.borders[3], b.borders[2]) {
		res = append(res, match{
			aSide:              "left",
			bSide:              "down",
			aRotation:          90,
			aFlippedVertical:   true,
			aFlippedHorizontal: false,
		})
	}
	if bordersMatch(a.flippedBorders[3], b.borders[2]) {
		res = append(res, match{
			aSide:              "left",
			bSide:              "down",
			aRotation:          90,
			aFlippedVertical:   false,
			aFlippedHorizontal: false,
		})
	}
	if bordersMatch(a.borders[3], b.borders[3]) {
		res = append(res, match{
			aSide:              "left",
			bSide:              "left",
			aRotation:          0,
			aFlippedVertical:   false,
			aFlippedHorizontal: true,
		})
	}
	if bordersMatch(a.flippedBorders[3], b.borders[3]) {
		res = append(res, match{
			aSide:              "left",
			bSide:              "left",
			aRotation:          0,
			aFlippedVertical:   true,
			aFlippedHorizontal: true,
		})
	}

	return res
}

func loadTiles() []tile {
	rawTiles := strings.Split(input, "\n\n")
	res := make([]tile, len(rawTiles))
	for i, rawTile := range rawTiles {
		tileLines := strings.Split(rawTile, "\n")
		// Tile 2311:
		tileID, err := strconv.Atoi(tileLines[0][5:9])
		if err != nil {
			panic(fmt.Sprintf("invalid tile id %q, (%q)", tileLines[0], tileLines[0][5:9]))
		}
		res[i] = newTile(tileID, tileLines[1:])
	}

	return res
}

func MultiplyFourCorners() {
	tiles := loadTiles()
	//fmt.Println(tiles)

	// each tile is a vertex in a graph.
	// we need a graph where exactly 100 vertices have 4 edges each (somewhere in the middle of the matrix),
	// 40 have 3 edges each (somewhere along the edges the matrix),
	// and 4 have 2 edges each (the 4 corners of the matrix).
	// If we find this to be correct, then the 4 tiles with only 2 matches are the corners
	// The data revealed exactly 126 matches (and a max of 1 match between any 2 tiles, i.e only one way to connect them),
	// which is what we needed.

	matchesPerID := map[int][]int{}
	for i, tileA := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			tileB := tiles[j]
			matches := tileA.matches(tileB)
			lenMatches := len(matches)
			if lenMatches == 0 {
				continue
			}
			if lenMatches > 1 {
				panic("broke our assumption that every 2 tiles can only connect in one specific way")
			}
			matchesPerID[tileA.id] = append(matchesPerID[tileA.id], tileB.id)
			matchesPerID[tileB.id] = append(matchesPerID[tileB.id], tileA.id)
			fmt.Printf("a: %d b: %d matches: %#v\n", tileA.id, tileB.id, matches)
		}
	}

	cornerIDs := []int{}
	for tileID, matchingIDs := range matchesPerID {
		if len(matchingIDs) == 2 {
			cornerIDs = append(cornerIDs, tileID)
		}
	}
	fmt.Printf("%#v", cornerIDs)

	// confirmed ; we got exactly 4 IDs with 2 matches
	product := 1
	for _, titleID := range cornerIDs {
		product *= titleID
	}
	fmt.Println("product of corner tile IDS: ", product)
}

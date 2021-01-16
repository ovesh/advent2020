package day11

import "fmt"

func getVisibleSeatsForDirection(rowAdvance, colAdvance int, state []string, fromRow, fromCol int) string {
	res := ""
	row, col := fromRow, fromCol
	rowLen := len(state[0])
	for {
		row += rowAdvance
		col += colAdvance
		if row >= len(state) || row < 0 || col >= rowLen || col < 0 {
			return res
		}
		res += string(state[row][col])
	}
}

func visibleSeatIsOccupied(visibleSeats string) bool {
	for _, seat := range visibleSeats {
		if string(seat) == "L" {
			return false
		}
		if string(seat) == "#" {
			return true
		}
	}
	// all Floor
	return false
}

func allVisibleSeatsAreEmpty(state []string, row, col int) bool {
	// up
	if visibleSeatIsOccupied(getVisibleSeatsForDirection(-1, 0, state, row, col)) {
		return false
	}
	// up right
	if visibleSeatIsOccupied(getVisibleSeatsForDirection(-1, 1, state, row, col)) {
		return false
	}
	// right
	if visibleSeatIsOccupied(getVisibleSeatsForDirection(0, 1, state, row, col)) {
		return false
	}
	// right down
	if visibleSeatIsOccupied(getVisibleSeatsForDirection(1, 1, state, row, col)) {
		return false
	}
	// down
	if visibleSeatIsOccupied(getVisibleSeatsForDirection(1, 0, state, row, col)) {
		return false
	}
	// left down
	if visibleSeatIsOccupied(getVisibleSeatsForDirection(1, -1, state, row, col)) {
		return false
	}
	// left
	if visibleSeatIsOccupied(getVisibleSeatsForDirection(0, -1, state, row, col)) {
		return false
	}
	// left up
	if visibleSeatIsOccupied(getVisibleSeatsForDirection(-1, -1, state, row, col)) {
		return false
	}
	return true
}

func countOccupiedVisibleSeats(state []string, row, col int) int {
	count := 0
	// up
	if visibleSeatIsOccupied(getVisibleSeatsForDirection(-1, 0, state, row, col)) {
		count++
	}
	// up right
	if visibleSeatIsOccupied(getVisibleSeatsForDirection(-1, 1, state, row, col)) {
		count++
	}
	// right
	if visibleSeatIsOccupied(getVisibleSeatsForDirection(0, 1, state, row, col)) {
		count++
	}
	// right down
	if visibleSeatIsOccupied(getVisibleSeatsForDirection(1, 1, state, row, col)) {
		count++
	}
	// down
	if visibleSeatIsOccupied(getVisibleSeatsForDirection(1, 0, state, row, col)) {
		count++
	}
	// left down
	if visibleSeatIsOccupied(getVisibleSeatsForDirection(1, -1, state, row, col)) {
		count++
	}
	// left
	if visibleSeatIsOccupied(getVisibleSeatsForDirection(0, -1, state, row, col)) {
		count++
	}
	// left up
	if visibleSeatIsOccupied(getVisibleSeatsForDirection(-1, -1, state, row, col)) {
		count++
	}

	return count
}

func stepSeatSimulation2(initialState []string) []string {
	res := make([]string, len(initialState))
	copy(res, initialState)
	for row := range initialState {
		for col := range initialState[row] {
			if string(initialState[row][col]) == "." {
				continue
			}
			if allVisibleSeatsAreEmpty(initialState, row, col) {
				res[row] = res[row][0:col] + "#" + res[row][col+1:]
			}
			if string(initialState[row][col]) == "#" && countOccupiedVisibleSeats(initialState, row, col) > 4 {
				res[row] = res[row][0:col] + "L" + res[row][col+1:]
			}
		}
	}
	return res
}

func HowManyEmptySeats2() {
	for {
		afterStep := stepSeatSimulation2(seatsInitialState)
		if seatStatesEqual(seatsInitialState, afterStep) {
			fmt.Println("reached terminal state")
			count := countOccupiedSeats(afterStep)
			fmt.Println("total of ", count, " occupied seats")
			return
		}
		seatsInitialState = afterStep
	}
}

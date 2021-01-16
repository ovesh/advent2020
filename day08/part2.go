package day08

import "fmt"

func changeOne(startAt int, instructions []instruction) ([]instruction, int) {
	i := startAt
	for ; i < len(instructions); i++ {
		if instructions[i].command == "nop" || instructions[i].command == "jmp" {
			break
		}
	}
	var newCmd string
	if instructions[i].command == "nop" {
		newCmd = "jmp"
	} else {
		newCmd = "nop"
	}
	cloned := make([]instruction, len(instructions))
	copy(cloned, instructions)
	cloned[i].command = newCmd
	return cloned, i
}

func tryMutation(instructions []instruction) bool {
	pointer := 0
	accum := 0
	visited := map[int]bool{}
	stop := false
	for {
		pointer, accum, stop = step(pointer, instructions, accum, visited)
		if stop {
			fmt.Println("loop!")
			return true
		}
		if pointer == len(instructions) {
			fmt.Println("accum: ", accum)
			return false
		}
	}
}

func FixInfiniteLoop() {
	instructions := loadInstructions()
	for i := 0; i < len(instructions); i++ {
		curInstructions, changedAt := changeOne(i, instructions)
		i = changedAt
		fmt.Println("trying when mutating line ", i, ": ", instructions[i])
		looping := tryMutation(curInstructions)
		if !looping {
			return
		}
	}
}

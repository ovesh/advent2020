package day12

import (
	"fmt"
	"math"
)

func stepNavInstructionWaypoint(curX, curY, waypointX, waypointY int, instruction navInstruction) (newX, newY, newWaypointX, newWaypointY int) {
	switch instruction.command {
	case "F":
		newX := curX + instruction.value*waypointX
		newY := curY + instruction.value*waypointY
		return newX, newY, waypointX, waypointY
	case "N":
		return curX, curY, waypointX, waypointY + instruction.value
	case "E":
		return curX, curY, waypointX + instruction.value, waypointY
	case "S":
		return curX, curY, waypointX, waypointY - instruction.value
	case "W":
		return curX, curY, waypointX - instruction.value, waypointY
	case "L":
		switch (instruction.value + 360) % 360 {
		case 0:
			return curX, curY, waypointX, waypointY
		case 90:
			return curX, curY, -waypointY, waypointX
		case 180:
			return curX, curY, -waypointX, -waypointY
		case 270:
			return curX, curY, waypointY, -waypointX
		}
	case "R":
		switch (instruction.value + 360) % 360 {
		case 0:
			return curX, curY, waypointX, waypointY
		case 90:
			return curX, curY, waypointY, -waypointX
		case 180:
			return curX, curY, -waypointX, -waypointY
		case 270:
			return curX, curY, -waypointY, waypointX
		}
	}
	panic("unknown instruction " + instruction.command)
}

func ManhattanDistanceWaypoint() {
	instructions := loadNavigtionIstructions()
	x, y, waypointX, waypointY := 0, 0, 10, 1
	for _, instruction := range instructions {
		x, y, waypointX, waypointY = stepNavInstructionWaypoint(x, y, waypointX, waypointY, instruction)
		fmt.Println("instruction", instruction, x, y, waypointX, waypointY)
	}
	fmt.Println("final x y ", x, y)
	fmt.Println("Manhattan distance is ", math.Abs(float64(x))+math.Abs(float64(y)))
}

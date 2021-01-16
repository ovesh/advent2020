package day05

import "fmt"

func FindMySeat() {
	seats := loadSeats()
	takenSeats := make([]bool, 127*8)
	for _, s := range seats {
		id := seatID(s)
		takenSeats[id] = true
	}

	for i := range takenSeats {
		if !takenSeats[i] {
			if (i > 0 && !takenSeats[i-1]) || (i < 127*8-1 && !takenSeats[i+1]) {
				continue
			}
			fmt.Println("found seat! ", i)
		}
	}
}

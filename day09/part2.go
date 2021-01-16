package day09

import "fmt"

func trySum(targetSum int, in []int) (found bool, smallest int, largest int) {
	sum := 0
	smallest = in[0]
	largest = in[0]
	for _, i := range in {
		if i < smallest {
			smallest = i
		}
		if i > largest {
			largest = i
		}
		sum += i
		if sum == targetSum {
			return true, smallest, largest
		}
	}

	return false, -1, -1
}

func FindContiguousSum() {
	sum, idx := FirstIllegalNumber()
	for i := 0; i < len(xmasInput); i++ {
		found, smallest, largest := trySum(sum, xmasInput[i:idx])
		if found {
			fmt.Println("found smallest: ", smallest, ", largest: ", largest, ", sum: ", smallest+largest)
			return
		}
	}
}

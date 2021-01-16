package day02

import (
	"fmt"
)

func isValid2(record record) bool {
	first := rune(record.pwd[record.minTimes-1]) == record.char
	second := rune(record.pwd[record.maxTimes-1]) == record.char
	if first && !second {
		return true
	}
	if !first && second {
		return true
	}
	return false
}

func CountValidPasswords2() {
	records := loadDb()
	count := 0
	for _, r := range records {
		if isValid2(r) {
			count++
		}
	}
	fmt.Println("correct: ", count)
}

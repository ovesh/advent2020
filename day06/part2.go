package day06

import (
	"fmt"
	"strings"
)

func CustomsTotal2() {
	rawGroups := strings.Split(rawCustomsAnswers, "\n\n")
	total := 0
	for _, rawGroup := range rawGroups {
		members := strings.Split(rawGroup, "\n")
		lenMembers := len(members)
		answersToNumMembers := map[rune]int{}
		for _, member := range members {
			for _, answer := range member {
				if _, ok := answersToNumMembers[answer]; !ok {
					answersToNumMembers[answer] = 0
				}
				answersToNumMembers[answer] = answersToNumMembers[answer] + 1
			}
		}
		for _, v := range answersToNumMembers {
			if v == lenMembers {
				//fmt.Println("everybody answered correctly to ", string(k), "in group ", rawGroup)
				total += 1
			}
		}
	}

	fmt.Println("total: ", total)
}

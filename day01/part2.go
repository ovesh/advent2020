package day1

import "fmt"

func Find2020Multipler2() int {
	for i := 0; i < len(inputs); i++ {
		for j := i + 1; j < len(inputs); j++ {
			for k := j + 1; k < len(inputs); k++ {
				if inputs[i]+inputs[j]+inputs[k] == 2020 {
					fmt.Println("found them! ", inputs[i], inputs[j], inputs[k], inputs[i]*inputs[j]*inputs[k])
					return inputs[i] * inputs[j] * inputs[k]
				}
			}
		}
	}
	return 0
}

package day07

import "fmt"

func rulesByColor(rules []rule) map[string]rule {
	res := make(map[string]rule, len(rules))
	for _, r := range rules {
		res[r.outerColor] = r
	}
	return res
}

func howMany(color string, rulesByColor map[string]rule) int {
	res := 0
	rule := rulesByColor[color]
	for containedColor, amount := range rule.colorsToAmounts {
		res += amount
		res += amount * howMany(containedColor, rulesByColor)
	}

	return res
}

func HowManyInBag(color string) {
	rules := loadLuggageRules()
	byColor := rulesByColor(rules)
	res := howMany(color, byColor)
	fmt.Println("total: ", res)
}

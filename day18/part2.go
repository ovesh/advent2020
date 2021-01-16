package day18

import (
	"fmt"
	"strconv"
	"strings"
)

type expression interface {
	eval() int
}

type addition struct {
	a expression
	b expression
}

func (a addition) eval() int {
	return a.a.eval() + a.b.eval()
}

type mult struct {
	a expression
	b expression
}

func (m mult) eval() int {
	return m.a.eval() * m.b.eval()
}

type num struct {
	val string
}

func (n num) eval() int {
	v, err := strconv.Atoi(n.val)
	if err != nil {
		panic(fmt.Sprintf("%q is not a number", n.val))
	}
	return v
}

func removeSpaces(s string) string {
	res := ""
	for _, c := range s {
		if string(c) == " " {
			continue
		}
		res += string(c)
	}
	return res
}

func parseMathExpression(s string) expression {
	if len(s) == 1 && strings.Contains("1234567890", s) {
		return num{val: s}
	}

	for i := 0; i < len(s); i++ {
		sc := string(s[i])
		if sc == "(" {
			// skip
			matchingBrace := findMatchingBrace(s[i+1:])
			i = i + 1 + matchingBrace
			continue
		}
		if sc == "*" {
			return mult{
				a: parseMathExpression(s[0:i]),
				b: parseMathExpression(s[i+1:]),
			}
		}
	}

	// no top-level multiplication ; look for additions
	for i := 0; i < len(s); i++ {
		sc := string(s[i])
		if sc == "(" {
			// skip
			matchingBrace := findMatchingBrace(s[i+1:])
			i = matchingBrace
			continue
		}
		if sc == "+" {
			return addition{
				a: parseMathExpression(s[0:i]),
				b: parseMathExpression(s[i+1:]),
			}
		}
	}

	// no top level operations ; down to numbers and braces
	for i := 0; i < len(s); i++ {
		sc := string(s[i])
		if sc == "(" {
			// skip
			matchingBrace := findMatchingBrace(s[i+1:])
			return parseMathExpression(s[i+1 : i+1+matchingBrace])
		}
	}

	panic(fmt.Sprintf("I don't know what to do with %q", s))
}

func SumAllMathHomework2() {
	lines := strings.Split(mathHomeworkInput, "\n")
	sum := 0
	for _, line := range lines {
		exp := parseMathExpression(removeSpaces(line))

		result := exp.eval()
		fmt.Println(line, " = ", result)
		sum += result
	}
	fmt.Println("total sum:", sum)
}

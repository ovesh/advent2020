package day16

import (
	"fmt"
	"strings"
)

var myTicket1 = []int{11, 12, 13}
var myTicket = []int{67, 107, 59, 79, 53, 131, 61, 101, 71, 73, 137, 109, 157, 113, 173, 103, 83, 167, 149, 163}

func isTicketValid(tkt []int, validRanges map[string]rng) bool {
	for _, val := range tkt {
		oneLegalMatch := false
		for _, validRange := range validRanges {
			if validRange.isLegal(val) {
				oneLegalMatch = true
			}
		}
		if !oneLegalMatch {
			return false
		}
	}
	return true
}

func filterInvalidTickets(allTx [][]int, validRanges map[string]rng) [][]int {
	res := [][]int{}
	for _, tkt := range allTx {
		if isTicketValid(tkt, validRanges) {
			res = append(res, tkt)
		}
	}
	return res
}

func getPossibleIndexesPerField(tkt []int, validRanges map[string]rng) map[string]intSet {
	res := map[string]intSet{}
	for fieldIndex, fieldVal := range tkt {
		for key, validRange := range validRanges {
			if validRange.isLegal(fieldVal) {
				if _, ok := res[key]; !ok {
					res[key] = map[int]bool{}
				}
				res[key].add(fieldIndex)
			}
		}
	}
	return res
}

func removeDupes(skipCategory string, toRemove int, potentialIndexMappings map[string]intSet) {
	for cat := range potentialIndexMappings {
		if cat == skipCategory {
			continue
		}
		potentialIndexMappings[cat].remove(toRemove)
	}
}

func countIndexes(potentialIndexMappings map[string]intSet) int {
	res := 0
	for _, mapping := range potentialIndexMappings {
		res += len(mapping)
	}
	return res
}

type intSet map[int]bool

func (s intSet) add(i int) {
	s[i] = true
}
func (s intSet) remove(i int) {
	delete(s, i)
}

func (s intSet) contains(i int) bool {
	return s[i]
}
func (s intSet) onlyMember() int {
	if len(s) > 1 {
		panic("set has more than 1 member")
	}
	for k := range s {
		return k
	}
	panic("set is empty")
}

func (s intSet) intersect(other intSet) intSet {
	res := intSet(map[int]bool{})
	for i := range s {
		if other.contains(i) {
			res.add(i)
		}
	}
	return res
}

func intersectAll(sets ...intSet) intSet {
	union := intSet(map[int]bool{})
	for _, s := range sets {
		for i := range s {
			union.add(i)
		}
	}
	res := intSet(map[int]bool{})
	for i := range union {
		allInclude := true
		for _, s := range sets {
			if !s.contains(i) {
				allInclude = false
				break
			}
		}
		if allInclude {
			res.add(i)
		}
	}
	return res
}

func TicketDepartureFields() {
	validRanges := loadValidTicketFieldRanges(legalRangesRawInput)
	validTickets := filterInvalidTickets(nearbyTicketValues, validRanges)

	potentialIndexMappings := map[string][]intSet{}
	for _, tkt := range validTickets {
		possibleIndexesPerField := getPossibleIndexesPerField(tkt, validRanges)
		for fieldName, mappings := range possibleIndexesPerField {
			potentialIndexMappings[fieldName] = append(potentialIndexMappings[fieldName], mappings)
		}
	}

	narrowedIndexMappings := map[string]intSet{}
	for fieldName, mappings := range potentialIndexMappings {
		narrowedIndexMappings[fieldName] = intersectAll(mappings...)
	}

	clearedFields := map[string]bool{}
	//usedMappings := [][]int{}
	for countIndexes(narrowedIndexMappings) > 20 {
		for k, v := range narrowedIndexMappings {
			if len(v) == 1 && !clearedFields[k] {
				clearedFields[k] = true
				//usedMappings = append(usedMappings, v[0])
				removeDupes(k, v.onlyMember(), narrowedIndexMappings)
				//fmt.Println("remaining indexes: ", countIndexes(narrowedIndexMappings))
			}
		}
	}

	fieldToIndex := map[string]int{}
	for k, v := range narrowedIndexMappings {
		fieldToIndex[k] = v.onlyMember()
	}

	fmt.Println("final mapping:", fieldToIndex)

	product := 1
	fmt.Println("my ticket:")
	for fieldName, index := range fieldToIndex {
		val := myTicket[index]
		fmt.Println(fieldName, ": ", val)
		if strings.HasPrefix(fieldName, "departure") {
			product *= val
		}
	}
	fmt.Println("product of departure fields", product)
}

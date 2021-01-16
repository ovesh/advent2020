package day14

import (
	"fmt"
	"strconv"
	"strings"
)

type memInstruction2 struct {
	command    string
	address    int64
	addressVal int64
	turnOnMask mask
	origMask   string
}

func (m memInstruction2) String() string {
	if m.command == "mem" {
		return fmt.Sprint("{setmem: ", m.address, "=", m.addressVal, "}")
	}
	return fmt.Sprintf("{bitmask: and=%v or=%v}")
}

func getXToBitIndexes(msk string) map[int]int {
	res := map[int]int{}
	idx := 0
	for i, c := range msk {
		if string(c) == "X" {
			res[idx] = i
			idx++
		}
	}
	return res
}

func pow2(b int) int {
	res := 1
	for i := 0; i < b; i++ {
		res <<= 1
	}
	return res
}

func isBitOn(bitIdx int, num int) bool {
	return ((1 << bitIdx) & num) > 0
}

func turnNthBitOn(bitIdx int, num int64) int64 {
	return (1 << (35 - bitIdx)) | num
}
func turnNthBitOff(bitIdx int, num int64) int64 {
	msk := int64(0)
	for i := 0; i < 36; i++ {
		if i != bitIdx {
			msk++
		}
		msk <<= 1
	}
	msk >>= 1
	return msk & num
}

func generateAddressesFor(msk string, adrs int64) []int64 {
	xToBitIndexes := getXToBitIndexes(msk)
	numBits := len(xToBitIndexes)
	upTo := pow2(numBits)
	res := make([]int64, upTo)
	for i := 0; i < upTo; i++ {
		for bitIdx := 0; bitIdx < numBits; bitIdx++ {
			if isBitOn(bitIdx, i) {
				adrs = turnNthBitOn(xToBitIndexes[bitIdx], adrs)
			} else {
				adrs = turnNthBitOff(xToBitIndexes[bitIdx], adrs)
			}
		}
		res[i] = adrs
	}
	return res
}

func loadMemInstructions2() []memInstruction2 {
	memInsts := strings.Split(memInstructionsRawInput, "\n")
	res := make([]memInstruction2, len(memInsts))
	for i, memInst := range memInsts {
		start := memInst[:4]
		if start == "mem[" {
			groups := memSetRE.FindAllStringSubmatch(memInst, -1)
			address, err := strconv.ParseInt(groups[0][1], 10, 64)
			if err != nil {
				panic("failed to parse address " + memInst)
			}
			val, err := strconv.ParseInt(groups[0][2], 10, 64)
			if err != nil {
				panic("failed to parse address " + memInst)
			}
			res[i] = memInstruction2{
				command:    "mem",
				address:    address,
				addressVal: val,
			}
		}
		if start == "mask" {
			turnOnMask := mask(0)
			msk := bitmaskSetRE.FindAllStringSubmatch(memInst, -1)[0][1]
			for i, c := range msk {
				switch string(c) {
				case "1":
					turnOnMask++
					fmt.Println("added 1   mask ", turnOnMask, "for '1' in idx", i)
				}
				turnOnMask <<= 1
			}
			turnOnMask >>= 1
			fmt.Println("final turn-on mask ", turnOnMask)
			res[i] = memInstruction2{
				command:    "bitmask",
				turnOnMask: turnOnMask,
				origMask:   msk,
			}
		}
	}

	return res
}

func stepBitmaskInstruction2(
	inst memInstruction2,
	memory map[int64]int64,
	currentOrigMask string,
	currentTurnOnMask mask,
) (string, mask) {
	if inst.command == "mem" {
		turnedOnAddress := inst.address | int64(currentTurnOnMask)
		addresses := generateAddressesFor(currentOrigMask, turnedOnAddress)
		for _, adrs := range addresses {
			memory[adrs] = inst.addressVal
		}
		return currentOrigMask, currentTurnOnMask
	}
	return inst.origMask, inst.turnOnMask
}

func MemoryAddressesTotal2() {
	pointerToVal := map[int64]int64{}
	instructions := loadMemInstructions2()
	fmt.Printf("instructions: %v\n", instructions)
	origMask := ""
	turnOnMask := mask(0)
	for _, inst := range instructions {
		origMask, turnOnMask = stepBitmaskInstruction2(inst, pointerToVal, origMask, turnOnMask)
	}

	total := int64(0)
	fmt.Println("final memory: ", pointerToVal)
	for _, val := range pointerToVal {
		total += val
	}
	fmt.Println("got total:", total)
}

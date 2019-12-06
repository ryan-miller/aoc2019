package aoc2019

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertStringToSlice(s string) []int {
	stringArray := strings.Split(s, ",")
	integerArray := []int{}
	for _, i := range stringArray {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		integerArray = append(integerArray, j)
	}
	return integerArray
}

func ConvertSliceToString(intSlice []int) string {
	// this is messy
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(intSlice)), ","), "[]")
}

func FindOpCodes(s string) []int {
	intSlice := ConvertStringToSlice(s)
	opCodeSlice := []int{}

	for i := 0; i < len(intSlice); i++ {
		if i%4 == 0 {
			opCodeSlice = append(opCodeSlice, intSlice[i])
		}
	}

	return opCodeSlice
}

func GetIntCodeValue(intCode []int, position int) int {
	return intCode[position]
}

func SetIntCodeValue(intCode []int, position int, value int) {
	intCode[position] = value
}

func ProcessIntCode(initial string) string {

	intCodeSlice := ConvertStringToSlice(initial)

	for i := 0; i < len(intCodeSlice); i += 4 {

		opCodeValue := GetIntCodeValue(intCodeSlice, i)

		switch opCodeValue {
		case 1:
			firstValue := GetIntCodeValue(intCodeSlice, GetIntCodeValue(intCodeSlice, i+1))
			secondValue := GetIntCodeValue(intCodeSlice, GetIntCodeValue(intCodeSlice, i+2))
			sum := firstValue + secondValue
			SetIntCodeValue(intCodeSlice, GetIntCodeValue(intCodeSlice, i+3), sum)
		case 2:
			firstValue := GetIntCodeValue(intCodeSlice, GetIntCodeValue(intCodeSlice, i+1))
			secondValue := GetIntCodeValue(intCodeSlice, GetIntCodeValue(intCodeSlice, i+2))
			product := firstValue * secondValue
			SetIntCodeValue(intCodeSlice, GetIntCodeValue(intCodeSlice, i+3), product)
		case 99:
			fmt.Printf("Will End: %v\n", ConvertSliceToString(intCodeSlice))
			return ConvertSliceToString(intCodeSlice)
		default:
			fmt.Errorf("Unkown OP CODE: %s", strconv.Itoa(opCodeValue))
		}

	}

	return ConvertSliceToString(intCodeSlice)
}

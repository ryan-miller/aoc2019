package aoc2019

import (
	"reflect"
	"testing"
)

/*

1,0,0,0,99 becomes 2,0,0,0,99 (1 + 1 = 2).
2,3,0,3,99 becomes 2,3,0,6,99 (3 * 2 = 6).
2,4,4,5,99,0 becomes 2,4,4,5,99,9801 (99 * 99 = 9801).
1,1,1,4,99,5,6,0,99 becomes 30,1,1,4,2,5,6,0,99.

opcodes: 1, 2, 99
1. add
2. multiply
99. end of program

start at 1st position, do something, move 4 positions for next instruction
a, b, c, d
a is opcode
if a is 1 or 2 compute with b and c and store in d position.
if a is 99 stop.

*/

func TestStringToSlice(t *testing.T) {

	intCode := "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,10,1,19,1,5,19,23,1,23,5,27,2,27,10,31,1,5,31,35,2,35,6,39,1,6,39,43,2,13,43,47,2,9,47,51,1,6,51,55,1,55,9,59,2,6,59,63,1,5,63,67,2,67,13,71,1,9,71,75,1,75,9,79,2,79,10,83,1,6,83,87,1,5,87,91,1,6,91,95,1,95,13,99,1,10,99,103,2,6,103,107,1,107,5,111,1,111,13,115,1,115,13,119,1,13,119,123,2,123,13,127,1,127,6,131,1,131,9,135,1,5,135,139,2,139,6,143,2,6,143,147,1,5,147,151,1,151,2,155,1,9,155,0,99,2,14,0,0"
	finalCode := "490668,0,0,2,1,1,2,3,1,3,4,3,1,5,0,3,2,10,1,0,1,5,19,1,1,23,5,2,2,27,10,8,1,5,31,9,2,35,6,18,1,6,39,20,2,13,43,100,2,9,47,300,1,6,51,302,1,55,9,305,2,6,59,610,1,5,63,611,2,67,13,3055,1,9,71,3058,1,75,9,3061,2,79,10,12244,1,6,83,12246,1,5,87,12,247,1,6,91,12249,1,95,13,12254,1,10,99,12258,2,6,103,24516,1,107,5,24517,1,111,13,24522,1,115,13,24527,1,13,119,24532,2,123,13,122660,1,127,6,122662,1,131,9,122665,1,5,135,122666,2,139,6,245332,2,6,143,490664,1,5,147,490665,1,151,2,490665,1,9,155,0,99,2,14,0,0"
	t.Run("convert commas separated string to slice of integers", func(t *testing.T) {
		testStrings := []struct {
			s        string
			intSlice []int
		}{
			{"1,9,10,3", []int{1, 9, 10, 3}},
			{"2,3,11,0", []int{2, 3, 11, 0}},
			{"99", []int{99}},
			{"30,40,50", []int{30, 40, 50}},
		}

		for _, ts := range testStrings {
			got := ConvertStringToSlice(ts.s)
			want := ts.intSlice

			if !reflect.DeepEqual(got, want) {
				t.Errorf("want %v got %v", want, got)
			}
		}

	})

	t.Run("convert integer slice to comma separated string", func(t *testing.T) {
		testStrings := []struct {
			intSlice []int
			s        string
		}{
			{[]int{1, 9, 10, 3}, "1,9,10,3"},
			{[]int{2, 3, 11, 0}, "2,3,11,0"},
			{[]int{99}, "99"},
			{[]int{30, 40, 50}, "30,40,50"},
		}

		for _, ts := range testStrings {
			got := ConvertSliceToString(ts.intSlice)
			want := ts.s

			if !reflect.DeepEqual(got, want) {
				t.Errorf("want %v got %v", want, got)
			}
		}

	})

	t.Run("find opCodes", func(t *testing.T) {
		testString := "1,9,10,3,2,3,11,0,99,30,40,50"
		want := []int{1, 2, 99}

		got := FindOpCodes(testString)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v got %v", want, got)
		}
	})

	t.Run("process intCodes", func(t *testing.T) {
		testIntCodes := []struct {
			Initial string
			Final   string
		}{
			{"99", "99"},
			{"1,0,0,0,99", "2,0,0,0,99"},
			{"2,3,0,3,99", "2,3,0,6,99"},
			{"2,4,4,5,99,0", "2,4,4,5,99,9801"},
			{"1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99"},
			{intCode, finalCode},
		}

		for _, ts := range testIntCodes {
			got := ProcessIntCode(ts.Initial)
			want := ts.Final

			if got != want {
				t.Errorf("want %v got %v", want, got)
			}
		}
	})

}

// input: 1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,10,1,19,1,5,19,23,1,23,5,27,2,27,10,31,1,5,31,35,2,35,6,39,1,6,39,43,2,13,43,47,2,9,47,51,1,6,51,55,1,55,9,59,2,6,59,63,1,5,63,67,2,67,13,71,1,9,71,75,1,75,9,79,2,79,10,83,1,6,83,87,1,5,87,91,1,6,91,95,1,95,13,99,1,10,99,103,2,6,103,107,1,107,5,111,1,111,13,115,1,115,13,119,1,13,119,123,2,123,13,127,1,127,6,131,1,131,9,135,1,5,135,139,2,139,6,143,2,6,143,147,1,5,147,151,1,151,2,155,1,9,155,0,99,2,14,0,0
// stringSlice := strings.Split(s, ",")

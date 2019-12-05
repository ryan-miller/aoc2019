package aoc2019

func Sum(mass int) int {
	return mass/3 - 2
}

func SumAll(masses []int) int {

	s := 0
	for i := 0; i < len(masses); i++ {
		s += Sum(masses[i])
	}
	return s
}

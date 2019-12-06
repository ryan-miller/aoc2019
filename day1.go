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

func RecursiveSum(mass int) int {

	totalFuel := 0

	for mass > 0 {

		mass = Sum(mass)
		if mass > 0 {
			totalFuel += mass
		}
	}

	return totalFuel
}

func RecursiveSumAll(masses []int) int {
	totalFuel := 0
	for i := 0; i < len(masses); i++ {
		totalFuel += RecursiveSum(masses[i])
	}
	return totalFuel
}

package module01

/*
func GCD(a, b int) int {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	var common []int

	// save the factors as a map
	factorsA := make(map[int]int)
	factorsB := make(map[int]int)
	for _, n := range Factor(primes, a) {
		factorsA[n]++
	}
	for _, n := range Factor(primes, b) {
		factorsB[n]++
	}

	for k1, v1 := range factorsA {
		for k2, v2 := range factorsB {
			if k1 == k2 && v1 > 0 && v2 > 0 {
				common = append(common, k1)
				factorsA[k1]--
				factorsB[k2]--
				break
			}
		}
	}

	total := 1
	for _, n := range common {
		total *= n
	}
	return total

}
*/

func GCD(a, b int) int {
	// Euclidean algorithm
	if b == 0 {
		return a
	}
	a, b = b, a%b
	return GCD(a, b)
}

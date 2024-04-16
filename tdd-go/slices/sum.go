package main

func Reduce[T any](a []T, f func(T, T) T, initial T) T {
	var result = initial
	for _, v := range a {
		result = f(result, v)
	}
	return result
}

func Sum(numbers []int) int {
	add := func(acc, v int) int { return acc + v }
	return Reduce(numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	add := func(acc, v []int) []int { return append(acc, Sum(v)) }
	return Reduce(numbersToSum, add, []int{})
}

func SumAllTails(numbersToSum ...[]int) []int {
	add := func(acc, v []int) []int {
		if len(v) == 0 {
			return append(acc, 0)
		} else {
			return append(acc, Sum(v[1:]))
		}
	}
	return Reduce(numbersToSum, add, []int{})
}

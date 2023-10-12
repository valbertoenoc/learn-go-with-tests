package arrays

func Sum(numbers []int) int {
	s := 0
	for _, n := range numbers {
		s += n
	}

	return s
}

func SumAll(arraysToSum ...[]int) []int {
	var sums []int

	for _, arr := range arraysToSum {
		sums = append(sums, Sum(arr))
	}

	return sums
}

func SumAllTails(arraysToSum ...[]int) []int {
	var sums []int

	for _, arr := range arraysToSum {
		if len(arr) == 0 {
			sums = append(sums, 0)
		} else {
			tail := arr[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

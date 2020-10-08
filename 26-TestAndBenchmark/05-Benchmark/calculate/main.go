package calculate

//Sum will return summary of multiple int
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		if v < 10 {
			sum += v
		} else {
			sum = 0
		}
	}

	return sum
}

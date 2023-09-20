package main


func printFinonacci(limit int) int {
	var numOne = 0
	var numTwo = 1
	var result int
	if limit == 1 {
		return numOne
	} else if limit == 2 {
		return numTwo
	} else {

		for i := 3; i <= limit; i++ {
			result = numOne + numTwo
			numOne = numTwo
			numTwo = result
		}
	}
	return result
}

// fibonacci by recursion

func fibonnaciRecursion(limit int, one int, second int) int {
	var temp int
	if limit == 1 {
		return one
	} else if limit == 2 {
		return second
	} else {
		temp=second
		second = one + second
		one = temp
		
		return fibonnaciRecursion(limit-1, one, second)
	}
}

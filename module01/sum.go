package module01

import "fmt"

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	//total := 0
	//for _, v := range numbers {
	//	// total = total + v
	//	total += v
	//}
	//return total

	fmt.Println(numbers)
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + Sum(numbers[1:])
}

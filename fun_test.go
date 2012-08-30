// Steve Phillips / elimisteve
// 2012.01.27

package fun

import (
	"fmt"
	"testing"
)

// TestRange should use t.Error instead of fmt.Printf
func TestRange(t *testing.T) {
	const errStr = "Wanted %v, got %v\n"

	if returned, correct := Range(), []int{}; fmt.Sprintf("%v", returned) != fmt.Sprintf("%v", correct) {
		fmt.Printf(errStr, correct, returned)
	}
	if returned, correct := Range(5), []int{0, 1, 2, 3, 4}; fmt.Sprintf("%v", returned) != fmt.Sprintf("%v", correct) {
		fmt.Printf(errStr, correct, returned)
	}
	if returned, correct := Range(10, 15), []int{10, 11, 12, 13, 14}; fmt.Sprintf("%v", returned) != fmt.Sprintf("%v", correct) {
		fmt.Printf(errStr, correct, returned)
	}
	if returned, correct := Range(27, 36, 2), []int{27, 29, 31, 33, 35}; fmt.Sprintf("%v", returned) != fmt.Sprintf("%v", correct) {
		fmt.Printf(errStr, correct, returned)
	}
	if returned, correct := Range(-2, 12, 3), []int{-2, 1, 4, 7, 10}; fmt.Sprintf("%v", returned) != fmt.Sprintf("%v", correct) {
		fmt.Printf(errStr, correct, returned)
	}
}

func main() {
	t := new(testing.T)
	TestRange(t)
}

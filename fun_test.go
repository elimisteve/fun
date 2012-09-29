// Steve Phillips / elimisteve
// 2012.01.27

package fun

import (
	"fmt"
	"testing"
	"time"
)

const errStr = "Wanted %v, got %v\n"

// TestRange should use t.Error instead of fmt.Printf
func TestRange(t *testing.T) {
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

func TestRandom(t *testing.T) {
	const HEX_CHARSET = "0123456789abcdef"
	older := RandStrOfLen(5, HEX_CHARSET)
	time.Sleep(1 * time.Second)
	newer := RandStrOfLen(5, HEX_CHARSET)
	if older == newer {
		fmt.Printf(errStr, older, newer)
	}
}


func main() {
	t := new(testing.T)
	TestRange(t)
	TestRandom(t)
}

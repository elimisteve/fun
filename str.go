// Steve Phillips / elimisteve
// 2014.01.05

package fun

import "math/rand"

// RandomString produces a string of length `length` using characters
// from `alphabet`
func RandomString(alphabet string, length int) string {
	s := make([]byte, len(alphabet))
	alphabetLen := len(alphabet)
	for i := 0; i < length; i++ {
		s[i] = alphabet[rand.Intn(alphabetLen)]
	}
	return string(s)
}

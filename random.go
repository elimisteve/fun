// Steve Phillips / elimisteve
// 2012.09.26

package fun

import "math/rand"

func RandStrOfLen(length int, charset string) (str string) {
	for i := 0; i < length; i++ {
		ndx := rand.Intn(len(charset))
		str += charset[ndx]
	}
	return
}

package randgen

import "math/rand"

func Rand(min int, max int) int {
	return min + rand.Intn(max-min)
}

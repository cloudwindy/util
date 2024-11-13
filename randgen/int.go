package randgen

import "math/rand"

func Int(min int, max int) int {
	return min + rand.Intn(max-min)
}

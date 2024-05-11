package ko

import "math/rand"

func Get(sep string) string {
	return left[rand.Intn(len(left))] + sep + right[rand.Intn(len(right))] //nolint:gosec // G404: Use of weak random number generator (math/rand instead of crypto/rand)
}

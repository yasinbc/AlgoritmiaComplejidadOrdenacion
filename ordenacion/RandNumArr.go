package ordenacion

import (
	"math/rand"
)

func RanNumArr(size int) []int {
	array := rand.Perm(size)

	return array
}

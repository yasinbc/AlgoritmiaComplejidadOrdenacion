package NumerosAleatorios

import (
	"math/rand"
	"time"
)

func RandInt(lower, upper int) int {
	rand.Seed(time.Now().UnixNano())
	rng := upper - lower
	return rand.Intn(rng) + lower
}

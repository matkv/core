package random

import (
	"math/rand"
	"time"
)

func Int(max int) int {
	if max <= 0 {
		max = 1000
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

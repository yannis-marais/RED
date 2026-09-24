package world

import (
	"math/rand"
	"sync"
	"time"
)

var randomSeedOnce sync.Once

func SeedRandom() {
	randomSeedOnce.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})
}

func init() {
	SeedRandom()
}

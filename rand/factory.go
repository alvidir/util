package rand

import (
	"math/rand"
	"time"
)

// New builds a brand new randomizer for a given seed
func New(seed int64) Adapter {
	source := rand.NewSource(seed)
	return Adapter{rand.New(source)}
}

// Start build a brand new randomizer with creation time as seed
func Start() Adapter {
	seed := time.Now().UnixNano()
	return New(seed)
}

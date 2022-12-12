package random

import (
	"crypto/rand"
	"math/big"
)

func GetRandom() int64 {
	// Declaring the range using max and min
	var max, min int64 = 100, 1
	// Calculate the new Maxm we will be using
	bg := big.NewInt(max - min + 1)
	// Generate random Integer,
	n, err := rand.Int(rand.Reader, bg)
	if err != nil {
		panic(err)
	}
	// add n to min to support the passed in range
	return n.Int64() + min
}

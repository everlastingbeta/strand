package strand

import (
	mathrand "math/rand"
	"time"
)

// SeededBytes returns a seeded byte slice.
// size is the length of the byte slice that will be returned.
// charset is the string of characters that the byte slice should pick from.
// seed is the [opitonal] int64 value passed into the Default Source.  If no
// seed is provided then it will default to `time.Now().UnixNano()`.
func SeededBytes(size int, charset string, seed ...int64) []byte {
	if len(seed) == 0 {
		seed = append(seed, time.Now().UnixNano())
	}

	mathrand.Seed(seed[0])
	nonce := make([]byte, size)
	for i := range nonce {
		nonce[i] = charset[mathrand.Intn(len(charset))]
	}

	return nonce
}

// SeededString returns a seeded random string.
// size is the length of the string that will be returned.
// charset is the string of characters that the string should pick from.
// seed is the [opitonal] int64 value passed into the Default Source.  If no
// seed is provided then it will default to `time.Now().UnixNano()`.
func SeededString(size int, charset string, seed ...int64) string {
	return string(SeededBytes(size, charset, seed...))
}

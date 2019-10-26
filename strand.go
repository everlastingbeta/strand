package strand

import (
	"crypto/rand"
	mathrand "math/rand"
	"time"
)

const (
	// UppercaseAlphabet represents the uppercase letters in the English alphabet
	UppercaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// LowercaseAlphabet represents the lowercase letters in the English alphabet
	LowercaseAlphabet = "abcdefghijklmnopqrstuvwxyz"

	// Alphabet represents the combination of lowercase and uppercase letters
	// in the English alphabet
	Alphabet = LowercaseAlphabet + UppercaseAlphabet

	// Numbers represnts all of the digits required to create a numeric value
	Numbers = "0123456789"

	// AlphaNumeric represents the combination of Alphabet and Numbers
	AlphaNumeric = Alphabet + Numbers

	// Symbols represents a selection of special characters
	Symbols = "<>,\\./|?;:[]{}+=_-()*&^%$#@!~"

	// ALL represents the combination of Alphanumeric and Symbols
	ALL = AlphaNumeric + Symbols
)

// Bytes returns a secure random byte slice.
// size is the length of the byte slice that will be returned.
// charset is the string of characters that the byte slice should pick from.
func Bytes(size int, charset string) ([]byte, error) {
	nonce := make([]byte, size)
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	for i, b := range nonce {
		nonce[i] = charset[b%byte(len(charset))]
	}

	return nonce, nil
}

// String returns a secure random string.
// size is the length of the string that will be returned.
// charset is the string of characters that the string should pick from.
func String(size int, charset string) (string, error) {
	nonce, err := Bytes(size, charset)
	if err != nil {
		return "", err
	}
	return string(nonce), nil
}

// SeededBytes returns a seeded byte slice.
// size is the length of the byte slice that will be returned.
// charset is the string of characters that the byte slice should pick from.
// seed is the [opitonal] int64 value passed into the Default Source.  If no
// seed is provided then it will default to `time.Now().UnixNano()`.
//   - defaults to `time.Now().UnixNano()`
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

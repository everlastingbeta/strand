package strand

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

// SeededBytes returns a deterministic byte slice based on the provided seed.
// The function generates a sequence of bytes where each byte is selected from
// the provided charset.
//
// Parameters:
//   - size: the length of the byte slice to be returned.
//   - charset: the string of characters from which the bytes will be selected.
//   - seed: optional int64 value to initialize the random source. If omitted,
//     time.Now().UnixNano() will be used as the default seed.
//
// Returns a byte slice of the specified size with characters from the charset.
//
// Security Notice: This function uses math/rand/v2 which is NOT cryptographically
// secure. For security-sensitive applications, use Bytes() instead.
func SeededBytes(size int, charset string, seed ...int64) []byte {
	seedValue := time.Now().UnixNano()
	if len(seed) > 0 {
		seedValue = seed[0]
	}

	// Create a local random number generator with the specified seed
	// Using the v2 package which has simplified APIs
	rng := rand.New(rand.NewPCG(uint64(seedValue), uint64(seedValue>>32)))

	return generateSeededBytes(rng, size, charset)
}

// SeededBytesWithContext returns a deterministic byte slice like SeededBytes,
// but accepts a context for cancellation support.
//
// Parameters:
//   - ctx: context for cancellation support.
//   - size: the length of the byte slice to be returned.
//   - charset: the string of characters from which the bytes will be selected.
//   - seed: optional int64 value to initialize the random source. If omitted,
//     time.Now().UnixNano() will be used as the default seed.
//
// Returns a byte slice of the specified size or an error if the context is canceled.
//
// Security Notice: This function uses math/rand/v2 which is NOT cryptographically
// secure. For security-sensitive applications, use BytesWithContext() instead.
func SeededBytesWithContext(ctx context.Context, size int, charset string, seed ...int64) ([]byte, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("failed to created seeded bytes due to context ending early: %w", ctx.Err())
	default:
		return SeededBytes(size, charset, seed...), nil
	}
}

// SeededString returns a deterministic string based on the provided seed.
// This is a convenience wrapper around SeededBytes that converts the result to a string.
//
// Parameters:
//   - size: the length of the string to be returned.
//   - charset: the string of characters from which the string will be generated.
//   - seed: optional int64 value to initialize the random source. If omitted,
//     time.Now().UnixNano() will be used as the default seed.
//
// Returns a string of the specified size with characters from the charset.
//
// Security Notice: This function uses math/rand/v2 which is NOT cryptographically
// secure. For security-sensitive applications, use String() instead.
func SeededString(size int, charset string, seed ...int64) string {
	return string(SeededBytes(size, charset, seed...))
}

// SeededStringWithContext returns a deterministic string like SeededString,
// but accepts a context for cancellation support.
//
// Parameters:
//   - ctx: context for cancellation support.
//   - size: the length of the string to be returned.
//   - charset: the string of characters from which the string will be generated.
//   - seed: optional int64 value to initialize the random source. If omitted,
//     time.Now().UnixNano() will be used as the default seed.
//
// Returns a string of the specified size or an error if the context is canceled.
//
// Security Notice: This function uses math/rand/v2 which is NOT cryptographically
// secure. For security-sensitive applications, use StringWithContext() instead.
func SeededStringWithContext(ctx context.Context, size int, charset string, seed ...int64) (string, error) {
	bytes, err := SeededBytesWithContext(ctx, size, charset, seed...)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// generateSeededBytes is an internal helper function that takes a random source
// and generates a byte slice of the specified size using characters from the charset.
func generateSeededBytes(rng *rand.Rand, size int, charset string) []byte {
	if size <= 0 {
		return []byte{}
	}

	if len(charset) == 0 {
		return make([]byte, size)
	}

	charsetLen := len(charset)
	nonce := make([]byte, size)

	for i := range nonce {
		nonce[i] = charset[rng.IntN(charsetLen)]
	}

	return nonce
}

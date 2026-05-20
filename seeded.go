package strand

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

// SeededBytes returns a deterministic byte slice of the given size whose
// bytes are drawn from charset using math/rand/v2 seeded by seed.
//
// If seed is omitted, time.Now().UnixNano() is used. If size <= 0 an empty
// slice is returned; if charset is empty a zero-valued slice of size is
// returned.
//
// Not cryptographically secure: use Bytes for security-sensitive callers.
func SeededBytes(size int, charset string, seed ...int64) []byte {
	seedValue := time.Now().UnixNano()
	if len(seed) > 0 {
		seedValue = seed[0]
	}

	rng := rand.New(rand.NewPCG(uint64(seedValue), uint64(seedValue>>32)))

	return generateSeededBytes(rng, size, charset)
}

// SeededBytesWithContext is SeededBytes with context cancellation. It
// returns an error wrapping ctx.Err() if ctx is already done at call time.
func SeededBytesWithContext(ctx context.Context, size int, charset string, seed ...int64) ([]byte, error) {
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("failed to create seeded bytes due to context ending early: %w", err)
	}

	return SeededBytes(size, charset, seed...), nil
}

// SeededString returns a deterministic string. See SeededBytes for rules.
//
// Not cryptographically secure: use String for security-sensitive callers.
func SeededString(size int, charset string, seed ...int64) string {
	return string(SeededBytes(size, charset, seed...))
}

// SeededStringWithContext is SeededString with context cancellation.
func SeededStringWithContext(ctx context.Context, size int, charset string, seed ...int64) (string, error) {
	b, err := SeededBytesWithContext(ctx, size, charset, seed...)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

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

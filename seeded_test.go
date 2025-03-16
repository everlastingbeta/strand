package strand_test

import (
	"context"
	"testing"
	"time"

	"github.com/everlastingbeta/strand"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestSeededBytes verifies that the SeededBytes function correctly generates
// deterministic byte slices of the requested size using only characters from
// the specified charset. It also tests that using the same seed produces
// identical results.
func TestSeededBytes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string // Description of the test case
		charset string // Character set to use
		size    int    // Size of the output to generate
		seed    int64  // Seed for the random number generator
	}{
		{
			name:    "uppercase characters",
			charset: strand.UppercaseAlphabet,
			size:    10,
			seed:    42,
		},
		{
			name:    "lowercase characters",
			charset: strand.LowercaseAlphabet,
			size:    15,
			seed:    123,
		},
		{
			name:    "mixed case alphabet",
			charset: strand.Alphabet,
			size:    20,
			seed:    9999,
		},
		{
			name:    "numbers only",
			charset: strand.Numbers,
			size:    8,
			seed:    1234567,
		},
		{
			name:    "symbols only",
			charset: strand.Symbols,
			size:    12,
			seed:    987654,
		},
		{
			name:    "all characters",
			charset: strand.ALL,
			size:    25,
			seed:    55555,
		},
		{
			name:    "custom character set",
			charset: "\\\"/|!#$%^&*()_=~funset0Fdat@",
			size:    18,
			seed:    424242,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Test with explicit seed
			nonce := strand.SeededBytes(tt.size, tt.charset, tt.seed)
			assert.Len(t, nonce, tt.size)
			assert.True(t, onlyContains(string(nonce), tt.charset))

			// Deterministic behavior check - same seed should produce same output
			nonce2 := strand.SeededBytes(tt.size, tt.charset, tt.seed)
			assert.Equal(t, nonce, nonce2, "Same seed should produce same output")

			// Default seed (time-based)
			nonceDefault := strand.SeededBytes(tt.size, tt.charset)
			assert.Len(t, nonceDefault, tt.size)
			assert.True(t, onlyContains(string(nonceDefault), tt.charset))
		})
	}
}

// TestSeededString verifies that the SeededString function correctly generates
// deterministic strings of the requested size using only characters from the
// specified charset. It also tests that using the same seed produces identical results.
func TestSeededString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string // Description of the test case
		charset string // Character set to use
		size    int    // Size of the output to generate
		seed    int64  // Seed for the random number generator
	}{
		{
			name:    "uppercase characters",
			charset: strand.UppercaseAlphabet,
			size:    10,
			seed:    42,
		},
		{
			name:    "lowercase characters",
			charset: strand.LowercaseAlphabet,
			size:    15,
			seed:    123,
		},
		{
			name:    "mixed case alphabet",
			charset: strand.Alphabet,
			size:    20,
			seed:    9999,
		},
		{
			name:    "numbers only",
			charset: strand.Numbers,
			size:    8,
			seed:    1234567,
		},
		{
			name:    "symbols only",
			charset: strand.Symbols,
			size:    12,
			seed:    987654,
		},
		{
			name:    "all characters",
			charset: strand.ALL,
			size:    25,
			seed:    55555,
		},
		{
			name:    "custom character set",
			charset: "\\\"/|!#$%^&*()_=~funset0Fdat@",
			size:    18,
			seed:    424242,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Test with explicit seed
			str := strand.SeededString(tt.size, tt.charset, tt.seed)
			assert.Len(t, str, tt.size)
			assert.True(t, onlyContains(str, tt.charset))

			// Deterministic behavior check - same seed should produce same output
			str2 := strand.SeededString(tt.size, tt.charset, tt.seed)
			assert.Equal(t, str, str2, "Same seed should produce same output")

			// Default seed (time-based)
			strDefault := strand.SeededString(tt.size, tt.charset)
			assert.Len(t, strDefault, tt.size)
			assert.True(t, onlyContains(strDefault, tt.charset))
		})
	}
}

// TestSeededBytesWithContext verifies that SeededBytesWithContext honors context cancellation.
func TestSeededBytesWithContext(t *testing.T) {
	t.Parallel()

	t.Run("successful generation", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		result, err := strand.SeededBytesWithContext(ctx, 10, strand.Alphabet, 42)
		require.NoError(t, err)
		assert.Len(t, result, 10)
	})

	t.Run("respects context cancellation", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel the context immediately

		result, err := strand.SeededBytesWithContext(ctx, 10, strand.Alphabet, 42)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.ErrorIs(t, err, context.Canceled)
	})

	t.Run("respects context timeout", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
		defer cancel()

		// Sleep to ensure timeout occurs
		time.Sleep(1 * time.Millisecond)

		result, err := strand.SeededBytesWithContext(ctx, 10, strand.Alphabet, 42)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	})
}

// TestSeededStringWithContext verifies that SeededStringWithContext honors context cancellation.
func TestSeededStringWithContext(t *testing.T) {
	t.Parallel()

	t.Run("successful generation", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		result, err := strand.SeededStringWithContext(ctx, 10, strand.Alphabet, 42)
		require.NoError(t, err)
		assert.Len(t, result, 10)
	})

	t.Run("respects context cancellation", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel the context immediately

		result, err := strand.SeededStringWithContext(ctx, 10, strand.Alphabet, 42)
		require.Error(t, err)
		assert.Empty(t, result)
		assert.ErrorIs(t, err, context.Canceled)
	})

	t.Run("respects context timeout", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
		defer cancel()

		// Sleep to ensure timeout occurs
		time.Sleep(1 * time.Millisecond)

		result, err := strand.SeededStringWithContext(ctx, 10, strand.Alphabet, 42)
		require.Error(t, err)
		assert.Empty(t, result)
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	})
}

// TestSeededDeterminism specifically focuses on verifying the deterministic
// behavior of the seeded functions, ensuring that identical seeds produce
// identical outputs across multiple calls.
func TestSeededDeterminism(t *testing.T) {
	t.Parallel()

	t.Run("bytes with same seed are deterministic", func(t *testing.T) {
		t.Parallel()

		seed := int64(12345)
		charset := strand.Alphabet
		size := 20

		result1 := strand.SeededBytes(size, charset, seed)
		result2 := strand.SeededBytes(size, charset, seed)

		assert.Equal(t, result1, result2)
	})

	t.Run("strings with same seed are deterministic", func(t *testing.T) {
		t.Parallel()

		seed := int64(67890)
		charset := strand.Alphabet
		size := 20

		result1 := strand.SeededString(size, charset, seed)
		result2 := strand.SeededString(size, charset, seed)

		assert.Equal(t, result1, result2)
	})
}

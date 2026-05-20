package strand_test

import (
	"context"
	"testing"
	"time"

	"github.com/everlastingbeta/strand"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSeededBytes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		charset string
		size    int
		seed    int64
	}{
		{name: tcUppercase, charset: strand.UppercaseAlphabet, size: 10, seed: 42},
		{name: tcLowercase, charset: strand.LowercaseAlphabet, size: 15, seed: 123},
		{name: tcMixedCase, charset: strand.Alphabet, size: 20, seed: 9999},
		{name: tcNumbers, charset: strand.Numbers, size: 8, seed: 1234567},
		{name: tcSymbols, charset: strand.Symbols, size: 12, seed: 987654},
		{name: tcAll, charset: strand.ALL, size: 25, seed: 55555},
		{name: tcCustom, charset: customCharset, size: 18, seed: 424242},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			nonce := strand.SeededBytes(tt.size, tt.charset, tt.seed)
			assert.Len(t, nonce, tt.size)
			assert.True(t, onlyContains(string(nonce), tt.charset))

			nonce2 := strand.SeededBytes(tt.size, tt.charset, tt.seed)
			assert.Equal(t, nonce, nonce2, "same seed should produce same output")

			nonceDefault := strand.SeededBytes(tt.size, tt.charset)
			assert.Len(t, nonceDefault, tt.size)
			assert.True(t, onlyContains(string(nonceDefault), tt.charset))
		})
	}
}

func TestSeededString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		charset string
		size    int
		seed    int64
	}{
		{name: tcUppercase, charset: strand.UppercaseAlphabet, size: 10, seed: 42},
		{name: tcLowercase, charset: strand.LowercaseAlphabet, size: 15, seed: 123},
		{name: tcMixedCase, charset: strand.Alphabet, size: 20, seed: 9999},
		{name: tcNumbers, charset: strand.Numbers, size: 8, seed: 1234567},
		{name: tcSymbols, charset: strand.Symbols, size: 12, seed: 987654},
		{name: tcAll, charset: strand.ALL, size: 25, seed: 55555},
		{name: tcCustom, charset: customCharset, size: 18, seed: 424242},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			str := strand.SeededString(tt.size, tt.charset, tt.seed)
			assert.Len(t, str, tt.size)
			assert.True(t, onlyContains(str, tt.charset))

			str2 := strand.SeededString(tt.size, tt.charset, tt.seed)
			assert.Equal(t, str, str2, "same seed should produce same output")

			strDefault := strand.SeededString(tt.size, tt.charset)
			assert.Len(t, strDefault, tt.size)
			assert.True(t, onlyContains(strDefault, tt.charset))
		})
	}
}

func TestSeededBytesWithContext(t *testing.T) {
	t.Parallel()

	t.Run("successful generation", func(t *testing.T) {
		t.Parallel()

		result, err := strand.SeededBytesWithContext(t.Context(), 10, strand.Alphabet, 42)
		require.NoError(t, err)
		assert.Len(t, result, 10)
	})

	t.Run("respects context cancellation", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(t.Context())
		cancel()

		result, err := strand.SeededBytesWithContext(ctx, 10, strand.Alphabet, 42)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.ErrorIs(t, err, context.Canceled)
	})

	t.Run("respects context timeout", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(t.Context(), 1*time.Nanosecond)
		defer cancel()

		time.Sleep(1 * time.Millisecond)

		result, err := strand.SeededBytesWithContext(ctx, 10, strand.Alphabet, 42)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	})
}

func TestSeededStringWithContext(t *testing.T) {
	t.Parallel()

	t.Run("successful generation", func(t *testing.T) {
		t.Parallel()

		result, err := strand.SeededStringWithContext(t.Context(), 10, strand.Alphabet, 42)
		require.NoError(t, err)
		assert.Len(t, result, 10)
	})

	t.Run("respects context cancellation", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(t.Context())
		cancel()

		result, err := strand.SeededStringWithContext(ctx, 10, strand.Alphabet, 42)
		require.Error(t, err)
		assert.Empty(t, result)
		assert.ErrorIs(t, err, context.Canceled)
	})

	t.Run("respects context timeout", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(t.Context(), 1*time.Nanosecond)
		defer cancel()

		time.Sleep(1 * time.Millisecond)

		result, err := strand.SeededStringWithContext(ctx, 10, strand.Alphabet, 42)
		require.Error(t, err)
		assert.Empty(t, result)
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	})
}

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

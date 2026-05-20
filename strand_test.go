package strand_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/everlastingbeta/strand"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Shared test-case names and the custom charset literal used across
// strand_test.go and seeded_test.go. Extracted to satisfy goconst and
// keep table-test cases readable.
const (
	tcUppercase   = "uppercase characters"
	tcLowercase   = "lowercase characters"
	tcMixedCase   = "mixed case alphabet"
	tcNumbers     = "numbers only"
	tcSymbols     = "symbols only"
	tcAll         = "all characters"
	tcCustom      = "custom character set"
	customCharset = "\\\"/|!#$%^&*()_=~funset0Fdat@"
)

func onlyContains(value, characters string) bool {
	for _, letter := range value {
		if !strings.ContainsRune(characters, letter) {
			return false
		}
	}

	return true
}

func TestBytes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		charset string
		size    int
		wantErr bool
		errType error
	}{
		{name: tcUppercase, charset: strand.UppercaseAlphabet, size: 10},
		{name: tcLowercase, charset: strand.LowercaseAlphabet, size: 15},
		{name: tcMixedCase, charset: strand.Alphabet, size: 20},
		{name: tcNumbers, charset: strand.Numbers, size: 8},
		{name: tcSymbols, charset: strand.Symbols, size: 12},
		{name: tcAll, charset: strand.ALL, size: 25},
		{name: tcCustom, charset: customCharset, size: 18},
		{name: "invalid size", charset: strand.Alphabet, size: 0, wantErr: true, errType: strand.ErrInvalidSize},
		{name: "empty charset", charset: "", size: 10, wantErr: true, errType: strand.ErrEmptyCharset},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			nonce, err := strand.Bytes(tt.size, tt.charset)

			if tt.wantErr {
				require.Error(t, err)
				assert.ErrorIs(t, err, tt.errType)

				return
			}

			require.NoError(t, err)
			assert.Len(t, nonce, tt.size)
			assert.True(t, onlyContains(string(nonce), tt.charset))
		})
	}
}

func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		charset string
		size    int
		wantErr bool
		errType error
	}{
		{name: tcUppercase, charset: strand.UppercaseAlphabet, size: 10},
		{name: tcLowercase, charset: strand.LowercaseAlphabet, size: 15},
		{name: tcMixedCase, charset: strand.Alphabet, size: 20},
		{name: tcNumbers, charset: strand.Numbers, size: 8},
		{name: tcSymbols, charset: strand.Symbols, size: 12},
		{name: tcAll, charset: strand.ALL, size: 25},
		{name: tcCustom, charset: customCharset, size: 18},
		{name: "invalid size", charset: strand.Alphabet, size: 0, wantErr: true, errType: strand.ErrInvalidSize},
		{name: "empty charset", charset: "", size: 10, wantErr: true, errType: strand.ErrEmptyCharset},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			str, err := strand.String(tt.size, tt.charset)

			if tt.wantErr {
				require.Error(t, err)
				assert.ErrorIs(t, err, tt.errType)

				return
			}

			require.NoError(t, err)
			assert.Len(t, str, tt.size)
			assert.True(t, onlyContains(str, tt.charset))
		})
	}
}

func TestBytesWithContext(t *testing.T) {
	t.Parallel()

	t.Run("successful generation", func(t *testing.T) {
		t.Parallel()

		result, err := strand.BytesWithContext(t.Context(), 10, strand.Alphabet)
		require.NoError(t, err)
		assert.Len(t, result, 10)
	})

	t.Run("respects context cancellation", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(t.Context())
		cancel()

		result, err := strand.BytesWithContext(ctx, 10, strand.Alphabet)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.ErrorIs(t, err, context.Canceled)
	})

	t.Run("respects context timeout", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(t.Context(), 1*time.Nanosecond)
		defer cancel()

		time.Sleep(1 * time.Millisecond)

		result, err := strand.BytesWithContext(ctx, 10, strand.Alphabet)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	})
}

func TestStringWithContext(t *testing.T) {
	t.Parallel()

	t.Run("successful generation", func(t *testing.T) {
		t.Parallel()

		result, err := strand.StringWithContext(t.Context(), 10, strand.Alphabet)
		require.NoError(t, err)
		assert.Len(t, result, 10)
	})

	t.Run("respects context cancellation", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(t.Context())
		cancel()

		result, err := strand.StringWithContext(ctx, 10, strand.Alphabet)
		require.Error(t, err)
		assert.Empty(t, result)
		assert.ErrorIs(t, err, context.Canceled)
	})

	t.Run("respects context timeout", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithTimeout(t.Context(), 1*time.Nanosecond)
		defer cancel()

		time.Sleep(1 * time.Millisecond)

		result, err := strand.StringWithContext(ctx, 10, strand.Alphabet)
		require.Error(t, err)
		assert.Empty(t, result)
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	})
}

func TestMustBytes(t *testing.T) {
	t.Parallel()

	t.Run("successful generation", func(t *testing.T) {
		t.Parallel()

		assert.NotPanics(t, func() {
			result := strand.MustBytes(10, strand.Alphabet)
			assert.Len(t, result, 10)
		})
	})

	t.Run("panics on error", func(t *testing.T) {
		t.Parallel()

		assert.PanicsWithError(t, strand.ErrInvalidSize.Error(), func() {
			strand.MustBytes(0, strand.Alphabet)
		})
	})
}

func TestMustString(t *testing.T) {
	t.Parallel()

	t.Run("successful generation", func(t *testing.T) {
		t.Parallel()

		assert.NotPanics(t, func() {
			result := strand.MustString(10, strand.Alphabet)
			assert.Len(t, result, 10)
		})
	})

	t.Run("panics on error", func(t *testing.T) {
		t.Parallel()

		assert.PanicsWithError(t, strand.ErrInvalidSize.Error(), func() {
			strand.MustString(0, strand.Alphabet)
		})
	})
}

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

// onlyContains verifies that the given value contains only characters
// that are present in the characters string.
//
// This helper function is used to validate that generated random strings
// only include characters from the specified charset.
//
// Parameters:
//   - value: the string to check.
//   - characters: the set of allowed characters.
//
// Returns true if all characters in value are found in the characters string,
// false otherwise.
func onlyContains(value, characters string) bool {
	for _, letter := range value {
		if !strings.Contains(characters, string(letter)) {
			return false
		}
	}

	return true
}

// TestBytes verifies that the Bytes function correctly generates random
// byte slices of the requested size using only characters from the specified charset.
// It also tests error conditions for invalid inputs.
func TestBytes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string // Description of the test case
		charset string // Character set to use
		size    int    // Size of the output to generate
		wantErr bool   // Whether an error is expected
		errType error  // Expected error type (if wantErr is true)
	}{
		{
			name:    "uppercase characters",
			charset: strand.UppercaseAlphabet,
			size:    10,
		},
		{
			name:    "lowercase characters",
			charset: strand.LowercaseAlphabet,
			size:    15,
		},
		{
			name:    "mixed case alphabet",
			charset: strand.Alphabet,
			size:    20,
		},
		{
			name:    "numbers only",
			charset: strand.Numbers,
			size:    8,
		},
		{
			name:    "symbols only",
			charset: strand.Symbols,
			size:    12,
		},
		{
			name:    "all characters",
			charset: strand.ALL,
			size:    25,
		},
		{
			name:    "custom character set",
			charset: "\\\"/|!#$%^&*()_=~funset0Fdat@",
			size:    18,
		},
		{
			name:    "invalid size",
			charset: strand.Alphabet,
			size:    0,
			wantErr: true,
			errType: strand.ErrInvalidSize,
		},
		{
			name:    "empty charset",
			charset: "",
			size:    10,
			wantErr: true,
			errType: strand.ErrEmptyCharset,
		},
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

// TestString verifies that the String function correctly generates random
// strings of the requested size using only characters from the specified charset.
// It also tests error conditions for invalid inputs.
func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string // Description of the test case
		charset string // Character set to use
		size    int    // Size of the output to generate
		wantErr bool   // Whether an error is expected
		errType error  // Expected error type (if wantErr is true)
	}{
		{
			name:    "uppercase characters",
			charset: strand.UppercaseAlphabet,
			size:    10,
		},
		{
			name:    "lowercase characters",
			charset: strand.LowercaseAlphabet,
			size:    15,
		},
		{
			name:    "mixed case alphabet",
			charset: strand.Alphabet,
			size:    20,
		},
		{
			name:    "numbers only",
			charset: strand.Numbers,
			size:    8,
		},
		{
			name:    "symbols only",
			charset: strand.Symbols,
			size:    12,
		},
		{
			name:    "all characters",
			charset: strand.ALL,
			size:    25,
		},
		{
			name:    "custom character set",
			charset: "\\\"/|!#$%^&*()_=~funset0Fdat@",
			size:    18,
		},
		{
			name:    "invalid size",
			charset: strand.Alphabet,
			size:    0,
			wantErr: true,
			errType: strand.ErrInvalidSize,
		},
		{
			name:    "empty charset",
			charset: "",
			size:    10,
			wantErr: true,
			errType: strand.ErrEmptyCharset,
		},
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

// TestBytesWithContext verifies that BytesWithContext honors context cancellation.
func TestBytesWithContext(t *testing.T) {
	t.Parallel()

	t.Run("successful generation", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		result, err := strand.BytesWithContext(ctx, 10, strand.Alphabet)
		require.NoError(t, err)
		assert.Len(t, result, 10)
	})

	t.Run("respects context cancellation", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel the context immediately

		result, err := strand.BytesWithContext(ctx, 10, strand.Alphabet)
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

		result, err := strand.BytesWithContext(ctx, 10, strand.Alphabet)
		require.Error(t, err)
		assert.Nil(t, result)
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	})
}

// TestStringWithContext verifies that StringWithContext honors context cancellation.
func TestStringWithContext(t *testing.T) {
	t.Parallel()

	t.Run("successful generation", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		result, err := strand.StringWithContext(ctx, 10, strand.Alphabet)
		require.NoError(t, err)
		assert.Len(t, result, 10)
	})

	t.Run("respects context cancellation", func(t *testing.T) {
		t.Parallel()

		ctx, cancel := context.WithCancel(context.Background())
		cancel() // Cancel the context immediately

		result, err := strand.StringWithContext(ctx, 10, strand.Alphabet)
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

		result, err := strand.StringWithContext(ctx, 10, strand.Alphabet)
		require.Error(t, err)
		assert.Empty(t, result)
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	})
}

// TestMustBytes verifies that the MustBytes function correctly generates
// random byte slices and panics when expected for invalid inputs.
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

// TestMustString verifies that the MustString function correctly generates
// random strings and panics when expected for invalid inputs.
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

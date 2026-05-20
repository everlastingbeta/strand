package strand

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
)

// Errors returned by the strand package.
var (
	ErrInvalidSize   = errors.New("invalid size: must be greater than 0")
	ErrEmptyCharset  = errors.New("invalid charset: cannot be empty")
	ErrRandomFailure = errors.New("failed to generate random bytes")
)

// Predefined character sets for common use cases.
const (
	UppercaseAlphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowercaseAlphabet string = "abcdefghijklmnopqrstuvwxyz"
	Alphabet          string = LowercaseAlphabet + UppercaseAlphabet
	Numbers           string = "0123456789"
	AlphaNumeric      string = Alphabet + Numbers
	Symbols           string = "<>,\\./|?;:[]{}+=_-()*&^%$#@!~"
	ALL               string = AlphaNumeric + Symbols
)

// Bytes returns a cryptographically secure random byte slice of the given
// size whose bytes are drawn uniformly from charset.
//
// size must be > 0 and charset must be non-empty; otherwise ErrInvalidSize
// or ErrEmptyCharset is returned. Suitable for tokens, passwords, and keys.
func Bytes(size int, charset string) ([]byte, error) {
	return BytesWithContext(context.Background(), size, charset)
}

// BytesWithContext is Bytes with context cancellation. It returns an error
// wrapping ctx.Err() if ctx is already done at call time.
func BytesWithContext(ctx context.Context, size int, charset string) ([]byte, error) {
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("failed to create secure random bytes due to context ending early: %w", err)
	}

	if size <= 0 {
		return nil, ErrInvalidSize
	}

	if len(charset) == 0 {
		return nil, ErrEmptyCharset
	}

	nonce := make([]byte, size)
	if _, err := rand.Read(nonce); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrRandomFailure, err)
	}

	charsetLen := len(charset)
	for i, b := range nonce {
		nonce[i] = charset[int(b)%charsetLen]
	}

	return nonce, nil
}

// String returns a cryptographically secure random string of the given size
// whose characters are drawn uniformly from charset. See Bytes for parameter
// rules and error conditions.
func String(size int, charset string) (string, error) {
	return StringWithContext(context.Background(), size, charset)
}

// StringWithContext is String with context cancellation.
func StringWithContext(ctx context.Context, size int, charset string) (string, error) {
	nonce, err := BytesWithContext(ctx, size, charset)
	if err != nil {
		return "", err
	}

	return string(nonce), nil
}

// MustBytes is like Bytes but panics on error. Intended for callers that
// statically know the inputs are valid (e.g. package init).
func MustBytes(size int, charset string) []byte {
	b, err := Bytes(size, charset)
	if err != nil {
		panic(err)
	}

	return b
}

// MustString is like String but panics on error.
func MustString(size int, charset string) string {
	s, err := String(size, charset)
	if err != nil {
		panic(err)
	}

	return s
}

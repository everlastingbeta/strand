package strand

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
)

// Common error types for the strand package.
var (
	ErrInvalidSize   = errors.New("invalid size: must be greater than 0")
	ErrEmptyCharset  = errors.New("invalid charset: cannot be empty")
	ErrRandomFailure = errors.New("failed to generate random bytes")
)

const (
	// UppercaseAlphabet contains all uppercase letters in the English alphabet (A-Z).
	UppercaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// LowercaseAlphabet contains all lowercase letters in the English alphabet (a-z).
	LowercaseAlphabet = "abcdefghijklmnopqrstuvwxyz"

	// Alphabet combines both lowercase and uppercase letters in the English alphabet.
	// It is equivalent to LowercaseAlphabet + UppercaseAlphabet.
	Alphabet = LowercaseAlphabet + UppercaseAlphabet

	// Numbers contains all digits (0-9) used to create numeric values.
	Numbers = "0123456789"

	// AlphaNumeric combines all letters and numbers.
	// It is equivalent to Alphabet + Numbers.
	AlphaNumeric = Alphabet + Numbers

	// Symbols contains a selection of common special characters for use in
	// password and token generation.
	Symbols = "<>,\\./|?;:[]{}+=_-()*&^%$#@!~"

	// ALL combines all alphanumeric characters and symbols.
	// It is equivalent to AlphaNumeric + Symbols.
	ALL = AlphaNumeric + Symbols
)

// Bytes generates a cryptographically secure random byte slice using characters
// from the provided charset.
//
// Parameters:
//   - size: the length of the byte slice to be returned. Must be greater than 0.
//   - charset: the string of characters from which bytes will be selected. Cannot be empty.
//
// Returns:
//   - []byte: a randomly generated byte slice of the specified size.
//   - error: an error if random generation fails or if invalid parameters are provided.
//
// This function uses crypto/rand and is suitable for security-sensitive applications
// like generating tokens, passwords, and cryptographic keys.
func Bytes(size int, charset string) ([]byte, error) {
	return BytesWithContext(context.Background(), size, charset)
}

// BytesWithContext generates a cryptographically secure random byte slice using
// characters from the provided charset, with support for context cancellation.
//
// Parameters:
//   - ctx: context for cancellation support.
//   - size: the length of the byte slice to be returned. Must be greater than 0.
//   - charset: the string of characters from which bytes will be selected. Cannot be empty.
//
// Returns:
//   - []byte: a randomly generated byte slice of the specified size.
//   - error: an error if random generation fails, if invalid parameters are provided,
//     or if the context is canceled.
//
// This function uses crypto/rand and is suitable for security-sensitive applications.
func BytesWithContext(ctx context.Context, size int, charset string) ([]byte, error) {
	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("failed to created secure random bytes due to context ending early: %w", ctx.Err())
	default:
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

		charsetLen := byte(len(charset))
		for i, b := range nonce {
			nonce[i] = charset[b%charsetLen]
		}

		return nonce, nil
	}
}

// String generates a cryptographically secure random string using characters
// from the provided charset.
//
// Parameters:
//   - size: the length of the string to be returned. Must be greater than 0.
//   - charset: the string of characters from which the result will be generated. Cannot be empty.
//
// Returns:
//   - string: a randomly generated string of the specified size.
//   - error: an error if random generation fails or if invalid parameters are provided.
//
// This function uses crypto/rand and is suitable for security-sensitive applications
// like generating passwords, session tokens, and API keys.
func String(size int, charset string) (string, error) {
	return StringWithContext(context.Background(), size, charset)
}

// StringWithContext generates a cryptographically secure random string using
// characters from the provided charset, with support for context cancellation.
//
// Parameters:
//   - ctx: context for cancellation support.
//   - size: the length of the string to be returned. Must be greater than 0.
//   - charset: the string of characters from which the result will be generated. Cannot be empty.
//
// Returns:
//   - string: a randomly generated string of the specified size.
//   - error: an error if random generation fails, if invalid parameters are provided,
//     or if the context is canceled.
//
// This function uses crypto/rand and is suitable for security-sensitive applications.
func StringWithContext(ctx context.Context, size int, charset string) (string, error) {
	nonce, err := BytesWithContext(ctx, size, charset)
	if err != nil {
		return "", err
	}

	return string(nonce), nil
}

// MustBytes works like Bytes but panics on error instead of returning it.
//
// This function is useful when you know the inputs are valid and want a simpler API
// without error checking, such as in initialization code or when using constant values.
//
// Parameters:
//   - size: the length of the byte slice to be returned. Must be greater than 0.
//   - charset: the string of characters from which bytes will be selected. Cannot be empty.
//
// Returns a randomly generated byte slice of the specified size.
//
// Panics if an error occurs during generation or if invalid parameters are provided.
func MustBytes(size int, charset string) []byte {
	b, err := Bytes(size, charset)
	if err != nil {
		panic(err)
	}

	return b
}

// MustString works like String but panics on error instead of returning it.
//
// This function is useful when you know the inputs are valid and want a simpler API
// without error checking, such as in initialization code or when using constant values.
//
// Parameters:
//   - size: the length of the string to be returned. Must be greater than 0.
//   - charset: the string of characters from which the result will be generated. Cannot be empty.
//
// Returns a randomly generated string of the specified size.
//
// Panics if an error occurs during generation or if invalid parameters are provided.
func MustString(size int, charset string) string {
	s, err := String(size, charset)
	if err != nil {
		panic(err)
	}

	return s
}

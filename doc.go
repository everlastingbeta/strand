// Package strand generates random strings and byte slices using a
// caller-supplied character set.
//
// The package exposes two families of generators:
//
//   - Cryptographically secure (Bytes, String, MustBytes, MustString and
//     their *WithContext variants) backed by crypto/rand. Use these for
//     tokens, passwords, session IDs, and any other security-sensitive
//     output.
//   - Deterministic (SeededBytes, SeededString and their *WithContext
//     variants) backed by math/rand/v2 with a caller-supplied seed. Use
//     these for reproducible fixtures, tests, and other non-security uses.
//
// Predefined character sets — UppercaseAlphabet, LowercaseAlphabet,
// Alphabet, Numbers, AlphaNumeric, Symbols, and ALL — cover common cases;
// any string may also be passed as a custom charset.
//
// Errors
//
// The secure generators return ErrInvalidSize when size <= 0,
// ErrEmptyCharset when charset is empty, and an error wrapping
// ErrRandomFailure if the underlying entropy source fails. The Must*
// variants panic on the same conditions. Context-aware variants additionally
// return an error wrapping ctx.Err() if the context is already done.
//
// Basic usage
//
//	token, err := strand.String(32, strand.AlphaNumeric)
//	if err != nil {
//	    return err
//	}
//
//	// Deterministic output for tests.
//	fixture := strand.SeededString(16, strand.Alphabet, 42)
package strand

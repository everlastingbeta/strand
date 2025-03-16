package strand_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/everlastingbeta/strand"
)

// FuzzBytes uses Go's built-in fuzzing to test the Bytes function with
// various inputs. This helps identify edge cases and potential bugs that
// might not be covered by traditional tests.
func FuzzBytes(f *testing.F) {
	// Add seed corpus
	f.Add(10, strand.Alphabet)
	f.Add(1, strand.Numbers)
	f.Add(100, strand.Symbols)

	// Fuzz test
	f.Fuzz(func(t *testing.T, size int, charset string) {
		// Skip negative sizes or empty charsets as they should return errors
		if size <= 0 || len(charset) == 0 {
			return
		}

		// Only test with reasonable sizes to avoid excessive resource usage
		if size > 1000 {
			size = 1000
		}

		bytes, err := strand.Bytes(size, charset)
		if err != nil {
			t.Fatalf("Failed to generate bytes: %v", err)
		}

		// Verify the length
		if len(bytes) != size {
			t.Errorf("Expected length %d, got %d", size, len(bytes))
		}

		// Verify each byte is from the charset
		for _, b := range bytes {
			found := false

			for i := range len(charset) {
				if b == charset[i] {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("Byte %c not found in charset %s", b, charset)
			}
		}
	})
}

// FuzzString uses Go's built-in fuzzing to test the String function with
// various inputs. This helps identify edge cases and potential bugs that
// might not be covered by traditional tests.
func FuzzString(f *testing.F) {
	// Add seed corpus
	f.Add(10, strand.Alphabet)
	f.Add(1, strand.Numbers)
	f.Add(100, strand.Symbols)

	// Fuzz test
	f.Fuzz(func(t *testing.T, size int, charset string) {
		// Skip negative sizes or empty charsets as they should return errors
		if size <= 0 || len(charset) == 0 {
			return
		}

		// Only test with reasonable sizes to avoid excessive resource usage
		if size > 1000 {
			size = 1000
		}

		str, err := strand.String(size, charset)
		if err != nil {
			t.Fatalf("Failed to generate string: %v", err)
		}

		// Verify the length
		if len(str) != size {
			t.Errorf("Expected length %d, got %d", size, len(str))
		}

		// Verify each character is from the charset
		for _, ch := range str {
			if !strings.ContainsRune(charset, ch) {
				t.Errorf("Character %c not found in charset %s", ch, charset)
			}
		}
	})
}

// FuzzSeededDeterminism uses fuzzing to verify deterministic behavior
// of the seeded functions with various inputs.
func FuzzSeededDeterminism(f *testing.F) {
	// Add seed corpus
	f.Add(10, strand.Alphabet, int64(42))
	f.Add(20, strand.Numbers, int64(123))
	f.Add(30, strand.Symbols, int64(9999))

	// Fuzz test
	f.Fuzz(func(t *testing.T, size int, charset string, seed int64) {
		// Skip negative sizes or empty charsets
		if size <= 0 || len(charset) == 0 {
			return
		}

		// Only test with reasonable sizes to avoid excessive resource usage
		if size > 1000 {
			size = 1000
		}

		// Test bytes
		bytes1 := strand.SeededBytes(size, charset, seed)
		bytes2 := strand.SeededBytes(size, charset, seed)

		if len(bytes1) != size {
			t.Errorf("Expected bytes length %d, got %d", size, len(bytes1))
		}

		// Same seed should produce identical results
		if !bytes.Equal(bytes1, bytes2) {
			t.Errorf("SeededBytes not deterministic with seed %d", seed)
		}

		// Test strings
		str1 := strand.SeededString(size, charset, seed)
		str2 := strand.SeededString(size, charset, seed)

		if len(str1) != size {
			t.Errorf("Expected string length %d, got %d", size, len(str1))
		}

		// Same seed should produce identical results
		if str1 != str2 {
			t.Errorf("SeededString not deterministic with seed %d", seed)
		}
	})
}

package strand_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/everlastingbeta/strand"
)

const fuzzMaxSize = 1000

func FuzzBytes(f *testing.F) {
	f.Add(10, strand.Alphabet)
	f.Add(1, strand.Numbers)
	f.Add(100, strand.Symbols)

	f.Fuzz(func(t *testing.T, size int, charset string) {
		if size <= 0 || len(charset) == 0 {
			return
		}

		if size > fuzzMaxSize {
			size = fuzzMaxSize
		}

		out, err := strand.Bytes(size, charset)
		if err != nil {
			t.Fatalf("Bytes(%d, %q) failed: %v", size, charset, err)
		}

		if len(out) != size {
			t.Errorf("len = %d, want %d", len(out), size)
		}

		for _, b := range out {
			if strings.IndexByte(charset, b) < 0 {
				t.Errorf("byte %q not in charset %q", b, charset)
			}
		}
	})
}

func FuzzString(f *testing.F) {
	f.Add(10, strand.Alphabet)
	f.Add(1, strand.Numbers)
	f.Add(100, strand.Symbols)

	f.Fuzz(func(t *testing.T, size int, charset string) {
		if size <= 0 || len(charset) == 0 {
			return
		}

		if size > fuzzMaxSize {
			size = fuzzMaxSize
		}

		str, err := strand.String(size, charset)
		if err != nil {
			t.Fatalf("String(%d, %q) failed: %v", size, charset, err)
		}

		if len(str) != size {
			t.Errorf("len = %d, want %d", len(str), size)
		}

		for _, ch := range str {
			if !strings.ContainsRune(charset, ch) {
				t.Errorf("rune %q not in charset %q", ch, charset)
			}
		}
	})
}

func FuzzSeededDeterminism(f *testing.F) {
	f.Add(10, strand.Alphabet, int64(42))
	f.Add(20, strand.Numbers, int64(123))
	f.Add(30, strand.Symbols, int64(9999))

	f.Fuzz(func(t *testing.T, size int, charset string, seed int64) {
		if size <= 0 || len(charset) == 0 {
			return
		}

		if size > fuzzMaxSize {
			size = fuzzMaxSize
		}

		b1 := strand.SeededBytes(size, charset, seed)
		b2 := strand.SeededBytes(size, charset, seed)

		if len(b1) != size {
			t.Errorf("SeededBytes len = %d, want %d", len(b1), size)
		}

		if !bytes.Equal(b1, b2) {
			t.Errorf("SeededBytes not deterministic with seed %d", seed)
		}

		s1 := strand.SeededString(size, charset, seed)
		s2 := strand.SeededString(size, charset, seed)

		if len(s1) != size {
			t.Errorf("SeededString len = %d, want %d", len(s1), size)
		}

		if s1 != s2 {
			t.Errorf("SeededString not deterministic with seed %d", seed)
		}
	})
}

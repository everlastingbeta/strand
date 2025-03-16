package strand_test

import (
	"context"
	"testing"

	"github.com/everlastingbeta/strand"
)

// BenchmarkBytes measures the performance of the Bytes function
// with various character set sizes and output lengths.
//
// These benchmarks help identify performance bottlenecks and allow
// for comparing the impact of implementation changes over time.
func BenchmarkBytes(b *testing.B) {
	sizes := []int{8, 16, 32, 64, 128, 256}
	charsets := []struct {
		name    string
		charset string
	}{
		{"Alphabet", strand.Alphabet},
		{"AlphaNumeric", strand.AlphaNumeric},
		{"ALL", strand.ALL},
	}

	for _, size := range sizes {
		for _, cs := range charsets {
			b.Run(cs.name+"_"+string(rune(size)), func(b *testing.B) {
				b.ReportAllocs()

				for range b.N {
					_, _ = strand.Bytes(size, cs.charset)
				}
			})
		}
	}
}

// BenchmarkBytesWithContext measures the performance of the BytesWithContext function
// with various character set sizes and output lengths.
func BenchmarkBytesWithContext(b *testing.B) {
	sizes := []int{8, 16, 32, 64, 128}
	charsets := []struct {
		name    string
		charset string
	}{
		{"Alphabet", strand.Alphabet},
		{"AlphaNumeric", strand.AlphaNumeric},
		{"ALL", strand.ALL},
	}

	ctx := context.Background()

	for _, size := range sizes {
		for _, cs := range charsets {
			b.Run(cs.name+"_"+string(rune(size)), func(b *testing.B) {
				b.ReportAllocs()

				for range b.N {
					_, _ = strand.BytesWithContext(ctx, size, cs.charset)
				}
			})
		}
	}
}

// BenchmarkString measures the performance of the String function
// with various character set sizes and output lengths.
//
// These benchmarks help identify performance bottlenecks and allow
// for comparing the impact of implementation changes over time.
func BenchmarkString(b *testing.B) {
	sizes := []int{8, 16, 32, 64, 128}
	charsets := []struct {
		name    string
		charset string
	}{
		{"Alphabet", strand.Alphabet},
		{"AlphaNumeric", strand.AlphaNumeric},
		{"ALL", strand.ALL},
	}

	for _, size := range sizes {
		for _, cs := range charsets {
			b.Run(cs.name+"_"+string(rune(size)), func(b *testing.B) {
				b.ReportAllocs()

				for range b.N {
					_, _ = strand.String(size, cs.charset)
				}
			})
		}
	}
}

// BenchmarkStringWithContext measures the performance of the StringWithContext function
// with various character set sizes and output lengths.
func BenchmarkStringWithContext(b *testing.B) {
	sizes := []int{8, 16, 32, 64, 128}
	charsets := []struct {
		name    string
		charset string
	}{
		{"Alphabet", strand.Alphabet},
		{"AlphaNumeric", strand.AlphaNumeric},
		{"ALL", strand.ALL},
	}

	ctx := context.Background()

	for _, size := range sizes {
		for _, cs := range charsets {
			b.Run(cs.name+"_"+string(rune(size)), func(b *testing.B) {
				b.ReportAllocs()

				for range b.N {
					_, _ = strand.StringWithContext(ctx, size, cs.charset)
				}
			})
		}
	}
}

// BenchmarkSeededBytes measures the performance of the SeededBytes function
// with various character set sizes and output lengths.
//
// A constant seed is used to ensure consistent benchmark results across runs.
// These benchmarks help identify performance bottlenecks and allow for comparing
// seeded operations against their cryptographically secure counterparts.
func BenchmarkSeededBytes(b *testing.B) {
	sizes := []int{8, 16, 32, 64, 128, 256, 512, 1024}
	charsets := []struct {
		name    string
		charset string
	}{
		{"Alphabet", strand.Alphabet},
		{"AlphaNumeric", strand.AlphaNumeric},
		{"ALL", strand.ALL},
	}

	seed := int64(42)

	for _, size := range sizes {
		for _, cs := range charsets {
			b.Run(cs.name+"_"+string(rune(size)), func(b *testing.B) {
				b.ReportAllocs()

				for range b.N {
					_ = strand.SeededBytes(size, cs.charset, seed)
				}
			})
		}
	}
}

// BenchmarkSeededBytesWithContext measures the performance of the SeededBytesWithContext function
// with various character set sizes and output lengths.
func BenchmarkSeededBytesWithContext(b *testing.B) {
	sizes := []int{8, 16, 32, 64, 128}
	charsets := []struct {
		name    string
		charset string
	}{
		{"Alphabet", strand.Alphabet},
		{"AlphaNumeric", strand.AlphaNumeric},
		{"ALL", strand.ALL},
	}

	ctx := context.Background()
	seed := int64(42)

	for _, size := range sizes {
		for _, cs := range charsets {
			b.Run(cs.name+"_"+string(rune(size)), func(b *testing.B) {
				b.ReportAllocs()

				for range b.N {
					_, _ = strand.SeededBytesWithContext(ctx, size, cs.charset, seed)
				}
			})
		}
	}
}

// BenchmarkSeededString measures the performance of the SeededString function
// with various character set sizes and output lengths.
//
// A constant seed is used to ensure consistent benchmark results across runs.
// These benchmarks help identify performance bottlenecks and allow for comparing
// seeded operations against their cryptographically secure counterparts.
func BenchmarkSeededString(b *testing.B) {
	sizes := []int{8, 16, 32, 64, 128, 256, 512, 1024}
	charsets := []struct {
		name    string
		charset string
	}{
		{"Alphabet", strand.Alphabet},
		{"AlphaNumeric", strand.AlphaNumeric},
		{"ALL", strand.ALL},
	}

	seed := int64(42)

	for _, size := range sizes {
		for _, cs := range charsets {
			b.Run(cs.name+"_"+string(rune(size)), func(b *testing.B) {
				b.ReportAllocs()

				for range b.N {
					_ = strand.SeededString(size, cs.charset, seed)
				}
			})
		}
	}
}

// BenchmarkSeededStringWithContext measures the performance of the SeededStringWithContext function
// with various character set sizes and output lengths.
func BenchmarkSeededStringWithContext(b *testing.B) {
	sizes := []int{8, 16, 32, 64, 128}
	charsets := []struct {
		name    string
		charset string
	}{
		{"Alphabet", strand.Alphabet},
		{"AlphaNumeric", strand.AlphaNumeric},
		{"ALL", strand.ALL},
	}

	ctx := context.Background()
	seed := int64(42)

	for _, size := range sizes {
		for _, cs := range charsets {
			b.Run(cs.name+"_"+string(rune(size)), func(b *testing.B) {
				b.ReportAllocs()

				for range b.N {
					_, _ = strand.SeededStringWithContext(ctx, size, cs.charset, seed)
				}
			})
		}
	}
}

// BenchmarkMustBytes measures the performance of the MustBytes function.
func BenchmarkMustBytes(b *testing.B) {
	size := 32
	charset := strand.Alphabet

	b.Run("MustBytes", func(b *testing.B) {
		b.ReportAllocs()

		for range b.N {
			_ = strand.MustBytes(size, charset)
		}
	})
}

// BenchmarkMustString measures the performance of the MustString function.
func BenchmarkMustString(b *testing.B) {
	size := 32
	charset := strand.Alphabet

	b.Run("MustString", func(b *testing.B) {
		b.ReportAllocs()

		for range b.N {
			_ = strand.MustString(size, charset)
		}
	})
}

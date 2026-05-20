package strand_test

import (
	"strconv"
	"testing"

	"github.com/everlastingbeta/strand"
)

// A constant seed keeps seeded-benchmark results comparable across runs.
const benchSeed int64 = 42

// Sub-benchmark name prefixes used by every charset-parameterized benchmark.
const (
	csAlphabet     = "Alphabet"
	csAlphaNumeric = "AlphaNumeric"
	csAll          = "ALL"
)

func BenchmarkBytes(b *testing.B) {
	sizes := []int{8, 16, 32, 64, 128, 256}
	charsets := []struct {
		name    string
		charset string
	}{
		{csAlphabet, strand.Alphabet},
		{csAlphaNumeric, strand.AlphaNumeric},
		{csAll, strand.ALL},
	}

	for _, size := range sizes {
		for _, cs := range charsets {
			b.Run(cs.name+"_"+strconv.Itoa(size), func(b *testing.B) {
				b.ReportAllocs()

				for b.Loop() {
					_, _ = strand.Bytes(size, cs.charset)
				}
			})
		}
	}
}

func BenchmarkString(b *testing.B) {
	sizes := []int{8, 16, 32, 64, 128}
	charsets := []struct {
		name    string
		charset string
	}{
		{csAlphabet, strand.Alphabet},
		{csAlphaNumeric, strand.AlphaNumeric},
		{csAll, strand.ALL},
	}

	for _, size := range sizes {
		for _, cs := range charsets {
			b.Run(cs.name+"_"+strconv.Itoa(size), func(b *testing.B) {
				b.ReportAllocs()

				for b.Loop() {
					_, _ = strand.String(size, cs.charset)
				}
			})
		}
	}
}

func BenchmarkSeededBytes(b *testing.B) {
	sizes := []int{8, 16, 32, 64, 128, 256, 512, 1024}
	charsets := []struct {
		name    string
		charset string
	}{
		{csAlphabet, strand.Alphabet},
		{csAlphaNumeric, strand.AlphaNumeric},
		{csAll, strand.ALL},
	}

	for _, size := range sizes {
		for _, cs := range charsets {
			b.Run(cs.name+"_"+strconv.Itoa(size), func(b *testing.B) {
				b.ReportAllocs()

				for b.Loop() {
					_ = strand.SeededBytes(size, cs.charset, benchSeed)
				}
			})
		}
	}
}

func BenchmarkSeededString(b *testing.B) {
	sizes := []int{8, 16, 32, 64, 128, 256, 512, 1024}
	charsets := []struct {
		name    string
		charset string
	}{
		{csAlphabet, strand.Alphabet},
		{csAlphaNumeric, strand.AlphaNumeric},
		{csAll, strand.ALL},
	}

	for _, size := range sizes {
		for _, cs := range charsets {
			b.Run(cs.name+"_"+strconv.Itoa(size), func(b *testing.B) {
				b.ReportAllocs()

				for b.Loop() {
					_ = strand.SeededString(size, cs.charset, benchSeed)
				}
			})
		}
	}
}

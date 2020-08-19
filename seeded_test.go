package strand_test

import (
	mathrand "math/rand"
	"testing"
	"time"

	"github.com/everlastingbeta/strand"
	"github.com/stretchr/testify/assert"
)

func TestSeededBytes(t *testing.T) {
	assert := assert.New(t)

	seededRand := mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	tests := []struct {
		Name    string
		Charset string
		Size    int
	}{
		{
			Name:    "will return a seeded-byte slice with uppercase characters",
			Charset: strand.UppercaseAlphabet,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a seeded-byte slice with lowercase characters",
			Charset: strand.LowercaseAlphabet,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a seeded-byte slice with both uppercase and lowercase characters",
			Charset: strand.Alphabet,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a seeded-byte slice with numbers",
			Charset: strand.Numbers,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a seeded-byte slice with symbols",
			Charset: strand.Symbols,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a seeded-byte slice with all possible characters",
			Charset: strand.ALL,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a seeded-byte slice with the customized set of characters",
			Charset: "\\\"/|!#$%^&*()_=~funset0Fdat@",
			Size:    seededRand.Intn(20) + 1,
		},
	}
	for _, test := range tests {
		nonce := strand.SeededString(test.Size, test.Charset)

		if !assert.Equal(test.Size, len(nonce), test.Name) {
			t.FailNow()
		}

		assert.True(onlyContains(nonce, test.Charset), test.Name)
	}
}

func TestSeededString(t *testing.T) {
	assert := assert.New(t)

	seededRand := mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	tests := []struct {
		Name    string
		Charset string
		Size    int
	}{
		{
			Name:    "will return a seeded-string with uppercase characters",
			Charset: strand.UppercaseAlphabet,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a seeded-string with lowercase characters",
			Charset: strand.LowercaseAlphabet,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a seeded-string with both uppercase and lowercase characters",
			Charset: strand.Alphabet,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a seeded-string with numbers",
			Charset: strand.Numbers,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a seeded-string with symbols",
			Charset: strand.Symbols,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a seeded-string with all possible characters",
			Charset: strand.ALL,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a seeded-string with the customized set of characters",
			Charset: "\\\"/|!#$%^&*()_=~funset0Fdat@",
			Size:    seededRand.Intn(20) + 1,
		},
	}
	for _, test := range tests {
		nonce := strand.SeededString(test.Size, test.Charset)

		if !assert.Equal(test.Size, len(nonce), test.Name) {
			t.FailNow()
		}

		assert.True(onlyContains(nonce, test.Charset), test.Name)
	}
}

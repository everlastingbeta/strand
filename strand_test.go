package strand_test

import (
	mathrand "math/rand"
	"strings"
	"testing"
	"time"

	"github.com/everlastingbeta/strand"
	"github.com/stretchr/testify/assert"
)

// onlyContains is a simple helper method that allows for us to verify that
// the generated value only contains the characters that we had expected
func onlyContains(value, characters string) bool {
	for _, letter := range value {
		if !strings.Contains(characters, string(letter)) {
			return false
		}
	}
	return true
}

func TestBytes(t *testing.T) {
	assert := assert.New(t)

	seededRand := mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	tests := []struct {
		Name    string
		Charset string
		Size    int
	}{
		{
			Name:    "will return a byte slice with uppercase characters",
			Charset: strand.UppercaseAlphabet,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a byte slice with lowercase characters",
			Charset: strand.LowercaseAlphabet,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a byte slice with both uppercase and lowercase characters",
			Charset: strand.Alphabet,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a byte slice with numbers",
			Charset: strand.Numbers,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a byte slice with symbols",
			Charset: strand.Symbols,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a byte slice with all possible characters",
			Charset: strand.ALL,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a byte slice with the customized set of characters",
			Charset: "\\\"/|!#$%^&*()_=~funset0Fdat@",
			Size:    seededRand.Intn(20) + 1,
		},
	}
	for _, test := range tests {
		nonce, err := strand.Bytes(test.Size, test.Charset)
		if !assert.NoError(err, test.Name) {
			t.FailNow()
		}

		if !assert.Equal(test.Size, len(nonce), test.Name) {
			t.FailNow()
		}

		assert.True(onlyContains(string(nonce), test.Charset), test.Name)
	}
}

func TestString(t *testing.T) {
	assert := assert.New(t)

	seededRand := mathrand.New(mathrand.NewSource(time.Now().UnixNano()))
	tests := []struct {
		Name    string
		Charset string
		Size    int
	}{
		{
			Name:    "will return a string with uppercase characters",
			Charset: strand.UppercaseAlphabet,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a string with lowercase characters",
			Charset: strand.LowercaseAlphabet,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a string with both uppercase and lowercase characters",
			Charset: strand.Alphabet,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a string with numbers",
			Charset: strand.Numbers,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a string with symbols",
			Charset: strand.Symbols,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a string with all possible characters",
			Charset: strand.ALL,
			Size:    seededRand.Intn(20) + 1,
		}, {
			Name:    "will return a string with the customized set of characters",
			Charset: "\\\"/|!#$%^&*()_=~funset0Fdat@",
			Size:    seededRand.Intn(20) + 1,
		},
	}
	for _, test := range tests {
		nonce, err := strand.String(test.Size, test.Charset)
		if !assert.NoError(err, test.Name) {
			t.FailNow()
		}

		if !assert.Equal(test.Size, len(nonce), test.Name) {
			t.FailNow()
		}

		assert.True(onlyContains(nonce, test.Charset), test.Name)
	}
}

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

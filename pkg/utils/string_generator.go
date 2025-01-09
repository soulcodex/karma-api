package utils

import (
	"math/rand"
	"time"
)

type StringGenerator interface {
	Generate(size int) string
}

type RandomStringGenerator struct{}

func NewRandomStringGenerator() *RandomStringGenerator {
	return &RandomStringGenerator{}
}

func (rsg *RandomStringGenerator) Generate(size int) string {
	return randomString(size)
}

type FixedStringGenerator struct {
	value string
}

func NewFixedStringGenerator(value string) *FixedStringGenerator {
	return &FixedStringGenerator{
		value: value,
	}
}

func (fsg *FixedStringGenerator) Generate(size int) string {
	return fsg.value
}

func randomString(size int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.NewSource(time.Now().UnixNano())

	b := make([]rune, size)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

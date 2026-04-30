// Package common for common utils
package common

import (
	"math/rand"
	"strings"
)

func GetMapKeys[T any](m map[string]T) []string {
	total := len(m)
	keys := make([]string, total)
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func RandKey(chars int) string {
	const keyAlphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyAlphabetLen = len(keyAlphabet)

	sb := strings.Builder{}
	sb.Grow(chars)

	for range chars {
		//nolint:gosec
		c := keyAlphabet[rand.Int63()%int64(keyAlphabetLen)]
		sb.WriteByte(c)
	}

	return sb.String()
}

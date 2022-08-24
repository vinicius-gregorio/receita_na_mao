package util

import (
	"math/rand"
	"strings"
)

func RandomString(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	k := len(chars)

	for i := 0; i < length; i++ {
		c := chars[rand.Intn(k)]

		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomName() string {
	return RandomString(999)
}

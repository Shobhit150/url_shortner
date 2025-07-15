package utils

import (
	"math/rand"
)

const slugChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateSlug() string {
	slugLength := 6
	slug := make([] byte, slugLength)
	for i:= 0; i<slugLength; i++ {
		slug[i] = slugChars[rand.Intn(len(slugChars))] 
	}
	return string(slug)
}
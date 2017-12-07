package test

import (
	"testing"
	"math/rand"
)

var lowercase = []rune("abcdefghijklmnopqrstuvwxyz")
var uppercase = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var characters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
var lettersOnly = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomAlphaNum(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = characters[rand.Intn(len(characters))]
	}
	return string(b)
}

func RandomAlphaLower(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = lowercase[rand.Intn(len(lowercase))]
	}
	return string(b)
}

func RandomAlphaUpper(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = uppercase[rand.Intn(len(uppercase))]
	}
	return string(b)
}

func RandomAlphaMixed(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = lettersOnly[rand.Intn(len(lettersOnly))]
	}
	return string(b)
}

func RandomEmail() string {
	before := make([]rune, 8)
	middle := make([]rune, 8)
	end := make([]rune, 3)

	for i := range before {
		before[i] = characters[rand.Intn(len(characters))]
	}

	for i := range middle {
		middle[i] = characters[rand.Intn(len(characters))]
	}

	for i := range end {
		end[i] = characters[rand.Intn(len(characters))]
	}

	return string(before) + "@" + string(middle) + "." + string(end)
}


func OutExpCheck(t *testing.T, output, expected interface{}) {
	InOutExpCheckMsg(t, nil, output, expected, "")
}

func InOutExpCheck(t *testing.T, input, output, expected interface{}) {
	InOutExpCheckMsg(t, input, output, expected, "")
}

func InOutExpCheckMsg(t *testing.T, input, output, expected interface{}, message string) {
	if output != expected {
		if input != nil {
			t.Errorf("Value check failed. Input: %v. Output: %v. Expected: %v.", input, output, expected)
		} else {
			t.Errorf("Value check failed. Output: %v. Expected: %v.", output, expected)
		}
	}

	if message != "" {
		t.Log(message)
	}
}

func ErrorCheck(t *testing.T, output, expected error) {
	ErrorCheckMessage(t, output, expected, "")
}

func ErrorCheckMessage(t *testing.T, output, expected error, message string) {
	if output != expected {
		t.Errorf("Error check failed. Output: %v. Expected: %v.", output, expected)
	}

	if message != "" {
		t.Log(message)
	}
}


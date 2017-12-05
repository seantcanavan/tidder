package tools

import (
	"math/rand"
	"strings"
	"strconv"
	"bytes"
	"regexp"
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

func IsAllUpper(input string) bool {
	return containsCase(input, true, true)
}

func IsAllLower(input string) bool {
	return containsCase(input, false, true)
}

func ContainsUpper(input string) bool {
	return containsCase(input, true, false)
}

func ContainsLower(input string) bool {
	return containsCase(input, false, false)
}

func ContainsNumeral(input string) bool {

	for i := 0; i < len(input); i++ {
		if _, err := strconv.ParseInt(string(input[i]),10,64); err == nil {
			return true
		}
	}

	return false
}

func StrConcat(values ...string) string {
	var buffer bytes.Buffer

	for _, element := range values {
		buffer.WriteString(element)
	}

	return buffer.String()
}

func StrTrimConcat(values ...string) string {
	var buffer bytes.Buffer

	for _, element := range values {
		buffer.WriteString(strings.TrimSpace(element))
	}

	return buffer.String()
}

func containsCase(input string, upper bool, exclusive bool) bool {

	if len(input) < 1 {
		return false
	}

	if upper {
		if exclusive {
			return strings.ToUpper(input) == input
		} else {
			return strings.ToLower(input) != input
		}
	} else {
		if exclusive {
			return strings.ToLower(input) == input
		} else {
			return strings.ToUpper(input) != input
		}
	}

	return false
}

func IsValidEmail(emailAddress string) bool {
	match, err := regexp.MatchString("[a-zA-Z0-9]{1,}@[a-zA-Z0-9]{1,}\\.[a-zA-Z0-9]{1,}", emailAddress)
	return err == nil && match
}

func IsValidUserName(userName string) bool {
	match, err := regexp.MatchString("^[a-zA-Z0-9]{1,}$", userName)
	return err == nil && match
}

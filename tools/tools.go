package tools

import (
	"strings"
	"strconv"
	"bytes"
	"regexp"
)


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
	match, err := regexp.MatchString("^[a-zA-Z0-9]{1,}@[a-zA-Z0-9]{1,}\\.[a-zA-Z0-9]{1,}$", emailAddress)
	return err == nil && match
}

func IsValidUserName(userName string) bool {
	match, err := regexp.MatchString("^[a-zA-Z0-9]{1,}$", userName)
	return err == nil && match
}

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}
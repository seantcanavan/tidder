package tools

import (
	"testing"
	"strings"
)

func TestRandomAlphaNum(t *testing.T) {
	tables := []struct {
		input int
	} {
		{3},
		{5},
		{7},
		{9},
	}

	for _, table := range tables {
		result := RandomAlphaNum(table.input)
		if len(result) != table.input {
			t.Errorf("Result was incorrect. Expected: %v. Got: %v.", table.input, len(result))
		}
	}
}

func TestRandomAlphaLower(t *testing.T) {
	tables := []struct {
		input int
	} {
		{3},
		{5},
		{7},
		{9},
	}

	for _, table := range tables {
		result := RandomAlphaLower(table.input)

		if len(result) != table.input {
			t.Errorf("Result was incorrect. Expected: %v. Got: %v.", table.input, len(result))
		}

		if ContainsUpper(result) {
			t.Errorf("Result was incorrect. Input: %v should not contain upper case.", result)
		}

		if ContainsNumeral(result) {
			t.Errorf("Result was incorrect. Input: %v should not contain numerical value.", result)
		}
	}
}

func TestRandomAlphaUpper(t *testing.T) {
	tables := []struct {
		input int
	} {
		{3},
		{5},
		{7},
		{9},
	}

	for _, table := range tables {
		result := RandomAlphaUpper(table.input)

		if len(result) != table.input {
			t.Errorf("Result was incorrect. Expected: %v. Got: %v.", table.input, len(result))
		}

		if ContainsLower(result) {
			t.Errorf("Result was incorrect. Input: %v should not contain upper case.", result)
		}

		if ContainsNumeral(result) {
			t.Errorf("Result was incorrect. Input: %v should not contain numerical value.", result)
		}
	}
}

func TestRandomAlphaMixed(t *testing.T) {
	tables := []struct {
		input int
	} {
		{3},
		{5},
		{7},
		{9},
	}

	for _, table := range tables {
		result := RandomAlphaMixed(table.input)

		if len(result) != table.input {
			t.Errorf("Result was incorrect. Expected: %v. Got: %v.", table.input, len(result))
		}

		if ContainsNumeral(result) {
			t.Errorf("Result was incorrect. Input: %v should not contain numerical value.", result)
		}
	}
}

func TestRandomEmail(t *testing.T) {
	for i := 0; i < 5; i++ {
		result := RandomEmail()
		lastAt := strings.LastIndex(result, "@")
		lastPeriod := strings.LastIndex(result, ".")

		atCount := strings.Count(result, "@")
		periodCount := strings.Count(result, ".")

		if lastPeriod <= lastAt {
			t.Errorf("Result was incorrect. @ should come before . in email address. Result: %v.", result)
		}

		if atCount != 1 {
			t.Errorf("Should encounter only one @ in email. Result: %v.", result)
		}

		if periodCount != 1 {
			t.Errorf("Should encounter only one . in email. Result: %v.", result)
		}
	}
}

func TestIsAllUpper(t *testing.T) {
	tables := []struct {
		input string
		result bool
	} {
		{"UPPERCASE", true},
		{"UPPER CASE", true},
		{"CRUISE CONTROL", true},
		{"CRUISE CONTRoL", false},
		{"lowercase", false},
	}

	for _, table := range tables {
		result := IsAllUpper(table.input)
		if result != table.result {
			t.Errorf("Result was incorrect. Expected: %v. Got: %v. Input: %v.", table.result, result, table.input)
		}
	}
}

func TestIsAllLower(t *testing.T) {
	tables := []struct {
		input string
		result bool
	} {
		{"UPPERCASE", false},
		{"CRUISE CONTROL", false},
		{"CRUISE CONTRoL", false},
		{"lowercase", true},
		{"lower case", true},
	}

	for _, table := range tables {
		result := IsAllLower(table.input)
		if result != table.result {
			t.Errorf("Result was incorrect. Expected: %v. Got: %v. Input: %v.", table.result, result, table.input)
		}
	}
}

func TestContainsUpper(t *testing.T) {
	tables := []struct {
		input string
		result bool
	} {
		{"UPPERCASE", true},
		{"CRUISE CONTROL", true},
		{"CRUISE CONTRoL", true},
		{"lowercase", false},
		{"lower case", false},
	}

	for _, table := range tables {
		result := ContainsUpper(table.input)
		if result != table.result {
			t.Errorf("Result was incorrect. Expected: %v. Got: %v. Input: %v.", table.result, result, table.input)
		}
	}
}

func TestContainsLower(t *testing.T) {
	tables := []struct {
		input string
		result bool
	} {
		{"UPPERCASE", false},
		{"CRUISE CONTROL", false},
		{"CRUISE CONTRoL", true},
		{"lowercase", true},
		{"lower case", true},
	}

	for _, table := range tables {
		result := ContainsLower(table.input)
		if result != table.result {
			t.Errorf("Result was incorrect. Expected: %v. Got: %v. Input: %v.", table.result, result, table.input)
		}
	}
}

func TestContainsNumeral(t *testing.T) {
	tables := []struct {
		input string
		result bool
	} {
		{"UPPER12CASE", true},
		{"CRUISE 1 CONTROL", true},
		{"CRUISE 2 CONTRoL", true},
		{"lower12case", true},
		{"lower case1", true},
		{"UPPERCASE", false},
		{"CRUISE CONTROL", false},
		{"CRUISE CONTRoL", false},
		{"lowercase", false},
		{"lower case", false},
	}

	for _, table := range tables {
		result := ContainsNumeral(table.input)
		if result != table.result {
			t.Errorf("Result was incorrect. Expected: %v. Got: %v. Input: %v.", table.result, result, table.input)
		}
	}
}

func TestStrConcat(t *testing.T) {
	tables := []struct {
		input []string
		result string
	} {
		{[]string{"one", "two", "three"}, "onetwothree"},
		{[]string{"three", "four", "five"}, "threefourfive"},
	}

	for _, table := range tables {
		result := StrConcat(table.input...)

		if result != table.result {
			t.Errorf("Result was incorrect. Expected: %v. Got: %v. Input: %v.", table.result, result, table.input)
		}
	}
}

func TestStrTrimConcat(t *testing.T) {
	tables := []struct {
		input []string
		result string
	} {
		{[]string{"one ", " two", " three "}, "onetwothree"},
		{[]string{" three", "four ", " five "}, "threefourfive"},
	}

	for _, table := range tables {
		result := StrTrimConcat(table.input...)

		if result != table.result {
			t.Errorf("Result was incorrect. Expected: %v. Got: %v. Input: %v.", table.result, result, table.input)
		}
	}
}

func TestIsValidEmail(t *testing.T) {
	tables := []struct {
		input string
		boolResult bool
	} {
		{"1@2.3", true},
		{"seantcanavan@github.com", true},
		{"teststudent@co.uk", true},
		{"test.student.org@com", false},
	}

	for _, table := range tables {
		boolResult := IsValidEmail(table.input)

		if boolResult != table.boolResult {
			t.Errorf("Result was incorrect. Expected: %v. Got: %v. Input: %v.", table.boolResult, boolResult, table.input)
		}
	}
}

func TestIsValidUserName(t *testing.T) {
	tables := []struct {
		input string
		boolResult bool
	} {
		{"1@2.3", false},
		{"seantcanavan@github.com", false},
		{"teststudent@co.uk", false},
		{"test.student.org@com", false},
		{"3987423894723894728934", true},
		{"adsffioqquupoivPOIFiofPOIFU", true},
		{"404f9f6ygg1k4hjg879sdfsfdlkj", true},
	}

	for _, table := range tables {
		boolResult := IsValidUserName(table.input)

		if boolResult != table.boolResult {
			t.Errorf("Result was incorrect. Expected: %v. Got: %v. Input: %v.", table.boolResult, boolResult, table.input)
		}
	}
}
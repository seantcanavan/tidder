package tools

import (
	"testing"
	"strings"
	"github.com/seantcanavan/tidder/test"
	"github.com/satori/go.uuid"
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
		result := test.RandomAlphaNum(table.input)
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
		result := test.RandomAlphaLower(table.input)

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
		result := test.RandomAlphaUpper(table.input)

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
		result := test.RandomAlphaMixed(table.input)

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
		result := test.RandomEmail()
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
		input    string
		expected bool
	} {
		{"UPPERCASE", true},
		{"UPPER CASE", true},
		{"CRUISE CONTROL", true},
		{"CRUISE CONTRoL", false},
		{"lowercase", false},
	}

	for _, table := range tables {
		result := IsAllUpper(table.input)
		if result != table.expected {
			t.Errorf("Result was incorrect. Expected: %v. Got: %v. Input: %v.", table.expected, result, table.input)
		}
	}
}

func TestIsAllLower(t *testing.T) {
	tables := []struct {
		input    string
		expected bool
	} {
		{"UPPERCASE", false},
		{"CRUISE CONTROL", false},
		{"CRUISE CONTRoL", false},
		{"lowercase", true},
		{"lower case", true},
	}

	for _, table := range tables {
		result := IsAllLower(table.input)
		if result != table.expected {
			t.Errorf("Result was incorrect. Expected: %v. Got: %v. Input: %v.", table.expected, result, table.input)
		}
	}
}

func TestContainsUpper(t *testing.T) {
	tables := []struct {
		input    string
		expected bool
	} {
		{"UPPERCASE", true},
		{"CRUISE CONTROL", true},
		{"CRUISE CONTRoL", true},
		{"lowercase", false},
		{"lower case", false},
	}

	for _, table := range tables {
		result := ContainsUpper(table.input)
		if result != table.expected {
			t.Errorf("Result was incorrect. Expected: %v. Got: %v. Input: %v.", table.expected, result, table.input)
		}
	}
}

func TestContainsLower(t *testing.T) {
	tables := []struct {
		input    string
		expected bool
	} {
		{"UPPERCASE", false},
		{"CRUISE CONTROL", false},
		{"CRUISE CONTRoL", true},
		{"lowercase", true},
		{"lower case", true},
	}

	for _, table := range tables {
		result := ContainsLower(table.input)
		if result != table.expected {
			t.Errorf("Result was incorrect. Expected: %v. Got: %v. Input: %v.", table.expected, result, table.input)
		}
	}
}

func TestContainsNumeral(t *testing.T) {
	tables := []struct {
		input    string
		expected bool
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
		if result != table.expected {
			t.Errorf("Result was incorrect. Expected: %v. Got: %v. Input: %v.", table.expected, result, table.input)
		}
	}
}

func TestStrConcat(t *testing.T) {
	tables := []struct {
		input    []string
		expected string
	} {
		{[]string{"one", "two", "three"}, "onetwothree"},
		{[]string{"three", "four", "five"}, "threefourfive"},
	}

	for _, table := range tables {
		result := StrConcat(table.input...)

		if result != table.expected {
			t.Errorf("Result was incorrect. Expected: %v. Got: %v. Input: %v.", table.expected, result, table.input)
		}
	}
}

func TestStrTrimConcat(t *testing.T) {
	tables := []struct {
		input    []string
		expected string
	} {
		{[]string{"one ", " two", " three "}, "onetwothree"},
		{[]string{" three", "four ", " five "}, "threefourfive"},
	}

	for _, table := range tables {
		result := StrTrimConcat(table.input...)

		if result != table.expected {
			t.Errorf("Result was incorrect. Expected: %v. Got: %v. Input: %v.", table.expected, result, table.input)
		}
	}
}

func TestIsValidEmail(t *testing.T) {
	tables := []struct {
		input    string
		expected bool
	} {
		{"1@2.3", true},
		{"seantcanavan@github.com", true},
		{"teststudent@co.uk", true},
		{"test.student.org@com", false},
	}

	for _, table := range tables {
		result := IsValidEmail(table.input)
		test.InOutExpCheck(t, table.input, result, table.expected)
	}
}

func TestIsValidUserName(t *testing.T) {
	tables := []struct {
		input    string
		expected bool
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
		result := IsValidUserName(table.input)
		test.InOutExpCheck(t, table.input, result, table.expected)
	}
}

func TestIsValidUUID(t *testing.T) {
	uuid01 := uuid.NewV4().String()
	uuid02 := uuid.NewV4().String()
	uuid03 := uuid.NewV4().String()
	uuid04 := uuid.NewV4().String()
	uuid05 := uuid.NewV4().String()
	uuid06 := uuid.NewV4().String()
	uuid07 := test.RandomAlphaNum(20)
	uuid08 := test.RandomAlphaNum(20)
	uuid09 := test.RandomAlphaNum(20)
	uuid10 := test.RandomAlphaNum(20)

	tables := []struct {
		input    string
		expected bool
	} {
		{uuid01, true},
		{uuid02, true},
		{uuid03, true},
		{uuid04, true},
		{uuid05, true},
		{uuid06, true},
		{uuid07, false},
		{uuid08, false},
		{uuid09, false},
		{uuid10, false},
	}

	for _, table := range tables {
		result := IsValidUUID(table.input)
		test.InOutExpCheck(t, table.input, result, table.expected)
	}
}
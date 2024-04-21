package password

import (
	"testing"
)

func TestValidatePasswordError(t *testing.T) {
	testcases := []struct {
		given    string
		expected string
	}{
		{
			given:    "abcdefgh",
			expected: "the password must contain at least 2 numbers",
		},
		{
			given:    "abcdef12",
			expected: "the password must contain at least one capital letter",
		},
		{
			given:    "abcdeF12",
			expected: "the password must contain at least one special character",
		},
	}

	for _, testcase := range testcases {
		t.Run("", func(t *testing.T) {
			actual := ValidatePassword(testcase.given)

			if testcase.expected != actual.Error() {
				t.Errorf("given a password %s expected error is %v but actual error is %v", testcase.given, testcase.expected, actual)
			}
		})
	}
}

func TestValidatePasswordValid(t *testing.T) {
	testcases := []struct {
		given string
	}{
		{
			given: "@bcdeF12",
		},
	}

	for _, testcase := range testcases {
		t.Run("", func(t *testing.T) {
			actual := ValidatePassword(testcase.given)

			if actual != nil {
				t.Errorf("given a password %s should valid but actual error is %v", testcase.given, actual)
			}
		})
	}
}

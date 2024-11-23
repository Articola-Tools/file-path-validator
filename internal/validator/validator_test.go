package validator

import (
	"testing"
)

func TestIsSnakeCase(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid snake_case",
			input:    "snake_case",
			expected: true,
		},
		{
			name:     "Valid snake_case with numbers",
			input:    "snake_case_123",
			expected: true,
		},
		{
			name:     "Valid single word",
			input:    "snake",
			expected: true,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "Invalid: PascalCase",
			input:    "PascalCase",
			expected: false,
		},
		{
			name:     "Invalid: camelCase",
			input:    "camelCase",
			expected: false,
		},
		{
			name:     "Invalid: Contains uppercase",
			input:    "snake_Case",
			expected: false,
		},
		{
			name:     "Invalid: Starts with underscore",
			input:    "_snake_case",
			expected: false,
		},
		{
			name:     "Invalid: Ends with underscore",
			input:    "snake_case_",
			expected: false,
		},
		{
			name:     "Invalid: Double underscore",
			input:    "snake__case",
			expected: false,
		},
		{
			name:     "Invalid: Contains hyphen",
			input:    "snake-case",
			expected: false,
		},
		{
			name:     "Invalid: Contains space",
			input:    "snake case",
			expected: false,
		},
		{
			name:     "Invalid: Contains special characters",
			input:    "snake@case",
			expected: false,
		},
		{
			name:     "Valid: Long snake_case",
			input:    "this_is_a_very_long_snake_case_string_with_numbers_123",
			expected: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			result := IsSnakeCase(testCase.input)

			if result != testCase.expected {
				t.Errorf("IsSnakeCase(%q) = %v, want %v", testCase.input, result, testCase.expected)
			}
		})
	}
}

func TestIsPascalCase(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty string", "", false},
		{"Single lowercase letest_caseer", "a", false},
		{"Single uppercase letter", "A", true},
		{"Valid PascalCase", "PascalCase", true},
		{"Valid PascalCase with numbers", "Pascal123Case", true},
		{"camelCase", "camelCase", false},
		{"snake_case", "snake_case", false},
		{"UPPERCASE", "UPPERCASE", true},
		{"PascalCase with space", "Pascal Case", false},
		{"PascalCase with underscore", "Pascal_Case", false},
		{"PascalCase with hyphen", "Pascal-Case", false},
		{"PascalCase starting with number", "1PascalCase", false},
		{"Complex valid PascalCase", "ThisIsAValidPascalCaseString123", true},
		{"Complex invalid PascalCase", "ThisIsNot_APascalCaseString", false},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			result := IsPascalCase(testCase.input)

			if result != testCase.expected {
				t.Errorf("IsPascalCase(%q) = %v, want %v", testCase.input, result, testCase.expected)
			}
		})
	}
}

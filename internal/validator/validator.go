package validator

import (
	"strings"
	"unicode"
)

// IsSnakeCase checks if the given string is in snake_case format.
//
// Parameters:
//
//	stringToCheck: The string to check
//
// Returns:
//
//	bool: true if the string is in snake_case, false otherwise
//
// snake_case is defined as a string that:
//  1. Contains only lowercase letters, digits, and underscores
//  2. Does not start or end with an underscore
//  3. Does not contain consecutive underscores
//
// Examples:
//
//	IsSnakeCase("snake_case")           // returns true
//	IsSnakeCase("snake_case_with_123")  // returns true
//	IsSnakeCase("SnakeCase")            // returns false
//	IsSnakeCase("snake-case")           // returns false
//	IsSnakeCase("snake__case")          // returns false
//	IsSnakeCase("_snake_case")          // returns false
//	IsSnakeCase("snake_case_")          // returns false
//
// Note:
//
//	This function considers an empty string as valid snake_case.
func IsSnakeCase(stringToCheck string) bool {
	if stringToCheck == "" {
		return true
	}

	if strings.HasPrefix(stringToCheck, "_") || strings.HasSuffix(stringToCheck, "_") {
		return false
	}

	prevChar := rune('a')

	for _, char := range stringToCheck {
		if char == '_' {
			if prevChar == '_' {
				return false
			}
		} else if !unicode.IsLower(char) && !unicode.IsDigit(char) {
			return false
		}

		prevChar = char
	}

	return true
}

// IsPascalCase checks if the given string is in PascalCase format.
//
// PascalCase is defined as a string that:
//  1. Starts with an uppercase letter
//  2. Contains only letters and digits
//  3. Has no spaces or other punctuation
//
// Parameters:
//
//	stringToCheck: The string to check
//
// Returns:
//
//	bool: true if the string is in PascalCase, false otherwise
//
// Examples:
//
//	IsPascalCase("PascalCase")  // returns true
//	IsPascalCase("pascalCase")  // returns false
//	IsPascalCase("Pascal_Case") // returns false
//	IsPascalCase("Pascal123")   // returns true
//	IsPascalCase("")            // returns false
//
// Note:
//
//	This function considers strings with numbers as valid PascalCase
//	as long as they start with an uppercase letter.
func IsPascalCase(stringToCheck string) bool {
	if len(stringToCheck) == 0 {
		return false
	}

	if !unicode.IsUpper(rune(stringToCheck[0])) {
		return false
	}

	for _, r := range stringToCheck[1:] {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

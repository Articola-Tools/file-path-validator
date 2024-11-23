package processor

import (
	"reflect"
	"testing"
)

func TestIsHiddenFile(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		filename string
		expected bool
	}{
		{
			name:     "Hidden file",
			filename: ".hidden_file",
			expected: true,
		},
		{
			name:     "Hidden file with extension",
			filename: ".hidden.txt",
			expected: true,
		},
		{
			name:     "Hidden file in path",
			filename: "/path/to/.hidden",
			expected: true,
		},
		{
			name:     "Normal file",
			filename: "normal_file",
			expected: false,
		},
		{
			name:     "Normal file with extension",
			filename: "normal.txt",
			expected: false,
		},
		{
			name:     "Normal file in path",
			filename: "/path/to/normal.txt",
			expected: false,
		},
		{
			name:     "Current directory",
			filename: ".",
			expected: false,
		},
		{
			name:     "Parent directory",
			filename: "..",
			expected: false,
		},
		{
			name:     "Hidden file starting with two dots",
			filename: "..hidden",
			expected: true,
		},
		{
			name:     "Empty string",
			filename: "",
			expected: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			result := IsHiddenFile(testCase.filename)

			if result != testCase.expected {
				t.Errorf("IsHiddenFile(%q) = %v, want %v", testCase.filename, result, testCase.expected)
			}
		})
	}
}

func TestGetFileNameWithoutExtension(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple file name with extension",
			input:    "file.txt",
			expected: "file",
		},
		{
			name:     "File name with multiple dots",
			input:    "file.name.with.dots.txt",
			expected: "file.name.with.dots",
		},
		{
			name:     "File name without extension",
			input:    "filename",
			expected: "filename",
		},
		{
			name:     "Hidden file with extension (Unix-style)",
			input:    ".hidden.txt",
			expected: ".hidden",
		},
		{
			name:     "File path with directories (Unix-style)",
			input:    "/path/to/file.txt",
			expected: "file",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			result := GetFileNameWithoutExtension(testCase.input)

			if result != testCase.expected {
				t.Errorf("GetFileNameWithoutExtension(%q) = %q, want %q", testCase.input, result, testCase.expected)
			}
		})
	}
}

func TestGetFoldersFromPath(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Unix-style absolute path",
			input:    "/path/to/some/file.txt",
			expected: []string{"path", "to", "some"},
		},
		{
			name:     "Relative path",
			input:    "relative/path/to/file.txt",
			expected: []string{"relative", "path", "to"},
		},
		{
			name:     "File in current directory",
			input:    "file.txt",
			expected: []string{},
		},
		{
			name:     "Path with multiple slashes",
			input:    "//path///to////file.txt",
			expected: []string{"path", "to"},
		},
		{
			name:     "Path with dots",
			input:    "/path/./to/../to/./file.txt",
			expected: []string{"path", "to"},
		},
		{
			name:     "Empty string",
			input:    "",
			expected: []string{},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			result := GetFoldersFromPath(testCase.input)

			if len(result) == 0 && len(testCase.expected) == 0 {
				return
			}

			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("GetFoldersFromPath(%q) = %v, want %v", testCase.input, result, testCase.expected)
			}
		})
	}
}

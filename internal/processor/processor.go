package processor

import (
	"path/filepath"
	"strings"
)

// IsHiddenFile checks if a file is hidden (Unix-style).
//
// Parameters:
//
//	filename: A string representing the file name or path.
//
// Returns:
//
//	bool: true if the file is hidden, false otherwise.
//
// In Unix-like systems, hidden files are denoted by a leading dot (.) in the file name.
// This function considers a file hidden if:
//  1. The base name of the file (without path) starts with a dot.
//  2. The file is not "." or ".." (current and parent directory).
//
// Examples:
//
//	IsHiddenFile(".hidden_file")     // returns true
//	IsHiddenFile("/path/to/.hidden") // returns true
//	IsHiddenFile("normal_file")      // returns false
//	IsHiddenFile(".")                // returns false
//	IsHiddenFile("..")               // returns false
//
// Note:
//
//	This function only checks the file name and does not access the file system.
//	It does not check for any system-specific or extended attributes that might
//	be used to hide files on certain operating systems.
func IsHiddenFile(filename string) bool {
	baseName := filepath.Base(filename)

	return strings.HasPrefix(baseName, ".") && baseName != "." && baseName != ".."
}

// GetFileNameWithoutExtension extracts the file name from a given file path
// and removes the file extension.
//
// Parameters:
//
//	path: A string representing the file path. This can be a full path or just a file name.
//	      It supports both Unix-style and Windows-style paths.
//
// Returns:
//
//	string: The file name without its extension.
//
// The function handles the following cases:
//   - Files with extensions (e.g., "file.txt" becomes "file")
//   - Files with multiple dots (e.g., "file.name.with.dots.txt" becomes "file.name.with.dots")
//   - Files without extensions (returned as-is)
//   - Paths with directories (directories are stripped, only file name is returned)
//
// Examples:
//
//	GetFileNameWithoutExtension("/path/to/file.txt")                  // returns "file"
//	GetFileNameWithoutExtension("C:\\Users\\file.name.with.dots.txt") // returns "file.name.with.dots"
//	GetFileNameWithoutExtension("just_filename")                      // returns "just_filename"
//	GetFileNameWithoutExtension("/path/to/folder/")                   // returns "folder"
//
// Note:
//
//	This function does not check if the file actually exists in the file system.
//	It only performs string manipulations on the provided path.
func GetFileNameWithoutExtension(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

// GetFoldersFromPath extracts all folder names from a given file path.
//
// Parameters:
//
//	path: A string representing the file path. This can be an absolute or relative path.
//	      It supports both Unix-style and Windows-style paths.
//
// Returns:
//
//	[]string: An array of folder names in the order they appear in the path.
//	          If the path is empty or contains no folders, an empty array is returned.
//
// The function handles the following cases:
//   - Absolute paths (e.g., "/path/to/file.txt" or "C:\path\to\file.txt")
//   - Relative paths (e.g., "path/to/file.txt")
//
// The function normalizes the path, removing any redundant separators or relative path components.
// The last component of the path (must always be a file name) is always excluded from the result.
//
// Examples:
//
//	GetFoldersFromPath("/path/to/some/file.txt")                   // returns ["path", "to", "some"]
//	GetFoldersFromPath("C:\\Users\\Username\\Documents\\file.txt") // returns ["C:", "Users", "Username", "Documents"]
//	GetFoldersFromPath("file.txt")                                 // returns []
//
// Note:
//
//	This function does not check if the folders actually exist in the file system.
//	It only performs string manipulations on the provided path.
func GetFoldersFromPath(path string) []string {
	cleanPath := filepath.Clean(path)
	components := strings.Split(cleanPath, string(filepath.Separator))

	var folders []string

	for _, component := range components[:len(components)-1] {
		if component != "" {
			folders = append(folders, component)
		}
	}

	return folders
}

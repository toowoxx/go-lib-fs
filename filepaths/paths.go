package filepaths

import (
	"path/filepath"
	"strings"
)

// Components returns all components of a given path. If the path begins
// with the path separator, i.e. is an absolute path, two empty strings are returned.
// The returning list contains only one string if `pathToSplit` is empty.
// The slice returned can never be empty.
func Components(pathToSplit string) []string {
	return strings.Split(pathToSplit, string(filepath.Separator))
}

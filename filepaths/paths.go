package filepaths

import (
	"path/filepath"
	"strings"
)

// Components returns all components of a given path. If the path begins
// with the path separator, i.e. is an absolute path, the first string
// is empty. The returning list is empty if `pathToSplit` is empty.
func Components(pathToSplit string) []string {
	return strings.Split(pathToSplit, string(filepath.Separator))
}

// Package fs contains common file system related utility
// functions useful in many situations
package fs

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// DirectoryExists returns true if the path exists and
// there was no error
func DirectoryExists(path string) bool {
	stat, err := os.Stat(path)
	if err == nil {
		return stat.IsDir()
	}
	return false
}

// FileExists returns true if the path exists and
// there was no error
func FileExists(path string) bool {
	stat, err := os.Stat(path)
	if err == nil {
		return !stat.IsDir()
	}
	return false
}

// CreateFile creates an
func CreateFile(path string) error {
	var f *os.File
	var err error
	if f, err = os.Create(path); err != nil {
		return errors.New("could not create file: " + err.Error())
	}
	if err = f.Close(); err != nil {
		return errors.New("could not close created file: " + err.Error())
	}
	return nil
}

// CountLines returns the number of newline characters
// in the file specified by the path.
func CountLines(path string) (int, error) {
	r, err := os.Open(path)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("could not open file %s: %v", path, err))
	}

	buf := make([]byte, 1*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil
		case err != nil:
			return count, errors.New(fmt.Sprintf("error during read of %s: %v", path, err))
		}
	}
}

func AddExtension(filename string, ext string) string {
	return fmt.Sprintf("%s.%s", filename, ext)
}

func ReplaceExtension(filename string, ext string) string {
	return fmt.Sprintf("%s.%s", strings.TrimSuffix(filename, filepath.Ext(filename)), ext)
}

func RemoveExtension(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

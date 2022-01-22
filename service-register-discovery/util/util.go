package util

import (
	"fmt"
	"path/filepath"
)

// Ext return file extension
// e.g. Ext("abc.json") => "json"
func Ext(filename string, checkList ...string) (ext string, err error) {
	ext = filepath.Ext(filename)
	if len(ext) <= 1 {
		return ext, fmt.Errorf("filename: %s requires valid extension", filename)
	}

	ext = ext[1:]
	if len(checkList) > 0 && !StringInSlice(ext, checkList) {
		return ext, fmt.Errorf("Unsupported Config Type %s", ext)
	}

	return ext, nil
}

// StringInSlice return true if string in the slice `list`
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

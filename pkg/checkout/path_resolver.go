package checkout

import (
	"path/filepath"
	"strings"
)

func SecondLastAndLast[S ~[]string](slice S) string {
	return filepath.Join(slice[len(slice)-2], slice[len(slice)-1])
}

func DetermineFolderFromURL(url string) string {
	return filepath.Join("repos", strings.ReplaceAll(SecondLastAndLast(strings.Split(url, "/")), ".git", ""))
}

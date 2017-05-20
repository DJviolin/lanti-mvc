package middlewares

// https://github.com/siongui/userpages/blob/master/content/articles/2017/02/13/go-template-parse-all-files-in-directory%25en.rst

/*
USAGE:
	import (
		mw "github.com/djviolin/lanti-mvc/src/middlewares"
	)

	mw.ParseDirectory("./views", "index")
*/

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

// GetAllFilePathsInDirectory : recursively get all file paths in dir + sub-dirs
func GetAllFilePathsInDirectory(dirpath string) ([]string, error) {
	var paths []string
	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".html") && !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return paths, nil
}

// ParseDirectory : recursively parse all files in dir + sub-dirs
func ParseDirectory(dirpath string, file string) (*template.Template, error) {
	paths, err := GetAllFilePathsInDirectory(dirpath)
	if err != nil {
		return nil, err
	}
	//fmt.Println(paths) // logging
	t := template.New(file)
	return t.ParseFiles(paths...)
}

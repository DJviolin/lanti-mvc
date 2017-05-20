package middlewares

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
)

// GetAllFilePathsInDirectory : Recursively get all file paths in directory, including sub-directories.
// https://github.com/siongui/userpages/blob/master/content/articles/2017/02/13/go-template-parse-all-files-in-directory%25en.rst
func GetAllFilePathsInDirectory(dirpath string) ([]string, error) {
	var paths []string
	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return paths, nil
}

// ParseDirectory : Recursively parse all files in directory, including sub-directories.
// https://github.com/siongui/userpages/blob/master/content/articles/2017/02/13/go-template-parse-all-files-in-directory%25en.rst
//func ParseDirectory(dirpath string) (*template.Template, error) {
func ParseDirectory(dirpath string, filename string) (*template.Template, error) {
	paths, err := GetAllFilePathsInDirectory(dirpath)
	if err != nil {
		return nil, err
	}
	//fmt.Println(paths)          // logging
	t := template.New(filename)
	return t.ParseFiles(paths...)
}

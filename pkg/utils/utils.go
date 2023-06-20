//Package util contains useful functions for handling 
//files and directories while performing the CRUD ops.
//Currently, these functions include extensions of 
//I/O utility functions for efficiently performing the
// required operations as per our use case

//These functions are called in other packages in place
//of the usual ioutil or os functions and although most
//return the values of the concrete types from the original
// functions, the set of parameter received are variable to 
//increase the efficiency for our controllers.

package utils

import(
	"fmt"
	"os"
	"io/fs"
	"io/ioutil"
	"encoding/json"
)

//
func DeleteFile(path string) {
	switch  fi, err := os.Stat(path); {
	case fi==nil, err!=nil:
		fmt.Println("unable to find file or directory name")
	case fi.Mode().IsDir():
		os.RemoveAll(path)
		return
	case fi.Mode().IsRegular(): 
		os.RemoveAll(path)
		return
	}
	return
}

func ParseFile(path string, x interface{}) {
	b, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	if err := json.Unmarshal(b, x); err != nil {
		fmt.Println(err)
	}
	return
}

func ReadDir(dir string)([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	return files, err
}

func WriteFile(dir string, path string, user interface{}) {
	tmpPath := path + ".tmp"
	os.MkdirAll(dir, 0755)
	b, _ := json.MarshalIndent(user, "", "\t")
	b = append(b, byte('\n'))
	if err := os.WriteFile(tmpPath, b, 0644); err != nil{
		fmt.Println(err)
	}
	os.Rename(tmpPath, path)
	return
}
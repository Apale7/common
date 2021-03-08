package utils

import (
	"io/fs"
	"io/ioutil"
	"os"
)

/*
CreateTempFile
usage:
	f, stat, effect, err := CreateTempFile(dir, pattern)
	if err != nil {
		//......
	}
	defer effect()

*/
func CreateTempFile(dir, pattern string) (file *os.File, stat fs.FileInfo, effect func(), err error) {
	file, err = ioutil.TempFile(dir, pattern)
	if err != nil {
		return
	}
	stat, err = os.Stat(file.Name())
	if err != nil {
		return
	}
	effect = func() {
		file.Close()
		os.Remove(file.Name())
	}
	return
}

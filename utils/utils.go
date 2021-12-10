package utils

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetInput(path string) []string {
	filePath, _ := filepath.Abs(path)
	dat, err := ioutil.ReadFile(filePath)
	Check(err)
	return strings.Split(string(dat), "\n")
}

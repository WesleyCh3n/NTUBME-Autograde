package utils

import (
	"io/ioutil"
	"strings"
)

func SedFile(path string, oldStr string, newStr string) (err error) {
	read, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	newContents := strings.Replace(string(read), oldStr, newStr, -1)

	if err = ioutil.WriteFile(path, []byte(newContents), 0); err != nil {
		return err
	}
	return nil
}

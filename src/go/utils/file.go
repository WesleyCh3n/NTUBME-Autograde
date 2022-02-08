package utils

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

func ReplaceStr(path string, oldStr string, newStr string) (err error) {
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

// insert line
// https://siongui.github.io/2017/01/30/go-insert-line-or-string-to-file/
func InsertStringToFile(path, str string, index int) (err error) {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fileContent := ""
	if index == -1 {
		index = len(lines) - 1
	}
	for i, line := range lines {
		fileContent += line
		fileContent += "\n"
		if i == index {
			fileContent += str
		}
	}

	return ioutil.WriteFile(path, []byte(fileContent), 0644)
}

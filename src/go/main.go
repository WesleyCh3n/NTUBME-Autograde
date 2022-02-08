package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

type Template struct {
	Ag struct {
		Hw      string   `yaml:"Homework"`
		Tar     []string `yaml:"AdditionalTar"`
		VarType []string `yaml:"VariableType"`
		Test    []struct {
			Input  []string `yaml:"input"`
			Answer []struct {
				L  string `yaml:"L"`
				Op string `yaml:"op"`
				R  string `yaml:"R"`
			}
		} `yaml:"Test"`
	} `yaml:"Autograde"`
}

var HwNumPro string

func Parser(f []byte) {
	t := Template{}
	if err := yaml.Unmarshal(f, &t); err != nil {
		log.Fatalf("error: %v", err)
	}

	num, _ := strconv.Atoi(regexp.MustCompile("[0-9]+").FindString(t.Ag.Hw))
	problem := strings.ToUpper(regexp.MustCompile("[a-z]").FindString(t.Ag.Hw))

	HwNumPro = fmt.Sprintf("%02d%s", num, problem)
	fmt.Println(color.GreenString("-- Homework number:"), HwNumPro)

	if len(t.Ag.Tar) <= 0 {
		color.Yellow("-- No additional file included")
	} else {
		fmt.Println(color.GreenString("-- Additional files:"), t.Ag.Tar)
	}

	if len(t.Ag.VarType) <= 0 {
		color.Yellow("-- No global answer declared")
	} else {
		fmt.Println(color.GreenString("-- Global answer Type:"), t.Ag.VarType)
	}

	if len(t.Ag.Test) <= 0 {
		color.Yellow("-- No test case")
	} else {
		for i, test := range t.Ag.Test {
			fmt.Println(color.CyanString("-- Testing:"), i)
			fmt.Println(test.Input)
			for _, answer := range test.Answer {
				fmt.Println(strings.ReplaceAll(answer.L, "ans", "answer") +
					answer.Op +
					answer.R)
			}
		}
	}
}

func main() {
	// Read file
	yamlFile, err := ioutil.ReadFile("./sample.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Parse file
	Parser(yamlFile)
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"

	"ga/utils"

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

func GetTemplate() {
	temp_files := []struct {
		Filename string
		Url      string
	}{
		{
			"./autograde-Makefile",
			"https://github.com/WesleyCh3n/NTUBME-Autograde/raw/main/src/Makefile",
		},
		{
			"./gtest.cpp",
			"https://github.com/WesleyCh3n/NTUBME-Autograde/raw/main/src/gtest.cpp",
		},
		{
			"./score.py",
			"https://github.com/WesleyCh3n/NTUBME-Autograde/raw/main/src/score.py",
		},
	}
	for _, temp := range temp_files {
		if err := utils.DownloadFile(temp.Filename, temp.Url); err != nil {
			log.Fatalf("error: %v", err)
		}
	}
}

func Parser(f []byte) {
	t := Template{}
	if err := yaml.Unmarshal(f, &t); err != nil {
		log.Fatalf("error: %v", err)
	}

	num, _ := strconv.Atoi(regexp.MustCompile("[0-9]+").FindString(t.Ag.Hw))
	problem := strings.ToUpper(regexp.MustCompile("[a-z]").FindString(t.Ag.Hw))

	HwNumPro = fmt.Sprintf("%02d%s", num, problem)
	utils.LogInfo("-- Homework number:", HwNumPro)

	if len(t.Ag.Tar) <= 0 {
		utils.LogWarn("-- No additional file included", "")
	} else {
		utils.LogInfo("-- Additional files:", t.Ag.Tar)
	}

	if len(t.Ag.VarType) <= 0 {
		utils.LogWarn("-- No global answer declared", "")
	} else {
		utils.LogInfo("-- Global answer Type:", t.Ag.VarType)
	}

	if len(t.Ag.Test) <= 0 {
		utils.LogWarn("-- No test case", "")
	} else {
		for i, test := range t.Ag.Test {
			utils.LogCyan("-- Testing:", i)

			fmt.Println(strings.Join(test.Input, " "))
			for _, answer := range test.Answer {
				fmt.Println(strings.ReplaceAll(answer.L, "ans", "answer") +
					answer.Op +
					answer.R)
			}
		}
	}

	func() {
		utils.SedFile("./autograde-Makefile", "{{HW_NUM}}", HwNumPro)
		utils.SedFile("./autograde-Makefile", "{{N_TEST}}", strconv.Itoa(len(t.Ag.Test)))
		// utils.SedFile("./autograde-Makefile", "{{INPUTS}}", strconv.Itoa(len(t.Ag.Test)))
	}()
}

func main() {
	// Download template
	GetTemplate()

	// Read file
	yamlFile, err := ioutil.ReadFile("./sample.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Parse file
	Parser(yamlFile)
}

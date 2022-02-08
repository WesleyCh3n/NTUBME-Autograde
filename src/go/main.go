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

func GetOp(op string) string {
	switch op {
	case "=":
		return "ASSERT_EQ"
	case ">":
		return "ASSERT_GT"
	case "<":
		return "ASSERT_LT"
	case "!=":
		return "ASSERT_NE"
	case ">=":
		return "ASSERT_GE"
	case "<=":
		return "ASSERT_LE"
	case "&=":
		return "ASSERT_STREQ"
	case "&?":
		return "ASSERT_STRNE"
	default:
		return ""
	}
}

func createPackage(f []byte) {
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

	testInputs := []string{} // test input for score.py as user input
	testCases := ""          // test cases for gtest.cpp

	if len(t.Ag.Test) <= 0 {
		utils.LogWarn("-- No test case", "")
	} else {
		for i, test := range t.Ag.Test {
			utils.LogCyan("-- Testing:", i)

			testInputs = append(testInputs, strings.Join(test.Input, " "))
			utils.LogCyan("     Input:", test.Input)

			testContent := ""
			for _, answer := range test.Answer {
				testContent += fmt.Sprintf("\t%s(%s, %s);\n",
					GetOp(answer.Op),
					strings.ReplaceAll(answer.L, "ans", "answer"),
					answer.R)
			}

			utils.LogCyan("     GoogleTest:\n", testContent[:len(testContent)-1])
			testCases += fmt.Sprintf("\nTEST(GoogleTest, test%d){\n%s}",
				i+1,
				testContent)
		}
	}

	utils.LogInfo("-- Create autograde-Makefile", "")
	utils.ReplaceStr("./autograde-Makefile", "{{HW_NUM}}", HwNumPro)
	utils.ReplaceStr("./autograde-Makefile", "{{N_TEST}}", strconv.Itoa(len(t.Ag.Test)))
	utils.ReplaceStr("./autograde-Makefile", "{{INPUTS}}", fmt.Sprintf("\"%s\"", strings.Join(testInputs, ";")))

	utils.LogInfo("-- Create gtest.cpp", "")
	varTypes := ""
	for i, v := range t.Ag.VarType {
		varTypes += fmt.Sprintf("extern %s answer%d;\n", v, i+1)
	}
	utils.InsertStringToFile("./gtest.cpp", varTypes, 2)
	utils.InsertStringToFile("./gtest.cpp", testCases, -1)
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
	createPackage(yamlFile)
}

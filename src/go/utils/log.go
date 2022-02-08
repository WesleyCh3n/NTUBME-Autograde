package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func LogInfo(prefix string, str interface{}) {
	fmt.Println(color.GreenString(prefix), str)
}

func LogWarn(prefix string, str interface{}) {
	fmt.Println(color.YellowString(prefix), str)
}

func LogCyan(prefix string, str interface{}) {
	fmt.Println(color.CyanString(prefix), str)
}

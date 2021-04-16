package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/liangyaopei/checker"
)

type customStruct struct {
	Inner struct {
		Str string
	}
}

func main() {
	checkFn := func(exprValue interface{}) (bool, string) {
		res, ok := exprValue.(string)
		if !ok {
			res = reflect.ValueOf(exprValue).String()
		}
		if !strings.Contains(res, "checker") {
			return false, "String should contain checker."
		}
		return true, ""
	}

	rule := checker.Custom("Inner.Str", checkFn)

	validator := checker.NewChecker()
	validator.Add(rule, "wrong Custom")

	custom := customStruct{}
	custom.Inner.Str = "checker custom function"
	isValid, _, _ := validator.Check(&custom)
	fmt.Printf("isValid:%v", isValid)
	fmt.Println()

	custom.Inner.Str = "abc"
	isValid, prompt, errMsg := validator.Check(&custom)

	custom.Inner.Str = "checker custom function"

	custom.Inner.Str = "abc"
	isValid, prompt, errMsg = validator.Check(&custom)
	fmt.Printf("prompt:%v\n", prompt)
	fmt.Printf("errMsg:%v\n", errMsg)
}

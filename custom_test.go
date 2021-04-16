package checker

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustom(t *testing.T) {
	type customStruct struct {
		Inner struct {
			Str string
		}
	}

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

	rule := Custom("Inner.Str", checkFn)

	validator := NewChecker()
	validator.Add(rule, "wrong Custom")

	custom := customStruct{}
	custom.Inner.Str = "checker custom function"
	isValid, _, _ := validator.Check(&custom)
	assert.Equal(t, true, isValid)

	custom.Inner.Str = "abc"
	isValid, prompt, errMsg := validator.Check(&custom)
	assert.Equal(t, false, isValid)
	assert.Equal(t, "wrong Custom", prompt)
	assert.Equal(t, "[Custom]:Inner.Str String should contain checker.", errMsg)
}

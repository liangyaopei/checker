package checker

import (
	"fmt"
	"reflect"
	"testing"
)

type customizedRule struct {
	fieldExpr string

	name string
}

func (r customizedRule) Check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.fieldExpr)
	}
	if exprValue != nil {
		return false,
			fmt.Sprintf("[%s]:'%s' should not be nil", r.name, r.fieldExpr)
	}
	return true, ""
}

type customizedStruct struct {
	NilPtr *int
}

func TestCustomizedRule(t *testing.T) {
	ch := NewChecker()

	fieldExpr := "NilPtr"
	customRule := customizedRule{
		fieldExpr: fieldExpr,
		name:      "customizedRule",
	}

	ch.Add(customRule, "invalid ptr")

	param := customizedStruct{
		NilPtr: nil,
	}

	isValid, prompt, errMsg := ch.Check(param)
	if !isValid {
		t.Errorf("errMsg:%s,prompt:%s", errMsg, prompt)
		return
	}
	t.Logf("valid nil pointer")
}

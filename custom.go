package checker

import (
	"fmt"
	"reflect"
)

type customRule struct {
	baseRule
	checkFn func(exprValue interface{}) (bool, string)
}

func (r *customRule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r *customRule) Check(param interface{}) (bool, string) {
	exprValue, kind := fetchField(param, r)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:%s cannot be found", r.name, r.getCompleteFieldExpr())
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:%s is nil", r.name, r.getCompleteFieldExpr())
	}
	isValid, errMsg := r.checkFn(exprValue)
	if !isValid {
		return false,
			fmt.Sprintf("[%s]:%s %s",
				r.name, r.getCompleteFieldExpr(), errMsg)
	}
	return true, ""
}

func Custom(fieldExpr string, checkFn func(exprValue interface{}) (bool, string)) *customRule {
	return &customRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "Custom",
		},
		checkFn,
	}
}

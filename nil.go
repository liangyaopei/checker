package checker

import (
	"fmt"
	"reflect"
)

type nilRule struct {
	baseRule
}

func (r *nilRule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r nilRule) Check(param interface{}) (bool, string) {
	exprValue, kind := fetchField(param, &r)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:%s cannot be found", r.name, r.getCompleteFieldExpr())
	}
	if exprValue != nil {
		return false,
			fmt.Sprintf("[%s]:%s is not nil", r.name, r.getCompleteFieldExpr())
	}
	return true, ""
}

// Nil checks if the field is nil
func Nil(fieldExpr string) *nilRule {
	return &nilRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "Nil",
		},
	}
}

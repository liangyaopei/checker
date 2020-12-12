package checker

import (
	"fmt"
	"reflect"
)

type nonZeroRule struct {
	fieldExpr string
	name      string
}

func (r nonZeroRule) Check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.fieldExpr)
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is nil", r.name, r.fieldExpr)
	}
	refVal := reflect.ValueOf(exprValue)
	if refVal.IsZero() {
		return false,
			fmt.Sprintf("[%s]:'%s' should not be zero value", r.name, r.fieldExpr)
	}
	return true, ""
}

func NewNonZeroRule(fieldExpr string) Rule {
	return nonZeroRule{
		fieldExpr: fieldExpr,
		name:      "nonZeroRule",
	}
}

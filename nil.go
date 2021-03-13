package checker

import (
	"fmt"
	"reflect"
)

type nilRule struct {
	fieldExpr string
	name      string
}

func (r nilRule) Check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.fieldExpr)
	}
	if exprValue != nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is not nil", r.name, r.fieldExpr)
	}
	return true, ""
}

// Nil checks if the field is nil
func Nil(fieldExpr string) Rule {
	return nilRule{
		fieldExpr: fieldExpr,
		name:      "nilRule",
	}
}

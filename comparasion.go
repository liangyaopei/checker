package checker

import (
	"fmt"
	"reflect"
)

type eqRuleString struct {
	fieldExpr  string
	equivalent string
}

func (r eqRuleString) check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if exprValue == nil {
		return false,
			fmt.Sprintf("[eqRuleString]:'%s' is nil", r.fieldExpr)
	}
	if kind != reflect.String {
		return false,
			fmt.Sprintf("[eqRuleString]:'%s' should be kind string,actual is %v", r.fieldExpr, kind)
	}
	exprValueStr := exprValue.(string)
	if exprValueStr != r.equivalent {
		return false,
			fmt.Sprintf("[eqRuleString]:'%s' should be %s,actual is %s", r.fieldExpr, r.equivalent, exprValueStr)
	}
	return true, ""
}

func NewEqRuleString(filedExpr string, equivalent string) Rule {
	return eqRuleString{
		fieldExpr:  filedExpr,
		equivalent: equivalent,
	}
}

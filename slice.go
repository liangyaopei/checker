package checker

import (
	"fmt"
	"reflect"
)

type sliceRule struct {
	fieldExpr string
	innerRule Rule
}

func (r sliceRule) check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if exprValue == nil {
		return false,
			fmt.Sprintf("[sliceRule]:'%s' is nil", exprValue)
	}

	if kind != reflect.Slice && kind != reflect.Array {
		return false,
			fmt.Sprintf("[sliceRule]:'%s' should be slice/array,actual is %v", r.fieldExpr, kind)
	}

	sliceValue := reflect.ValueOf(exprValue)
	for i := 0; i < sliceValue.Len(); i++ {
		idxValue := sliceValue.Index(i)
		isValid, msg := r.innerRule.check(idxValue.Interface())
		if !isValid {
			return isValid, msg
		}
	}
	return true, ""
}

func NewSliceRule(fieldExpr string, innerRule Rule) Rule {
	return sliceRule{
		fieldExpr: fieldExpr,
		innerRule: innerRule,
	}
}

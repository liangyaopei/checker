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

type lengthRule struct {
	fieldExpr string
	lenLimit  int
}

func (r lengthRule) check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if exprValue == nil {
		return false,
			fmt.Sprintf("[lengthRule]:'%s' is nil", exprValue)
	}

	if kind != reflect.Slice && kind != reflect.Array &&
		kind != reflect.String && kind != reflect.Map {
		return false,
			fmt.Sprintf("[lengthRule]:'%s' should be slice/array/string/map,actual is %v", r.fieldExpr, kind)
	}

	lenValue := reflect.ValueOf(exprValue)
	if lenValue.Len() > r.lenLimit {
		return false,
			fmt.Sprintf("[lengthRule]:'%s' length should be less than or equal to %d,actual is %d",
				r.fieldExpr, r.lenLimit, lenValue.Len())

	}
	return true, ""
}

func NewLengthRule(fieldExpr string, lenLimit int) Rule {
	return lengthRule{
		fieldExpr: fieldExpr,
		lenLimit:  lenLimit,
	}
}

package checker

import (
	"fmt"
	"reflect"
)

type sliceRule struct {
	fieldExpr string
	innerRule Rule

	name string
}

func (r sliceRule) check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.fieldExpr)
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is nil", r.name, r.fieldExpr)
	}

	if kind != reflect.Slice && kind != reflect.Array {
		return false,
			fmt.Sprintf("[%s]:'%s' should be slice/array,actual is %v",
				r.name, r.fieldExpr, kind)
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
		name:      "sliceRule",
	}
}

type lengthRule struct {
	fieldExpr string
	le        int
	ge        int

	name string
}

func (r lengthRule) check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.fieldExpr)
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is nil", r.name, r.fieldExpr)
	}

	if kind != reflect.Slice && kind != reflect.Array &&
		kind != reflect.String && kind != reflect.Map {
		return false,
			fmt.Sprintf("[%s]:'%s' should be slice/array/string/map,actual is %v",
				r.name, r.fieldExpr, kind)
	}

	lenValue := reflect.ValueOf(exprValue)
	length := lenValue.Len()
	if length < r.ge || length > r.le {
		return false,
			fmt.Sprintf("[%s]:'%s' length should be between %d and %d,actual is %d",
				r.name, r.fieldExpr, r.le, r.ge, length)

	}
	return true, ""
}

func NewLengthRule(fieldExpr string, ge int, le int) Rule {
	return lengthRule{
		fieldExpr: fieldExpr,
		ge:        ge,
		le:        le,
		name:      "lengthRule",
	}
}

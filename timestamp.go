package checker

import (
	"fmt"
	"reflect"
	"time"
)

type timestampStrRule struct {
	fieldExpr string
	layout    string

	name string
}

func (r timestampStrRule) check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.fieldExpr)
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is nil", r.name, r.fieldExpr)
	}
	if kind != reflect.String {
		return false,
			fmt.Sprintf("[%s]:'%s' should be kind string,actual is %v",
				r.name, r.fieldExpr, kind)
	}
	exprValueStr := exprValue.(string)
	_, err := time.Parse(r.layout, exprValueStr)
	if err != nil {
		return false,
			fmt.Sprintf("[%s]:'%s' should be format %s,actual is %s",
				r.name, r.fieldExpr, r.layout, exprValueStr)
	}
	return true, ""
}

func NewTimestampStrRule(filedExpr string, layout string) Rule {
	return timestampStrRule{
		fieldExpr: filedExpr,
		layout:    layout,
		name:      "timestampStrRule",
	}
}

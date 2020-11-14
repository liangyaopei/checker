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

type eqRuleInt struct {
	fieldExpr  string
	equivalent int
}

func (r eqRuleInt) check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if exprValue == nil {
		return false,
			fmt.Sprintf("[eqRuleInt]:'%s' is nil", r.fieldExpr)
	}
	if kind != reflect.Int8 && kind != reflect.Int16 && kind != reflect.Int32 &&
		kind != reflect.Int64 && kind != reflect.Int {
		return false,
			fmt.Sprintf("[eqRuleInt]:'%s' should be kind int8/int16/int32/int64/int,actual is %v", r.fieldExpr, kind)
	}
	exprValueInt := exprValue.(int)
	if exprValueInt != r.equivalent {
		return false,
			fmt.Sprintf("[eqRuleInt]:'%s' should be %d,actual is %d", r.fieldExpr, r.equivalent, exprValueInt)
	}
	return true, ""
}

func NewEqRuleInt(filedExpr string, equivalent int) Rule {
	return eqRuleInt{
		fieldExpr:  filedExpr,
		equivalent: equivalent,
	}
}

type eqRuleUint struct {
	fieldExpr  string
	equivalent uint
}

func (r eqRuleUint) check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if exprValue == nil {
		return false,
			fmt.Sprintf("[eqRuleUint]:'%s' is nil", r.fieldExpr)
	}
	if kind != reflect.Uint8 && kind != reflect.Uint16 && kind != reflect.Uint32 &&
		kind != reflect.Uint64 && kind != reflect.Uint {
		return false,
			fmt.Sprintf("[eqRuleUint]:'%s' should be kind uint8/uint16/uint32/uint64/uint,actual is %v", r.fieldExpr, kind)
	}
	exprValueUint := exprValue.(uint)
	if exprValueUint != r.equivalent {
		return false,
			fmt.Sprintf("[eqRuleUint]:'%s' should be %d,actual is %d", r.fieldExpr, r.equivalent, exprValueUint)
	}
	return true, ""
}

func NewEqRuleUint(filedExpr string, equivalent uint) Rule {
	return eqRuleUint{
		fieldExpr:  filedExpr,
		equivalent: equivalent,
	}
}

type eqRuleFloat struct {
	fieldExpr  string
	equivalent float64
}

func (r eqRuleFloat) check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if exprValue == nil {
		return false,
			fmt.Sprintf("[eqRuleFloat]:'%s' is nil", r.fieldExpr)
	}
	if kind != reflect.Float32 && kind != reflect.Float64 {
		return false,
			fmt.Sprintf("[eqRuleFloat]:'%s' should be kind float32/float64,actual is %v", r.fieldExpr, kind)
	}
	exprValueFloat := exprValue.(float64)
	if exprValueFloat != r.equivalent {
		return false,
			fmt.Sprintf("[eqRuleFloat]:'%s' should be %d,actual is %d", r.fieldExpr, r.equivalent, exprValueFloat)
	}
	return true, ""
}

func NewEqRuleFloat(filedExpr string, equivalent float64) Rule {
	return eqRuleFloat{
		fieldExpr:  filedExpr,
		equivalent: equivalent,
	}
}

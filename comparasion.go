package checker

import (
	"fmt"
	"reflect"
)

type eqRuleString struct {
	fieldExpr  string
	equivalent string
	name       string
}

func (r eqRuleString) check(param interface{}) (bool, string) {
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
		name:       "eqRuleString",
	}
}

type eqRuleInt struct {
	fieldExpr  string
	equivalent int
	name       string
}

func (r eqRuleInt) check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.fieldExpr)
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is nil", r.name, r.fieldExpr)
	}
	if kind != reflect.Int8 && kind != reflect.Int16 && kind != reflect.Int32 &&
		kind != reflect.Int64 && kind != reflect.Int {
		return false,
			fmt.Sprintf("[%s]:'%s' should be kind int8/int16/int32/int64/int,actual is %v",
				r.name, r.fieldExpr, kind)
	}
	exprValueInt := exprValue.(int)
	if exprValueInt != r.equivalent {
		return false,
			fmt.Sprintf("[%s]:'%s' should be %d,actual is %d",
				r.name, r.fieldExpr, r.equivalent, exprValueInt)
	}
	return true, ""
}

func NewEqRuleInt(filedExpr string, equivalent int) Rule {
	return eqRuleInt{
		fieldExpr:  filedExpr,
		equivalent: equivalent,
		name:       "eqRuleInt",
	}
}

type eqRuleUint struct {
	fieldExpr  string
	equivalent uint
	name       string
}

func (r eqRuleUint) check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.fieldExpr)
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is nil", r.name, r.fieldExpr)
	}
	if kind != reflect.Uint8 && kind != reflect.Uint16 && kind != reflect.Uint32 &&
		kind != reflect.Uint64 && kind != reflect.Uint {
		return false,
			fmt.Sprintf("[%s]:'%s' should be kind uint8/uint16/uint32/uint64/uint,actual is %v",
				r.name, r.fieldExpr, kind)
	}
	exprValueUint := exprValue.(uint)
	if exprValueUint != r.equivalent {
		return false,
			fmt.Sprintf("[%s]:'%s' should be %d,actual is %d",
				r.name, r.fieldExpr, r.equivalent, exprValueUint)
	}
	return true, ""
}

func NewEqRuleUint(filedExpr string, equivalent uint) Rule {
	return eqRuleUint{
		fieldExpr:  filedExpr,
		equivalent: equivalent,
		name:       "eqRuleUint",
	}
}

type eqRuleFloat struct {
	fieldExpr  string
	equivalent float64
	name       string
}

func (r eqRuleFloat) check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.fieldExpr)
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is nil", r.name, r.fieldExpr)
	}

	if kind != reflect.Float32 && kind != reflect.Float64 {
		return false,
			fmt.Sprintf("[%s]:'%s' should be kind float32/float64,actual is %v",
				r.name, r.fieldExpr, kind)
	}
	exprValueFloat := exprValue.(float64)
	if exprValueFloat != r.equivalent {
		return false,
			fmt.Sprintf("[%s]:'%s' should be %f,actual is %f",
				r.name, r.fieldExpr, r.equivalent, exprValueFloat)
	}
	return true, ""
}

func NewEqRuleFloat(filedExpr string, equivalent float64) Rule {
	return eqRuleFloat{
		fieldExpr:  filedExpr,
		equivalent: equivalent,
		name:       "eqRuleFloat",
	}
}

type notEqRuleString struct {
	fieldExpr    string
	inequivalent string
	name         string
}

func (r notEqRuleString) check(param interface{}) (bool, string) {
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
	if exprValueStr != r.inequivalent {
		return false,
			fmt.Sprintf("[%s]:'%s' should not be %s,actual is %s",
				r.name, r.fieldExpr, r.inequivalent, exprValueStr)
	}
	return true, ""
}

func NewNotEqRuleString(filedExpr string, inequivalent string) Rule {
	return notEqRuleString{
		fieldExpr:    filedExpr,
		inequivalent: inequivalent,
		name:         "notEqRuleString",
	}
}

type notEqRuleInt struct {
	fieldExpr    string
	inequivalent int
	name         string
}

func (r notEqRuleInt) check(param interface{}) (bool, string) {
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
	exprValueInt := exprValue.(int)
	if exprValueInt != r.inequivalent {
		return false,
			fmt.Sprintf("[%s]:'%s' should not be %d,actual is %d",
				r.name, r.fieldExpr, r.inequivalent, exprValueInt)
	}
	return true, ""
}

func NewNotEqRuleInt(filedExpr string, inequivalent int) Rule {
	return notEqRuleInt{
		fieldExpr:    filedExpr,
		inequivalent: inequivalent,
		name:         "notEqRuleInt",
	}
}

type notEqRuleUint struct {
	fieldExpr    string
	inequivalent uint
	name         string
}

func (r notEqRuleUint) check(param interface{}) (bool, string) {
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
	exprValueUint := exprValue.(uint)
	if exprValueUint != r.inequivalent {
		return false,
			fmt.Sprintf("[%s]:'%s' should not be %d,actual is %d",
				r.name, r.fieldExpr, r.inequivalent, exprValueUint)
	}
	return true, ""
}

func NewNotEqRuleUint(filedExpr string, inequivalent uint) Rule {
	return notEqRuleUint{
		fieldExpr:    filedExpr,
		inequivalent: inequivalent,
		name:         "notEqRuleUint",
	}
}

type notEqRuleFloat struct {
	fieldExpr    string
	inequivalent float64
	name         string
}

func (r notEqRuleFloat) check(param interface{}) (bool, string) {
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
	exprValueFloat := exprValue.(float64)
	if exprValueFloat != r.inequivalent {
		return false,
			fmt.Sprintf("[%s]:'%s' should not be %f,actual is %f",
				r.name, r.fieldExpr, r.inequivalent, exprValueFloat)
	}
	return true, ""
}

func NewNotEqRuleFloat(filedExpr string, inequivalent float64) Rule {
	return notEqRuleFloat{
		fieldExpr:    filedExpr,
		inequivalent: inequivalent,
		name:         "notEqRuleFloat",
	}
}

type rangeRuleInt struct {
	fieldExpr string
	ge        int
	le        int
	name      string
}

func (r rangeRuleInt) check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.fieldExpr)
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is nil", r.name, r.fieldExpr)
	}

	if kind != reflect.Int8 && kind != reflect.Int16 && kind != reflect.Int32 &&
		kind != reflect.Int64 && kind != reflect.Int {
		return false,
			fmt.Sprintf("[%s]:'%s' should be kind int8/int16/int32/int64/int,actual is %v",
				r.name, r.fieldExpr, kind)
	}
	exprValueInt := exprValue.(int)
	if exprValueInt > r.le || exprValueInt < r.ge {
		return false,
			fmt.Sprintf("[%s]:'%s' should be between %d and %d,actual is %d",
				r.name, r.fieldExpr, r.ge, r.le, exprValueInt)
	}
	return true, ""
}

func NewRangeRuleInt(filedExpr string, le int, ge int) Rule {
	return rangeRuleInt{
		fieldExpr: filedExpr,
		ge:        le,
		le:        ge,
		name:      "rangeRuleInt",
	}
}

type rangeRuleUint struct {
	fieldExpr string
	ge        uint
	le        uint
	name      string
}

func (r rangeRuleUint) check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.fieldExpr)
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is nil", r.name, r.fieldExpr)
	}

	if kind != reflect.Uint8 && kind != reflect.Uint16 && kind != reflect.Uint32 &&
		kind != reflect.Uint64 && kind != reflect.Uint {
		return false,
			fmt.Sprintf("[%s]:'%s' should be kind uint8/uint16/uint32/uint64/uint,actual is %v",
				r.name, r.fieldExpr, kind)
	}
	exprValueUint := exprValue.(uint)
	if exprValueUint > r.le || exprValueUint < r.ge {
		return false,
			fmt.Sprintf("[%s]:'%s' should be betweeen %d and %d ,actual is %d",
				r.name, r.fieldExpr, r.ge, r.le, exprValueUint)
	}
	return true, ""
}

func NewRangeRuleUint(filedExpr string, le uint, ge uint) Rule {
	return rangeRuleUint{
		fieldExpr: filedExpr,
		ge:        le,
		le:        ge,
		name:      "rangeRuleUint",
	}
}

type rangeRuleFloat struct {
	fieldExpr string
	le        float64
	ge        float64
	name      string
}

func (r rangeRuleFloat) check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.fieldExpr)
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is nil", r.name, r.fieldExpr)
	}
	if kind != reflect.Float32 && kind != reflect.Float64 {
		return false,
			fmt.Sprintf("[%s]:'%s' should be kind float32/float64,actual is %v",
				r.name, r.fieldExpr, kind)
	}
	exprValueFloat := exprValue.(float64)
	if exprValueFloat > r.ge || exprValueFloat < r.le {
		return false,
			fmt.Sprintf("[%s]:'%s' should be between %f and %f,actual is %f",
				r.name, r.fieldExpr, r.le, r.ge, exprValueFloat)
	}
	return true, ""
}

func NewRangeRuleFloat(filedExpr string, le float64, ge float64) Rule {
	return rangeRuleFloat{
		fieldExpr: filedExpr,
		le:        le,
		ge:        ge,
		name:      "rangeRuleFloat",
	}
}

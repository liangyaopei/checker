package checker

import (
	"fmt"
	"reflect"
	"time"
)

type eqRuleString struct {
	fieldExpr  string
	equivalent string
	name       string
}

func (r eqRuleString) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := getStrField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if exprValueStr != r.equivalent {
		return false,
			fmt.Sprintf("[eqRuleString]:'%s' should be %s,actual is %s", r.fieldExpr, r.equivalent, exprValueStr)
	}
	return true, ""
}

// NewEqRuleString is the validation function for validating if the field's value is equal to given string.
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

func (r eqRuleInt) Check(param interface{}) (bool, string) {
	exprValueInt, isValid, errMsg := getIntField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if exprValueInt != r.equivalent {
		return false,
			fmt.Sprintf("[%s]:'%s' should be %d,actual is %d",
				r.name, r.fieldExpr, r.equivalent, exprValueInt)
	}
	return true, ""
}

// NewEqRuleInt is the validation function for validating if the field's value is equal to given int.
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

func (r eqRuleUint) Check(param interface{}) (bool, string) {
	exprValueUint, isValid, errMsg := getUintField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if exprValueUint != r.equivalent {
		return false,
			fmt.Sprintf("[%s]:'%s' should be %d,actual is %d",
				r.name, r.fieldExpr, r.equivalent, exprValueUint)
	}
	return true, ""
}

// NewEqRuleUint is the validation function for validating if the field's value is equal to given uint.
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

func (r eqRuleFloat) Check(param interface{}) (bool, string) {
	exprValueFloat, isValid, errMsg := getFloatField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if exprValueFloat != r.equivalent {
		return false,
			fmt.Sprintf("[%s]:'%s' should be %f,actual is %f",
				r.name, r.fieldExpr, r.equivalent, exprValueFloat)
	}
	return true, ""
}

// NewEqRuleFloat is the validation function for validating if the field's value is equal to given float.
func NewEqRuleFloat(filedExpr string, equivalent float64) Rule {
	return eqRuleFloat{
		fieldExpr:  filedExpr,
		equivalent: equivalent,
		name:       "eqRuleFloat",
	}
}

type eqRuleTimestamp struct {
	fieldExpr  string
	equivalent time.Time
	name       string
}

func (r eqRuleTimestamp) Check(param interface{}) (bool, string) {
	tsVal, isValid, errMsg := getTimeField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if !tsVal.Equal(r.equivalent) {
		return false,
			fmt.Sprintf("[%s]:'%s' should be %v,actual is %v",
				r.name, r.fieldExpr, r.equivalent, tsVal)
	}
	return true, ""
}

// NewEqRuleFloat is the validation function for validating if the field's value is equal to given timestamp.
func NewEqRuleTimestamp(filedExpr string, equivalent time.Time) Rule {
	return eqRuleTimestamp{
		fieldExpr:  filedExpr,
		equivalent: equivalent,
		name:       "eqRuleTimestamp",
	}
}

type eqRuleTimestampStr struct {
	fieldExpr  string
	layout     string
	equivalent time.Time
	name       string
}

func (r eqRuleTimestampStr) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := getStrField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	exprValTime, err := time.Parse(r.layout, exprValueStr)
	if err != nil {
		return false,
			fmt.Sprintf("[%s]:'%s' should be format %s,actual is %s",
				r.name, r.fieldExpr, r.layout, exprValueStr)
	}

	if !exprValTime.Equal(r.equivalent) {
		return false,
			fmt.Sprintf("[%s]:'%s' should be %v,actual is %v",
				r.name, r.fieldExpr, r.equivalent.Format(r.layout), exprValueStr)
	}
	return true, ""
}

// NewEqRuleTimestampStr is the validation function for validating if the field's value string and is equal to given timestamp.
func NewEqRuleTimestampStr(filedExpr string, layout string, equivalent time.Time) Rule {
	return eqRuleTimestampStr{
		fieldExpr:  filedExpr,
		layout:     layout,
		equivalent: equivalent,
		name:       "eqRuleTimestampStr",
	}
}

type neRuleString struct {
	fieldExpr    string
	inequivalent string
	name         string
}

func (r neRuleString) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := getStrField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if exprValueStr == r.inequivalent {
		return false,
			fmt.Sprintf("[%s]:'%s' should not be %s,actual is %s",
				r.name, r.fieldExpr, r.inequivalent, exprValueStr)
	}
	return true, ""
}

// NewNeRuleString is the validation function for validating if the field's value is not equal to given string.
func NewNeRuleString(filedExpr string, inequivalent string) Rule {
	return neRuleString{
		fieldExpr:    filedExpr,
		inequivalent: inequivalent,
		name:         "neRuleString",
	}
}

type neRuleInt struct {
	fieldExpr    string
	inequivalent int
	name         string
}

func (r neRuleInt) Check(param interface{}) (bool, string) {
	exprValueInt, isValid, errMsg := getIntField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if exprValueInt == r.inequivalent {
		return false,
			fmt.Sprintf("[%s]:'%s' should not be %d,actual is %d",
				r.name, r.fieldExpr, r.inequivalent, exprValueInt)
	}
	return true, ""
}

// NewNeRuleInt is the validation function for validating if the field's value is not equal to given int.
func NewNeRuleInt(filedExpr string, inequivalent int) Rule {
	return neRuleInt{
		fieldExpr:    filedExpr,
		inequivalent: inequivalent,
		name:         "neRuleInt",
	}
}

type neRuleUint struct {
	fieldExpr    string
	inequivalent uint
	name         string
}

func (r neRuleUint) Check(param interface{}) (bool, string) {
	exprValueUint, isValid, errMsg := getUintField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if exprValueUint == r.inequivalent {
		return false,
			fmt.Sprintf("[%s]:'%s' should not be %d,actual is %d",
				r.name, r.fieldExpr, r.inequivalent, exprValueUint)
	}
	return true, ""
}

// NewNeRuleUint is the validation function for validating if the field's value is not equal to given uint.
func NewNeRuleUint(filedExpr string, inequivalent uint) Rule {
	return neRuleUint{
		fieldExpr:    filedExpr,
		inequivalent: inequivalent,
		name:         "neRuleUint",
	}
}

type neRuleFloat struct {
	fieldExpr    string
	inequivalent float64
	name         string
}

func (r neRuleFloat) Check(param interface{}) (bool, string) {
	exprValueFloat, isValid, errMsg := getFloatField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if exprValueFloat == r.inequivalent {
		return false,
			fmt.Sprintf("[%s]:'%s' should not be %f,actual is %f",
				r.name, r.fieldExpr, r.inequivalent, exprValueFloat)
	}
	return true, ""
}

// NewNeRuleFloat is the validation function for validating if the field's value is not equal to given float.
func NewNeRuleFloat(filedExpr string, inequivalent float64) Rule {
	return neRuleFloat{
		fieldExpr:    filedExpr,
		inequivalent: inequivalent,
		name:         "neRuleFloat",
	}
}

type neTimestamp struct {
	fieldExpr    string
	inequivalent time.Time
	name         string
}

func (r neTimestamp) Check(param interface{}) (bool, string) {
	tsVal, isValid, errMsg := getTimeField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if tsVal.Equal(r.inequivalent) {
		return false,
			fmt.Sprintf("[%s]:'%s' should not be %v,actual is %v",
				r.name, r.fieldExpr, r.inequivalent, tsVal)
	}
	return true, ""
}

// NewNeRuleTimestamp is the validation function for validating if the field's value is not equal to given timestamp.
func NewNeRuleTimestamp(filedExpr string, inequivalent time.Time) Rule {
	return neTimestamp{
		fieldExpr:    filedExpr,
		inequivalent: inequivalent,
		name:         "neTimestamp",
	}
}

type neRuleTimestampStr struct {
	fieldExpr    string
	layout       string
	inequivalent time.Time
	name         string
}

func (r neRuleTimestampStr) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := getStrField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	exprValTime, err := time.Parse(r.layout, exprValueStr)
	if err != nil {
		return false,
			fmt.Sprintf("[%s]:'%s' should be format %s,actual is %s",
				r.name, r.fieldExpr, r.layout, exprValueStr)
	}

	if exprValTime.Equal(r.inequivalent) {
		return false,
			fmt.Sprintf("[%s]:'%s' should not be %v,actual is %v",
				r.name, r.fieldExpr, r.inequivalent.Format(r.layout), exprValueStr)
	}
	return true, ""
}

// NewNeRuleTimestampStr is the validation function for validating if the field's value string and is not equal to given timestamp.
func NewNeRuleTimestampStr(filedExpr string, layout string, inequivalent time.Time) Rule {
	return neRuleTimestampStr{
		fieldExpr:    filedExpr,
		layout:       layout,
		inequivalent: inequivalent,
		name:         "neRuleTimestampStr",
	}
}

type rangeRuleInt struct {
	fieldExpr string
	ge        int
	le        int
	name      string
}

func (r rangeRuleInt) Check(param interface{}) (bool, string) {
	exprValueInt, isValid, errMsg := getIntField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if exprValueInt > r.le || exprValueInt < r.ge {
		return false,
			fmt.Sprintf("[%s]:'%s' should be between %d and %d,actual is %d",
				r.name, r.fieldExpr, r.ge, r.le, exprValueInt)
	}
	return true, ""
}

// NewRangeRuleInt is the validation function for validating if the field's value is in given range.
func NewRangeRuleInt(filedExpr string, ge int, le int) Rule {
	return rangeRuleInt{
		fieldExpr: filedExpr,
		ge:        ge,
		le:        le,
		name:      "rangeRuleInt",
	}
}

type rangeRuleUint struct {
	fieldExpr string
	ge        uint
	le        uint
	name      string
}

func (r rangeRuleUint) Check(param interface{}) (bool, string) {
	exprValueUint, isValid, errMsg := getUintField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if exprValueUint > r.le || exprValueUint < r.ge {
		return false,
			fmt.Sprintf("[%s]:'%s' should be betweeen %d and %d ,actual is %d",
				r.name, r.fieldExpr, r.ge, r.le, exprValueUint)
	}
	return true, ""
}

// NewRangeRuleUint is the validation function for validating if the field's value is in given range.
func NewRangeRuleUint(filedExpr string, ge uint, le uint) Rule {
	return rangeRuleUint{
		fieldExpr: filedExpr,
		ge:        ge,
		le:        le,
		name:      "rangeRuleUint",
	}
}

type rangeRuleFloat struct {
	fieldExpr string
	le        float64
	ge        float64
	name      string
}

func (r rangeRuleFloat) Check(param interface{}) (bool, string) {
	exprValueFloat, isValid, errMsg := getFloatField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if exprValueFloat > r.ge || exprValueFloat < r.le {
		return false,
			fmt.Sprintf("[%s]:'%s' should be between %f and %f,actual is %f",
				r.name, r.fieldExpr, r.le, r.ge, exprValueFloat)
	}
	return true, ""
}

// NewRangeRuleFloat is the validation function for validating if the field's value is in given range.
func NewRangeRuleFloat(filedExpr string, ge float64, le float64) Rule {
	return rangeRuleFloat{
		fieldExpr: filedExpr,
		le:        ge,
		ge:        le,
		name:      "rangeRuleFloat",
	}
}

type rangeRuleTimestamp struct {
	fieldExpr string
	le        time.Time
	ge        time.Time
	name      string
}

func (r rangeRuleTimestamp) Check(param interface{}) (bool, string) {
	tsVal, isValid, errMsg := getTimeField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if tsVal.Before(r.ge) || tsVal.After(r.le) {
		return false,
			fmt.Sprintf("[%s]:'%s' should be between %v and %v,actual is %v",
				r.name, r.fieldExpr, r.le, r.ge, tsVal)
	}
	return true, ""
}

// NewRangeRuleTimestamp is the validation function for validating if the field's value is in given range.
func NewRangeRuleTimestamp(filedExpr string, ge time.Time, le time.Time) Rule {
	return rangeRuleTimestamp{
		fieldExpr: filedExpr,
		le:        le,
		ge:        ge,
		name:      "rangeRuleTimestamp",
	}
}

type rangeRuleTimestampStr struct {
	fieldExpr string
	layout    string
	le        time.Time
	ge        time.Time
	name      string
}

func (r rangeRuleTimestampStr) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := getStrField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	exprValTime, err := time.Parse(r.layout, exprValueStr)
	if err != nil {
		return false,
			fmt.Sprintf("[%s]:'%s' should be format %s,actual is %s",
				r.name, r.fieldExpr, r.layout, exprValueStr)
	}

	if exprValTime.Before(r.ge) || exprValTime.After(r.le) {
		return false,
			fmt.Sprintf("[%s]:'%s' should be between %v and %v,actual is %v",
				r.name, r.fieldExpr, r.le.Format(r.layout), r.ge.Format(r.layout), exprValueStr)
	}
	return true, ""
}

// NewRangeRuleTimestampStr is the validation function for validating if the field's value is in given range.
func NewRangeRuleTimestampStr(filedExpr string, layout string, ge time.Time, le time.Time) Rule {
	return rangeRuleTimestampStr{
		fieldExpr: filedExpr,
		layout:    layout,
		le:        le,
		ge:        ge,
		name:      "rangeRuleTimestampStr",
	}
}

func getIntField(param interface{}, fieldExpr string, name string) (int, bool, string) {
	exprValue, kind := fetchFieldInStruct(param, fieldExpr)
	if kind == reflect.Invalid {
		return 0, false,
			fmt.Sprintf("[%s]:'%s' cannot be found", name, fieldExpr)
	}
	if exprValue == nil {
		return 0, false,
			fmt.Sprintf("[%s]:'%s' is nil", name, fieldExpr)
	}
	if kind != reflect.Int8 && kind != reflect.Int16 && kind != reflect.Int32 &&
		kind != reflect.Int64 && kind != reflect.Int {
		return 0, false,
			fmt.Sprintf("[%s]:'%s' should be kind int8/int16/int32/int64/int,actual is %v",
				name, fieldExpr, kind)
	}
	return exprValue.(int), true, ""
}

func getUintField(param interface{}, fieldExpr string, name string) (uint, bool, string) {
	exprValue, kind := fetchFieldInStruct(param, fieldExpr)
	if kind == reflect.Invalid {
		return 0, false,
			fmt.Sprintf("[%s]:'%s' cannot be found", name, fieldExpr)
	}
	if exprValue == nil {
		return 0, false,
			fmt.Sprintf("[%s]:'%s' is nil", name, fieldExpr)
	}
	if kind != reflect.Uint8 && kind != reflect.Uint16 && kind != reflect.Uint32 &&
		kind != reflect.Uint64 && kind != reflect.Uint {
		return 0, false,
			fmt.Sprintf("[%s]:'%s' should be kind uint8/uint16/uint32/uint64/uint,actual is %v",
				name, fieldExpr, kind)
	}
	return exprValue.(uint), true, ""
}

func getFloatField(param interface{}, fieldExpr string, name string) (float64, bool, string) {
	exprValue, kind := fetchFieldInStruct(param, fieldExpr)
	if kind == reflect.Invalid {
		return 0.0, false,
			fmt.Sprintf("[%s]:'%s' cannot be found", name, fieldExpr)
	}
	if exprValue == nil {
		return 0.0, false,
			fmt.Sprintf("[%s]:'%s' is nil", name, fieldExpr)
	}

	if kind != reflect.Float32 && kind != reflect.Float64 {
		return 0.0, false,
			fmt.Sprintf("[%s]:'%s' should be kind float32/float64,actual is %v",
				name, fieldExpr, kind)
	}
	return exprValue.(float64), true, ""
}

func getTimeField(param interface{}, fieldExpr string, name string) (time.Time, bool, string) {
	exprValue, kind := fetchFieldInStruct(param, fieldExpr)
	if kind == reflect.Invalid {
		return time.Time{}, false,
			fmt.Sprintf("[%s]:'%s' cannot be found", name, fieldExpr)
	}
	if exprValue == nil {
		return time.Time{}, false,
			fmt.Sprintf("[%s]:'%s' is nil", name, fieldExpr)
	}
	tsVal, ok := exprValue.(time.Time)
	if !ok {
		return time.Time{}, false,
			fmt.Sprintf("[%s]:'%s' should be time.Time,actual is %v",
				name, fieldExpr, reflect.TypeOf(exprValue).String())
	}
	return tsVal, true, ""
}

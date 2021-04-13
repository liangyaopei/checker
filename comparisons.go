package checker

import (
	"fmt"
	"time"
)

// <=======================Comparison rule about string ===============>

type eqRuleString struct {
	BaseRule
	equivalent string
}

func (r eqRuleString) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := fetchFieldStr(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if exprValueStr != r.equivalent {
		return false,
			fmt.Sprintf("[eqRuleString]:'%s' should be %s,actual is %s", r.fieldExpr, r.equivalent, exprValueStr)
	}
	return true, ""
}

// EqStr is the validation function for validating if the field's value is equal to given string.
func EqStr(filedExpr string, equivalent string) *eqRuleString {
	return &eqRuleString{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "EqStr",
		},
		equivalent,
	}
}

type neRuleString struct {
	fieldExpr    string
	inequivalent string
	name         string
}

func (r neRuleString) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := fetchFieldStr(param, r.fieldExpr, r.name)
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

// NeStr is the validation function for validating if the field's value is not equal to given string.
func NeStr(filedExpr string, inequivalent string) Rule {
	return neRuleString{
		fieldExpr:    filedExpr,
		inequivalent: inequivalent,
		name:         "neRuleString",
	}
}

// <=======================Comparison rule about int ===============>

type eqRuleInt struct {
	fieldExpr  string
	equivalent int
	name       string
}

func (r eqRuleInt) Check(param interface{}) (bool, string) {
	exprValueInt, isValid, errMsg := fetchFieldInt(param, r.fieldExpr, r.name)
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

// EqInt is the validation function for validating if the field's value is equal to given int.
func EqInt(filedExpr string, equivalent int) Rule {
	return eqRuleInt{
		fieldExpr:  filedExpr,
		equivalent: equivalent,
		name:       "eqRuleInt",
	}
}

type neRuleInt struct {
	fieldExpr    string
	inequivalent int
	name         string
}

func (r neRuleInt) Check(param interface{}) (bool, string) {
	exprValueInt, isValid, errMsg := fetchFieldInt(param, r.fieldExpr, r.name)
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

// NeInt is the validation function for validating if the field's value is not equal to given int.
func NeInt(filedExpr string, inequivalent int) Rule {
	return neRuleInt{
		fieldExpr:    filedExpr,
		inequivalent: inequivalent,
		name:         "neRuleInt",
	}
}

type rangeRuleInt struct {
	fieldExpr string
	ge        int
	le        int
	name      string
}

func (r rangeRuleInt) Check(param interface{}) (bool, string) {
	exprValueInt, isValid, errMsg := fetchFieldInt(param, r.fieldExpr, r.name)
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

// RangeInt is the validation function for validating if the field's value is in given range.
func RangeInt(filedExpr string, ge int, le int) Rule {
	return rangeRuleInt{
		fieldExpr: filedExpr,
		ge:        ge,
		le:        le,
		name:      "rangeRuleInt",
	}
}

// <=======================Comparison rule about uint ===============>

type eqRuleUint struct {
	fieldExpr  string
	equivalent uint
	name       string
}

func (r eqRuleUint) Check(param interface{}) (bool, string) {
	exprValueUint, isValid, errMsg := fetchFieldUint(param, r.fieldExpr, r.name)
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

// EqUint is the validation function for validating if the field's value is equal to given uint.
func EqUint(filedExpr string, equivalent uint) Rule {
	return eqRuleUint{
		fieldExpr:  filedExpr,
		equivalent: equivalent,
		name:       "eqRuleUint",
	}
}

type neRuleUint struct {
	fieldExpr    string
	inequivalent uint
	name         string
}

func (r neRuleUint) Check(param interface{}) (bool, string) {
	exprValueUint, isValid, errMsg := fetchFieldUint(param, r.fieldExpr, r.name)
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

// NeUint is the validation function for validating if the field's value is not equal to given uint.
func NeUint(filedExpr string, inequivalent uint) Rule {
	return neRuleUint{
		fieldExpr:    filedExpr,
		inequivalent: inequivalent,
		name:         "neRuleUint",
	}
}

type rangeRuleUint struct {
	fieldExpr string
	ge        uint
	le        uint
	name      string
}

func (r rangeRuleUint) Check(param interface{}) (bool, string) {
	exprValueUint, isValid, errMsg := fetchFieldUint(param, r.fieldExpr, r.name)
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

// RangeUint is the validation function for validating if the field's value is in given range.
func RangeUint(filedExpr string, ge uint, le uint) Rule {
	return rangeRuleUint{
		fieldExpr: filedExpr,
		ge:        ge,
		le:        le,
		name:      "rangeRuleUint",
	}
}

// <=======================Comparison rule about float64 ===============>

type eqRuleFloat struct {
	fieldExpr  string
	equivalent float64
	name       string
}

func (r eqRuleFloat) Check(param interface{}) (bool, string) {
	exprValueFloat, isValid, errMsg := fetchFieldFloat(param, r.fieldExpr, r.name)
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

// EqFloat is the validation function for validating if the field's value is equal to given float.
func EqFloat(filedExpr string, equivalent float64) Rule {
	return eqRuleFloat{
		fieldExpr:  filedExpr,
		equivalent: equivalent,
		name:       "eqRuleFloat",
	}
}

type neRuleFloat struct {
	fieldExpr    string
	inequivalent float64
	name         string
}

func (r neRuleFloat) Check(param interface{}) (bool, string) {
	exprValueFloat, isValid, errMsg := fetchFieldFloat(param, r.fieldExpr, r.name)
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

// NeFloat is the validation function for validating if the field's value is not equal to given float.
func NeFloat(filedExpr string, inequivalent float64) Rule {
	return neRuleFloat{
		fieldExpr:    filedExpr,
		inequivalent: inequivalent,
		name:         "neRuleFloat",
	}
}

type rangeRuleFloat struct {
	fieldExpr string
	le        float64
	ge        float64
	name      string
}

func (r rangeRuleFloat) Check(param interface{}) (bool, string) {
	exprValueFloat, isValid, errMsg := fetchFieldFloat(param, r.fieldExpr, r.name)
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

// RangeFloat is the validation function for validating if the field's value is in given range.
func RangeFloat(filedExpr string, ge float64, le float64) Rule {
	return rangeRuleFloat{
		fieldExpr: filedExpr,
		le:        ge,
		ge:        le,
		name:      "rangeRuleFloat",
	}
}

// <=======================Comparison rule about time.Time ===============>

type eqRuleTime struct {
	fieldExpr  string
	equivalent time.Time
	name       string
}

func (r eqRuleTime) Check(param interface{}) (bool, string) {
	tsVal, isValid, errMsg := fetchFieldTime(param, r.fieldExpr, r.name)
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

// EqTime is the validation function for validating if the field's value is equal to given timestamp.
func EqTime(filedExpr string, equivalent time.Time) Rule {
	return eqRuleTime{
		fieldExpr:  filedExpr,
		equivalent: equivalent,
		name:       "eqRuleTime",
	}
}

type neRuleTime struct {
	fieldExpr    string
	inequivalent time.Time
	name         string
}

func (r neRuleTime) Check(param interface{}) (bool, string) {
	tsVal, isValid, errMsg := fetchFieldTime(param, r.fieldExpr, r.name)
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

// NeTime is the validation function for validating if the field's value is not equal to given timestamp.
func NeTime(filedExpr string, inequivalent time.Time) Rule {
	return neRuleTime{
		fieldExpr:    filedExpr,
		inequivalent: inequivalent,
		name:         "neRuleTime",
	}
}

type eqRuleTimeStr struct {
	fieldExpr  string
	layout     string
	equivalent time.Time
	name       string
}

func (r eqRuleTimeStr) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := fetchFieldStr(param, r.fieldExpr, r.name)
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

// EqTimeStr is the validation function for validating if the field's value string and is equal to given timestamp.
func EqTimeStr(filedExpr string, layout string, equivalent time.Time) Rule {
	return eqRuleTimeStr{
		fieldExpr:  filedExpr,
		layout:     layout,
		equivalent: equivalent,
		name:       "eqRuleTimeStr",
	}
}

type neRuleTimeStr struct {
	fieldExpr    string
	layout       string
	inequivalent time.Time
	name         string
}

func (r neRuleTimeStr) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := fetchFieldStr(param, r.fieldExpr, r.name)
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

// NeTimeStr is the validation function for validating if the field's value string and is not equal to given timestamp.
func NeTimeStr(filedExpr string, layout string, inequivalent time.Time) Rule {
	return neRuleTimeStr{
		fieldExpr:    filedExpr,
		layout:       layout,
		inequivalent: inequivalent,
		name:         "neRuleTimeStr",
	}
}

type rangeRuleTime struct {
	fieldExpr string
	le        time.Time
	ge        time.Time
	name      string
}

func (r rangeRuleTime) Check(param interface{}) (bool, string) {
	tsVal, isValid, errMsg := fetchFieldTime(param, r.fieldExpr, r.name)
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

// RangeTime is the validation function for validating if the field's value is in given range.
func RangeTime(filedExpr string, ge time.Time, le time.Time) Rule {
	return rangeRuleTime{
		fieldExpr: filedExpr,
		le:        le,
		ge:        ge,
		name:      "rangeRuleTime",
	}
}

type rangeRuleTimeStr struct {
	fieldExpr string
	layout    string
	le        time.Time
	ge        time.Time
	name      string
}

func (r rangeRuleTimeStr) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := fetchFieldStr(param, r.fieldExpr, r.name)
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

// RangeTimeStr is the validation function for validating if the field's value is in given range.
func RangeTimeStr(filedExpr string, layout string, ge time.Time, le time.Time) Rule {
	return rangeRuleTimeStr{
		fieldExpr: filedExpr,
		layout:    layout,
		le:        le,
		ge:        ge,
		name:      "rangeRuleTimeStr",
	}
}

// <=======================Comparison rule about Comparable ===============>

type eqRuleComp struct {
	fieldExpr  string
	equivalent Comparable
	name       string
}

func (r eqRuleComp) Check(param interface{}) (bool, string) {
	exprValueComp, isValid, errMsg := fetchFieldComparable(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if !exprValueComp.EqualTo(r.equivalent) {
		return false,
			fmt.Sprintf("[%s]:'%s' should be %v,actual is %v",
				r.name, r.fieldExpr, r.equivalent, exprValueComp)

	}
	return true, ""
}

// EqComp checks if the field is Comparable
// and is equal to given Comparable variable
func EqComp(filedExpr string, equivalent Comparable) Rule {
	return eqRuleComp{
		fieldExpr:  filedExpr,
		equivalent: equivalent,
		name:       "eqRuleComp",
	}
}

type neRuleComp struct {
	fieldExpr    string
	inequivalent Comparable
	name         string
}

func (r neRuleComp) Check(param interface{}) (bool, string) {
	exprValueComp, isValid, errMsg := fetchFieldComparable(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if exprValueComp.EqualTo(r.inequivalent) {
		return false,
			fmt.Sprintf("[%s]:'%s' should not be equal to %v,actual is %v",
				r.name, r.fieldExpr, r.inequivalent, exprValueComp)

	}
	return true, ""
}

// NeComp checks if the field is Comparable
// and is not equal to given Comparable variable
func NeComp(filedExpr string, inequivalent Comparable) Rule {
	return neRuleComp{
		fieldExpr:    filedExpr,
		inequivalent: inequivalent,
		name:         "neRuleComp",
	}
}

type rangeRuleComp struct {
	fieldExpr string
	le        Comparable
	ge        Comparable
	name      string
}

func (r rangeRuleComp) Check(param interface{}) (bool, string) {
	exprValueComp, isValid, errMsg := fetchFieldComparable(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if exprValueComp.LessThan(r.ge) || !exprValueComp.LessThan(r.le) {
		return false, fmt.Sprintf("[%s]:'%s' should be between %v and %v,actual is %v",
			r.name, r.fieldExpr, r.ge, r.le, exprValueComp)
	}
	return true, ""
}

// RangeComp checks if the value
// is between [ge,le]
func RangeComp(filedExpr string, ge Comparable, le Comparable) Rule {
	return rangeRuleComp{
		fieldExpr: filedExpr,
		ge:        ge,
		le:        le,
		name:      "rangeRuleComp",
	}
}

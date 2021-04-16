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

func (r *eqRuleString) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r eqRuleString) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := fetchFieldStr(param, &r)
	if !isValid {
		return false, errMsg
	}
	if exprValueStr != r.equivalent {
		return false,
			fmt.Sprintf("[%s]:%s should be %s,actual is %s",
				r.name, r.getCompleteFieldExpr(), r.equivalent, exprValueStr)
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
	BaseRule
	inequivalent string
}

func (r *neRuleString) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r neRuleString) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := fetchFieldStr(param, &r)
	if !isValid {
		return false, errMsg
	}
	if exprValueStr == r.inequivalent {
		return false,
			fmt.Sprintf("[%s]:%s should not be %s,actual is %s",
				r.name, r.getCompleteFieldExpr(), r.inequivalent, exprValueStr)
	}
	return true, ""
}

// NeStr is the validation function for validating if the field's value is not equal to given string.
func NeStr(filedExpr string, inequivalent string) *neRuleString {
	return &neRuleString{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "NeStr",
		},
		inequivalent,
	}
}

// <=======================Comparison rule about int ===============>

type eqRuleInt struct {
	BaseRule
	equivalent int
}

func (r *eqRuleInt) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r eqRuleInt) Check(param interface{}) (bool, string) {
	exprValueInt, isValid, errMsg := fetchFieldInt(param, &r)
	if !isValid {
		return false, errMsg
	}
	if exprValueInt != r.equivalent {
		return false,
			fmt.Sprintf("[%s]:%s should be %d,actual is %d",
				r.name, r.getCompleteFieldExpr(), r.equivalent, exprValueInt)
	}
	return true, ""
}

// EqInt is the validation function for validating if the field's value is equal to given int.
func EqInt(filedExpr string, equivalent int) *eqRuleInt {
	return &eqRuleInt{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "EqInt",
		},
		equivalent,
	}
}

type neRuleInt struct {
	BaseRule
	inequivalent int
}

func (r *neRuleInt) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r neRuleInt) Check(param interface{}) (bool, string) {
	exprValueInt, isValid, errMsg := fetchFieldInt(param, &r)
	if !isValid {
		return false, errMsg
	}
	if exprValueInt == r.inequivalent {
		return false,
			fmt.Sprintf("[%s]:%s should not be %d,actual is %d",
				r.name, r.getCompleteFieldExpr(), r.inequivalent, exprValueInt)
	}
	return true, ""
}

// NeInt is the validation function for validating if the field's value is not equal to given int.
func NeInt(filedExpr string, inequivalent int) *neRuleInt {
	return &neRuleInt{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "NeInt",
		},
		inequivalent,
	}
}

type rangeRuleInt struct {
	BaseRule
	ge int
	le int
}

func (r *rangeRuleInt) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r rangeRuleInt) Check(param interface{}) (bool, string) {
	exprValueInt, isValid, errMsg := fetchFieldInt(param, &r)
	if !isValid {
		return false, errMsg
	}
	if exprValueInt > r.le || exprValueInt < r.ge {
		return false,
			fmt.Sprintf("[%s]:%s should be between %d and %d,actual is %d",
				r.name, r.getCompleteFieldExpr(), r.ge, r.le, exprValueInt)
	}
	return true, ""
}

// RangeInt is the validation function for validating if the field's value is in given range.
func RangeInt(filedExpr string, ge int, le int) *rangeRuleInt {
	return &rangeRuleInt{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "RangeInt",
		},
		ge,
		le,
	}
}

// <=======================Comparison rule about uint ===============>

type eqRuleUint struct {
	BaseRule
	equivalent uint
}

func (r *eqRuleUint) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r eqRuleUint) Check(param interface{}) (bool, string) {
	exprValueUint, isValid, errMsg := fetchFieldUint(param, &r)
	if !isValid {
		return false, errMsg
	}
	if exprValueUint != r.equivalent {
		return false,
			fmt.Sprintf("[%s]:%s should be %d,actual is %d",
				r.name, r.getCompleteFieldExpr(), r.equivalent, exprValueUint)
	}
	return true, ""
}

// EqUint is the validation function for validating if the field's value is equal to given uint.
func EqUint(filedExpr string, equivalent uint) *eqRuleUint {
	return &eqRuleUint{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "EqUint",
		},
		equivalent,
	}
}

type neRuleUint struct {
	BaseRule
	inequivalent uint
}

func (r *neRuleUint) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r neRuleUint) Check(param interface{}) (bool, string) {
	exprValueUint, isValid, errMsg := fetchFieldUint(param, &r)
	if !isValid {
		return false, errMsg
	}
	if exprValueUint == r.inequivalent {
		return false,
			fmt.Sprintf("[%s]:%s should not be %d,actual is %d",
				r.name, r.getCompleteFieldExpr(), r.inequivalent, exprValueUint)
	}
	return true, ""
}

// NeUint is the validation function for validating if the field's value is not equal to given uint.
func NeUint(filedExpr string, inequivalent uint) *neRuleUint {
	return &neRuleUint{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "NeUint",
		},
		inequivalent,
	}
}

type rangeRuleUint struct {
	BaseRule
	ge uint
	le uint
}

func (r *rangeRuleUint) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r rangeRuleUint) Check(param interface{}) (bool, string) {
	exprValueUint, isValid, errMsg := fetchFieldUint(param, &r)
	if !isValid {
		return false, errMsg
	}
	if exprValueUint > r.le || exprValueUint < r.ge {
		return false,
			fmt.Sprintf("[%s]:%s should be betweeen %d and %d ,actual is %d",
				r.name, r.getCompleteFieldExpr(), r.ge, r.le, exprValueUint)
	}
	return true, ""
}

// RangeUint is the validation function for validating if the field's value is in given range.
func RangeUint(filedExpr string, ge uint, le uint) *rangeRuleUint {
	return &rangeRuleUint{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "RangeUint",
		},
		ge,
		le,
	}
}

// <=======================Comparison rule about float64 ===============>

type eqRuleFloat struct {
	BaseRule
	equivalent float64
}

func (r *eqRuleFloat) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r eqRuleFloat) Check(param interface{}) (bool, string) {
	exprValueFloat, isValid, errMsg := fetchFieldFloat(param, &r)
	if !isValid {
		return false, errMsg
	}
	if exprValueFloat != r.equivalent {
		return false,
			fmt.Sprintf("[%s]:%s should be %f,actual is %f",
				r.name, r.getCompleteFieldExpr(), r.equivalent, exprValueFloat)
	}
	return true, ""
}

// EqFloat is the validation function for validating if the field's value is equal to given float.
func EqFloat(filedExpr string, equivalent float64) *eqRuleFloat {
	return &eqRuleFloat{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "EqFloat",
		},
		equivalent,
	}
}

type neRuleFloat struct {
	BaseRule
	inequivalent float64
}

func (r *neRuleFloat) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r neRuleFloat) Check(param interface{}) (bool, string) {
	exprValueFloat, isValid, errMsg := fetchFieldFloat(param, &r)
	if !isValid {
		return false, errMsg
	}
	if exprValueFloat == r.inequivalent {
		return false,
			fmt.Sprintf("[%s]:%s should not be %f,actual is %f",
				r.name, r.getCompleteFieldExpr(), r.inequivalent, exprValueFloat)
	}
	return true, ""
}

// NeFloat is the validation function for validating if the field's value is not equal to given float.
func NeFloat(filedExpr string, inequivalent float64) *neRuleFloat {
	return &neRuleFloat{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "NeFloat",
		},
		inequivalent,
	}
}

type rangeRuleFloat struct {
	BaseRule
	le float64
	ge float64
}

func (r *rangeRuleFloat) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r rangeRuleFloat) Check(param interface{}) (bool, string) {
	exprValueFloat, isValid, errMsg := fetchFieldFloat(param, &r)
	if !isValid {
		return false, errMsg
	}
	if exprValueFloat < r.ge || exprValueFloat > r.le {
		return false,
			fmt.Sprintf("[%s]:%s should be between %f and %f,actual is %f",
				r.name, r.getCompleteFieldExpr(), r.le, r.ge, exprValueFloat)
	}
	return true, ""
}

// RangeFloat is the validation function for validating if the field's value is in given range.
func RangeFloat(filedExpr string, ge float64, le float64) *rangeRuleFloat {
	return &rangeRuleFloat{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "RangeFloat",
		},
		le,
		ge,
	}
}

// <=======================Comparison rule about time.Time ===============>

type eqRuleTime struct {
	BaseRule
	equivalent time.Time
}

func (r *eqRuleTime) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r eqRuleTime) Check(param interface{}) (bool, string) {
	tsVal, isValid, errMsg := fetchFieldTime(param, &r)
	if !isValid {
		return false, errMsg
	}
	if !tsVal.Equal(r.equivalent) {
		return false,
			fmt.Sprintf("[%s]:%s should be %v,actual is %v",
				r.name, r.getCompleteFieldExpr(), r.equivalent, tsVal)
	}
	return true, ""
}

// EqTime is the validation function for validating if the field's value is equal to given timestamp.
func EqTime(filedExpr string, equivalent time.Time) *eqRuleTime {
	return &eqRuleTime{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "EqTime",
		},
		equivalent,
	}
}

type neRuleTime struct {
	BaseRule
	inequivalent time.Time
}

func (r *neRuleTime) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r neRuleTime) Check(param interface{}) (bool, string) {
	tsVal, isValid, errMsg := fetchFieldTime(param, &r)
	if !isValid {
		return false, errMsg
	}
	if tsVal.Equal(r.inequivalent) {
		return false,
			fmt.Sprintf("[%s]:%s should not be %v,actual is %v",
				r.name, r.getCompleteFieldExpr(), r.inequivalent, tsVal)
	}
	return true, ""
}

// NeTime is the validation function for validating if the field's value is not equal to given timestamp.
func NeTime(filedExpr string, inequivalent time.Time) *neRuleTime {
	return &neRuleTime{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "NeTime",
		},
		inequivalent,
	}
}

type eqRuleTimeStr struct {
	BaseRule
	layout     string
	equivalent time.Time
}

func (r *eqRuleTimeStr) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r eqRuleTimeStr) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := fetchFieldStr(param, &r)
	if !isValid {
		return false, errMsg
	}
	exprValTime, err := time.Parse(r.layout, exprValueStr)
	if err != nil {
		return false,
			fmt.Sprintf("[%s]:%s should be format %s,actual is %s",
				r.name, r.getCompleteFieldExpr(), r.layout, exprValueStr)
	}

	if !exprValTime.Equal(r.equivalent) {
		return false,
			fmt.Sprintf("[%s]:%s should be %v,actual is %v",
				r.name, r.getCompleteFieldExpr(), r.equivalent.Format(r.layout), exprValueStr)
	}
	return true, ""
}

// EqTimeStr is the validation function for validating if the field's value string and is equal to given timestamp.
func EqTimeStr(filedExpr string, layout string, equivalent time.Time) *eqRuleTimeStr {
	return &eqRuleTimeStr{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "EqTimeStr",
		},
		layout,
		equivalent,
	}
}

type neRuleTimeStr struct {
	BaseRule
	layout       string
	inequivalent time.Time
}

func (r *neRuleTimeStr) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r neRuleTimeStr) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := fetchFieldStr(param, &r)
	if !isValid {
		return false, errMsg
	}
	exprValTime, err := time.Parse(r.layout, exprValueStr)
	if err != nil {
		return false,
			fmt.Sprintf("[%s]:%s should be format %s,actual is %s",
				r.name, r.getCompleteFieldExpr(), r.layout, exprValueStr)
	}

	if exprValTime.Equal(r.inequivalent) {
		return false,
			fmt.Sprintf("[%s]:%s should not be %v,actual is %v",
				r.name, r.getCompleteFieldExpr(), r.inequivalent.Format(r.layout), exprValueStr)
	}
	return true, ""
}

// NeTimeStr is the validation function for validating if the field's value string and is not equal to given timestamp.
func NeTimeStr(filedExpr string, layout string, inequivalent time.Time) *neRuleTimeStr {
	return &neRuleTimeStr{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "NeTimeStr",
		},
		layout,
		inequivalent,
	}
}

type rangeRuleTime struct {
	BaseRule
	le time.Time
	ge time.Time
}

func (r *rangeRuleTime) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r rangeRuleTime) Check(param interface{}) (bool, string) {
	tsVal, isValid, errMsg := fetchFieldTime(param, &r)
	if !isValid {
		return false, errMsg
	}
	if tsVal.Before(r.ge) || tsVal.After(r.le) {
		return false,
			fmt.Sprintf("[%s]:%s should be between %v and %v,actual is %v",
				r.name, r.getCompleteFieldExpr(), r.le, r.ge, tsVal)
	}
	return true, ""
}

// RangeTime is the validation function for validating if the field's value is in given range.
func RangeTime(filedExpr string, ge time.Time, le time.Time) *rangeRuleTime {
	return &rangeRuleTime{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "RangeTime",
		},
		le,
		ge,
	}
}

type rangeRuleTimeStr struct {
	BaseRule
	layout string
	le     time.Time
	ge     time.Time
}

func (r *rangeRuleTimeStr) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r rangeRuleTimeStr) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := fetchFieldStr(param, &r)
	if !isValid {
		return false, errMsg
	}
	exprValTime, err := time.Parse(r.layout, exprValueStr)
	if err != nil {
		return false,
			fmt.Sprintf("[%s]:%s should be format %s,actual is %s",
				r.name, r.getCompleteFieldExpr(), r.layout, exprValueStr)
	}

	if exprValTime.Before(r.ge) || exprValTime.After(r.le) {
		return false,
			fmt.Sprintf("[%s]:%s should be between %v and %v,actual is %v",
				r.name, r.getCompleteFieldExpr(), r.le.Format(r.layout), r.ge.Format(r.layout), exprValueStr)
	}
	return true, ""
}

// RangeTimeStr is the validation function for validating if the field's value is in given range.
func RangeTimeStr(filedExpr string, layout string, ge time.Time, le time.Time) *rangeRuleTimeStr {
	return &rangeRuleTimeStr{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "RangeTimeStr",
		},
		layout,
		le,
		ge,
	}
}

// <=======================Comparison rule about Comparable ===============>

type eqRuleComp struct {
	BaseRule
	equivalent Comparable
}

func (r *eqRuleComp) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r eqRuleComp) Check(param interface{}) (bool, string) {
	exprValueComp, isValid, errMsg := fetchFieldComparable(param, &r)
	if !isValid {
		return false, errMsg
	}
	if !exprValueComp.EqualTo(r.equivalent) {
		return false,
			fmt.Sprintf("[%s]:%s should be %v,actual is %v",
				r.name, r.getCompleteFieldExpr(), r.equivalent, exprValueComp)

	}
	return true, ""
}

// EqComp checks if the field is Comparable
// and is equal to given Comparable variable
func EqComp(filedExpr string, equivalent Comparable) *eqRuleComp {
	return &eqRuleComp{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "EqComp",
		},
		equivalent,
	}
}

type neRuleComp struct {
	BaseRule
	inequivalent Comparable
}

func (r *neRuleComp) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r neRuleComp) Check(param interface{}) (bool, string) {
	exprValueComp, isValid, errMsg := fetchFieldComparable(param, &r)
	if !isValid {
		return false, errMsg
	}
	if exprValueComp.EqualTo(r.inequivalent) {
		return false,
			fmt.Sprintf("[%s]:%s should not be equal to %v,actual is %v",
				r.name, r.getCompleteFieldExpr(), r.inequivalent, exprValueComp)

	}
	return true, ""
}

// NeComp checks if the field is Comparable
// and is not equal to given Comparable variable
func NeComp(filedExpr string, inequivalent Comparable) *neRuleComp {
	return &neRuleComp{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "NeComp",
		},
		inequivalent,
	}
}

type rangeRuleComp struct {
	BaseRule
	le Comparable
	ge Comparable
}

func (r *rangeRuleComp) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r rangeRuleComp) Check(param interface{}) (bool, string) {
	exprValueComp, isValid, errMsg := fetchFieldComparable(param, &r)
	if !isValid {
		return false, errMsg
	}
	if exprValueComp.LessThan(r.ge) || !exprValueComp.LessThan(r.le) {
		return false, fmt.Sprintf("[%s]:%s should be between %v and %v,actual is %v",
			r.name, r.getCompleteFieldExpr(), r.ge, r.le, exprValueComp)
	}
	return true, ""
}

// RangeComp checks if the value
// is between [ge,le]
func RangeComp(filedExpr string, ge Comparable, le Comparable) *rangeRuleComp {
	return &rangeRuleComp{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "RangeComp",
		},
		le,
		ge,
	}
}

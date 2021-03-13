package checker

import (
	"fmt"
	"time"
)

type eqRuleString struct {
	fieldExpr  string
	equivalent string
	name       string
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
func EqStr(filedExpr string, equivalent string) Rule {
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

type eqRuleTimestamp struct {
	fieldExpr  string
	equivalent time.Time
	name       string
}

func (r eqRuleTimestamp) Check(param interface{}) (bool, string) {
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

// EqTs is the validation function for validating if the field's value is equal to given timestamp.
func EqTs(filedExpr string, equivalent time.Time) Rule {
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

// EqTsStr is the validation function for validating if the field's value string and is equal to given timestamp.
func EqTsStr(filedExpr string, layout string, equivalent time.Time) Rule {
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

type neTimestamp struct {
	fieldExpr    string
	inequivalent time.Time
	name         string
}

func (r neTimestamp) Check(param interface{}) (bool, string) {
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

// NeTs is the validation function for validating if the field's value is not equal to given timestamp.
func NeTs(filedExpr string, inequivalent time.Time) Rule {
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

// NeTsStr is the validation function for validating if the field's value string and is not equal to given timestamp.
func NeTsStr(filedExpr string, layout string, inequivalent time.Time) Rule {
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

type rangeRuleTimestamp struct {
	fieldExpr string
	le        time.Time
	ge        time.Time
	name      string
}

func (r rangeRuleTimestamp) Check(param interface{}) (bool, string) {
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

// RangeTs is the validation function for validating if the field's value is in given range.
func RangeTs(filedExpr string, ge time.Time, le time.Time) Rule {
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

// RangeTsStr is the validation function for validating if the field's value is in given range.
func RangeTsStr(filedExpr string, layout string, ge time.Time, le time.Time) Rule {
	return rangeRuleTimestampStr{
		fieldExpr: filedExpr,
		layout:    layout,
		le:        le,
		ge:        ge,
		name:      "rangeRuleTimestampStr",
	}
}

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

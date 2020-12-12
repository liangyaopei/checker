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

func (r eqRuleInt) Check(param interface{}) (bool, string) {
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

func (r eqRuleUint) Check(param interface{}) (bool, string) {
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

func (r eqRuleFloat) Check(param interface{}) (bool, string) {
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

type eqRuleTimestamp struct {
	fieldExpr  string
	equivalent time.Time
	name       string
}

func (r eqRuleTimestamp) Check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.fieldExpr)
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is nil", r.name, r.fieldExpr)
	}
	tsVal, ok := exprValue.(time.Time)
	if !ok {
		return false,
			fmt.Sprintf("[%s]:'%s' should be time.Time,actual is %v",
				r.name, r.fieldExpr, reflect.TypeOf(exprValue).String())
	}
	if !tsVal.Equal(r.equivalent) {
		return false,
			fmt.Sprintf("[%s]:'%s' should be %v,actual is %v",
				r.name, r.fieldExpr, r.equivalent, tsVal)
	}
	return true, ""
}

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

func NewEqRuleTimestampStr(filedExpr string, layout string, equivalent time.Time) Rule {
	return eqRuleTimestampStr{
		fieldExpr:  filedExpr,
		layout:     layout,
		equivalent: equivalent,
		name:       "eqRuleTimestampStr",
	}
}

type notEqRuleString struct {
	fieldExpr    string
	inequivalent string
	name         string
}

func (r notEqRuleString) Check(param interface{}) (bool, string) {
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

func (r notEqRuleInt) Check(param interface{}) (bool, string) {
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

func (r notEqRuleUint) Check(param interface{}) (bool, string) {
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

func (r notEqRuleFloat) Check(param interface{}) (bool, string) {
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

type notEqRuleTimestamp struct {
	fieldExpr    string
	inequivalent time.Time
	name         string
}

func (r notEqRuleTimestamp) Check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.fieldExpr)
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is nil", r.name, r.fieldExpr)
	}
	tsVal, ok := exprValue.(time.Time)
	if !ok {
		return false,
			fmt.Sprintf("[%s]:'%s' should be time.Time,actual is %v",
				r.name, r.fieldExpr, reflect.TypeOf(exprValue).String())
	}
	if tsVal.Equal(r.inequivalent) {
		return false,
			fmt.Sprintf("[%s]:'%s' should not be %v,actual is %v",
				r.name, r.fieldExpr, r.inequivalent, tsVal)
	}
	return true, ""
}

func NewNotEqRuleTimestamp(filedExpr string, inequivalent time.Time) Rule {
	return notEqRuleTimestamp{
		fieldExpr:    filedExpr,
		inequivalent: inequivalent,
		name:         "notEqRuleTimestamp",
	}
}

type notEqRuleTimestampStr struct {
	fieldExpr    string
	layout       string
	inequivalent time.Time
	name         string
}

func (r notEqRuleTimestampStr) Check(param interface{}) (bool, string) {
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

func NewNotEqRuleTimestampStr(filedExpr string, layout string, inequivalent time.Time) Rule {
	return notEqRuleTimestampStr{
		fieldExpr:    filedExpr,
		layout:       layout,
		inequivalent: inequivalent,
		name:         "eqRuleTimestampStr",
	}
}

type rangeRuleInt struct {
	fieldExpr string
	ge        int
	le        int
	name      string
}

func (r rangeRuleInt) Check(param interface{}) (bool, string) {
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
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.fieldExpr)
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is nil", r.name, r.fieldExpr)
	}
	tsVal, ok := exprValue.(time.Time)
	if !ok {
		return false,
			fmt.Sprintf("[%s]:'%s' should be time.Time,actual is %v",
				r.name, r.fieldExpr, reflect.TypeOf(exprValue).String())
	}
	if tsVal.Before(r.ge) || tsVal.After(r.le) {
		return false,
			fmt.Sprintf("[%s]:'%s' should be between %v and %v,actual is %v",
				r.name, r.fieldExpr, r.le, r.ge, tsVal)
	}
	return true, ""
}

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

func NewRangeRuleTimestampStr(filedExpr string, layout string, ge time.Time, le time.Time) Rule {
	return rangeRuleTimestampStr{
		fieldExpr: filedExpr,
		layout:    layout,
		le:        le,
		ge:        ge,
		name:      "rangeRuleTimestampStr",
	}
}

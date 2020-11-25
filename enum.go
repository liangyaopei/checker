package checker

import (
	"fmt"
	"reflect"
)

type enumRuleString struct {
	fieldExpr string
	set       map[string]struct{}
	enum      []string
	name      string
}

func (r enumRuleString) check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' should be in %v,actual is nil",
				r.name, r.fieldExpr, r.enum)
	}
	if kind != reflect.String {
		return false,
			fmt.Sprintf("[%s]:'%s' should be kind string,actual is %v",
				r.name, r.fieldExpr, kind)
	}
	exprValueStr := exprValue.(string)
	_, exist := r.set[exprValueStr]
	if !exist {
		return false,
			fmt.Sprintf("[%s]:'%s' should be in %v,actual is %s",
				r.name, r.fieldExpr, r.enum, exprValueStr)
	}
	return true, ""
}

func NewEnumRuleString(filedExpr string, enum []string) Rule {
	set := make(map[string]struct{}, len(enum))
	for i := 0; i < len(enum); i++ {
		set[enum[i]] = struct{}{}
	}
	return enumRuleString{
		fieldExpr: filedExpr,
		set:       set,
		enum:      enum,
		name:      "enumRuleString",
	}
}

type enumRuleInt struct {
	fieldExpr string
	set       map[int]struct{}
	enum      []int
	name      string
}

func (r enumRuleInt) check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' should be in %v,actual is nil",
				r.name, r.fieldExpr, r.enum)
	}
	if kind != reflect.Int8 && kind != reflect.Int16 && kind != reflect.Int32 &&
		kind != reflect.Int64 && kind != reflect.Int {
		return false,
			fmt.Sprintf("[%s]:'%s' should be kind int8/int16/int32/int64/int,actual is %v",
				r.name, r.fieldExpr, kind)
	}
	exprValueInt := exprValue.(int)
	_, exist := r.set[exprValueInt]
	if !exist {
		return false,
			fmt.Sprintf("[%s]:'%s' should be in %v,actual is %d",
				r.name, r.fieldExpr, r.enum, exprValueInt)
	}
	return true, ""
}

func NewEnumRuleInt(filedExpr string, enum []int) Rule {
	set := make(map[int]struct{}, len(enum))
	for i := 0; i < len(enum); i++ {
		set[enum[i]] = struct{}{}
	}
	return enumRuleInt{
		fieldExpr: filedExpr,
		set:       set,
		enum:      enum,
		name:      "enumRuleInt",
	}
}

type enumRuleUint struct {
	fieldExpr string
	set       map[uint]struct{}
	enum      []uint
	name      string
}

func (r enumRuleUint) check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' should be in %v,actual is nil",
				r.name, r.fieldExpr, r.enum)
	}
	if kind != reflect.Uint8 && kind != reflect.Uint16 && kind != reflect.Uint32 &&
		kind != reflect.Uint64 && kind != reflect.Uint {
		return false,
			fmt.Sprintf("[%s]:'%s' should be kind uint8/uint16/uint32/uint64/uint,actual is %v",
				r.name, r.fieldExpr, kind)
	}
	exprValueUInt := exprValue.(uint)
	_, exist := r.set[exprValueUInt]
	if !exist {
		return false,
			fmt.Sprintf("[%s]:'%s' should be in %v,actual is %d",
				r.name, r.fieldExpr, r.enum, exprValueUInt)
	}
	return true, ""
}

func NewEnumRuleUint(filedExpr string, enum []uint) Rule {
	set := make(map[uint]struct{}, len(enum))
	for i := 0; i < len(enum); i++ {
		set[enum[i]] = struct{}{}
	}
	return enumRuleUint{
		fieldExpr: filedExpr,
		set:       set,
		enum:      enum,
		name:      "enumRuleUint",
	}
}

type enumRuleFloat struct {
	fieldExpr string
	set       map[float64]struct{}
	enum      []float64
	name      string
}

func (r enumRuleFloat) check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' should be in %v,actual is nil",
				r.name, r.fieldExpr, r.enum)
	}
	if kind != reflect.Float64 && kind != reflect.Float32 {
		return false,
			fmt.Sprintf("[%s]:'%s' should be kind float32/float64,actual is %v",
				r.name, r.fieldExpr, kind)
	}
	exprValueFloat := exprValue.(float64)
	_, exist := r.set[exprValueFloat]
	if !exist {
		return false,
			fmt.Sprintf("[%s]:'%s' should be in %v,actual is %d",
				r.name, r.fieldExpr, r.enum, exprValueFloat)
	}
	return true, ""
}

func NewEnumRuleFloat(filedExpr string, enum []float64) Rule {
	set := make(map[float64]struct{}, len(enum))
	for i := 0; i < len(enum); i++ {
		set[enum[i]] = struct{}{}
	}
	return enumRuleFloat{
		fieldExpr: filedExpr,
		set:       set,
		enum:      enum,
		name:      "enumRuleFloat",
	}
}

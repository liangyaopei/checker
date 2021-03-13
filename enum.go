package checker

import (
	"fmt"
)

type enumRuleString struct {
	fieldExpr string
	set       map[string]struct{}
	enum      []string
	name      string
}

func (r enumRuleString) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := fetchFieldStr(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	_, exist := r.set[exprValueStr]
	if !exist {
		return false,
			fmt.Sprintf("[%s]:'%s' should be in %v,actual is %s",
				r.name, r.fieldExpr, r.enum, exprValueStr)
	}
	return true, ""
}

// InStr checks if the filed is string and is in enum
func InStr(filedExpr string, enum ...string) Rule {
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

func (r enumRuleInt) Check(param interface{}) (bool, string) {
	exprValueInt, isValid, errMsg := fetchFieldInt(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	_, exist := r.set[exprValueInt]
	if !exist {
		return false,
			fmt.Sprintf("[%s]:'%s' should be in %v,actual is %d",
				r.name, r.fieldExpr, r.enum, exprValueInt)
	}
	return true, ""
}

// InInt checks if the filed is int and is in enum
func InInt(filedExpr string, enum ...int) Rule {
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

func (r enumRuleUint) Check(param interface{}) (bool, string) {
	exprValueUint, isValid, errMsg := fetchFieldUint(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	_, exist := r.set[exprValueUint]
	if !exist {
		return false,
			fmt.Sprintf("[%s]:'%s' should be in %v,actual is %d",
				r.name, r.fieldExpr, r.enum, exprValueUint)
	}
	return true, ""
}

// InUint checks if the filed is uint and is in enum
func InUint(filedExpr string, enum ...uint) Rule {
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

func (r enumRuleFloat) Check(param interface{}) (bool, string) {
	exprValueFloat, isValid, errMsg := fetchFieldFloat(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	_, exist := r.set[exprValueFloat]
	if !exist {
		return false,
			fmt.Sprintf("[%s]:'%s' should be in %v,actual is %f",
				r.name, r.fieldExpr, r.enum, exprValueFloat)
	}
	return true, ""
}

// InFloat checks if the filed is float and is in enum
func InFloat(filedExpr string, enum ...float64) Rule {
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

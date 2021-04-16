package checker

import (
	"fmt"
)

type enumRuleString struct {
	BaseRule
	set  map[string]struct{}
	enum []string
}

func (r *enumRuleString) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r enumRuleString) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := fetchFieldStr(param, &r)
	if !isValid {
		return false, errMsg
	}
	_, exist := r.set[exprValueStr]
	if !exist {
		return false,
			fmt.Sprintf("[%s]:%s should be in %v,actual is %s",
				r.name, r.getCompleteFieldExpr(), r.enum, exprValueStr)
	}
	return true, ""
}

// InStr checks if the filed is string and is in enum
func InStr(filedExpr string, enum ...string) *enumRuleString {
	set := make(map[string]struct{}, len(enum))
	for i := 0; i < len(enum); i++ {
		set[enum[i]] = struct{}{}
	}
	return &enumRuleString{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "InStr",
		},
		set,
		enum,
	}
}

type enumRuleInt struct {
	BaseRule
	set  map[int]struct{}
	enum []int
}

func (r *enumRuleInt) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r enumRuleInt) Check(param interface{}) (bool, string) {
	exprValueInt, isValid, errMsg := fetchFieldInt(param, &r)
	if !isValid {
		return false, errMsg
	}
	_, exist := r.set[exprValueInt]
	if !exist {
		return false,
			fmt.Sprintf("[%s]:%s should be in %v,actual is %d",
				r.name, r.getCompleteFieldExpr(), r.enum, exprValueInt)
	}
	return true, ""
}

// InInt checks if the filed is int and is in enum
func InInt(filedExpr string, enum ...int) *enumRuleInt {
	set := make(map[int]struct{}, len(enum))
	for i := 0; i < len(enum); i++ {
		set[enum[i]] = struct{}{}
	}
	return &enumRuleInt{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "InInt",
		},
		set,
		enum,
	}
}

type enumRuleUint struct {
	BaseRule
	set  map[uint]struct{}
	enum []uint
}

func (r *enumRuleUint) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r enumRuleUint) Check(param interface{}) (bool, string) {
	exprValueUint, isValid, errMsg := fetchFieldUint(param, &r)
	if !isValid {
		return false, errMsg
	}
	_, exist := r.set[exprValueUint]
	if !exist {
		return false,
			fmt.Sprintf("[%s]:%s should be in %v,actual is %d",
				r.name, r.getCompleteFieldExpr(), r.enum, exprValueUint)
	}
	return true, ""
}

// InUint checks if the filed is uint and is in enum
func InUint(filedExpr string, enum ...uint) *enumRuleUint {
	set := make(map[uint]struct{}, len(enum))
	for i := 0; i < len(enum); i++ {
		set[enum[i]] = struct{}{}
	}
	return &enumRuleUint{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "InUint",
		},
		set,
		enum,
	}
}

type enumRuleFloat struct {
	BaseRule
	set  map[float64]struct{}
	enum []float64
}

func (r *enumRuleFloat) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r enumRuleFloat) Check(param interface{}) (bool, string) {
	exprValueFloat, isValid, errMsg := fetchFieldFloat(param, &r)
	if !isValid {
		return false, errMsg
	}
	_, exist := r.set[exprValueFloat]
	if !exist {
		return false,
			fmt.Sprintf("[%s]:%s should be in %v,actual is %f",
				r.name, r.getCompleteFieldExpr(), r.enum, exprValueFloat)
	}
	return true, ""
}

// InFloat checks if the filed is float and is in enum
func InFloat(filedExpr string, enum ...float64) *enumRuleFloat {
	set := make(map[float64]struct{}, len(enum))
	for i := 0; i < len(enum); i++ {
		set[enum[i]] = struct{}{}
	}
	return &enumRuleFloat{
		BaseRule{
			fieldExpr: filedExpr,
			name:      "InFloat",
		},
		set,
		enum,
	}
}

package checker

import (
	"fmt"
	"reflect"
	"strings"
)

// Rule represents the limitation
// of param should obey
type Rule interface {
	Check(param interface{}) (bool, string)
}

type andRule struct {
	rules []Rule
}

func (r andRule) Check(param interface{}) (bool, string) {
	for i := 0; i < len(r.rules); i++ {
		isValid, msg := r.rules[i].Check(param)
		if !isValid {
			return isValid, msg
		}
	}
	return true, ""
}

// NewAndRule accepts slice of rules
// is passed when all rules passed
func NewAndRule(rules []Rule) Rule {
	return andRule{
		rules: rules,
	}
}

type orRule struct {
	rules []Rule
}

func (r orRule) Check(param interface{}) (bool, string) {
	messages := make([]string, 0, len(r.rules))
	for i := 0; i < len(r.rules); i++ {
		isValid, msg := r.rules[i].Check(param)
		if isValid {
			return true, ""
		}
		messages = append(messages, msg)
	}
	return false,
		fmt.Sprintf("%s, at least one ot them should be true",
			strings.Join(messages, " or "))
}

// NewOrRule accepts slice of rules
// is failed when all rules failed
func NewOrRule(rules []Rule) Rule {
	return orRule{
		rules: rules,
	}
}

type notRule struct {
	innerRule Rule
}

func (r notRule) Check(param interface{}) (bool, string) {
	isInnerValid, errMsg := r.innerRule.Check(param)
	isValid := !isInnerValid
	if !isValid {
		return false,
			fmt.Sprintf("[notRule]:{%s}", errMsg)
	}
	return true, ""
}

// NewNotRule returns the opposite if innerRule
func NewNotRule(innerRule Rule) Rule {
	return notRule{
		innerRule: innerRule,
	}
}

func fetchFieldInStruct(param interface{}, filedExpr string) (interface{}, reflect.Kind) {
	pValue := reflect.ValueOf(param)
	if filedExpr == "" {
		return param, pValue.Kind()
	}

	exprs := strings.Split(filedExpr, ".")
	for i := 0; i < len(exprs); i++ {
		expr := exprs[i]
		if pValue.Kind() == reflect.Ptr {
			pValue = pValue.Elem()
		}
		if !pValue.IsValid() || pValue.Kind() != reflect.Struct {
			return nil, reflect.Invalid
		}
		pValue = pValue.FieldByName(expr)
	}

	// last field is pointer
	if pValue.Kind() == reflect.Ptr {
		if pValue.IsNil() {
			return nil, reflect.Ptr
		}
		pValue = pValue.Elem()
	}

	if !pValue.IsValid() {
		return nil, reflect.Invalid
	}
	return pValue.Interface(), pValue.Kind()
}

// FetchFieldInStruct can be used outside tht package
func FetchFieldInStruct(param interface{}, filedExpr string) (interface{}, reflect.Kind) {
	return fetchFieldInStruct(param, filedExpr)
}

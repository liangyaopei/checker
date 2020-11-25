package checker

import (
	"fmt"
	"reflect"
	"strings"
)

type Rule interface {
	check(param interface{}) (bool, string)
}

type andRule struct {
	rules []Rule
}

func (r andRule) check(param interface{}) (bool, string) {
	for i := 0; i < len(r.rules); i++ {
		isValid, msg := r.rules[i].check(param)
		if !isValid {
			return isValid, msg
		}
	}
	return true, ""
}

func NewAndRule(rules []Rule) Rule {
	return andRule{
		rules: rules,
	}
}

type orRule struct {
	rules []Rule
}

func (r orRule) check(param interface{}) (bool, string) {
	messages := make([]string, 0, len(r.rules))
	for i := 0; i < len(r.rules); i++ {
		isValid, msg := r.rules[i].check(param)
		if isValid {
			return true, ""
		}
		messages = append(messages, msg)
	}
	return false,
		fmt.Sprintf("%s, at least one ot them should be true",
			strings.Join(messages, " or "))
}

func NewOrRule(rules []Rule) Rule {
	return orRule{
		rules: rules,
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
	return pValue.Interface(), pValue.Kind()
}

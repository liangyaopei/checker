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
		} else {
			messages = append(messages, msg)
		}
	}
	return false,
		fmt.Sprintf("%s, at least one ot them should be true", strings.Join(messages, " or "))
}

func fetchFieldInStruct(param interface{}, filedExpr string) (interface{}, reflect.Kind) {
	pValue := reflect.ValueOf(param)
	if filedExpr == "" {
		return param, pValue.Kind()
	}
	if pValue.Kind() != reflect.Struct {
		return nil, reflect.Invalid
	}
	exprs := strings.Split(filedExpr, ",")
	for i := 0; i < len(exprs); i++ {
		expr := exprs[i]
		exprValue := pValue.FieldByName(expr)
		if !exprValue.IsValid() {
			return nil, reflect.Invalid
		}
		pValue = exprValue
	}
	return pValue.Interface(), pValue.Kind()
}

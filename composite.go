package checker

import (
	"fmt"
	"reflect"
	"strings"
)

type fieldRule struct {
	BaseRule
	subRule Rule
}

func (r *fieldRule) Prompt(prompt string) {
	r.prompt = prompt
}

func (r *fieldRule) Check(param interface{}) (bool, string) {
	exprValue, kind := fetchField(param, r.fieldExpr)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.getAbsoluteFieldExpr())
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is nil", r.name, r.getAbsoluteFieldExpr())
	}

	r.subRule.setUpperFieldExpr(r.getAbsoluteFieldExpr())

	isValid, errLog := r.subRule.Check(exprValue)
	if !isValid {
		if subPrompt := r.subRule.getPrompt(); subPrompt != "" {
			r.Prompt(subPrompt)
		}
	}
	return isValid, errLog
}

// Field applies rule to fieldExpr
func Field(fieldExpr string, subRule Rule) *fieldRule {
	return &fieldRule{
		BaseRule{
			fieldExpr: fieldExpr,
			name:      "Field",
		},
		subRule,
	}
}

type andRule struct {
	BaseRule
	rules []Rule
}

func (r *andRule) Prompt(prompt string) {
	r.prompt = prompt
}

func (r *andRule) Check(param interface{}) (bool, string) {
	for i := 0; i < len(r.rules); i++ {
		r.rules[i].setUpperFieldExpr(r.getAbsoluteFieldExpr())

		isValid, msg := r.rules[i].Check(param)
		if !isValid {
			if subPrompt := r.rules[i].getPrompt(); subPrompt != "" {
				r.rules[i].Prompt(subPrompt)
			}
			return isValid, msg
		}
	}
	return true, ""
}

// And accepts slice of rules
// is passed when all rules passed
func And(rules ...Rule) *andRule {
	return &andRule{
		BaseRule{
			name: "And",
		},
		rules,
	}
}

type orRule struct {
	BaseRule
	rules []Rule
}

func (r *orRule) Prompt(prompt string) {
	r.prompt = prompt
}

func (r *orRule) Check(param interface{}) (bool, string) {
	messages := make([]string, 0, len(r.rules))
	for i := 0; i < len(r.rules); i++ {
		r.rules[i].setUpperFieldExpr(r.getAbsoluteFieldExpr())

		isValid, msg := r.rules[i].Check(param)
		if isValid {
			if subPrompt := r.rules[i].getPrompt(); subPrompt != "" {
				r.rules[i].Prompt(subPrompt)
			}
			return true, ""
		}
		messages = append(messages, msg)
	}
	return false,
		fmt.Sprintf("%s, at least one ot them should be true",
			strings.Join(messages, " or "))
}

// Or accepts slice of rules
// is failed when all rules failed
func Or(rules ...Rule) *orRule {
	return &orRule{
		BaseRule{
			name: "Or",
		},
		rules,
	}
}

type notRule struct {
	BaseRule
	innerRule Rule
}

func (r *notRule) Prompt(prompt string) {
	r.prompt = prompt
}

func (r *notRule) Check(param interface{}) (bool, string) {
	r.innerRule.setUpperFieldExpr(r.getAbsoluteFieldExpr())

	isInnerValid, errMsg := r.innerRule.Check(param)
	isValid := !isInnerValid
	if !isValid {
		if subPrompt := r.innerRule.getPrompt(); subPrompt != "" {
			r.Prompt(subPrompt)
		}
		return false,
			fmt.Sprintf("[notRule]:{%s}", errMsg)
	}
	return true, ""
}

// Not returns the opposite if innerRule
func Not(innerRule Rule) *notRule {
	return &notRule{
		BaseRule{
			name: "Not",
		},
		innerRule,
	}
}

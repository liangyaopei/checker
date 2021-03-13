package checker

import (
	"fmt"
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

// And accepts slice of rules
// is passed when all rules passed
func And(rules ...Rule) Rule {
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

// Or accepts slice of rules
// is failed when all rules failed
func Or(rules ...Rule) Rule {
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

// Not returns the opposite if innerRule
func Not(innerRule Rule) Rule {
	return notRule{
		innerRule: innerRule,
	}
}

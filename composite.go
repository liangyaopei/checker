package checker

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type fieldRule struct {
	baseRule
	subRule Rule
}

func (r *fieldRule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r fieldRule) Check(param interface{}) (bool, string) {
	exprValue, kind := fetchField(param, &r)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:%s cannot be found", r.name, r.getCompleteFieldExpr())
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:%s is nil", r.name, r.getCompleteFieldExpr())
	}

	r.subRule.setUpperFieldExpr(r.getCompleteFieldExpr())
	r.subRule.setCache(r.getCache())

	isValid, errLog := r.subRule.Check(exprValue)
	if !isValid {
		if prompt := r.getPrompt(); prompt != "" {
			r.Prompt(prompt)
		}
	}
	return isValid, errLog
}

// Field applies rule to fieldExpr
func Field(fieldExpr string, subRule Rule) *fieldRule {
	return &fieldRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "Field",
			lowerRule: subRule,
		},
		subRule,
	}
}

type lengthRule struct {
	baseRule
	le int
	ge int
}

func (r *lengthRule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r lengthRule) Check(param interface{}) (bool, string) {
	exprValue, kind := fetchField(param, &r)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:%s cannot be found", r.name, r.getCompleteFieldExpr())
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:%s is nil", r.name, r.getCompleteFieldExpr())
	}

	if kind != reflect.Slice && kind != reflect.Array &&
		kind != reflect.String && kind != reflect.Map {
		return false,
			fmt.Sprintf("[%s]:%s should be slice/array/string/map,actual is %s",
				r.name, r.getCompleteFieldExpr(), kind.String())
	}

	lenValue := reflect.ValueOf(exprValue)
	length := lenValue.Len()
	if length < r.ge || length > r.le {
		return false,
			fmt.Sprintf("[%s]:%s length should be between %d and %d,actual is %d",
				r.name, r.getCompleteFieldExpr(), r.le, r.ge, length)

	}
	return true, ""
}

// Length checks if the filed is slice/array/string/map
// and its length is between [ge,le]
func Length(fieldExpr string, ge int, le int) *lengthRule {
	return &lengthRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "Length",
		},
		le,
		ge,
	}
}

type arrayRule struct {
	baseRule
	innerRule Rule
}

func (r *arrayRule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r *arrayRule) Check(param interface{}) (bool, string) {
	exprValue, kind := fetchField(param, r)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:%s cannot be found", r.name, r.getCompleteFieldExpr())
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:%s is nil", r.name, r.getCompleteFieldExpr())
	}

	if kind != reflect.Slice && kind != reflect.Array {
		return false,
			fmt.Sprintf("[%s]:%s should be slice/array,actual is %s",
				r.name, r.getCompleteFieldExpr(), kind.String())
	}

	arrayValue := reflect.ValueOf(exprValue)
	for i := 0; i < arrayValue.Len(); i++ {
		r.innerRule.setUpperFieldExpr(r.getCompleteFieldExpr() + "[" + strconv.Itoa(i) + "]")

		idxValue := arrayValue.Index(i)
		isValid, msg := r.innerRule.Check(idxValue.Interface())
		if !isValid {
			return false, msg
		}
	}
	return true, ""
}

// Array checks if the filed is slice/array
// and its elements satisfy the innerRule
func Array(fieldExpr string, innerRule Rule) *arrayRule {
	return &arrayRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "Array",
			lowerRule: innerRule,
		},
		innerRule,
	}
}

// mapRule checks map's key and value
// keyRule or valueRule can be nil
type mapRule struct {
	baseRule
	keyRule   Rule
	valueRule Rule
}

func (r *mapRule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r *mapRule) Check(param interface{}) (bool, string) {
	exprValue, kind := fetchField(param, r)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:%s cannot be found", r.name, r.getCompleteFieldExpr())
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:%s is nil", r.name, r.getCompleteFieldExpr())
	}
	if kind != reflect.Map {
		return false,
			fmt.Sprintf("[%s]:%s should be kind map,actual is %v",
				r.name, r.getCompleteFieldExpr(), kind)
	}

	refVal := reflect.ValueOf(exprValue)

	for _, key := range refVal.MapKeys() {

		keyStr := getKeyStr(key)
		// check key
		if r.keyRule != nil {
			r.keyRule.setUpperFieldExpr(r.getCompleteFieldExpr() + "[" + keyStr + "]")
			r.keyRule.setCache(r.getCache())
			r.lowerRule = r.keyRule

			isValid, errMsg := r.keyRule.Check(key.Interface())
			if !isValid {
				if prompt := r.getPrompt(); prompt != "" {
					r.Prompt(prompt)
				}
				return false, errMsg
			}
		}

		// check value
		if r.valueRule != nil {
			value := refVal.MapIndex(key)

			r.valueRule.setUpperFieldExpr(r.getCompleteFieldExpr() + "[" + keyStr + "] 's value")
			r.valueRule.setCache(r.getCache())
			r.lowerRule = r.valueRule

			isValid, errMsg := r.valueRule.Check(value.Interface())
			if !isValid {
				if prompt := r.getPrompt(); prompt != "" {
					r.Prompt(prompt)
				}
				return false, errMsg
			}
		}
	}

	return true, ""
}

func getKeyStr(value reflect.Value) string {
	res := ""
	switch value.Kind() {
	case reflect.Bool:
		res = strconv.FormatBool(value.Bool())
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int,
		reflect.Int64:
		res = strconv.FormatInt(value.Int(), 10)
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint,
		reflect.Uint64:
		res = strconv.FormatUint(value.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		res = strconv.FormatFloat(value.Float(), 'f', -1, 64)
	default:
		res = value.String()
	}
	return res
}

// Map returns mapRule
// keyRule or valueRule can be nil
// if no need to check key/value
func Map(fieldExpr string, keyRule Rule, valueRule Rule) *mapRule {
	return &mapRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "Map",
		},
		keyRule,
		valueRule,
	}
}

type andRule struct {
	baseRule
	rules []Rule
}

func (r *andRule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r *andRule) Check(param interface{}) (bool, string) {
	for i := 0; i < len(r.rules); i++ {

		r.rules[i].setUpperFieldExpr(r.getCompleteFieldExpr())
		r.lowerRule = r.rules[i]
		r.rules[i].setCache(r.getCache())

		isValid, msg := r.rules[i].Check(param)
		if !isValid {
			if prompt := r.getPrompt(); prompt != "" {
				r.Prompt(prompt)
			}
			return isValid, msg
		}
	}
	return true, ""
}

// And accepts slice of rules
// It passed when all rules passed
func And(rules ...Rule) *andRule {
	return &andRule{
		baseRule{
			name: "And",
		},
		rules,
	}
}

type orRule struct {
	baseRule
	rules []Rule
}

func (r *orRule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r *orRule) Check(param interface{}) (bool, string) {
	messages := make([]string, 0, len(r.rules))
	prompts := make([]string, 0, len(r.rules))

	for i := 0; i < len(r.rules); i++ {
		r.rules[i].setUpperFieldExpr(r.getCompleteFieldExpr())
		r.lowerRule = r.rules[i]
		r.rules[i].setCache(r.getCache())

		isValid, msg := r.rules[i].Check(param)
		if isValid {
			return true, ""
		}
		if prompt := r.getPrompt(); prompt != "" {
			prompts = append(prompts, prompt)
		}
		messages = append(messages, msg)
	}
	prompt := strings.Join(prompts, " or ")
	r.Prompt(prompt)
	return false,
		fmt.Sprintf("%s, at least one ot them should be true",
			strings.Join(messages, " or "))
}

// Or accepts slice of rules
// It failed when all rules failed
func Or(rules ...Rule) *orRule {
	return &orRule{
		baseRule{
			name: "Or",
		},
		rules,
	}
}

type notRule struct {
	baseRule
	innerRule Rule
}

func (r *notRule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r *notRule) Check(param interface{}) (bool, string) {
	r.innerRule.setUpperFieldExpr(r.getCompleteFieldExpr())

	isInnerValid, errMsg := r.innerRule.Check(param)
	isValid := !isInnerValid
	if !isValid {
		if prompt := r.getPrompt(); prompt != "" {
			r.Prompt(prompt)
		}
		return false,
			fmt.Sprintf("[%s]:[%s] should not be true", r.name, errMsg)
	}
	return true, ""
}

// Not returns the opposite if innerRule
func Not(innerRule Rule) *notRule {
	return &notRule{
		baseRule{
			name:      "Not",
			lowerRule: innerRule,
		},
		innerRule,
	}
}

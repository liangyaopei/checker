package checker

import "reflect"

// Rule represents the restriction of param should satisfy
type Rule interface {
	Check(param interface{}) (bool, string)

	getPrompt() string
	Prompt(prompt string) Rule

	getName() string

	getFieldExpr() string
	getCompleteFieldExpr() string
	setUpperFieldExpr(expr string)

	setCache(m map[string]valueKindPair)
	getCache() map[string]valueKindPair
}

type valueKindPair struct {
	value interface{}
	kind  reflect.Kind
}

type BaseRule struct {
	fieldExpr      string
	upperFieldExpr string

	name   string
	prompt string

	lowerRule  Rule
	fieldCache map[string]valueKindPair
}

func (r BaseRule) getName() string {
	return r.name
}

func (r *BaseRule) SetPrompt(prompt string) {
	r.prompt = prompt
}

func (r BaseRule) getPrompt() string {
	lowerRulePrompt := ""
	if r.lowerRule != nil {
		lowerRulePrompt = r.lowerRule.getPrompt()
	}
	if lowerRulePrompt != "" {
		return lowerRulePrompt
	}
	return r.prompt
}

func (r BaseRule) getFieldExpr() string {
	return r.fieldExpr
}

func (r BaseRule) getCompleteFieldExpr() string {
	if r.upperFieldExpr == "" {
		return r.fieldExpr
	}
	if r.fieldExpr == "" {
		return r.upperFieldExpr
	}
	return r.upperFieldExpr + "." + r.fieldExpr
}

func (r *BaseRule) setUpperFieldExpr(expr string) {
	r.upperFieldExpr = expr
}

func (r *BaseRule) setCache(m map[string]valueKindPair) {
	r.fieldCache = m
}

func (r BaseRule) getCache() map[string]valueKindPair {
	return r.fieldCache
}

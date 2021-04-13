package checker

// Rule represents the restriction
// of param should obey
type Rule interface {
	Check(param interface{}) (bool, string)
	Prompt(prompt string)

	getPrompt() string

	setUpperFieldExpr(expr string)
	getAbsoluteFieldExpr() string

	setLowerRule(r Rule)
}

type BaseRule struct {
	name           string
	fieldExpr      string
	upperFieldExpr string

	prompt    string
	lowerRule Rule
}

func (r *BaseRule) getPrompt() string {
	return r.prompt
}

func (r *BaseRule) setUpperFieldExpr(upperFieldExpr string) {
	r.upperFieldExpr = upperFieldExpr
}

func (r BaseRule) getAbsoluteFieldExpr() string {
	if r.fieldExpr == "" {
		return r.upperFieldExpr
	}

	if r.upperFieldExpr == "" {
		return r.fieldExpr
	}
	return r.upperFieldExpr + "." + r.fieldExpr
}

func (r *BaseRule) setLowerRule(lowerRule Rule) {
	r.lowerRule = lowerRule
}

package checker

import (
	"fmt"
	"reflect"
	"time"
)

// Comparable is the interface
// for sturcts to be compared in
// CrossComp
type Comparable interface {
	EqualTo(other interface{}) bool
	LessThan(other interface{}) bool
}

type operand interface {
	Symbol() string
}

type eqOperand struct{}

func (op eqOperand) Symbol() string {
	return "="
}

type neOperand struct{}

func (op neOperand) Symbol() string {
	return "!="
}

type gtOperand struct{}

func (op gtOperand) Symbol() string {
	return ">"
}

type geOperand struct{}

func (op geOperand) Symbol() string {
	return ">="
}

type ltOperand struct{}

func (op ltOperand) Symbol() string {
	return "<"
}

type leOperand struct{}

func (op leOperand) Symbol() string {
	return "<="
}

// cross field comparison operand variables
var (
	CrossFieldEq = eqOperand{}
	CrossFieldNe = neOperand{}
	CrossFieldGt = gtOperand{}
	CrossFieldGe = geOperand{}
	CrossFieldLt = ltOperand{}
	CrossFieldLe = leOperand{}
)

type ruleWrapper struct {
	BaseRule
}

func (r ruleWrapper) Check(param interface{}) (bool, string) {
	panic("implement me")
}

func (r *ruleWrapper) Prompt(prompt string) Rule {
	panic("implement me")
}

type crossFieldCompareRule struct {
	op operand

	ruleLeft  ruleWrapper
	ruleRight ruleWrapper

	name           string
	prompt         string
	upperFieldExpr string
}

func (r crossFieldCompareRule) getPrompt() string {
	return r.prompt
}

func (r *crossFieldCompareRule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r crossFieldCompareRule) getName() string {
	return r.name
}

func (r crossFieldCompareRule) getFieldExpr() string {
	panic("crossFieldCompareRule.getFieldExpr should not be called")
}

func (r crossFieldCompareRule) getCompleteFieldExpr() string {
	panic("crossFieldCompareRule.getCompleteFieldExpr should not be called")
}

func (r crossFieldCompareRule) setUpperFieldExpr(expr string) {
	r.ruleLeft.setUpperFieldExpr(expr)
	r.ruleRight.setUpperFieldExpr(expr)
}

func (r crossFieldCompareRule) setCache(m map[string]valueKindPair) {
	r.ruleLeft.setCache(m)
	r.ruleRight.setCache(m)
}

func (r crossFieldCompareRule) getCache() map[string]valueKindPair {
	return r.ruleLeft.getCache()
}

func (r crossFieldCompareRule) Check(param interface{}) (bool, string) {
	exprValueLeft, kind := fetchField(param, &r.ruleLeft)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:%s cannot be found", r.name, r.ruleLeft.getCompleteFieldExpr())
	}
	if exprValueLeft == nil {
		return false,
			fmt.Sprintf("[%s]:%s is nil", r.name, r.ruleLeft.getCompleteFieldExpr())
	}

	exprValueRight, kind := fetchField(param, &r.ruleRight)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:%s cannot be found", r.name, r.ruleRight.getCompleteFieldExpr())
	}
	if exprValueRight == nil {
		return false,
			fmt.Sprintf("[%s]:%s is nil", r.name, r.ruleRight.getCompleteFieldExpr())
	}

	refValLeft := reflect.ValueOf(exprValueLeft)
	refValRight := reflect.ValueOf(exprValueRight)

	if refValLeft.Type() != refValRight.Type() {
		return false,
			fmt.Sprintf("[%s]:type incompatibility,%s is %s, %s is %s", r.name,
				r.ruleLeft.getCompleteFieldExpr(), refValLeft.Type().String(),
				r.ruleRight.getCompleteFieldExpr(), refValRight.Type().String())
	}
	isValid := false
	switch exprValueLeft.(type) {
	case int:
		isValid = compareInt(refValLeft.Int(), refValRight.Int(), r.op)
	case uint:
		isValid = compareUInt(refValLeft.Uint(), refValRight.Uint(), r.op)
	case float64:
		isValid = compareFloat64(refValLeft.Float(), refValRight.Float(), r.op)
	case string:
		isValid = compareString(refValLeft.String(), refValRight.String(), r.op)
	case time.Time:
		timeLeft, _ := exprValueLeft.(time.Time)
		timeRight, _ := exprValueRight.(time.Time)
		isValid = compareTime(timeLeft, timeRight, r.op)
	case Comparable:
		comparableLeft := exprValueLeft.(Comparable)
		comparableRight := exprValueRight.(Comparable)
		isValid = compareComparable(comparableLeft, comparableRight, r.op)
	default:
		isValid = false
	}
	if !isValid {
		return false,
			fmt.Sprintf("[%s]:%s does not %s %s",
				r.name, r.ruleLeft.getCompleteFieldExpr(), r.op.Symbol(), r.ruleRight.getCompleteFieldExpr())
	}
	return true, ""
}

// CrossComp checks if left and right fields
// are same types and satisfy the comparison operand
func CrossComp(fieldExprLeft string, fieldExprRight string, op operand) *crossFieldCompareRule {
	ruleLeft := ruleWrapper{
		BaseRule{
			fieldExpr: fieldExprLeft,
			name:      "CrossComp",
		},
	}
	ruleRight := ruleWrapper{
		BaseRule{
			fieldExpr: fieldExprRight,
			name:      "CrossComp",
		},
	}
	return &crossFieldCompareRule{
		ruleLeft:  ruleLeft,
		ruleRight: ruleRight,
		op:        op,
		name:      "CrossComp",
	}
}

func compareInt(left int64, right int64, op operand) bool {
	switch op.(type) {
	case eqOperand:
		return left == right
	case neOperand:
		return left != right
	case gtOperand:
		return left > right
	case geOperand:
		return left >= right
	case ltOperand:
		return left < right
	case leOperand:
		return left <= right
	default:
		return false
	}
}

func compareUInt(left uint64, right uint64, op operand) bool {
	switch op.(type) {
	case eqOperand:
		return left == right
	case neOperand:
		return left != right
	case gtOperand:
		return left > right
	case geOperand:
		return left >= right
	case ltOperand:
		return left < right
	case leOperand:
		return left <= right
	default:
		return false
	}
}

func compareFloat64(left float64, right float64, op operand) bool {
	switch op.(type) {
	case eqOperand:
		return left == right
	case neOperand:
		return left != right
	case gtOperand:
		return left > right
	case geOperand:
		return left >= right
	case ltOperand:
		return left < right
	case leOperand:
		return left <= right
	default:
		return false
	}
}

func compareString(left string, right string, op operand) bool {
	switch op.(type) {
	case eqOperand:
		return left == right
	case neOperand:
		return left != right
	case gtOperand:
		return left > right
	case geOperand:
		return left >= right
	case ltOperand:
		return left < right
	case leOperand:
		return left <= right
	default:
		return false
	}
}

func compareTime(left time.Time, right time.Time, op operand) bool {
	switch op.(type) {
	case eqOperand:
		return left.Equal(right)
	case neOperand:
		return !left.Equal(right)
	case gtOperand:
		return left.After(right)
	case geOperand:
		return left.After(right) || left.Equal(right)
	case ltOperand:
		return left.Before(right)
	case leOperand:
		return left.Before(right) || left.Equal(right)
	default:
		return false
	}
}

func compareComparable(left Comparable, right Comparable, op operand) bool {
	switch op.(type) {
	case eqOperand:
		return left.EqualTo(right)
	case neOperand:
		return !left.EqualTo(right)
	case gtOperand:
		return !left.LessThan(right)
	case geOperand:
		return !left.LessThan(right) || left.EqualTo(right)
	case ltOperand:
		return left.LessThan(right)
	case leOperand:
		return left.LessThan(right) || left.EqualTo(right)
	default:
		return false
	}
}

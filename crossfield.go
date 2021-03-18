package checker

import (
	"fmt"
	"reflect"
	"time"
)

// Comparable is the interface
// for sturcts to be compared in
// CrossComparable
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

type crossFieldCompareRule struct {
	fieldExprLeft  string
	fieldExprRight string
	op             operand

	name string
}

func (r crossFieldCompareRule) Check(param interface{}) (bool, string) {
	exprValueLeft, kind := fetchField(param, r.fieldExprLeft)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.fieldExprLeft)
	}
	if exprValueLeft == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is nil", r.name, r.fieldExprLeft)
	}

	exprValueRight, kind := fetchField(param, r.fieldExprRight)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.fieldExprRight)
	}
	if exprValueRight == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is nil", r.name, r.fieldExprRight)
	}

	refValLeft := reflect.ValueOf(exprValueLeft)
	refValRight := reflect.ValueOf(exprValueRight)

	if refValLeft.Type() != refValRight.Type() {
		return false,
			fmt.Sprintf("[%s]:type incompatibility,'%s' is %s,'%s' is %s", r.name,
				r.fieldExprLeft, refValLeft.Type().String(),
				r.fieldExprRight, refValRight.Type().String())
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
			fmt.Sprintf("[%s]:'%s' does not %s '%s'",
				r.name, r.fieldExprLeft, r.op.Symbol(), r.fieldExprRight)
	}
	return true, ""
}

// CrossComparable checks if left and right fields
// are same types and satisfy the comparison operand
func CrossComparable(fieldExprLeft string, fieldExprRight string, op operand) Rule {
	return crossFieldCompareRule{
		fieldExprLeft:  fieldExprLeft,
		fieldExprRight: fieldExprRight,
		op:             op,
		name:           "crossFieldCompareRule",
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

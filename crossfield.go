package checker

import (
	"fmt"
	"reflect"
	"time"
)

type crossFiledOp int

const (
	CrossFiledEq = iota + 1
	CrossFiledNe
	CrossFiledGt
	CrossFiledGe
	CrossFiledLt
	CrossFiledLe
)

var (
	opStrMap = map[crossFiledOp]string{
		CrossFiledEq: "=",
		CrossFiledNe: "!=",
		CrossFiledGt: ">",
		CrossFiledGe: ">=",
		CrossFiledLt: "<",
		CrossFiledLe: "<=",
	}
)

type crossFieldCompareRule struct {
	fieldExprLeft  string
	fieldExprRight string
	op             crossFiledOp

	name string
}

func (r crossFieldCompareRule) check(param interface{}) (bool, string) {
	exprValueLeft, kind := fetchFieldInStruct(param, r.fieldExprLeft)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.fieldExprLeft)
	}
	if exprValueLeft == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is nil", r.name, r.fieldExprLeft)
	}

	exprValueRight, kind := fetchFieldInStruct(param, r.fieldExprRight)
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
	default:
		isValid = false
	}
	if !isValid {
		return false,
			fmt.Sprintf("[%s]:'%s' does not %s '%s'",
				r.name, r.fieldExprLeft, opStrMap[r.op], r.fieldExprRight)
	}
	return true, ""
}

func NewCrossFieldCompareRule(fieldExprLeft string, fieldExprRight string, op crossFiledOp) Rule {
	return crossFieldCompareRule{
		fieldExprLeft:  fieldExprLeft,
		fieldExprRight: fieldExprRight,
		op:             op,
		name:           "crossFieldCompareRule",
	}
}

func compareInt(left int64, right int64, op crossFiledOp) bool {
	switch op {
	case CrossFiledEq:
		return left == right
	case CrossFiledNe:
		return left != right
	case CrossFiledGt:
		return left > right
	case CrossFiledGe:
		return left >= right
	case CrossFiledLt:
		return left < right
	case CrossFiledLe:
		return left <= right
	default:
		return false
	}
}

func compareUInt(left uint64, right uint64, op crossFiledOp) bool {
	switch op {
	case CrossFiledEq:
		return left == right
	case CrossFiledNe:
		return left != right
	case CrossFiledGt:
		return left > right
	case CrossFiledGe:
		return left >= right
	case CrossFiledLt:
		return left < right
	case CrossFiledLe:
		return left <= right
	default:
		return false
	}
}

func compareFloat64(left float64, right float64, op crossFiledOp) bool {
	switch op {
	case CrossFiledEq:
		return left == right
	case CrossFiledNe:
		return left != right
	case CrossFiledGt:
		return left > right
	case CrossFiledGe:
		return left >= right
	case CrossFiledLt:
		return left < right
	case CrossFiledLe:
		return left <= right
	default:
		return false
	}
}

func compareString(left string, right string, op crossFiledOp) bool {
	switch op {
	case CrossFiledEq:
		return left == right
	case CrossFiledNe:
		return left != right
	case CrossFiledGt:
		return left > right
	case CrossFiledGe:
		return left >= right
	case CrossFiledLt:
		return left < right
	case CrossFiledLe:
		return left <= right
	default:
		return false
	}
}

func compareTime(left time.Time, right time.Time, op crossFiledOp) bool {
	switch op {
	case CrossFiledEq:
		return left.Equal(right)
	case CrossFiledNe:
		return !left.Equal(right)
	case CrossFiledGt:
		return left.After(right)
	case CrossFiledGe:
		return left.After(right) || left.Equal(right)
	case CrossFiledLt:
		return left.Before(right)
	case CrossFiledLe:
		return left.Before(right) || left.Equal(right)
	default:
		return false
	}
}

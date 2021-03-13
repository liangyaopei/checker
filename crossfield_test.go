package checker

import (
	"testing"
	"time"
)

type crossStruct struct {
	Int1 int
	Int2 int

	Uint1 uint
	Uint2 uint

	Float1 float64
	Float2 float64

	Date1 time.Time
	Date2 time.Time
}

func TestCrossFieldSimple(t *testing.T) {
	layout := "2006-01-02"
	date1, _ := time.Parse(layout, "2020-12-12")
	date2, _ := time.Parse(layout, "2021-12-12")

	crossChecker := NewChecker()

	crossRuleInt := CrossComparable("Int1", "Int2", CrossFieldNe)
	crossChecker.Add(crossRuleInt, "invalid int1 and int2")

	crossRuleUInt := CrossComparable("Uint1", "Uint2", CrossFieldGt)
	crossChecker.Add(crossRuleUInt, "invalid uint1 and uint2")

	crossRuleFloat := CrossComparable("Float1", "Float2", CrossFieldLt)
	crossChecker.Add(crossRuleFloat, "invalid Float1 and Float2")

	crossRuleTime := CrossComparable("Date1", "Date2", CrossFieldLe)
	crossChecker.Add(crossRuleTime, "invalid Date1 and Date2")

	cross := crossStruct{
		Int1:   1,
		Int2:   2,
		Uint1:  10,
		Uint2:  9,
		Float1: 3.14,
		Float2: 6.28,
		Date1:  date1,
		Date2:  date2,
	}

	isValid, prompt, errMsg := crossChecker.Check(cross)
	if !isValid {
		t.Errorf("errMsg:%s,prompt:%s", errMsg, prompt)
		return
	}
	t.Logf("valid cross struct")
}

type innerInt struct {
	Val int
}

func (i innerInt) EqualTo(other interface{}) bool {
	otherInt, ok := other.(innerInt)
	if !ok {
		return false
	}
	return i.Val == otherInt.Val
}

func (i innerInt) LessThan(other interface{}) bool {
	otherInt, ok := other.(innerInt)
	if !ok {
		return false
	}
	return i.Val < otherInt.Val
}

type crossInner struct {
	Int1 innerInt
	Int2 innerInt
}

func TestCrossFieldCustom(t *testing.T) {

	crossChecker := NewChecker()

	crossRuleInt := CrossComparable("Int1", "Int2", CrossFieldGe)
	crossChecker.Add(crossRuleInt, "invalid int1 and int2")

	cross := crossInner{
		Int1: innerInt{
			Val: 10,
		},
		Int2: innerInt{
			Val: 2,
		},
	}

	isValid, prompt, errMsg := crossChecker.Check(cross)
	if !isValid {
		t.Errorf("errMsg:%s,prompt:%s", errMsg, prompt)
		return
	}
	t.Logf("valid cross struct")
}

package _checker

import (
	"testing"
	"time"

	"github.com/liangyaopei/checker"
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

func TestCrossField(t *testing.T) {
	layout := "2006-01-02"
	date1, _ := time.Parse(layout, "2020-12-12")
	date2, _ := time.Parse(layout, "2021-12-12")

	crossChecker := checker.NewChecker()

	crossRuleInt := checker.NewCrossFieldCompareRule("Int1", "Int2", checker.CrossFiledNe)
	crossChecker.Add(crossRuleInt, "invalid int1 and int2")

	crossRuleUInt := checker.NewCrossFieldCompareRule("Uint1", "Uint2", checker.CrossFiledGt)
	crossChecker.Add(crossRuleUInt, "invalid uint1 and uint2")

	crossRuleFloat := checker.NewCrossFieldCompareRule("Float1", "Float2", checker.CrossFiledLt)
	crossChecker.Add(crossRuleFloat, "invalid Float1 and Float2")

	crossRuleTime := checker.NewCrossFieldCompareRule("Date1", "Date2", checker.CrossFiledLt)
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

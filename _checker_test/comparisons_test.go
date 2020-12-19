package _checker

import (
	"testing"
	"time"

	"github.com/liangyaopei/checker"
)

type comparison struct {
	Int     int
	Uint    uint
	Float   float64
	String  string
	TimeStr string
	Time    time.Time
}

func TestComparison(t *testing.T) {
	cChecker := checker.NewChecker()

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, "2020-12-12")
	endDate, _ := time.Parse(layout, "2020-12-13")

	// equal rules
	intEqRule := checker.NewEqRuleInt("Int", 10)
	cChecker.Add(intEqRule, "invalid Int")

	uintEqRule := checker.NewEqRuleUint("Uint", 58)
	cChecker.Add(uintEqRule, "invalid Uint")

	floatRule := checker.NewEqRuleFloat("Float", 3.14)
	cChecker.Add(floatRule, "invalid Float")

	strRule := checker.NewEqRuleString("String", "string")
	cChecker.Add(strRule, "invalid String")

	timeStrRule := checker.NewEqRuleTimestampStr("TimeStr", layout, startDate)
	cChecker.Add(timeStrRule, "invalid TimeStr")

	timeRule := checker.NewEqRuleTimestamp("Time", startDate)
	cChecker.Add(timeRule, "invalid Time")

	// not equal rules
	intNeRule := checker.NewNeRuleInt("Int", 20)
	cChecker.Add(intNeRule, "invalid Int")

	uintNeRule := checker.NewNeRuleUint("Uint", 158)
	cChecker.Add(uintNeRule, "invalid Uint")

	floatNeRule := checker.NewNeRuleFloat("Float", 6.28)
	cChecker.Add(floatNeRule, "invalid Float")

	strNeRule := checker.NewNeRuleString("String", "string12")
	cChecker.Add(strNeRule, "invalid String")

	timeNeStrRule := checker.NewNeRuleTimestampStr("TimeStr", layout, endDate)
	cChecker.Add(timeNeStrRule, "invalid TimeStr")

	timeNeRule := checker.NewNeRuleTimestamp("Time", endDate)
	cChecker.Add(timeNeRule, "invalid Time")

	// range rules
	intRangeRule := checker.NewRangeRuleInt("Int", 1, 20)
	cChecker.Add(intRangeRule, "invalid Int")

	uintRangeRule := checker.NewRangeRuleUint("Uint", 10, 258)
	cChecker.Add(uintRangeRule, "invalid Uint")

	floatRangeRule := checker.NewRangeRuleFloat("Float", 3.14, 6.28)
	cChecker.Add(floatRangeRule, "invalid Float")

	timeRangeStrRule := checker.NewRangeRuleTimestampStr("TimeStr", layout, startDate, endDate)
	cChecker.Add(timeRangeStrRule, "invalid TimeStr")

	timeRangeRule := checker.NewRangeRuleTimestamp("Time", startDate, endDate)
	cChecker.Add(timeRangeRule, "invalid Time")

	comp := comparison{
		Int:     10,
		Uint:    58,
		Float:   3.14,
		String:  "string",
		TimeStr: "2020-12-12",
		Time:    startDate,
	}

	isValid, prompt, errMsg := cChecker.Check(comp)
	if !isValid {
		t.Errorf("errMsg:%s,prompt:%s", errMsg, prompt)
		return
	}
	t.Logf("valid comparsion")
}

type comp2 struct {
	InnerInt *innerInt
}

func TestComparisonComparable(t *testing.T) {
	cChecker := checker.NewChecker()

	equivalent := innerInt{
		Val: 100,
	}

	inequivalent := innerInt{
		Val: 120,
	}

	ge := innerInt{
		Val: 80,
	}
	le := innerInt{
		Val: 180,
	}

	eqRule := checker.NewEqRuleComp("InnerInt", equivalent)
	cChecker.Add(eqRule, "invalid InnerInt[eq]")

	neRule := checker.NewNeRuleComp("InnerInt", inequivalent)
	cChecker.Add(neRule, "invalid InnerInt[ne]")

	rangeRule := checker.NewRangeRuleComp("InnerInt", ge, le)
	cChecker.Add(rangeRule, "invalid InnerInt[range]")

	param := comp2{
		InnerInt: &innerInt{Val: 100},
		//InnerInt: nil,
	}

	isValid, prompt, errMsg := cChecker.Check(param)
	if !isValid {
		t.Errorf("errMsg:%s,prompt:%s", errMsg, prompt)
		return
	}
	t.Logf("valid comparable comparsion")
}

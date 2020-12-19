package checker

import (
	"testing"
	"time"
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
	cChecker := NewChecker()

	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, "2020-12-12")
	endDate, _ := time.Parse(layout, "2020-12-13")

	// equal rules
	intEqRule := NewEqRuleInt("Int", 10)
	cChecker.Add(intEqRule, "invalid Int")

	uintEqRule := NewEqRuleUint("Uint", 58)
	cChecker.Add(uintEqRule, "invalid Uint")

	floatRule := NewEqRuleFloat("Float", 3.14)
	cChecker.Add(floatRule, "invalid Float")

	strRule := NewEqRuleString("String", "string")
	cChecker.Add(strRule, "invalid String")

	timeStrRule := NewEqRuleTimestampStr("TimeStr", layout, startDate)
	cChecker.Add(timeStrRule, "invalid TimeStr")

	timeRule := NewEqRuleTimestamp("Time", startDate)
	cChecker.Add(timeRule, "invalid Time")

	// not equal rules
	intNeRule := NewNeRuleInt("Int", 20)
	cChecker.Add(intNeRule, "invalid Int")

	uintNeRule := NewNeRuleUint("Uint", 158)
	cChecker.Add(uintNeRule, "invalid Uint")

	floatNeRule := NewNeRuleFloat("Float", 6.28)
	cChecker.Add(floatNeRule, "invalid Float")

	strNeRule := NewNeRuleString("String", "string12")
	cChecker.Add(strNeRule, "invalid String")

	timeNeStrRule := NewNeRuleTimestampStr("TimeStr", layout, endDate)
	cChecker.Add(timeNeStrRule, "invalid TimeStr")

	timeNeRule := NewNeRuleTimestamp("Time", endDate)
	cChecker.Add(timeNeRule, "invalid Time")

	// range rules
	intRangeRule := NewRangeRuleInt("Int", 1, 20)
	cChecker.Add(intRangeRule, "invalid Int")

	uintRangeRule := NewRangeRuleUint("Uint", 10, 258)
	cChecker.Add(uintRangeRule, "invalid Uint")

	floatRangeRule := NewRangeRuleFloat("Float", 3.14, 6.28)
	cChecker.Add(floatRangeRule, "invalid Float")

	timeRangeStrRule := NewRangeRuleTimestampStr("TimeStr", layout, startDate, endDate)
	cChecker.Add(timeRangeStrRule, "invalid TimeStr")

	timeRangeRule := NewRangeRuleTimestamp("Time", startDate, endDate)
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
	cChecker := NewChecker()

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

	eqRule := NewEqRuleComp("InnerInt", equivalent)
	cChecker.Add(eqRule, "invalid InnerInt[eq]")

	neRule := NewNeRuleComp("InnerInt", inequivalent)
	cChecker.Add(neRule, "invalid InnerInt[ne]")

	rangeRule := NewRangeRuleComp("InnerInt", ge, le)
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

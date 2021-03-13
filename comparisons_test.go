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
	intEqRule := EqInt("Int", 10)
	cChecker.Add(intEqRule, "invalid Int")

	uintEqRule := EqUint("Uint", 58)
	cChecker.Add(uintEqRule, "invalid Uint")

	floatRule := EqFloat("Float", 3.14)
	cChecker.Add(floatRule, "invalid Float")

	strRule := EqStr("String", "string")
	cChecker.Add(strRule, "invalid String")

	timeStrRule := EqTsStr("TimeStr", layout, startDate)
	cChecker.Add(timeStrRule, "invalid TimeStr")

	timeRule := EqTs("Time", startDate)
	cChecker.Add(timeRule, "invalid Time")

	// not equal rules
	intNeRule := NeInt("Int", 20)
	cChecker.Add(intNeRule, "invalid Int")

	uintNeRule := NeUint("Uint", 158)
	cChecker.Add(uintNeRule, "invalid Uint")

	floatNeRule := NeFloat("Float", 6.28)
	cChecker.Add(floatNeRule, "invalid Float")

	strNeRule := NeStr("String", "string12")
	cChecker.Add(strNeRule, "invalid String")

	timeNeStrRule := NeTsStr("TimeStr", layout, endDate)
	cChecker.Add(timeNeStrRule, "invalid TimeStr")

	timeNeRule := NeTs("Time", endDate)
	cChecker.Add(timeNeRule, "invalid Time")

	// range rules
	intRangeRule := RangeInt("Int", 1, 20)
	cChecker.Add(intRangeRule, "invalid Int")

	uintRangeRule := RangeUint("Uint", 10, 258)
	cChecker.Add(uintRangeRule, "invalid Uint")

	floatRangeRule := RangeFloat("Float", 3.14, 6.28)
	cChecker.Add(floatRangeRule, "invalid Float")

	timeRangeStrRule := RangeTsStr("TimeStr", layout, startDate, endDate)
	cChecker.Add(timeRangeStrRule, "invalid TimeStr")

	timeRangeRule := RangeTs("Time", startDate, endDate)
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

	eqRule := EqComp("InnerInt", equivalent)
	cChecker.Add(eqRule, "invalid InnerInt[eq]")

	neRule := NeComp("InnerInt", inequivalent)
	cChecker.Add(neRule, "invalid InnerInt[ne]")

	rangeRule := RangeComp("InnerInt", ge, le)
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

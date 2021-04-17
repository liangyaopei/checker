package checker

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

	timeStrRule := EqTimeStr("TimeStr", layout, startDate)
	cChecker.Add(timeStrRule, "invalid TimeStr")

	timeRule := EqTime("Time", startDate)
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

	timeNeStrRule := NeTimeStr("TimeStr", layout, endDate)
	cChecker.Add(timeNeStrRule, "invalid TimeStr")

	timeNeRule := NeTime("Time", endDate)
	cChecker.Add(timeNeRule, "invalid Time")

	// range rules
	intRangeRule := RangeInt("Int", 1, 20)
	cChecker.Add(intRangeRule, "invalid Int")

	uintRangeRule := RangeUint("Uint", 10, 258)
	cChecker.Add(uintRangeRule, "invalid Uint")

	floatRangeRule := RangeFloat("Float", 3.14, 6.28)
	cChecker.Add(floatRangeRule, "invalid Float")

	timeRangeStrRule := RangeTimeStr("TimeStr", layout, startDate, endDate)
	cChecker.Add(timeRangeStrRule, "invalid TimeStr")

	timeRangeRule := RangeTime("Time", startDate, endDate)
	cChecker.Add(timeRangeRule, "invalid Time")

	comp := comparison{
		Int:     10,
		Uint:    58,
		Float:   3.14,
		String:  "string",
		TimeStr: "2020-12-12",
		Time:    startDate,
	}

	isValid, _, _ := cChecker.Check(comp)
	assert.Equal(t, true, isValid)
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

	isValid, _, _ := cChecker.Check(param)
	assert.Equal(t, true, isValid)
}

func TestMapRuleSimple(t *testing.T) {
	mChecker := NewChecker()

	keyRangeRule := RangeInt("", 1, 10)
	valueEnumRule := InInt("", 8, 9, 10)

	mapRule := Map("", keyRangeRule, valueEnumRule)
	mChecker.Add(mapRule, "invalid map")

	m := map[int]int{
		1: 8,
		2: 9,
		3: 10,
	}

	isValid, _, _ := mChecker.Check(m)
	assert.Equal(t, true, isValid)
}

type keyStruct struct {
	Key int
}

type valueStruct struct {
	Value int
}

type mapStruct struct {
	Map map[keyStruct]valueStruct
}

func TestMapRuleStruct(t *testing.T) {
	mChecker := NewChecker()

	keyRangeRule := RangeInt("Key", 1, 10)
	valueEnumRule := InInt("Value", 8, 9, 10)

	mapRule := Map("Map", keyRangeRule, valueEnumRule)
	mChecker.Add(mapRule, "invalid map")

	kvMap := make(map[keyStruct]valueStruct)
	keys := []keyStruct{{1}, {2}, {3}}
	for _, key := range keys {
		kvMap[key] = valueStruct{Value: 9}
	}
	m := mapStruct{
		kvMap,
	}

	isValid, _, _ := mChecker.Check(m)
	assert.Equal(t, true, isValid)

}

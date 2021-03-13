package checker

import (
	"testing"
)

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

	isValid, prompt, errMsg := mChecker.Check(m)
	if !isValid {
		t.Errorf("errMsg:%s,prompt:%s", errMsg, prompt)
		return
	}
	t.Logf("valid map")
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

	isValid, prompt, errMsg := mChecker.Check(m)
	if !isValid {
		t.Errorf("errMsg:%s,prompt:%s", errMsg, prompt)
		return
	}
	t.Logf("valid map")

}

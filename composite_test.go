package checker

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayInner(t *testing.T) {

	type twoDimArr struct {
		Arr struct {
			Arr2 [][]int
		}
	}

	arrayS := twoDimArr{}
	arrayS.Arr.Arr2 = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	prompt := "元素不能等于1"
	rule := Field("Arr",
		Array("Arr2",
			Array("", NeInt("", 1)).Prompt(prompt),
		),
	)

	ch := NewChecker()
	ch.Add(rule, "")

	isValid, errPrompt, errMsg := ch.Check(arrayS)
	assert.Equal(t, false, isValid, "")
	assert.Equal(t, prompt, errPrompt, "wrong errPrompt")
	msg := "[NeInt]:Arr.Arr2[0][0] should not be 1,actual is 1"
	assert.Equal(t, msg, errMsg, "wrong errMsg")
}

func TestMap(t *testing.T) {
	type MapStruct struct {
		InnerMap struct {
			Map map[int]string
		}
	}

	mapStruct := MapStruct{}
	mapStruct.InnerMap.Map = map[int]string{
		0: "orange",
		1: "apple",
		2: "watermelon",
	}
	rule := Field("InnerMap", And(
		Length("Map", 0, 5).Prompt("长度必须在[2,5]"),
		Map("Map",
			RangeInt("", 0, 2).Prompt("key范围必须在[0,2]"),
			InStr("", "orange", "apple", "watermelon").Prompt("错误的水果"),
		),
	))
	ch := NewChecker()
	ch.Add(rule, "")
	isValid, errPrompt, _ := ch.Check(mapStruct)
	assert.Equal(t, true, isValid, errPrompt)

	mapStruct.InnerMap.Map = map[int]string{
		3: "watermelon",
	}
	prompt := "key范围必须在[0,2]"
	msg := "[RangeInt]:InnerMap.Map[3] should be between 0 and 2,actual is 3"
	isValid, errPrompt, errMsg := ch.Check(mapStruct)
	assert.Equal(t, false, isValid, errPrompt)
	assert.Equal(t, prompt, errPrompt, "wrong errPrompt")
	assert.Equal(t, msg, errMsg, "wrong errMsg")

	mapStruct.InnerMap.Map = map[int]string{
		2: "pineapple",
	}
	isValid, errPrompt, errMsg = ch.Check(mapStruct)
	prompt = "错误的水果"
	msg = "[InStr]:InnerMap.Map[2] 's value should be in [orange apple watermelon],actual is pineapple"
	assert.Equal(t, prompt, errPrompt, "wrong errPrompt")
	assert.Equal(t, msg, errMsg, "wrong errMsg")

}

func TestGetKeyStr(t *testing.T) {
	cases := []interface{}{
		1.2,
		1,
		uint(3),
		"abc",
	}
	expecteds := []string{
		"1.2",
		"1",
		"3",
		"abc",
	}
	for i := 0; i < len(cases); i++ {
		aCase := cases[i]
		expected := expecteds[i]
		value := reflect.ValueOf(aCase)
		res := getKeyStr(value)
		assert.Equal(t, expected, res)
	}
}

package checker

import (
	"testing"
)

func TestNonStructStr(t *testing.T) {
	email := "yaopei.liang@foxmail.com"

	nonStructChecker := NewChecker()
	emailRule := NewEmailRule("")
	nonStructChecker.Add(emailRule, "invalid email")

	isValid, prompt, errMsg := nonStructChecker.Check(email)
	if !isValid {
		t.Errorf("errMsg:%s,prompt:%s", errMsg, prompt)
		return
	}
	t.Logf("valid email")
}

func TestNonStructArr(t *testing.T) {
	arr := []int{1, 2, 3}

	nonStructChecker := NewChecker()
	rangeRule := NewRangeRuleInt("", 0, 4)
	sliceRule := NewSliceRule("", rangeRule)
	nonStructChecker.Add(sliceRule, "invalid array")

	isValid, prompt, errMsg := nonStructChecker.Check(arr)
	if !isValid {
		t.Errorf("errMsg:%s,prompt:%s", errMsg, prompt)
		return
	}
	t.Logf("valid array")
}

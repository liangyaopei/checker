package checker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonStructStr(t *testing.T) {
	email := "yaopei.liang@foxmail.com"

	nonStructChecker := NewChecker()
	emailRule := Email("")
	nonStructChecker.Add(emailRule, "invalid email")

	isValid, _, _ := nonStructChecker.Check(email)
	assert.Equal(t, true, isValid)
}

func TestNonStructArr(t *testing.T) {
	arr := []int{1, 2, 3}

	nonStructChecker := NewChecker()
	rangeRule := RangeInt("", 0, 4)
	sliceRule := Array("", rangeRule)
	nonStructChecker.Add(sliceRule, "invalid array")

	isValid, _, _ := nonStructChecker.Check(arr)
	assert.Equal(t, true, isValid)

}

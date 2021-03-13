package checker

import (
	"testing"
)

type list struct {
	Name *string
	Next *list
}

func TestListEmptyPrtField(t *testing.T) {
	name := "list"
	node1 := list{Name: &name, Next: nil}
	lists := list{Name: &name, Next: &node1}

	listChecker := NewChecker()
	nameRule := Length("Next.Name", 1, 20)
	listChecker.Add(nameRule, "invalid info name")

	isValid, prompt, errMsg := listChecker.Check(lists)
	if !isValid {
		t.Logf("prompt:%s", prompt)
		t.Logf("errMsg:%s", errMsg)
		return
	}
	t.Log("pass check")
}

func TestNilList(t *testing.T) {
	listChecker := NewChecker()
	nameRule := Length("Next.Name", 1, 20)
	listChecker.Add(nameRule, "invalid info name")

	isValid, prompt, errMsg := listChecker.Check(nil)
	if !isValid {
		t.Logf("prompt:%s", prompt)
		t.Logf("errMsg:%s", errMsg)
		return
	}
	t.Log("pass check")
}

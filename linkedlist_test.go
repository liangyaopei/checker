package checker

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	isValid, _, _ := listChecker.Check(lists)
	assert.Equal(t, true, isValid)
}

func TestNilList(t *testing.T) {
	listChecker := NewChecker()
	nameRule := Length("Next.Name", 1, 20)
	listChecker.Add(nameRule, "invalid info name")

	isValid, _, _ := listChecker.Check(nil)
	assert.Equal(t, false, isValid)
}

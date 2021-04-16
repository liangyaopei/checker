package checker

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type sizeInt int

type field struct {
	Name       string
	Email      *string
	LinkedList Node
	Salary     float64
	Age        uint
	Size       sizeInt
	BirthDate  time.Time
	Comp       Comparable
}

type Node struct {
	Val  int
	Next *Node
}

func TestFetchFieldInStruct(t *testing.T) {
	email := "yaopei.liang@foxmail.com"
	name := "yaopei"
	linkedlist := Node{
		Val: 1,
		Next: &Node{
			Val: 2,
			Next: &Node{
				Val:  3,
				Next: nil,
			},
		},
	}
	f := field{Name: name, Email: &email, LinkedList: linkedlist}

	fName, _ := fetchField(f, &ruleWrapper{BaseRule{fieldExpr: "Name"}})
	fEmail, _ := fetchField(f, &ruleWrapper{BaseRule{fieldExpr: "Email"}})
	node1Val, _ := fetchField(f, &ruleWrapper{BaseRule{fieldExpr: "LinkedList.Val"}})
	node2Val, _ := fetchField(f, &ruleWrapper{BaseRule{fieldExpr: "LinkedList.Next.Val"}})
	node3Val, _ := fetchField(f, &ruleWrapper{BaseRule{fieldExpr: "LinkedList.Next.Next.Val"}})
	nodeNotNil, _ := fetchField(f, &ruleWrapper{BaseRule{fieldExpr: "LinkedList.Next.Next"}})
	nodeNil, _ := fetchField(f, &ruleWrapper{BaseRule{fieldExpr: "LinkedList.Next.Next.Next.Val"}})

	assert.Equal(t, name, fName, "error name")
	assert.Equal(t, email, fEmail, "error email")
	assert.Equal(t, 1, node1Val, "error node1Val")
	assert.Equal(t, 2, node2Val, "error node2Val")
	assert.Equal(t, 3, node3Val, "error node3Val")
	assert.NotNil(t, nodeNotNil, "error nodeNotNil")
	assert.Equal(t, nil, nodeNil, "error nodeNil")
}

func TestFetchFieldStr(t *testing.T) {
	email := "yaopei.liang@foxmail.com"
	name := "yaopei"

	f := field{Name: name, Email: &email}
	fName, _, _ := fetchFieldStr(f, &ruleWrapper{BaseRule{fieldExpr: "Name", name: "ruleWrapper"}})
	fEmail, _ := fetchField(f, &ruleWrapper{BaseRule{fieldExpr: "Email"}})
	assert.Equal(t, fName, name, "error name")
	assert.Equal(t, fEmail, email, "error email")
}

func TestFetchFieldStrErrPrompt(t *testing.T) {
	email := "yaopei.liang@foxmail.com"
	name := "yaopei"
	f := field{Name: name, Email: &email}
	_, isValid, errMsg := fetchFieldStr(f, &ruleWrapper{BaseRule{fieldExpr: "Name.Name", name: "ruleWrapper"}})

	assert.Equal(t, false, isValid, "wrong fetchField")
	assert.Equal(t, "[ruleWrapper]:Name.Name cannot be found", errMsg, "wrong errMsg")

}

func TestFetchFieldInt(t *testing.T) {
	size := 100
	f := field{Size: sizeInt(size)}
	fSize, _, errMsg := fetchFieldInt(f, &ruleWrapper{BaseRule{fieldExpr: "Size", name: "ruleWrapper"}})
	assert.Equal(t, fSize, size, errMsg)
}

func TestFetchFieldUInt(t *testing.T) {
	var age uint = 35
	f := field{Age: 35}
	fAge, _, errMsg := fetchFieldUint(f, &ruleWrapper{BaseRule{fieldExpr: "Age", name: "ruleWrapper"}})
	assert.Equal(t, fAge, age, errMsg)
}

func TestFetchFieldFloat(t *testing.T) {
	salary := 100.0
	f := field{Salary: salary}
	fSalary, _, errMsg := fetchFieldFloat(f, &ruleWrapper{BaseRule{fieldExpr: "Salary", name: "ruleWrapper"}})
	assert.Equal(t, fSalary, salary, errMsg)
}

func TestFetchFieldTime(t *testing.T) {
	birthDate, _ := time.Parse("2006-01-02", "2020-03-01")
	f := field{BirthDate: birthDate}
	fBirthDate, _, errMsg := fetchFieldTime(f, &ruleWrapper{BaseRule{fieldExpr: "BirthDate", name: "ruleWrapper"}})
	assert.Equal(t, fBirthDate, birthDate, errMsg)
}

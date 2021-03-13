package checker

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type field struct {
	Name       string
	Email      *string
	LinkedList Node
	Salary     float64
	Age        uint
	Size       int
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

	fName, _ := fetchFieldInStruct(f, "Name")
	fEmail, _ := fetchFieldInStruct(f, "Email")
	node1Val, _ := fetchFieldInStruct(f, "LinkedList.Val")
	node2Val, _ := fetchFieldInStruct(f, "LinkedList.Next.Val")
	node3Val, _ := fetchFieldInStruct(f, "LinkedList.Next.Next.Val")
	nodeNotNil, _ := fetchFieldInStruct(f, "LinkedList.Next.Next")
	nodeNil, _ := fetchFieldInStruct(f, "LinkedList.Next.Next.Next.Val")

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

	fName, _, _ := fetchFieldStr(f, "Name", "")
	fEmail, _, _ := fetchFieldStr(f, "Email", "")

	assert.Equal(t, name, fName, "error name")
	assert.Equal(t, email, fEmail, "error email")
}

func TestFetchFieldInt(t *testing.T) {
	size := 100
	f := field{Size: size}
	fSize, _, _ := fetchFieldInt(f, "Size", "")
	assert.Equal(t, size, fSize, "error fSize")
}

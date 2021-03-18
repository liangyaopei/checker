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

	fName, _ := fetchField(f, "Name")
	fEmail, _ := fetchField(f, "Email")
	node1Val, _ := fetchField(f, "LinkedList.Val")
	node2Val, _ := fetchField(f, "LinkedList.Next.Val")
	node3Val, _ := fetchField(f, "LinkedList.Next.Next.Val")
	nodeNotNil, _ := fetchField(f, "LinkedList.Next.Next")
	nodeNil, _ := fetchField(f, "LinkedList.Next.Next.Next.Val")

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
	f := field{Size: sizeInt(size)}
	fSize, _, _ := fetchFieldInt(f, "Size", "")
	assert.Equal(t, size, fSize, "error fSize")
}

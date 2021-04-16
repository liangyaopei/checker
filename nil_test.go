package checker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type zero struct {
	Int *int
}

func TestNonZeroRule(t *testing.T) {
	nonZeroChecker := NewChecker()

	iRule := Nil("Int")
	nonZeroChecker.Add(iRule, "invalid Int")

	z := zero{
		Int: nil,
	}

	isValid, _, _ := nonZeroChecker.Check(z)
	assert.Equal(t, true, isValid, "failed TestNonZeroRule")

}

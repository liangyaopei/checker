package checker

import (
	"testing"
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

	isValid, prompt, errMsg := nonZeroChecker.Check(z)
	if !isValid {
		t.Errorf("errMsg:%s,prompt:%s", errMsg, prompt)
		return
	}
	t.Logf("valid nil")
}

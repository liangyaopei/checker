package checker

import (
	"testing"
	"time"
)

type zero struct {
	Int   int
	Float float64
	Str   string
	Date  time.Time
}

func TestNonZeroRule(t *testing.T) {
	nonZeroChecker := NewChecker()

	iRule := NewNonZeroRule("Int")
	nonZeroChecker.Add(iRule, "invalid Int")

	fRule := NewNonZeroRule("Float")
	nonZeroChecker.Add(fRule, "invalid Float")

	sRule := NewNonZeroRule("Str")
	nonZeroChecker.Add(sRule, "invalid Str")

	dRule := NewNonZeroRule("Date")
	nonZeroChecker.Add(dRule, "invalid Date")

	date, _ := time.Parse("2006-01-01", "2020-12-01")
	z := zero{
		Int:   1,
		Float: 3.0,
		Str:   "checker",
		Date:  date,
	}

	isValid, prompt, errMsg := nonZeroChecker.Check(z)
	if !isValid {
		t.Errorf("errMsg:%s,prompt:%s", errMsg, prompt)
		return
	}
	t.Logf("valid zero")
}

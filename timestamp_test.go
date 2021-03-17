package checker

import (
	"testing"
	"time"
)

type timestamp struct {
	Date         string
	StartDate    *time.Time
	StartDateStr string
	RangeDate    time.Time
	RangeDateStr string
}

func TestRuleTimeStampStr(t *testing.T) {
	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, "2020-12-12")
	endDate, _ := time.Parse(layout, "2020-12-31")
	rangeDate, _ := time.Parse(layout, "2020-12-20")

	tsChecker := NewChecker()

	tsRule := isDatetime("Date", layout)
	tsChecker.Add(tsRule, "invalid Date")

	tsEqRule := EqTime("StartDate", startDate)
	tsChecker.Add(tsEqRule, "invalid StartDate")

	tsStrRule := EqTimeStr("StartDateStr", layout, startDate)
	tsChecker.Add(tsStrRule, "invalid StartDateStr")

	rangeTsRule := RangeTime("RangeDate", startDate, endDate)
	tsChecker.Add(rangeTsRule, "invalid RangeDate")

	rangeTsStrRule := RangeTimeStr("RangeDateStr",
		layout, startDate, endDate)
	tsChecker.Add(rangeTsStrRule, "invalid RangeDateStr")

	startDateTs, _ := time.Parse(layout, "2020-12-12")
	ts := timestamp{
		Date:         "2020-12-01",
		StartDate:    &startDateTs,
		StartDateStr: "2020-12-12",
		RangeDate:    rangeDate,
		RangeDateStr: "2020-12-15",
	}

	isValid, prompt, errMsg := tsChecker.Check(ts)
	if !isValid {
		t.Errorf("errMsg:%s,prompt:%s", errMsg, prompt)
		return
	}
	t.Logf("valid Date")
}

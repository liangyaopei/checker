package checker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type profile struct {
	// Info is pointer filed
	Info      *basicInfo
	Companies []company
}

type basicInfo struct {
	// 1 <= len <= 20
	Name string
	// 18 <= age <= 80
	Age int
	// 1<= len <= 64
	Email string
}

type company struct {
	// frontend,backend
	Position string
	// frontend: html,css,javascript
	// backend: C,Cpp,Java,Golang
	// SkillStack 'length is between [1,3]
	Skills []string
}

func getPassedProfile() profile {
	companies := []company{
		{
			Position: "frontend",
			Skills:   []string{"html", "css"},
		},
		{
			Position: "backend",
			Skills:   []string{"C", "Golang"},
		},
	}
	info := basicInfo{Name: "liang", Age: 24, Email: "yaopei.liang@foxmail.com"}
	return profile{
		Info:      &info,
		Companies: companies,
	}
}

func getFailedProfile() profile {
	companies := []company{
		{
			Position: "frontend",
			Skills:   []string{"Golang", "css"},
		},
		{
			Position: "backend",
			Skills:   []string{"C", "Golang"},
		},
	}
	info := basicInfo{Name: "liang", Age: 24, Email: "yaopei.liang@foxmail.com"}
	return profile{
		Info:      &info,
		Companies: companies,
	}
}

func getProfileChecker() Checker {
	profileChecker := NewChecker()

	rule :=
		And(
			Length("Info.Name", 1, 20),
			RangeInt("Info.Age", 18, 80),
			Array("Companies",
				And(
					Length("Skills", 1, 3),
					Or(
						And(
							EqStr("Position", "frontend"),
							Array("Skills",
								InStr("", "html", "css", "javascript"),
							),
						),
						And(
							EqStr("Position", "backend"),
							Array("Skills",
								InStr("", "C", "CPP", "Java", "Golang"),
							),
						)),
				)),
		)
	profileChecker.Add(rule, "invalid companies")

	return profileChecker
}

func TestProfileCheckerPassed(t *testing.T) {
	profile := getPassedProfile()
	profileChecker := getProfileChecker()
	isValid, prompt, errMsg := profileChecker.Check(profile)
	if !isValid {
		t.Logf("prompt:%s", prompt)
		t.Logf("errMsg:%s", errMsg)
		return
	}
	t.Log("pass check")
}

func TestProfileCheckerFailed(t *testing.T) {
	profile := getFailedProfile()
	profileChecker := getProfileChecker()
	isValid, _, _ := profileChecker.Check(profile)
	assert.Equal(t, false, isValid, "error failed checker")
}

type Item struct {
	Info typeInfo
}

type typeInfo struct {
	Type        string
	Range       []string
	Unit        string
	Granularity string
}

func TestField(t *testing.T) {
	items := []Item{
		{
			Info: typeInfo{
				Type:  "range",
				Range: []string{"2020-01-01", "2021-01-01"},
			},
		},
		{
			Info: typeInfo{
				Type:        "last",
				Granularity: "day",
				Unit:        "7",
			},
		},
	}

	rule := Field("Info",
		Or(
			And(
				EqStr("Type", "range"),
				Length("Range", 2, 2),
				Array("Range", isDatetime("", "2006-01-02")),
			),
			And(
				EqStr("Type", "last"),
				InStr("Granularity", "day", "week", "month"),
				Number("Unit"),
			),
		),
	)
	itemChecker := NewChecker()
	itemChecker.Add(rule, "wrong item")

	for _, item := range items {
		isValid, _, errMsg := itemChecker.Check(item)
		assert.Equal(t, true, isValid, errMsg)
	}
}

package checker

import (
	"testing"
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

	infoNameRule := NewLengthRule("Info.Name", 1, 20)
	profileChecker.Add(infoNameRule, "invalid info name")

	infoAgeRule := NewRangeRuleInt("Info.Age", 18, 80)
	profileChecker.Add(infoAgeRule, "invalid info age")

	infoEmailRule := NewAndRule(
		NewLengthRule("Info.Email", 1, 64),
		NewEmailRule("Info.Email"),
	)
	profileChecker.Add(infoEmailRule, "invalid info email")

	companyLenRule := NewLengthRule("Companies", 1, 3)
	profileChecker.Add(companyLenRule, "invalid companies len")

	frontendRule := NewAndRule(
		NewEqRuleString("Position", "frontend"),
		NewSliceRule("Skills",
			NewEnumRuleString("", "html", "css", "javascript"),
		),
	)
	backendRule := NewAndRule(
		NewEqRuleString("Position", "backend"),
		NewSliceRule("Skills",
			NewEnumRuleString("", "C", "CPP", "Java", "Golang"),
		),
	)
	companiesSliceRule := NewSliceRule("Companies",
		NewAndRule(
			NewLengthRule("Skills", 1, 3),
			NewOrRule(frontendRule, backendRule),
		))
	profileChecker.Add(companiesSliceRule, "invalid skill item")

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
	isValid, prompt, errMsg := profileChecker.Check(profile)
	if !isValid {
		t.Logf("prompt:%s", prompt)
		t.Logf("errMsg:%s", errMsg)
		return
	}
	t.Log("pass check")
}

package _checker_test

import (
	"testing"

	"github.com/liangyaopei/checker"
)

type profile struct {
	Info *basicInfo
	//Info      basicInfo
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

func getProfileChecker() checker.Checker {
	profileChecker := checker.NewChecker()

	infoNameRule := checker.NewLengthRule("Info.Name", 1, 20)
	profileChecker.Add(infoNameRule, "invalid info name")

	infoAgeRule := checker.NewRangeRuleInt("Info.Age", 18, 80)
	profileChecker.Add(infoAgeRule, "invalid info age")

	infoEmailRule := checker.NewAndRule([]checker.Rule{
		checker.NewLengthRule("Info.Email", 1, 64),
		checker.NewEmailRule("Info.Email"),
	})
	profileChecker.Add(infoEmailRule, "invalid info email")

	companyLenRule := checker.NewLengthRule("Companies", 1, 3)
	profileChecker.Add(companyLenRule, "invalid companies len")

	frontendRule := checker.NewAndRule([]checker.Rule{
		checker.NewEqRuleString("Position", "frontend"),
		checker.NewSliceRule("Skills",
			checker.NewEnumRuleString("", []string{"html", "css", "javascript"}),
		),
	})
	backendRule := checker.NewAndRule([]checker.Rule{
		checker.NewEqRuleString("Position", "backend"),
		checker.NewSliceRule("Skills",
			checker.NewEnumRuleString("", []string{"C", "CPP", "Java", "Golang"}),
		),
	})
	companiesSliceRule := checker.NewSliceRule("Companies",
		checker.NewAndRule([]checker.Rule{
			checker.NewLengthRule("Skills", 1, 3),
			checker.NewOrRule([]checker.Rule{
				frontendRule, backendRule,
			}),
		}))
	profileChecker.Add(companiesSliceRule, "invalid skill item")

	return profileChecker
}

func TestChecker(t *testing.T) {
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

func TestChecker2(t *testing.T) {
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

type list struct {
	Name *string
	Next *list
}

func TestList(t *testing.T) {
	name := "list"
	//node1 := list{Name: nil}
	node2 := list{Name: nil, Next: nil}
	lists := list{Name: &name, Next: &node2}

	listChecker := checker.NewChecker()
	nameRule := checker.NewLengthRule("Next.Name", 1, 20)
	listChecker.Add(nameRule, "invalid info name")

	isValid, prompt, errMsg := listChecker.Check(lists)
	if !isValid {
		t.Logf("prompt:%s", prompt)
		t.Logf("errMsg:%s", errMsg)
		return
	}
	t.Log("pass check")
}

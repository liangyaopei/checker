# Overview
This repo provides flexible, configurable validation method when validating struct in Golang.

# Install
```
go get -u github.com/liangyaopei/checker
```

# Description
## Rule
Users can customize their various validation rules attached every filed in struct,
 including complex,composite rules.
Below rules all accepts a `fieldExpr` to locate the field in struct, and:
- `NewAndRule` accepts an array of rules, it passed when all rules are passed.
- `NewOrRule` accepts an array of rules, it passed when any rule is passed.
- `NewSliceRule` accepts slice/array input and a rule, applies the to every element in array/slice.
- `NewLengthRule` accepts slice/array/string/map, it checks the length of input is less than or equal to limit.
- `NewEnumRuleXXX` accepts a slice, it checks whether value of related struct filed exist in input enum slice. It has
  `NewEnumRuleString`,`NewEnumRuleInt`, `NewEnumRuleUint`,`NewEnumRuleFloat`.
- `NewEqRuleXXX` accepts a value, it checks whether value of related struct filed is equal to value. It 
    has `NewEnumRuleString`, `NewEnumRuleInt`, `NewEnumRuleUint`,`NewEnumRuleFloat`.
- `NewNotEqRuleXXX` is similar to `NewEqRuleXXX`. It checks whether value of related struct filed is not equal to value.
- `NewRangeRuleXXX` accepts min and max value, it checks whether value of related struct filed is `min<=v<=max`. It has
   `NewRangeRuleInt`, `NewRangeRuleUint`,`NewRangeRuleFloat`.

## Checker
`NewChecker` returns a `Checker` interface. `Checker` has two methods:
- `Add(rule Rule, prompt string)` add rules, `prompt` is the prompt when the rule fails.
- `Check(param interface{}) (bool, string, string)` means checker use all added rules to check/validate struct.
`bool` tells whether all rules are passed or not, the first `string` is prompt of the rule, the second `string`
tells detailed information about why the value of field in struct fails.

# Example
Below example is taken from `./_checker_test/checker_test.go`.

Consider below `profile` struct. 
In `basicInfo` struct, `Name` has length limit, `Age` has length limit,`Email` has length limit and format limitation.
In `company` struct, `Position` can only be `frontend` or `backend`.
- if `Position` is equal to `frontend`, elem in `Skills` can only be `html,css,javascript`.
- if `Position` is equal to `backend`, elem in `Skills` can only be `C,Cpp,Java,Golang`.
- besides, `Skills` has length limit.
```go
type profile struct {
	Info      basicInfo
	Companies []company
}

type basicInfo struct {
	// len <= 20
	Name string
	// 18 <= age <= 80
	Age int
	// len <= 64
	Email string
}

type company struct {
	// frontend,backend
	Position string
	// frontend: html,css,javascript
	// backend: C,Cpp,Java,Golang
	// SkillStack has length limit 3
	Skills []string
}
```
rule and checker can be defined as follow:
```go
profileChecker := checker.NewChecker()

	infoNameRule := checker.NewLengthRule("Info.Name", 20)
	profileChecker.Add(infoNameRule, "invalid info name")

	infoAgeRule := checker.NewRangeRuleInt("Info.Age", 18, 80)
	profileChecker.Add(infoAgeRule, "invalid info age")

	infoEmailRule := checker.NewAndRule([]checker.Rule{
		checker.NewLengthRule("Info.Email", 64),
		checker.NewEmailRule("Info.Email"),
	})
	profileChecker.Add(infoEmailRule, "invalid info email")

	companyLenRule := checker.NewLengthRule("Companies", 3)
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
			checker.NewLengthRule("Skills", 3),
			checker.NewOrRule([]checker.Rule{
				frontendRule, backendRule,
			}),
		}))
	profileChecker.Add(companiesSliceRule, "invalid skill item")
```
We can see that the rules is more readable, flexible,configurable 
than simply using `if/else` statement.
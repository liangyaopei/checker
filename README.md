# Overview
Checker provides flexible, configurable validation method when validating struct in Golang.

# Install
```
go get -u github.com/liangyaopei/checker
```

# Description
## Rule
`Rule` is an interface, it checks whether the param is valid or not.
```go
type Rule interface {
	check(param interface{}) (bool, string)
}
```
All rules implementation accepts a `fieldExpr` to locate the field in struct.
`fieldExpr` is a string with format `Field1.Filed2`.
When filed is a pointer, first will check the field is nil or not; if not, get the pointer's
underlying value as param to validate.
All rule implementations include:
- Composition rule:
    - `NewAndRule` accepts an array of rules, it passed when all the rules  passed.
    - `NewOrRule` accepts an array of rules, it passed when one of the rules passed.
- Slice Rule: 
    - `NewSliceRule` accepts slice/array input and a rule, applies the to every element in array/slice.
- Enum Rule:
    - `NewEnumRule*` accepts a slice, it checks whether value of related struct filed exist in input enum slice. 
    Implementations include `NewEnumRuleString`,`NewEnumRuleInt`, `NewEnumRuleUint`,`NewEnumRuleFloat`.
- Comparison Rule:
    - `NewRqRule*` accepts a value, it checks whether value of related struct filed is equal to value.
- Regex Rule:
    - accepts a regex expression to check whether string field satisfies the regex expression.
    - Implementations include `NewEmailRule`,`NewAlphaNumericRule`,etc.
    

## Checker
`NewChecker` returns a `Checker` interface. `Checker` has two methods:
- `Add(rule Rule, prompt string)` add rules, `prompt` is the prompt when the rule fails.
- `Check(param interface{}) (bool, string, string)` means checker use all added rules to check/validate struct.
`bool` tells whether all rules are passed or not, the first `string` is prompt of the rule, the second `string`
tells detailed information about why the value of field in struct fails.

# Example
Below example is from `./_checker_test/checker_test.go`.

Consider below `profile` struct. 
- `Info` field
    - `Info` is not nil
    - `Name` has length limit between 1 and 20
    - `Age` has range limit between 18 and 80
    - `Email` has length limit between 1 and 64 , is  format of email.
-  `Companies` field
    - `Position` can only be `frontend` or `backend`
        - if `Position` is equal to `frontend`, elem in `Skills` can only be `html,css,javascript`.
        - if `Position` is equal to `backend`, elem in `Skills` can only be `C,Cpp,Java,Golang`.
 - besides, `Skills` has length limit between 1 and 3.
```go
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
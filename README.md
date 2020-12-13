# Checker
[中文版本](README_zh.md)

`Checker` is a parameter validation package, it can replace [gopkg.in/go-playground/validator.v10](https://godoc.org/gopkg.in/go-playground/validator.v10). `Checker` can be use in struct/non-struct validation, including cross field validation in struct, elements validation in Slice/Array/Map, and provides customized validation rule.

## Installation

```
go get -u github.com/liangyaopei/checker
```



## Usage

Examples are in [_checker_test](_checker_test).

The main principle is, every validation rule is a `Rule` interface, `Rule` validates parameter, returns `isValid` and error log.

`Checker` is a validatior, it adds `Rule` and error prompt on related fileld in sturtc.



For example, [non-struct parameter validation](_checker_test/nonstruct_test.go), `fieldExpr` is empty string.

```go
email := "abc@examplecom"

nonStructChecker := checker.NewChecker()

emailRule := checker.NewEmailRule("")
nonStructChecker.Add(emailRule, "invalid email")

isValid, prompt, errMsg := nonStructChecker.Check(email)
```



[struct parameter validation](_checker_test/timestamp_test.go)

```go
type timestamp struct {
	StartDateStr string
}

layout := "2006-01-02"
startDate, _ := time.Parse(layout, "2020-12-12")

tsChecker := checker.NewChecker()
tsStrRule := checker.NewEqRuleTimestampStr("StartDateStr", layout, startDate)
tsChecker.Add(tsStrRule, "invalid StartDateStr")

ts := timestamp{
	StartDateStr: "2020-12-12",
}
isValid, prompt, errMsg := tsChecker.Check(ts)
```

[customized validation rule](_checker_test/customized_rule_test.go), only implements `Rule` interface.



## Rule realted to corresponding tag in validator.v10

### Cross filed comparasion

| tag           | Rule                                                         |
| ------------- | ------------------------------------------------------------ |
| eqfield       | `NewCrossFieldCompareRule("Int1", "Int2", checker.CrossFiledEq)` |
| fieldcontains | `NewEnumRuleInt("Value", []int{8, 9, 10})`                   |
| fieldexcludes | `Not(checker.NewEnumRuleInt("Value", []int{8, 9, 10}))`      |
| gtfield       | `NewCrossFieldCompareRule("Int1", "Int2", CrossFiledGt)`     |
| gtefield      | `NewCrossFieldCompareRule("Int1", "Int2", checker.CrossFiledGe)` |
| nefield       | `NewCrossFieldCompareRule("Int1", "Int2", checker.CrossFiledNe)` |



### Strings

| tag      | Rule                           |
| -------- | ------------------------------ |
| alpha    | `NewAlphaRule("Field")`        |
| alphanum | `NewAlphaNumericRule("Field")` |
| email    | `NewEmailRule("Email")`        |
| isbn10   | `NewISBN10Rule("Field")`       |
| isbn10   | `NewISBN13Rule("Field")`       |

etc, regrex expression for string rule, can use `NewRegexRule(fieldExpr string, regexExpr string)`



### Comparasion

| Tag            | Rule                                                      |
| -------------- | --------------------------------------------------------- |
| eq             | `NewEqRuleInt(filedExpr string, equivalent int)` ...      |
| gt, gte,lt,lte | `NewRangeRuleInt(filedExpr string, ge int, le int)` ...   |
| ne             | `NewNotEqRuleInt(filedExpr string, inequivalent int)` ... |



### Others

| Tag                             | Rule                                                         |
| ------------------------------- | ------------------------------------------------------------ |
| len                             | `NewLengthRule(fieldExpr string, ge int, le int)`            |
| required_if, required_without等 | 通过 `NewAndRule(rules []Rule) Rule`, `NewOrRule(rules []Rule)`, `NewNotRule(innerRule Rule)`的组合实现 |



### easy for checker, hard for validatior

The main drawback of `validator` is,  validation rule is attached to fields in struct via tag, which is intrusive, and hard to read the validation logic.

1.  validation sturct of third party

```go
package thrid_party

type Param struct{
  Age `validate:"min=18,max=80"`
}
```

In user's package, try to change min to 20, `validator` can not change the validation rule, as we cannot change the struct layout outside our packages.

```go
package main

func validate(p thrid_party.Param)(isValid bool){
  ....
}

```

If use `checker`, the rule is simpler:

```go
rule := checker.NewRangeRuleInt("Age", 20, 80)
checker.Add(rule, "invlaid age")
```

因为`checker`的校验规则与结构体解耦，因此，修改校验规则非常简单。

Because validation rule of `checker` is decoupled from struct, which makes changes validation rule easy.

2. validate the length of linkedlist

The example is [here](_checker_test/linkedlist_test.go).

```go
type list struct {
	Name *string
  Next *list `validate:"nonzero"`
}
```

To validate the length of linkedlist, requiring the first node's `Next` cannot be nil. `validator` cannot do this, for the same tag is attached to the same field.

If use `checker`，

```go
	name := "list"
	node1 := list{Name: &name, Next: nil}
	lists := list{Name: &name, Next: &node1}

	listChecker := checker.NewChecker()
	nameRule := checker.NewLengthRule("Next.Name", 1, 20)
	listChecker.Add(nameRule, "invalid info name")
```

Length can be defined by `Next.Name`
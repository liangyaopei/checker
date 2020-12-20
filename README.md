# Checker
[![Go Report Card](https://goreportcard.com/badge/github.com/liangyaopei/checker)](https://goreportcard.com/report/github.com/liangyaopei/checker)
[![GoDoc](https://godoc.org/github.com/liangyaopei/checker?status.svg)](http://godoc.org/github.com/liangyaopei/checker)
[![Go Reference](https://pkg.go.dev/badge/github.com/liangyaopei/checker.svg)](https://pkg.go.dev/github.com/liangyaopei/checker)
[![Build Status](https://travis-ci.com/liangyaopei/checker.svg?branch=master)](https://travis-ci.com/liangyaopei/checker)
![License](https://img.shields.io/dub/l/vibe-d.svg)
[![Coverage Status](https://coveralls.io/repos/github/liangyaopei/checker/badge.svg?branch=master)](https://coveralls.io/github/liangyaopei/checker?branch=master)

[中文版本](README_zh.md)

`Checker` is a parameter validation package, it can replace [gopkg.in/go-playground/validator.v10](https://godoc.org/gopkg.in/go-playground/validator.v10). `Checker` can be use in struct/non-struct validation, including cross field validation in struct, elements validation in Slice/Array/Map, and provides customized validation rule.

## Requirements

Go 1.13 or above.

## Installation

```
go get -u github.com/liangyaopei/checker
```



## Usage

Examples are in test files which have suffix `test`.

The main principle is, every validation rule is a `Rule` interface, `Rule` validates parameter, returns `isValid` and error log.

`Checker` is a validatior, it adds `Rule` and error prompt on related fileld in struct.



For example, [non-struct parameter validation](nonstruct_test.go), `fieldExpr` is empty string.

```go
email := "abc@examplecom"

nonStructChecker := checker.NewChecker()

emailRule := checker.NewEmailRule("")
nonStructChecker.Add(emailRule, "invalid email")

isValid, prompt, errMsg := nonStructChecker.Check(email)
```



[struct parameter validation](timestamp_test.go)

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

[customized validation rule](customized_rule_test.go), only implements `Rule` interface.



## Rule realted to corresponding tag in validator.v10

### Cross filed comparasion

| tag           | Rule                                                         |
| ------------- | ------------------------------------------------------------ |
| eqfield       | `NewCrossFieldCompareRule("Int1", "Int2", checker.CrossFieldEq)` |
| fieldcontains | `NewEnumRuleInt("Value", []int{8, 9, 10})`                   |
| fieldexcludes | `Not(checker.NewEnumRuleInt("Value", []int{8, 9, 10}))`      |
| gtfield       | `NewCrossFieldCompareRule("Int1", "Int2", CrossFieldGt)`     |
| gtefield      | `NewCrossFieldCompareRule("Int1", "Int2", checker.CrossFieldGe)` |
| nefield       | `NewCrossFieldCompareRule("Int1", "Int2", checker.CrossFieldNe)` |



### Strings

| tag              | Rule                                  |      |
| ---------------- | ------------------------------------- | ---- |
| alpha            | `NewAlphaRule("Field")`               |      |
| alphanum         | `NewAlphaNumericRule("Field")`        |      |
| email            | `NewEmailRule("Email")`               |      |
| isbn10           | `NewISBN10Rule("Field")`              |      |
| isbn10           | `NewISBN13Rule("Field")`              |      |
| isbn             | `NewISBNRuke("Field")`                |      |
| ip               | `NewIPRule(FieldExpr)`                |      |
| ipv4             | `NewIPv4Rule("IPv4")`                 |      |
| ipv6             | `NewIPv6Rule("IPv6")`                 |      |
| uri              | `NewURIRule("URL")`                   |      |
| url              | `NewURLRule("URL")`                   |      |
| url_encoded      | `NewURLEncodedRule("Field")`          |      |
| html             | `NewHTMLRule("Field")`                |      |
| html_encoded     | `NewHTMLEncodedRule("Field")`         |      |
| hostname         | `NewHostNameRule("Field")`            |      |
| hostname_rfc1123 | `NewHostNameRFC1123Rule("Field")`     |      |
| json             | `NewIsJSONRule("Field")`              |      |
| dir              | `NewIsDirRule("Field")`               |      |
| datetime         | `NewIsDatetimeRule("Field","layout")` |      |

etc, regrex expression for string rule, can use `NewRegexRule(fieldExpr string, regexExpr string)`



### Comparasion

| Tag            | Rule                                                      |
| -------------- | --------------------------------------------------------- |
| eq             | `NewEqRuleInt(filedExpr string, equivalent int)` ...      |
| gt, gte,lt,lte | `NewRangeRuleInt(filedExpr string, ge int, le int)` ...   |
| ne             | `NewNotEqRuleInt(filedExpr string, inequivalent int)` ... |



### Slice/Array/Map

| Rule                                                         | Usage                                                        |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| `NewSliceRule(fieldExpr string, innerRule Rule) `            | apply `innerRule` to every elements in slice/array           |
| `NewMapRule(fieldExpr string, keyRule Rule, valueRule Rule)` | apply `keyRule` to keys in map, and `ValueRule` to valuse in map. `keyRule` or `ValueRule` can be nil. |



### Others

| Tag                             | Rule                                                         |
| ------------------------------- | ------------------------------------------------------------ |
| len                             | `NewLengthRule(fieldExpr string, ge int, le int)`            |
| required_if, required_without,etc | implemented by compositions of `NewAndRule(rules []Rule) Rule`, `NewOrRule(rules []Rule)`, `NewNotRule(innerRule Rule)` |



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
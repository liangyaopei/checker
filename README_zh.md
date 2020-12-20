# Checker
[![Go Report Card](https://goreportcard.com/badge/github.com/liangyaopei/checker)](https://goreportcard.com/report/github.com/liangyaopei/checker)
[![GoDoc](https://godoc.org/github.com/liangyaopei/checker?status.svg)](http://godoc.org/github.com/liangyaopei/checker)
[![Go Reference](https://pkg.go.dev/badge/github.com/liangyaopei/checker.svg)](https://pkg.go.dev/github.com/liangyaopei/checker)
[![Build Status](https://travis-ci.com/liangyaopei/checker.svg?branch=master)](https://travis-ci.com/liangyaopei/checker)
![License](https://img.shields.io/dub/l/vibe-d.svg)
[![Coverage Status](https://coveralls.io/repos/github/liangyaopei/checker/badge.svg?branch=master)](https://coveralls.io/github/liangyaopei/checker?branch=master)

[English Version](README.md)


`Checker`是Golang的参数校验的包，它可以完全替代[gopkg.in/go-playground/validator.v10](https://godoc.org/gopkg.in/go-playground/validator.v10)。`Checker`用于结构体或者非结构的参数校验，包括结构体中不同字段比较的校验，Slice/Array/Map中的元素校验，还提供自定义的校验规则。

## Go版本

Go 1.13 或以上.


## 安装
```
go get -u github.com/liangyaopei/checker
```

## 使用
使用的例子都在`test`为后缀的测试文件。
主要思想是，每个校验规则都是一个`Rule`，`Rule`对参数进行校验，返回是否合法以及错误日志。
`Checker`是校验器，在结构体的字段上添加`Rule`和错误提示。

例如，[非结构体的参数校验](nonstruct_test.go)，`fieldExpr`传空字符串。
```go
email := "abc@examplecom"

nonStructChecker := checker.NewChecker()

emailRule := checker.NewEmailRule("")
nonStructChecker.Add(emailRule, "invalid email")

isValid, prompt, errMsg := nonStructChecker.Check(email)
```

[结构体的参数校验](timestamp_test.go)。
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

[自定义校验规则](customized_rule_test.go),只要实现`Rule`接口即可。


## 与validator.v10的tag对应的Rule

### 跨字段的比较

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

等等，字符串自定义的正则表达式，可以使用`NewRegexRule(fieldExpr string, regexExpr string)`



### 比较



| Tag            | Rule                                                      |
| -------------- | --------------------------------------------------------- |
| eq             | `NewEqRuleInt(filedExpr string, equivalent int)` ...      |
| gt, gte,lt,lte | `NewRangeRuleInt(filedExpr string, ge int, le int)` ...   |
| ne             | `NewNotEqRuleInt(filedExpr string, inequivalent int)` ... |



### 其他

| Tag                             | Rule                                                         |
| ------------------------------- | ------------------------------------------------------------ |
| len                             | `NewLengthRule(fieldExpr string, ge int, le int)`            |
| required_if, required_without等 | 通过 `NewAndRule(rules []Rule) Rule`, `NewOrRule(rules []Rule)`, `NewNotRule(innerRule Rule)`的组合实现 |



## checker容易做，validator难做

`validator`主要的缺点是，把校验规则以标签的形式写在结构体字段上，这用很强的侵入性，并且不易于阅读校验逻辑。

1. 校验第三方包下的结构体

```go
package thrid_party

type Param struct{
  Age `validate:"min=18,max=80"`
}
```

在自己的代码包下,将min改为20，这个时候`validator`将无法添加校验规则。

```go
package main

func validate(p thrid_party.Param)(isValid bool){
  ....
}

```

而使用`checker`，只需要改为：

```go
rule := checker.NewRangeRuleInt("Age", 20, 80)
checker.Add(rule, "invlaid age")
```

因为`checker`的校验规则与结构体解耦，因此，修改校验规则非常简单。

2. 校验链表长度

这个例子在[这里](_checker_test/linkedlist_test.go)

```go
type list struct {
	Name *string
  Next *list `validate:"nonzero"`
}
```

要校验链表的长度，要求前几个节点的`Next`不为空，`validator`不能做到，因为自引用的结构体，同样的标签适用于相同的字段。

如果使用`checker`，

```go
	name := "list"
	node1 := list{Name: &name, Next: nil}
	lists := list{Name: &name, Next: &node1}

	listChecker := checker.NewChecker()
	nameRule := checker.NewLengthRule("Next.Name", 1, 20)
	listChecker.Add(nameRule, "invalid info name")
```

通过`Next.Name`可以指定链表的长度。
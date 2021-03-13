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
使用`Add`添加规则时，`fieldExpr`有三种情况：
- `fieldExpr`为空字符串，这时会直接校验值。
- `fieldExpr`为单个字段，这时会先取字段的值，再校验。
- `fieldExpr`为点(.)分割的字段，先按照`.`的层级关系取值，再校验。

```go
type Item struct {
	Info typeInfo
}

// typeInfo.Type = "range",typeInfo.Type的长度为2，元素都是格式符合"2006-01-02"
// typeInfo.Type = "last",typeInfo.Type的长度为1，元素是正整数，Granularity只能是day/week/month之一
type typeInfo struct {
	Type        string
	Range       []string
	Unit        string
	Granularity string
}

// 规则如下
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
```
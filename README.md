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

When use `Add` to add rule，`fieldExpr` has three situations：
- `fieldExpr` is empty，validate the value directly.
- `fieldExpr` is single field，fetch value in struct, then validate.
- `fieldExpr` is string separated by `.`, fetch value in struct according hierarchy of struct, then validate.


```go
type Item struct {
	Info typeInfo
}

// typeInfo.Type = "range", length of typeInfo.Type is 2，elements meets format of "2006-01-02"
// typeInfo.Type = "last", length of typeInfo.Typeis 1，elements meets of format positive integer，
// Granularity must be one of day/week/month
type typeInfo struct {
	Type        string
	Range       []string
	Unit        string
	Granularity string
}

// here is the rule
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
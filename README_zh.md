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

按字段取值时，如果字段是指针，就取指针的值校验；如果是空指针，则视为没有通过校验。

来自`checker_test.go`的例子：
```go
// Item.Email需要符合电子邮箱的格式
type Item struct {
	Info  typeInfo
	Email string
}

type typeStr string
// Item.Info.Type = "range",typeInfo.Type的长度为2，元素都是格式符合"2006-01-02"
// Item.Info.Type = "last",typeInfo.Type的长度为1，元素是正整数，Granularity只能是day/week/month之一
type typeInfo struct {
	Type        typeStr
	Range       []string
	Unit        string
	Granularity string
}


// 规则如下
rule := And(
		Email("Email"),
		Field("Info",
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
		),
	)
itemChecker := NewChecker()
// 校验参数
itemChecker.Add(rule, "wrong item")
```

## 规则
`Rule`是一个接口，它有很多的实现。`Rule`的实现可以分为复合规则和单个规则。



### 复合规则

复合规则包含其他的规则。

| 名字                                                       | 作用                                           |
| ---------------------------------------------------------- | ---------------------------------------------- |
| `Field(fieldExpr string, rule Rule) Rule`                  | 对字段使用`rule`校验                           |
| `And(rules ...Rule) Rule`                                  | 需要所有的规则都通过                           |
| `Or(rules ...Rule) Rule`                                   | 需要由一个规则通过                             |
| `Not(innerRule Rule) Rule`                                 | 对规则取反                                     |
| `Array(fieldExpr string, innerRule Rule) Rule`             | 对数组的每一个元素使用规则                     |
| `Map(fieldExpr string, keyRule Rule, valueRule Rule) Rule` | 对map的key/value使用keyRule和valueRule进行校验 |



### 单个规则

单个规则可分为比较型，枚举型，格式型等。

#### 比较型规则

比较型规则分为单个字段比较规则，多个字段比较规则。



单个字段比较规则包括：

| 名字                                              |
| ------------------------------------------------- |
| `EqInt(filedExpr string, equivalent int) Rule`    |
| `NeInt(filedExpr string, inequivalent int) Rule`  |
| `RangeInt(filedExpr string, ge int, le int) Rule` |

以及`uint`, `string`，`float`，`time.Time` , `Comparable`的实现。

多个字段比较规则

| 名字                                                         |
| ------------------------------------------------------------ |
| `CrossComparable(fieldExprLeft string, fieldExprRight string, op operand) Rule` |

`fieldExprLeft`，`fieldExprRight`用来定位参加比较的字段，`op`是运算操作符，包括相等/不等/大于等。

``CrossComparable`支持的字段类型包括`int`\`uint`\`float`\`string`\`time.Time`\`Comparable`。



#### 枚举型规则

枚举型包括

| 名字                                              |
| ------------------------------------------------- |
| `InStr(filedExpr string, enum ...string) Rule`    |
| `InInt(filedExpr string, enum ...int) Rule`       |
| `InUint(filedExpr string, enum ...uint) Rule`     |
| `InFloat(filedExpr string, enum ...float64) Rule` |



#### 格式型规则

格式型规则包括

| 名字                            |
| ------------------------------- |
| `Email(fieldExpr string) Rule`  |
| `Number(fieldExpr string) Rule` |
| `URL(fieldExpr string) Rule`    |
| `Ip(fieldExpr string) Rule`     |

等等



#### 自定义规则

除了以上已有规则，用户还可以通过实现`Rule`接口，实现特殊的规则。

下面的例子来自`customized_rule_test.go`, 来校验`fieldExpr`是否为空指针(同样的功能可以使用`Nil`规则实现)。

```go
type customizedRule struct {
	fieldExpr string

	name string
}

func (r customizedRule) Check(param interface{}) (bool, string) {
	exprValue, kind := fetchField(param, r.fieldExpr)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.fieldExpr)
	}
	if exprValue != nil {
		return false,
			fmt.Sprintf("[%s]:'%s' should not be nil", r.name, r.fieldExpr)
	}
	return true, ""
}

ch := NewChecker()
ch.Add(customRule, "invalid ptr")
```



## Checker

`Checker`是一个接口

- `Add(rule Rule, prompt string)`： 添加规则，和没有通过规则是的错误提示。
- `Check(param interface{}) (bool, string, string)`: 校验参数，依次返回是否通过校验，错误提示，错误日志。错误日志包含哪个字段没有通过哪个规则的信息。
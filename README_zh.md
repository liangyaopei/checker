`Checker`是一个Golang中参数校验的包，它可以替代
[gopkg.in/go-playground/validator.v10](https://godoc.org/gopkg.in/go-playground/validator.v10)。
`Checker`用于结构体或者非结构的参数校验，包括结构体中不同字段比较的校验，还提供自定义的校验规则。

# 安装
```
go get -u github.com/liangyaopei/checker
```

# 描述&例子


# 与validator.v10的比较
不同包下的结构体，校验标签不能定制
链表的例子，内嵌结构体，validator不使用
适用于非结构体的例子
标签与rule的比较,标签难以理解和记忆

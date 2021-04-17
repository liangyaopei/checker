package main

import (
	"fmt"

	"github.com/liangyaopei/checker"
)

type Item struct {
	Info  typeInfo
	Email string
}

type typeStr string
type typeInfo struct {
	Type        typeStr
	Range       []string
	Unit        string
	Granularity string
}

func getItems() []Item {
	items := []Item{
		{
			Email: "yaopei.liang@foxmail.com",
			Info: typeInfo{
				Type:        "last",
				Granularity: "day",
				Unit:        "7",
			},
		},
		{
			Email: "yaopei.liang@foxmail.com",
			Info: typeInfo{
				Type:  "range",
				Range: []string{"2020-01-01"},
			},
		},
		{
			Email: "yaopei.liang@foxmail.com",
			Info: typeInfo{
				Type:        "last",
				Granularity: "day",
				Unit:        "seven",
			},
		},
		{
			Email: "liangyaopei.com",
			Info: typeInfo{
				Type:        "last",
				Granularity: "day",
				Unit:        "7",
			},
		},
	}
	return items
}

func main() {
	rule := checker.And(
		checker.Email("Email").Prompt("Wrong email format"),
		checker.Field("Info",
			checker.Or(
				checker.And(
					checker.EqStr("Type", "range"),
					checker.Length("Range", 2, 2).Prompt("Range's length should be 2"),
					checker.Array("Range", checker.Time("", "2006-01-02")).
						Prompt("Range's element should be time format"),
				),
				checker.And(
					checker.EqStr("Type", "last"),
					checker.InStr("Granularity", "day", "week", "month"),
					checker.Number("Unit").Prompt("Unit should be number format"),
				),
			),
		),
	)

	validator := checker.NewChecker()
	validator.Add(rule, "wrong parameter")

	for idx, item := range getItems() {
		isValid, prompt, errMsg := validator.Check(item)
		fmt.Printf("item[%d].isValid:%v\n", idx, isValid)
		fmt.Printf("item[%d].prompt:%s\n", idx, prompt)
		fmt.Printf("item[%d].errMsg:%s\n", idx, errMsg)
		fmt.Println()
	}
}

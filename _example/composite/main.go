package main

import (
	"fmt"

	"github.com/liangyaopei/checker"
)

type arrayMap struct {
	Arr [][]map[int]float64
}

func main() {
	param := arrayMap{
		Arr: [][]map[int]float64{
			[]map[int]float64{
				map[int]float64{
					1: 1,
				},
				{
					2: 2,
				},
				{
					3: 3,
				},
			},
			{
				{
					4: 4,
				},
				{
					5: 5,
				},
				{
					6: 6,
				},
			},
			{
				{
					7: 7,
				},
				{
					8: 8,
				},
				{
					9: 9,
				},
			},
		}}
	rule := checker.Field("Arr",
		checker.Array("",
			checker.Array("",
				checker.Map("",
					checker.RangeInt("", 0, 20).Prompt("map的key在[0,20]"),
					checker.RangeFloat("", 0, 10).Prompt("map的value在[0,10]"),
				),
			),
		))

	validator := checker.NewChecker()
	validator.Add(rule, "wrong param")

	isValid, prompt, errMsg := validator.Check(param)
	fmt.Printf("isValid:%v\n", isValid)
	fmt.Printf("prompt:%s\n", prompt)
	fmt.Printf("errMsg:%s\n", errMsg)

	fmt.Println()

	param.Arr[0][0][-10] = -100
	isValid, prompt, errMsg = validator.Check(param)
	fmt.Printf("isValid:%v\n", isValid)
	fmt.Printf("prompt:%s\n", prompt)
	fmt.Printf("errMsg:%s\n", errMsg)

}

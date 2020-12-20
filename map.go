package checker

import (
	"fmt"
	"reflect"
)

// mapRule checks map's key and value
// keyRule or valueRule can be nil
type mapRule struct {
	fieldExpr string
	keyRule   Rule
	valueRule Rule

	name string
}

func (r mapRule) Check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.fieldExpr)
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is nil", r.name, r.fieldExpr)
	}
	if kind != reflect.Map {
		return false,
			fmt.Sprintf("[%s]:'%s' should be kind map,actual is %v",
				r.name, r.fieldExpr, kind)
	}

	refVal := reflect.ValueOf(exprValue)
	if r.keyRule != nil {
		for _, key := range refVal.MapKeys() {
			//fmt.Printf("key:%v", key)
			isValid, errMsg := r.keyRule.Check(key.Interface())
			if !isValid {
				return false,
					fmt.Sprintf("[%s]:'%s' key does not pass {%s}",
						r.name, r.fieldExpr, errMsg)
			}
		}
	}

	if r.valueRule != nil {
		for _, key := range refVal.MapKeys() {
			value := refVal.MapIndex(key)
			isValid, errMsg := r.valueRule.Check(value.Interface())
			if !isValid {
				return false,
					fmt.Sprintf("[%s]:'%s' value does not pass {%s}",
						r.name, r.fieldExpr, errMsg)
			}
		}
	}
	return true, ""
}

// NewMapRule returns mapRule
// keyRule or valueRule can be nil
// if no need to check key/value
func NewMapRule(fieldExpr string, keyRule Rule, valueRule Rule) Rule {
	return mapRule{
		fieldExpr: fieldExpr,
		keyRule:   keyRule,
		valueRule: valueRule,
		name:      "mapRule",
	}
}

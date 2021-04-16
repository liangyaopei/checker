package checker

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func fetchField(param interface{}, rule Rule) (interface{}, reflect.Kind) {
	completeExpr := rule.getCompleteFieldExpr()

	if rule.getCache() != nil {
		if valueKind, exist := rule.getCache()[completeExpr]; exist {
			return valueKind.value, valueKind.kind
		}
	}

	fieldExpr := rule.getFieldExpr()
	pValue := reflect.ValueOf(param)
	if fieldExpr == "" {
		return param, pValue.Kind()
	}

	exprs := strings.Split(fieldExpr, ".")
	for i := 0; i < len(exprs); i++ {
		expr := exprs[i]
		if pValue.Kind() == reflect.Ptr {
			pValue = pValue.Elem()
		}
		if !pValue.IsValid() || pValue.Kind() != reflect.Struct {
			return nil, reflect.Invalid
		}
		pValue = pValue.FieldByName(expr)
	}

	// last field is pointer
	if pValue.Kind() == reflect.Ptr {
		if pValue.IsNil() {
			return nil, reflect.Ptr
		}
		pValue = pValue.Elem()
	}

	if !pValue.IsValid() {
		return nil, reflect.Invalid
	}
	if cache := rule.getCache(); cache != nil {
		cache[completeExpr] = valueKindPair{
			value: pValue.Interface(),
			kind:  pValue.Kind(),
		}
	}
	//rule.getCache()[completeExpr] = valueKindPair{
	//	value: pValue.Interface(),
	//	kind:  pValue.Kind(),
	//}
	return pValue.Interface(), pValue.Kind()
}

func fetchFieldStr(param interface{}, rule Rule) (string, bool, string) {
	exprValue, kind := fetchField(param, rule)
	if kind == reflect.Invalid {
		return "", false,
			fmt.Sprintf("[%s]:%s cannot be found", rule.getName(), rule.getCompleteFieldExpr())
	}
	if exprValue == nil {
		return "", false,
			fmt.Sprintf("[%s]:%s is nil", rule.getName(), rule.getCompleteFieldExpr())
	}
	if kind != reflect.String {
		return "", false,
			fmt.Sprintf("[%s]:%s should be kind string,actual is %s",
				rule.getName(), rule.getCompleteFieldExpr(), kind.String())
	}
	res, ok := exprValue.(string)
	if !ok {
		res = reflect.ValueOf(exprValue).String()
	}
	return res, true, ""
}

func fetchFieldInt(param interface{}, rule Rule) (int, bool, string) {
	exprValue, kind := fetchField(param, rule)
	if kind == reflect.Invalid {
		return 0, false,
			fmt.Sprintf("[%s]:%s cannot be found", rule.getName(), rule.getCompleteFieldExpr())
	}
	if exprValue == nil {
		return 0, false,
			fmt.Sprintf("[%s]:%s is nil", rule.getName(), rule.getCompleteFieldExpr())
	}
	if kind != reflect.Int8 && kind != reflect.Int16 && kind != reflect.Int32 &&
		kind != reflect.Int64 && kind != reflect.Int {
		return 0, false,
			fmt.Sprintf("[%s]:%s should be kind int8/int16/int32/int64/int,actual is %s",
				rule.getName(), rule.getCompleteFieldExpr(), kind.String())
	}

	res, ok := exprValue.(int)
	if !ok {
		res = int(reflect.ValueOf(exprValue).Int())
	}
	return res, true, ""
}

func fetchFieldUint(param interface{}, rule Rule) (uint, bool, string) {
	exprValue, kind := fetchField(param, rule)
	if kind == reflect.Invalid {
		return 0, false,
			fmt.Sprintf("[%s]:%s cannot be found", rule.getName(), rule.getCompleteFieldExpr())
	}
	if exprValue == nil {
		return 0, false,
			fmt.Sprintf("[%s]:%s is nil", rule.getName(), rule.getCompleteFieldExpr())
	}
	if kind != reflect.Uint8 && kind != reflect.Uint16 && kind != reflect.Uint32 &&
		kind != reflect.Uint64 && kind != reflect.Uint {
		return 0, false,
			fmt.Sprintf("[%s]:%s should be kind uint8/uint16/uint32/uint64/uint,actual is %s",
				rule.getName(), rule.getCompleteFieldExpr(), kind.String())
	}
	res, ok := exprValue.(uint)
	if !ok {
		res = uint(reflect.ValueOf(exprValue).Uint())
	}
	return res, true, ""
}

func fetchFieldFloat(param interface{}, rule Rule) (float64, bool, string) {
	exprValue, kind := fetchField(param, rule)
	if kind == reflect.Invalid {
		return 0.0, false,
			fmt.Sprintf("[%s]:%s cannot be found", rule.getName(), rule.getCompleteFieldExpr())
	}
	if exprValue == nil {
		return 0.0, false,
			fmt.Sprintf("[%s]:%s is nil", rule.getName(), rule.getCompleteFieldExpr())
	}

	if kind != reflect.Float32 && kind != reflect.Float64 {
		return 0.0, false,
			fmt.Sprintf("[%s]:%s should be kind float32/float64,actual is %s",
				rule.getName(), rule.getCompleteFieldExpr(), kind.String())
	}
	res, ok := exprValue.(float64)
	if !ok {
		res = reflect.ValueOf(exprValue).Float()
	}
	return res, true, ""
}

func fetchFieldTime(param interface{}, rule Rule) (time.Time, bool, string) {
	exprValue, kind := fetchField(param, rule)
	if kind == reflect.Invalid {
		return time.Time{}, false,
			fmt.Sprintf("[%s]:%s cannot be found", rule.getName(), rule.getCompleteFieldExpr())
	}
	if exprValue == nil {
		return time.Time{}, false,
			fmt.Sprintf("[%s]:%s is nil", rule.getName(), rule.getCompleteFieldExpr())
	}
	tsVal, ok := exprValue.(time.Time)
	if !ok {
		return time.Time{}, false,
			fmt.Sprintf("[%s]:%s should be time.Time,actual is %v",
				rule.getName(), rule.getCompleteFieldExpr(), reflect.TypeOf(exprValue).String())
	}
	return tsVal, true, ""
}

func fetchFieldComparable(param interface{}, rule Rule) (Comparable, bool, string) {
	exprValue, kind := fetchField(param, rule)
	if kind == reflect.Invalid {
		return nil, false,
			fmt.Sprintf("[%s]:%s cannot be found", rule.getName(), rule.getCompleteFieldExpr())
	}
	if exprValue == nil {
		return nil, false,
			fmt.Sprintf("[%s]:%s is nil", rule.getName(), rule.getCompleteFieldExpr())
	}
	comp, ok := exprValue.(Comparable)
	if !ok {
		return nil, false,
			fmt.Sprintf("[%s]:%s should be type of checker.Comparable,actual is %v",
				rule.getName(), rule.getCompleteFieldExpr(), reflect.TypeOf(exprValue).String())
	}
	return comp, true, ""
}

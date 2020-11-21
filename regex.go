package checker

import (
	"fmt"
	"reflect"
	"regexp"
)

const (
	alphaRegexString        = "^[a-zA-Z]+$"
	alphaNumericRegexString = "^[a-zA-Z0-9]+$"
	numericRegexString      = "^[-+]?[0-9]+(?:\\.[0-9]+)?$"
	numberRegexString       = "^[0-9]+$"
	emailRegexString        = "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	iSBN10RegexString       = "^(?:[0-9]{9}X|[0-9]{10})$"
	iSBN13RegexString       = "^(?:(?:97(?:8|9))[0-9]{10})$"
)

func NewEmailRule(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(emailRegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   emailRegexString,
		regexObject: regexObject,
		ruleName:    "emailRule",
	}
}

func NewAlphaRule(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(alphaRegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   alphaRegexString,
		regexObject: regexObject,
		ruleName:    "alphaRule",
	}
}

func NewAlphaNumericRule(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(alphaNumericRegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   alphaNumericRegexString,
		regexObject: regexObject,
		ruleName:    "alphaNumericRule",
	}
}

func NewNumberRule(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(numberRegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   numberRegexString,
		regexObject: regexObject,
		ruleName:    "numberRule",
	}
}

func NewNumericRule(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(numericRegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   numericRegexString,
		regexObject: regexObject,
		ruleName:    "numericRule",
	}
}

func NewISBN10Rule(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(iSBN10RegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   iSBN10RegexString,
		regexObject: regexObject,
		ruleName:    "ISBN10Rule",
	}
}

func NewISBN13Rule(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(iSBN13RegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   iSBN13RegexString,
		regexObject: regexObject,
		ruleName:    "ISBN13Rule",
	}
}

type regexRule struct {
	fieldExpr   string
	regexExpr   string
	regexObject *regexp.Regexp
	ruleName    string
}

func (r regexRule) check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	ruleName := "regexRule"
	if r.ruleName != "" {
		ruleName = r.ruleName
	}

	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is nil", ruleName, r.fieldExpr)
	}
	if kind != reflect.String {
		return false,
			fmt.Sprintf("[%s]:'%s' should be kind string,actual is %v",
				ruleName, r.fieldExpr, kind)
	}
	exprValueStr := exprValue.(string)

	if !r.regexObject.MatchString(exprValueStr) {
		return false,
			fmt.Sprintf("[%s]:'%s' should macth regex expr %s,actual is %s",
				ruleName, r.fieldExpr, r.regexExpr, exprValueStr)
	}
	return true, ""
}

func NewRegexRule(fieldExpr string, regexExpr string) Rule {
	regexObject := regexp.MustCompile(regexExpr)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   regexExpr,
		regexObject: regexObject,
	}
}

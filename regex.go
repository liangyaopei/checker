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

	uRLEncodedRegexString  = `(%[A-Fa-f0-9]{2})`
	hTMLEncodedRegexString = `&#[x]?([0-9a-fA-F]{2})|(&gt)|(&lt)|(&quot)|(&amp)+[;]?`
	hTMLRegexString        = `<[/]?([a-zA-Z]+).*?>`

	hostnameRegexStringRFC952  = `^[a-zA-Z]([a-zA-Z0-9\-]+[\.]?)*[a-zA-Z0-9]$`                                      // https://tools.ietf.org/html/rfc952
	hostnameRegexStringRFC1123 = `^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?$` // accepts hostname starting with a digit https://tools.ietf.org/html/rfc1123
)

// NewEmailRule is the validation function for validating if the current field's value is a valid email address.
func NewEmailRule(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(emailRegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   emailRegexString,
		regexObject: regexObject,
		name:        "emailRule",
	}
}

// NewAlphaRule is the validation function for validating if the current field's value is a valid alpha value.
func NewAlphaRule(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(alphaRegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   alphaRegexString,
		regexObject: regexObject,
		name:        "alphaRule",
	}
}

// NewAlphaNumericRule is the validation function for validating if the current field's value is a valid alphanumeric value.
func NewAlphaNumericRule(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(alphaNumericRegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   alphaNumericRegexString,
		regexObject: regexObject,
		name:        "alphaNumericRule",
	}
}

// NewNumberRule is the validation function for validating if the current field's value is a valid number.
func NewNumberRule(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(numberRegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   numberRegexString,
		regexObject: regexObject,
		name:        "numberRule",
	}
}

// NewNumericRule is the validation function for validating if the current field's value is a valid numeric value.
func NewNumericRule(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(numericRegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   numericRegexString,
		regexObject: regexObject,
		name:        "numericRule",
	}
}

// NewURLEncodedRule is the validation function for validating if the current field's value is a valid encoded URL.
func NewURLEncodedRule(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(uRLEncodedRegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   uRLEncodedRegexString,
		regexObject: regexObject,
		name:        "URLEncodedRule",
	}
}

// NewHTMLEncodedRule is the validation function for validating if the current field's value is a valid encoded HTML.
func NewHTMLEncodedRule(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(hTMLEncodedRegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   hTMLEncodedRegexString,
		regexObject: regexObject,
		name:        "HTMLEncodedRule",
	}
}

// NewHTMLRule is the validation function for validating if the current field's value is a valid HTML.
func NewHTMLRule(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(hTMLRegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   hTMLRegexString,
		regexObject: regexObject,
		name:        "HTMLRule",
	}
}

// NewHostNameRule is the validation function for validating if the current field's value is a valid RFC953 hostname.
func NewHostNameRule(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(hostnameRegexStringRFC952)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   hostnameRegexStringRFC952,
		regexObject: regexObject,
		name:        "HostNameRule",
	}
}

// NewHostNameRFC1123Rule is the validation function for validating if the current field's value is a valid RFC1123 hostname.
func NewHostNameRFC1123Rule(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(hostnameRegexStringRFC1123)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   hostnameRegexStringRFC1123,
		regexObject: regexObject,
		name:        "HostNameRFC1123Rule",
	}
}

type regexRule struct {
	fieldExpr   string
	regexExpr   string
	regexObject *regexp.Regexp
	name        string
}

func (r regexRule) Check(param interface{}) (bool, string) {
	exprValue, kind := fetchFieldInStruct(param, r.fieldExpr)
	ruleName := "regexRule"
	if r.name != "" {
		ruleName = r.name
	}

	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.name, r.fieldExpr)
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is nil", r.name, r.fieldExpr)
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

// NewRegexRule is the validation function for validating if the current field's value satisfies regex pattern.
func NewRegexRule(fieldExpr string, regexExpr string) Rule {
	regexObject := regexp.MustCompile(regexExpr)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   regexExpr,
		regexObject: regexObject,
	}
}

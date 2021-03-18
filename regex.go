package checker

import (
	"fmt"
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

// Email is the validation function for validating if the current field's value is a valid email address.
func Email(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(emailRegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   emailRegexString,
		regexObject: regexObject,
		name:        "emailRule",
	}
}

// Alpha is the validation function for validating if the current field's value is a valid alpha value.
func Alpha(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(alphaRegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   alphaRegexString,
		regexObject: regexObject,
		name:        "alphaRule",
	}
}

// AlphaNumeric is the validation function for validating if the current field's value is a valid alphanumeric value.
func AlphaNumeric(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(alphaNumericRegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   alphaNumericRegexString,
		regexObject: regexObject,
		name:        "alphaNumericRule",
	}
}

// Number is the validation function for validating if the current field's value is a valid number.
func Number(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(numberRegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   numberRegexString,
		regexObject: regexObject,
		name:        "numberRule",
	}
}

// Numeric is the validation function for validating if the current field's value is a valid numeric value.
func Numeric(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(numericRegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   numericRegexString,
		regexObject: regexObject,
		name:        "numericRule",
	}
}

// URLEncoded is the validation function for validating if the current field's value is a valid encoded URL.
func URLEncoded(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(uRLEncodedRegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   uRLEncodedRegexString,
		regexObject: regexObject,
		name:        "URLEncodedRule",
	}
}

// HTMLEncoded is the validation function for validating if the current field's value is a valid encoded HTML.
func HTMLEncoded(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(hTMLEncodedRegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   hTMLEncodedRegexString,
		regexObject: regexObject,
		name:        "HTMLEncodedRule",
	}
}

// HTML is the validation function for validating if the current field's value is a valid HTML.
func HTML(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(hTMLRegexString)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   hTMLRegexString,
		regexObject: regexObject,
		name:        "HTMLRule",
	}
}

// HostName is the validation function for validating if the current field's value is a valid RFC953 hostname.
func HostName(fieldExpr string) Rule {
	regexObject := regexp.MustCompile(hostnameRegexStringRFC952)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   hostnameRegexStringRFC952,
		regexObject: regexObject,
		name:        "HostNameRule",
	}
}

// HostNameRFC1123 is the validation function for validating if the current field's value is a valid RFC1123 hostname.
func HostNameRFC1123(fieldExpr string) Rule {
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
	ruleName := "regexRule"
	if r.name != "" {
		ruleName = r.name
	}
	exprValueStr, isValid, errMsg := fetchFieldStr(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}

	if !r.regexObject.MatchString(exprValueStr) {
		return false,
			fmt.Sprintf("[%s]:'%s' should macth regex expr %s,actual is %s",
				ruleName, r.fieldExpr, r.regexExpr, exprValueStr)
	}
	return true, ""
}

// Regex is the validation function for validating if the current field's value satisfies regex pattern.
func Regex(fieldExpr string, regexExpr string) Rule {
	regexObject := regexp.MustCompile(regexExpr)
	return regexRule{
		fieldExpr:   fieldExpr,
		regexExpr:   regexExpr,
		regexObject: regexObject,
	}
}

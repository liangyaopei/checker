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

var (
	emailRegexObject           = regexp.MustCompile(emailRegexString)
	alphaRegexObject           = regexp.MustCompile(alphaRegexString)
	alphaNumericRegexObject    = regexp.MustCompile(alphaNumericRegexString)
	numberRegexObject          = regexp.MustCompile(numberRegexString)
	numericRegexObject         = regexp.MustCompile(numericRegexString)
	uRLEncodedRegexObject      = regexp.MustCompile(uRLEncodedRegexString)
	hTMLEncodedRegexObject     = regexp.MustCompile(hTMLEncodedRegexString)
	hTMLRegexObject            = regexp.MustCompile(hTMLRegexString)
	hostnameRegexObjectFC952   = regexp.MustCompile(hostnameRegexStringRFC952)
	hostnameRegexObjectRFC1123 = regexp.MustCompile(hostnameRegexStringRFC1123)
)

// Email is the validation function for validating if the current field's value is a valid email address.
func Email(fieldExpr string) *regexRule {
	return &regexRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "Email",
		},
		emailRegexString,
		emailRegexObject,
	}
}

// Alpha is the validation function for validating if the current field's value is a valid alpha value.
func Alpha(fieldExpr string) *regexRule {
	return &regexRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "Alpha",
		},
		alphaRegexString,
		alphaRegexObject,
	}
}

// AlphaNumeric is the validation function for validating if the current field's value is a valid alphanumeric value.
func AlphaNumeric(fieldExpr string) *regexRule {
	return &regexRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "AlphaNumeric",
		},
		alphaNumericRegexString,
		alphaNumericRegexObject,
	}
}

// Number is the validation function for validating if the current field's value is a valid number.
func Number(fieldExpr string) *regexRule {
	return &regexRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "Number",
		},
		numberRegexString,
		numberRegexObject,
	}
}

// Numeric is the validation function for validating if the current field's value is a valid numeric value.
func Numeric(fieldExpr string) *regexRule {
	return &regexRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "Numeric",
		},
		numericRegexString,
		numericRegexObject,
	}
}

// URLEncoded is the validation function for validating if the current field's value is a valid encoded URL.
func URLEncoded(fieldExpr string) *regexRule {
	return &regexRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "URLEncoded",
		},
		uRLEncodedRegexString,
		uRLEncodedRegexObject,
	}
}

// HTMLEncoded is the validation function for validating if the current field's value is a valid encoded HTML.
func HTMLEncoded(fieldExpr string) *regexRule {
	return &regexRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "HTMLEncoded",
		},
		hTMLEncodedRegexString,
		hTMLEncodedRegexObject,
	}
}

// HTML is the validation function for validating if the current field's value is a valid HTML.
func HTML(fieldExpr string) *regexRule {
	return &regexRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "HTML",
		},
		hTMLRegexString,
		hTMLRegexObject,
	}
}

// HostName is the validation function for validating if the current field's value is a valid RFC953 hostname.
func HostName(fieldExpr string) *regexRule {
	return &regexRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "HostName",
		},
		hostnameRegexStringRFC952,
		hostnameRegexObjectFC952,
	}
}

// HostNameRFC1123 is the validation function for validating if the current field's value is a valid RFC1123 hostname.
func HostNameRFC1123(fieldExpr string) Rule {
	return &regexRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "HostNameRFC1123",
		},
		hostnameRegexStringRFC1123,
		hostnameRegexObjectRFC1123,
	}
}

type regexRule struct {
	baseRule
	regexExpr   string
	regexObject *regexp.Regexp
}

func (r *regexRule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r regexRule) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := fetchFieldStr(param, &r)
	if !isValid {
		return false, errMsg
	}

	if !r.regexObject.MatchString(exprValueStr) {
		return false,
			fmt.Sprintf("[%s]:%s should match %s format, actual is %s",
				r.name, r.getCompleteFieldExpr(), r.name, exprValueStr)
	}
	return true, ""
}

// Regex is the validation function for validating if the current field's value satisfies regex pattern.
func Regex(fieldExpr string, regexExpr string) *regexRule {
	regexObject := regexp.MustCompile(regexExpr)
	return &regexRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "Regex",
		},
		regexExpr,
		regexObject,
	}
}

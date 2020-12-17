package checker

import (
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"reflect"
	"strings"
)

type urlRule struct {
	fieldExpr string

	name string
}

func (r urlRule) Check(param interface{}) (bool, string) {
	exprValStr, isValid, errMsg := getStrField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}

	var i int
	s := exprValStr
	// checks needed as of Go 1.6 because of change https://github.com/golang/go/commit/617c93ce740c3c3cc28cdd1a0d712be183d0b328#diff-6c2d018290e298803c0c9419d8739885L195
	// emulate browser and strip the '#' suffix prior to validation. see issue-#237
	if i = strings.Index(s, "#"); i > -1 {
		s = s[:i]
	}

	if len(s) == 0 {
		return false,
			fmt.Sprintf("[%s]:'%s' does not satisfy URL format",
				r.name, r.fieldExpr)
	}

	endPoint, err := url.ParseRequestURI(s)
	if err != nil || endPoint.Scheme == "" {
		return false,
			fmt.Sprintf("[%s]:'%s' does not satisfy url format,actual value is %s",
				r.name, r.fieldExpr, exprValStr)
	}
	return true, ""
}

// NewURLRule is the validation function for validating if the current field's value is a valid URL.
func NewURLRule(fieldExpr string) Rule {
	return urlRule{
		fieldExpr: fieldExpr,
		name:      "urlRule",
	}
}

type uriRule struct {
	fieldExpr string

	name string
}

func (r uriRule) Check(param interface{}) (bool, string) {
	exprValStr, isValid, errMsg := getStrField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}

	var i int
	s := exprValStr
	// checks needed as of Go 1.6 because of change https://github.com/golang/go/commit/617c93ce740c3c3cc28cdd1a0d712be183d0b328#diff-6c2d018290e298803c0c9419d8739885L195
	// emulate browser and strip the '#' suffix prior to validation. see issue-#237
	if i = strings.Index(s, "#"); i > -1 {
		s = s[:i]
	}

	if len(s) == 0 {
		return false,
			fmt.Sprintf("[%s]:'%s' does not satisfy URI format,actual value is %s",
				r.name, r.fieldExpr, exprValStr)
	}

	_, err := url.ParseRequestURI(s)
	if err != nil {
		return false,
			fmt.Sprintf("[%s]:'%s' does not satisfy url format",
				r.name, r.fieldExpr)
	}
	return true, ""
}

// NewURIRule is the validation function for validating if the current field's value is a valid URI.
func NewURIRule(fieldExpr string) Rule {
	return uriRule{
		fieldExpr: fieldExpr,
		name:      "uriRule",
	}
}

type ipv4Rule struct {
	fieldExpr string

	name string
}

func (r ipv4Rule) Check(param interface{}) (bool, string) {
	exprValStr, isValid, errMsg := getStrField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}

	ip := net.ParseIP(exprValStr)
	if ip == nil || ip.To4() == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' does not satisfy ipv4 format,actual value is %s",
				r.name, r.fieldExpr, exprValStr)
	}
	return true, ""
}

// NewIPv4Rule is the validation function for validating if a value is a valid v4 IP address.
func NewIPv4Rule(fieldExpr string) Rule {
	return ipv4Rule{
		fieldExpr: fieldExpr,
		name:      "ipv4Rule",
	}
}

type ipv6Rule struct {
	fieldExpr string

	name string
}

func (r ipv6Rule) Check(param interface{}) (bool, string) {
	exprValStr, isValid, errMsg := getStrField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	ip := net.ParseIP(exprValStr)

	if ip == nil || ip.To4() == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' does not satisfy ipv6 format,actual value is %s",
				r.name, r.fieldExpr, exprValStr)
	}
	return true, ""
}

// NewIPv6Rule is the validation function for validating if the field's value is a valid v6 IP address.
func NewIPv6Rule(fieldExpr string) Rule {
	return ipv6Rule{
		fieldExpr: fieldExpr,
		name:      "ipv6Rule",
	}
}

type ipRule struct {
	fieldExpr string

	name string
}

func (r ipRule) Check(param interface{}) (bool, string) {
	exprValStr, isValid, errMsg := getStrField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	ip := net.ParseIP(exprValStr)

	if ip == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' does not satisfy ip format,actual value is %s",
				r.name, r.fieldExpr, exprValStr)
	}
	return true, ""
}

// NewIPRule is the validation function for validating if the field's value is a valid v4 or v6 IP address.
func NewIPRule(fieldExpr string) Rule {
	return ipRule{
		fieldExpr: fieldExpr,
		name:      "ipRule",
	}
}

type startsWithRule struct {
	fieldExpr string
	prefix    string

	name string
}

func (r startsWithRule) Check(param interface{}) (bool, string) {
	exprValStr, isValid, errMsg := getStrField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}

	if !strings.HasPrefix(exprValStr, r.prefix) {
		return false,
			fmt.Sprintf("[%s]:'%s' does not has prefix %s",
				r.name, r.fieldExpr, r.prefix)
	}
	return true, ""
}

// NewStartsWithRule is the validation function for validating that the field's value starts with the text specified within the param.
func NewStartsWithRule(fieldExpr string, prefix string) Rule {
	return startsWithRule{
		fieldExpr: fieldExpr,
		prefix:    prefix,
		name:      "startsWithRule",
	}
}

type endsWithRule struct {
	fieldExpr string
	suffix    string

	name string
}

func (r endsWithRule) Check(param interface{}) (bool, string) {
	exprValStr, isValid, errMsg := getStrField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}

	if !strings.HasSuffix(exprValStr, r.suffix) {
		return false,
			fmt.Sprintf("[%s]:'%s' does not has suffix %s",
				r.name, r.fieldExpr, r.suffix)
	}
	return true, ""
}

// NewEndsWithRule is the validation function for validating that the field's value ends with the text specified within the param.
func NewEndsWithRule(fieldExpr string, suffix string) Rule {
	return endsWithRule{
		fieldExpr: fieldExpr,
		suffix:    suffix,
		name:      "startsWithRule",
	}
}

type isJSONRule struct {
	fieldExpr string

	name string
}

func (r isJSONRule) Check(param interface{}) (bool, string) {
	exprValStr, isValid, errMsg := getStrField(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if !json.Valid([]byte(exprValStr)) {
		return false,
			fmt.Sprintf("[%s]:'%s' does not satisfy JSON fromat",
				r.name, r.fieldExpr)
	}
	return true, ""
}

// NewIsJSONRule is the validation function for validating if the current field's value is a valid json string.
func NewIsJSONRule(fieldExpr string) Rule {
	return isJSONRule{
		fieldExpr: fieldExpr,
		name:      "isJSONRule",
	}
}

func getStrField(param interface{}, fieldExpr string, name string) (string, bool, string) {
	exprValue, kind := fetchFieldInStruct(param, fieldExpr)
	if kind == reflect.Invalid {
		return "", false,
			fmt.Sprintf("[%s]:'%s' cannot be found", name, fieldExpr)
	}
	if exprValue == nil {
		return "", false,
			fmt.Sprintf("[%s]:'%s' is nil", name, fieldExpr)
	}
	if kind != reflect.String {
		return "", false,
			fmt.Sprintf("[%s]:'%s' should be kind string,actual is %v",
				name, fieldExpr, kind)
	}
	return exprValue.(string), true, ""
}

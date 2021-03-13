package checker

import (
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	iSBN10RegexString = "^(?:[0-9]{9}X|[0-9]{10})$"
	iSBN13RegexString = "^(?:(?:97(?:8|9))[0-9]{10})$"
)

var (
	iSBN10Regex = regexp.MustCompile(iSBN10RegexString)
	iSBN13Regex = regexp.MustCompile(iSBN13RegexString)
)

type urlRule struct {
	fieldExpr string

	name string
}

func (r urlRule) Check(param interface{}) (bool, string) {
	exprValStr, isValid, errMsg := fetchFieldStr(param, r.fieldExpr, r.name)
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
	exprValStr, isValid, errMsg := fetchFieldStr(param, r.fieldExpr, r.name)
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
	exprValStr, isValid, errMsg := fetchFieldStr(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}

	ip := net.ParseIP(exprValStr)
	isValidIPv4 := ip != nil && ip.To4() != nil
	if !isValidIPv4 {
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
	exprValStr, isValid, errMsg := fetchFieldStr(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	ip := net.ParseIP(exprValStr)

	isValidIPv6 := ip != nil && ip.To4() == nil
	if !isValidIPv6 {
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
	exprValStr, isValid, errMsg := fetchFieldStr(param, r.fieldExpr, r.name)
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
	exprValStr, isValid, errMsg := fetchFieldStr(param, r.fieldExpr, r.name)
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
	exprValStr, isValid, errMsg := fetchFieldStr(param, r.fieldExpr, r.name)
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
	exprValStr, isValid, errMsg := fetchFieldStr(param, r.fieldExpr, r.name)
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

type isDirRule struct {
	fieldExpr string

	name string
}

func (r isDirRule) Check(param interface{}) (bool, string) {
	exprValStr, isValid, errMsg := fetchFieldStr(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	fileInfo, err := os.Stat(exprValStr)
	if err != nil || !fileInfo.IsDir() {
		return false,
			fmt.Sprintf("[%s]:'%s' does not satisfy dir fromat",
				r.name, r.fieldExpr)
	}
	return true, ""
}

// NewIsDirRule is the validation function for validating if the current field's value is a valid directory.
func NewIsDirRule(fieldExpr string) Rule {
	return isDirRule{
		fieldExpr: fieldExpr,
		name:      "isDirRule",
	}
}

type isDatetimeRule struct {
	fieldExpr string
	name      string

	layout string
}

func (r isDatetimeRule) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := fetchFieldStr(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	_, err := time.Parse(r.layout, exprValueStr)
	if err != nil {
		return false,
			fmt.Sprintf("[%s]:'%s' should be format %s,actual is %s",
				r.name, r.fieldExpr, r.layout, exprValueStr)
	}
	return true, ""
}

// NewIsDatetimeRule is the validation function for validating if the current field's value is a valid datetime string.
func NewIsDatetimeRule(fieldExpr string, layout string) Rule {
	return isDatetimeRule{
		fieldExpr: fieldExpr,
		name:      "isDatetimeRule",
		layout:    layout,
	}
}

type iSBN10Rule struct {
	fieldExpr string
	name      string
}

func (r iSBN10Rule) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := fetchFieldStr(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if !isISBN10(exprValueStr) {
		return false,
			fmt.Sprintf("[%s]:'%s' does not satisfy ISBN10 fromat",
				r.name, r.fieldExpr)
	}
	return true, ""
}

// NewIsISBN10Rule is the validation function for validating if the field's value is a valid v10 ISBN.
func NewIsISBN10Rule(fieldExpr string) Rule {
	return iSBN10Rule{
		fieldExpr: fieldExpr,
		name:      "IsISBN10Rule",
	}
}

type iSBN13Rule struct {
	fieldExpr string
	name      string
}

func (r iSBN13Rule) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := fetchFieldStr(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if !isISBN13(exprValueStr) {
		return false,
			fmt.Sprintf("[%s]:'%s' does not satisfy ISBN13 fromat",
				r.name, r.fieldExpr)
	}
	return true, ""
}

// NewIsISBN13Rule is the validation function for validating if the field's value is a valid v13 ISBN.
func NewIsISBN13Rule(fieldExpr string) Rule {
	return iSBN13Rule{
		fieldExpr: fieldExpr,
		name:      "IsISBN13Rule",
	}
}

type iSBNRule struct {
	fieldExpr string
	name      string
}

func (r iSBNRule) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := fetchFieldStr(param, r.fieldExpr, r.name)
	if !isValid {
		return false, errMsg
	}
	if !isISBN10(exprValueStr) && !isISBN13(exprValueStr) {
		return false,
			fmt.Sprintf("[%s]:'%s' does not satisfy ISBN fromat",
				r.name, r.fieldExpr)
	}
	return true, ""
}

// NewIsISBNRule is the validation function for validating if the field's value is a valid v10 or v13 ISBN.
func NewIsISBNRule(fieldExpr string) Rule {
	return iSBNRule{
		fieldExpr: fieldExpr,
		name:      "IsISBNRule",
	}
}

func isISBN10(exprValueStr string) bool {

	s := strings.Replace(strings.Replace(exprValueStr, "-", "", 3), " ", "", 3)

	if !iSBN10Regex.MatchString(s) {
		return false
	}

	var checksum int32
	var i int32

	for i = 0; i < 9; i++ {
		checksum += (i + 1) * int32(s[i]-'0')
	}

	if s[9] == 'X' {
		checksum += 10 * 10
	} else {
		checksum += 10 * int32(s[9]-'0')
	}

	return checksum%11 == 0
}

func isISBN13(exprValueStr string) bool {

	s := strings.Replace(strings.Replace(exprValueStr, "-", "", 4), " ", "", 4)

	if !iSBN13Regex.MatchString(s) {
		return false
	}

	var checksum int32
	var i int32

	factor := []int32{1, 3}

	for i = 0; i < 12; i++ {
		checksum += factor[i%2] * int32(s[i]-'0')
	}

	return (int32(s[12]-'0'))-((10-(checksum%10))%10) == 0
}

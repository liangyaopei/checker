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
	baseRule
}

func (r *urlRule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r urlRule) Check(param interface{}) (bool, string) {
	exprValStr, isValid, errMsg := fetchFieldStr(param, &r)
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
			fmt.Sprintf("[%s]:%s does not satisfy url format,actual value is %s",
				r.name, r.getCompleteFieldExpr(), exprValStr)
	}
	return true, ""
}

// URL is the validation function for validating if the current field's value is a valid URL.
func URL(fieldExpr string) *urlRule {
	return &urlRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "URL",
		},
	}
}

type uriRule struct {
	baseRule
}

func (r *uriRule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r uriRule) Check(param interface{}) (bool, string) {
	exprValStr, isValid, errMsg := fetchFieldStr(param, &r)
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
			fmt.Sprintf("[%s]:%s does not satisfy URI format,actual value is %s",
				r.name, r.getCompleteFieldExpr(), exprValStr)
	}

	_, err := url.ParseRequestURI(s)
	if err != nil {
		return false,
			fmt.Sprintf("[%s]:%s does not satisfy url format",
				r.name, r.getCompleteFieldExpr())
	}
	return true, ""
}

// URI is the validation function for validating if the current field's value is a valid URI.
func URI(fieldExpr string) *urlRule {
	return &urlRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "URI",
		},
	}
}

type ipv4Rule struct {
	baseRule
}

func (r *ipv4Rule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r ipv4Rule) Check(param interface{}) (bool, string) {
	exprValStr, isValid, errMsg := fetchFieldStr(param, &r)
	if !isValid {
		return false, errMsg
	}

	ip := net.ParseIP(exprValStr)
	isValidIPv4 := ip != nil && ip.To4() != nil
	if !isValidIPv4 {
		return false,
			fmt.Sprintf("[%s]:%s does not satisfy ipv4 format,actual value is %s",
				r.name, r.getCompleteFieldExpr(), exprValStr)
	}
	return true, ""
}

// IPv4 is the validation function for validating if a value is a valid v4 IP address.
func IPv4(fieldExpr string) *ipv4Rule {
	return &ipv4Rule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "IPv4",
		},
	}
}

type ipv6Rule struct {
	baseRule
}

func (r *ipv6Rule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r ipv6Rule) Check(param interface{}) (bool, string) {
	exprValStr, isValid, errMsg := fetchFieldStr(param, &r)
	if !isValid {
		return false, errMsg
	}
	ip := net.ParseIP(exprValStr)

	isValidIPv6 := ip != nil && ip.To4() == nil
	if !isValidIPv6 {
		return false,
			fmt.Sprintf("[%s]:%s does not satisfy ipv6 format,actual value is %s",
				r.name, r.getCompleteFieldExpr(), exprValStr)
	}
	return true, ""
}

// IPv6 is the validation function for validating if the field's value is a valid v6 IP address.
func IPv6(fieldExpr string) *ipv6Rule {
	return &ipv6Rule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "IPv6",
		},
	}
}

type ipRule struct {
	baseRule
}

func (r *ipRule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r ipRule) Check(param interface{}) (bool, string) {
	exprValStr, isValid, errMsg := fetchFieldStr(param, &r)
	if !isValid {
		return false, errMsg
	}
	ip := net.ParseIP(exprValStr)

	if ip == nil {
		return false,
			fmt.Sprintf("[%s]:%s does not satisfy ip format,actual value is %s",
				r.name, r.getCompleteFieldExpr(), exprValStr)
	}
	return true, ""
}

// Ip is the validation function for validating if the field's value is a valid v4 or v6 IP address.
func Ip(fieldExpr string) *ipRule {
	return &ipRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "fieldExpr",
		},
	}
}

type startsWithRule struct {
	baseRule
	prefix string
}

func (r *startsWithRule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r startsWithRule) Check(param interface{}) (bool, string) {
	exprValStr, isValid, errMsg := fetchFieldStr(param, &r)
	if !isValid {
		return false, errMsg
	}
	if !strings.HasPrefix(exprValStr, r.prefix) {
		return false,
			fmt.Sprintf("[%s]:%s does not has prefix %s",
				r.name, r.getCompleteFieldExpr(), r.prefix)
	}
	return true, ""
}

// StartsWith is the validation function for validating that the field's value starts with the text specified within the param.
func StartsWith(fieldExpr string, prefix string) *startsWithRule {
	return &startsWithRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "StartsWith",
		},
		prefix,
	}
}

type endsWithRule struct {
	baseRule
	suffix string
}

func (r *endsWithRule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r endsWithRule) Check(param interface{}) (bool, string) {
	exprValStr, isValid, errMsg := fetchFieldStr(param, &r)
	if !isValid {
		return false, errMsg
	}

	if !strings.HasSuffix(exprValStr, r.suffix) {
		return false,
			fmt.Sprintf("[%s]:%s does not has suffix %s",
				r.name, r.getCompleteFieldExpr(), r.suffix)
	}
	return true, ""
}

// EndsWith is the validation function for validating that the field's value ends with the text specified within the param.
func EndsWith(fieldExpr string, suffix string) *endsWithRule {
	return &endsWithRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "EndsWith",
		},
		suffix,
	}
}

type jsonRule struct {
	baseRule
}

func (r *jsonRule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r jsonRule) Check(param interface{}) (bool, string) {
	exprValStr, isValid, errMsg := fetchFieldStr(param, &r)
	if !isValid {
		return false, errMsg
	}
	if !json.Valid([]byte(exprValStr)) {
		return false,
			fmt.Sprintf("[%s]:%s does not satisfy JSON fromat",
				r.name, r.getCompleteFieldExpr())
	}
	return true, ""
}

// JSON is the validation function for validating if the current field's value is a valid json string.
func JSON(fieldExpr string) *jsonRule {
	return &jsonRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "isJSON",
		},
	}
}

type dirRule struct {
	baseRule
}

func (r *dirRule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r dirRule) Check(param interface{}) (bool, string) {
	exprValStr, isValid, errMsg := fetchFieldStr(param, &r)
	if !isValid {
		return false, errMsg
	}
	fileInfo, err := os.Stat(exprValStr)
	if err != nil || !fileInfo.IsDir() {
		return false,
			fmt.Sprintf("[%s]:%s does not satisfy dir fromat",
				r.name, r.getCompleteFieldExpr())
	}
	return true, ""
}

// Dir is the validation function for validating if the current field's value is a valid directory.
func Dir(fieldExpr string) *dirRule {
	return &dirRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "IsDir",
		},
	}
}

type timeRule struct {
	baseRule
	layout string
}

func (r *timeRule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r timeRule) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := fetchFieldStr(param, &r)
	if !isValid {
		return false, errMsg
	}
	_, err := time.Parse(r.layout, exprValueStr)
	if err != nil {
		return false,
			fmt.Sprintf("[%s]:%s should be format %s,actual is %s",
				r.name, r.getCompleteFieldExpr(), r.layout, exprValueStr)
	}
	return true, ""
}

// Time is the validation function for validating if the current field's value is a valid datetime string.
func Time(fieldExpr string, layout string) *timeRule {
	return &timeRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "Time",
		},
		layout,
	}
}

type iSBN10Rule struct {
	baseRule
}

func (r *iSBN10Rule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r iSBN10Rule) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := fetchFieldStr(param, &r)
	if !isValid {
		return false, errMsg
	}
	if !isISBN10(exprValueStr) {
		return false,
			fmt.Sprintf("[%s]:%s does not satisfy ISBN10 fromat",
				r.name, r.getCompleteFieldExpr())
	}
	return true, ""
}

// ISBN10 is the validation function for validating if the field's value is a valid v10 ISBN.
func ISBN10(fieldExpr string) *iSBN10Rule {
	return &iSBN10Rule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "ISBN10",
		},
	}
}

type iSBN13Rule struct {
	baseRule
}

func (r *iSBN13Rule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r iSBN13Rule) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := fetchFieldStr(param, &r)
	if !isValid {
		return false, errMsg
	}
	if !isISBN13(exprValueStr) {
		return false,
			fmt.Sprintf("[%s]:%s does not satisfy ISBN13 fromat",
				r.name, r.getCompleteFieldExpr())
	}
	return true, ""
}

// ISBN13 is the validation function for validating if the field's value is a valid v13 ISBN.
func ISBN13(fieldExpr string) *iSBN13Rule {
	return &iSBN13Rule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "ISBN13",
		},
	}
}

type iSBNRule struct {
	baseRule
}

func (r *iSBNRule) Prompt(prompt string) Rule {
	r.prompt = prompt
	return r
}

func (r iSBNRule) Check(param interface{}) (bool, string) {
	exprValueStr, isValid, errMsg := fetchFieldStr(param, &r)
	if !isValid {
		return false, errMsg
	}
	if !isISBN10(exprValueStr) && !isISBN13(exprValueStr) {
		return false,
			fmt.Sprintf("[%s]:%s does not satisfy ISBN fromat",
				r.name, r.getCompleteFieldExpr())
	}
	return true, ""
}

// ISBN is the validation function for validating if the field's value is a valid v10 or v13 ISBN.
func ISBN(fieldExpr string) *iSBNRule {
	return &iSBNRule{
		baseRule{
			fieldExpr: fieldExpr,
			name:      "ISBN",
		},
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

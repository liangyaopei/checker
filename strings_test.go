package checker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type param struct {
	URL        string
	IPv4       string
	IPv6       string
	HostName   string
	StartsWith string
	EndsWith   string
	ISBN10     string
	ISBN13     string
	Datetime   string
	JsonStr    string
}

func TestStringsRule(t *testing.T) {
	rChecker := NewChecker()

	urlRule := URL("URL")
	rChecker.Add(urlRule, "invalid url")

	ipv4Rule := IPv4("IPv4")
	rChecker.Add(ipv4Rule, "invalid ipv4")

	ipv6Rule := IPv6("IPv6")
	rChecker.Add(ipv6Rule, "invalid ipv6")

	hostnameRule := HostName("HostName")
	rChecker.Add(hostnameRule, "invalid hostname")

	startsWithRule := StartsWith("StartsWith", "Github")
	rChecker.Add(startsWithRule, "invalid startswith")

	ensWithRule := EndsWith("EndsWith", "lang")
	rChecker.Add(ensWithRule, "invalid endswith")

	iSBN10Rule := ISBN10("ISBN10")
	rChecker.Add(iSBN10Rule, "invalid isbn10")

	iSBN13Rule := ISBN13("ISBN13")
	rChecker.Add(iSBN13Rule, "invalid isbn13")

	iSBNRule := ISBN("ISBN13")
	rChecker.Add(iSBNRule, "invalid isbn")

	jsonRule := JSON("JsonStr")
	rChecker.Add(jsonRule, "invalid JsonStr")

	datetimeRule := Time("Datetime", "2006-01-02")
	rChecker.Add(datetimeRule, "invalid datetime")

	p := param{
		URL:        "https://github.com/",
		IPv4:       "14.215.177.38",
		IPv6:       "2001:cdba:0000:0000:0000:0000:3257:9652",
		HostName:   "www.liangyaopei.com",
		StartsWith: "Github",
		EndsWith:   "Golang",
		ISBN10:     "1-61729-085-8",
		ISBN13:     "978-3-16-148410-0",
		Datetime:   "2020-12-20",
		JsonStr:    `{"a":1}`,
	}
	isValid, _, _ := rChecker.Check(p)
	assert.Equal(t, isValid, true, "faield TestStringsRule")
}

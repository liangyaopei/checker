package _checker

import (
	"testing"

	"github.com/liangyaopei/checker"
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
	Directory  string
	Datetime   string
}

func TestStringsRule(t *testing.T) {
	rChecker := checker.NewChecker()

	urlRule := checker.NewURLRule("URL")
	rChecker.Add(urlRule, "invalid url")

	ipv4Rule := checker.NewIPv4Rule("IPv4")
	rChecker.Add(ipv4Rule, "invalid ipv4")

	ipv6Rule := checker.NewIPv6Rule("IPv6")
	rChecker.Add(ipv6Rule, "invalid ipv6")

	hostnameRule := checker.NewHostNameRule("HostName")
	rChecker.Add(hostnameRule, "invalid hostname")

	startsWithRule := checker.NewStartsWithRule("StartsWith", "Github")
	rChecker.Add(startsWithRule, "invalid startswith")

	ensWithRule := checker.NewEndsWithRule("EndsWith", "lang")
	rChecker.Add(ensWithRule, "invalid endswith")

	iSBN10Rule := checker.NewIsISBN10Rule("ISBN10")
	rChecker.Add(iSBN10Rule, "invalid isbn10")

	iSBN13Rule := checker.NewIsISBN13Rule("ISBN13")
	rChecker.Add(iSBN13Rule, "invalid isbn13")

	iSBNRule := checker.NewIsISBNRule("ISBN13")
	rChecker.Add(iSBNRule, "invalid isbn")

	isDirRule := checker.NewIsDirRule("Directory")
	rChecker.Add(isDirRule, "invalid directory")

	datetimeRule := checker.NewIsDatetimeRule("Datetime", "2006-01-02")
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
		Directory:  "/Users/liangyaopei",
		Datetime:   "2020-12-20",
	}
	isValid, prompt, errMsg := rChecker.Check(p)
	if !isValid {
		t.Errorf("errMsg:%s,prompt:%s", errMsg, prompt)
		return
	}
	t.Logf("valid param")
}

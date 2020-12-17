package _checker

import (
	"github.com/liangyaopei/checker"
	"testing"
)

type param struct {
	URL  string
	IPv4 string
}

func TestURLRule(t *testing.T) {
	rChecker := checker.NewChecker()

	urlRule := checker.NewURLRule("URL")
	rChecker.Add(urlRule, "invalid url")

	p := param{
		URL: "https://github.com/",
	}
	isValid, prompt, errMsg := rChecker.Check(p)
	if !isValid {
		t.Errorf("errMsg:%s,prompt:%s", errMsg, prompt)
		return
	}
	t.Logf("valid email")
}

func TestIPv4Rule(t *testing.T) {
	rChecker := checker.NewChecker()

	urlRule := checker.NewIPv4Rule("IPv4")
	rChecker.Add(urlRule, "invalid ipv4 address")

	p := param{
		IPv4: "14.215.177.38",
	}
	isValid, prompt, errMsg := rChecker.Check(p)
	if !isValid {
		t.Errorf("errMsg:%s,prompt:%s", errMsg, prompt)
		return
	}
	t.Logf("valid ipv4")
}

package checker

// Checker is the representation of
// validation object
type Checker interface {
	Add(rule Rule, prompt string)
	Check(param interface{}) (bool, string, string)
}

type ruleChecker struct {
	rules   []Rule
	prompts []string
}

func (c *ruleChecker) Add(rule Rule, prompt string) {
	c.rules = append(c.rules, rule)
	c.prompts = append(c.prompts, prompt)
}

func (c ruleChecker) Check(param interface{}) (bool, string, string) {
	for i := 0; i < len(c.rules); i++ {
		isValid, msg := c.rules[i].Check(param)
		if !isValid {
			return false, c.prompts[i], msg
		}
	}
	return true, "", ""
}

// NewChecker returns the
// Checker implementation
func NewChecker() Checker {
	return &ruleChecker{}
}

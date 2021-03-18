package checker

// Rule represents the restriction
// of param should obey
type Rule interface {
	Check(param interface{}) (bool, string)
}

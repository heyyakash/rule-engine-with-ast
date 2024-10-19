package helpers

import "errors"

func ValidateRule(rule string) error {
	if len(rule) == 0 {
		return errors.New("length of rule cannot be 0")
	}
	if rule[0] != '(' || rule[len(rule)-1] != ')' {
		return errors.New("rule should start with ( and end with )")
	}
	return nil
}

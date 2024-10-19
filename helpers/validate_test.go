package helpers_test

import (
	"testing"

	"github.com/heyyakash/rule-engine-with-ast/helpers"
)

func TestRuleValidation(t *testing.T) {
	if err := helpers.ValidateRule(""); err == nil {
		t.Error("Empty rule string validated")
	}

	if err := helpers.ValidateRule("age >= 30 AND gender = 'male'"); err == nil {
		t.Error("rule without ( and ) validated")
	}
}

package helpers

import "testing"

type testcase struct {
	Input  map[string]interface{}
	Expect bool
}

func TestEvaluation(t *testing.T) {
	testcases := []testcase{
		{
			Input: map[string]interface{}{
				"age":        34,
				"department": "Sales",
				"salary":     5000000,
				"experience": 8,
			},
			Expect: true,
		}, {
			Input: map[string]interface{}{
				"age":        25,
				"department": "Sales",
				"salary":     5000000,
				"experience": 8,
			},
			Expect: false,
		},
	}
	ast := MapToAST(TestOutputMap)
	for _, tc := range testcases {
		if Evaluate(ast, tc.Input) != tc.Expect {
			t.Error("Evaluation failing to evaluate the ast correctly")
		}
	}
}

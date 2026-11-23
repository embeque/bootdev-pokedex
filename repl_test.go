package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		// add more cases here
		{
			input:    "embeque-lukan frost",
			expected: []string{"embeque-lukan", "frost"},
		},
	}

	for i, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			// error and continue here
			t.Error("Length of the result and expected don't match")
			t.Fatalf("Test %d: slice length (got:%d, expected:%d)", i, len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Words don't match on the same index, length is same")
				t.Fatalf("Test %d: mismatch input (got:%s, expected:%s)", i, word, expectedWord)
			}
		}
	}
}

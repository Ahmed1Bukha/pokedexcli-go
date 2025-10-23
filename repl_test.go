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
		{
			input:    "test lmfao  ",
			expected: []string{"test", "lmfao"},
		},
		{
			input:    "NOOOO ggffGoNext  ",
			expected: []string{"NOOOO", "ggffGoNext"},
		},
		// add more cases here
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected){
			t.Errorf("The length of the output doesn't equal the length of the expected")
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord{
				t.Errorf("The words are not the same. Expected: %v, Actual:%v",expectedWord,word)
				return
			}
		}
	}
}
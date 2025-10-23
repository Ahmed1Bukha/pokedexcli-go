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
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "test lmfao  ",
			expected: []string{"test", "lmfao"},
		},
		{
			input:    "NO ggffGoNext  ",
			expected: []string{"no", "ggffgonext"},
		},
		// add more cases here
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("The length of the output doesn't equal the length of the expected. Actual: %v, expected:%v ",len(actual),len(c.expected))
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
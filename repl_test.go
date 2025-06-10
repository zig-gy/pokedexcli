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
			input:    " hello  world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "foo",
			expected: []string{"foo"},
		},
		{
			input:    "foo   bar   baz",
			expected: []string{"foo", "bar", "baz"},
		},
		{
			input:    "   foo   bar   ",
			expected: []string{"foo", "bar"},
		},
		{
			input:    "foo\tbar\nbaz",
			expected: []string{"foo", "bar", "baz"},
		},
		{
			input:    "foo  bar\tbaz",
			expected: []string{"foo", "bar", "baz"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Mismatch in array length")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("%s does not match %s", word, expectedWord)
			}
		}
	}
}
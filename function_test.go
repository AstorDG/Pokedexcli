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
			input:    " hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "I lovepokemon",
			expected: []string{"I", "lovepokemon"},
		},
		{
			input:    "pikachudiglit",
			expected: []string{"pikachudiglit"},
		},
	}

	for _, cas := range cases {
		actual := cleanInput(cas.input)
		if len(actual) != len(cas.expected) {
			t.Errorf("Output slice lenths don't match. Actual lenth: %v, Expected length: %v", len(actual), len(cas.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := cas.expected[i]
			if word != expectedWord {
				t.Errorf("Words don't match. Actual word: %v, Expected word: %v", word, expectedWord)
			}
		}
	}
}

package scrabble

import "testing"

type capTest struct {
	input    string
	expected bool
}

var testCases = []capTest{
	{"", false},
	{"1", false},
	{"a", true},
	{"A", true},
	{"ABC", true},
	{"abc", true},
	{"AbC", false},
	{"Abc", true},
}

func TestCapitalUse(t *testing.T) {
	for _, test := range testCases {
		if actual := CapitalUse(test.input); actual != test.expected {
			t.Errorf("Score(%q) expected %t, Actual %t", test.input, test.expected, actual)
		}
	}
}

func TestCapitalUse2(t *testing.T) {
	for _, test := range testCases {
		if actual := CapitalUse2(test.input); actual != test.expected {
			t.Errorf("Score(%q) expected %t, Actual %t", test.input, test.expected, actual)
		}
	}
}



func BenchmarkCapitalUse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			CapitalUse(test.input)
		}
	}
}

func BenchmarkCapitalUse2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			CapitalUse2(test.input)
		}
	}
}

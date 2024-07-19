package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := RomanToInt(tt.input)
			assert.Equal(t, tt.expected, got, tt.input)
		})
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"ice", "icecube", "icecream"}, "ice"},
		{[]string{"ab", "a"}, "a"},
	}
	for _, tt := range tests {
		testName := tt.input[0] + "..." + tt.input[len(tt.input)-1]
		t.Run(testName, func(t *testing.T) {
			got := LongestCommonPrefix(tt.input)
			assert.Equal(t, tt.expected, got, tt.input)
		})
	}
}

func TestIsValidParentheses(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"((])[]{}", false},
		{"(]", false},
		{"([)]", false},
		{"{[]}(())", true},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := IsValidParentheses(tt.input)
			assert.Equal(t, tt.expected, got, tt.input)
		})
	}
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1    *ListNode
		list2    *ListNode
		expected *ListNode
	}{
		{nil, nil, nil},
		{&ListNode{1, nil}, nil, &ListNode{1, nil}},
		{&ListNode{1, nil}, &ListNode{2, nil}, &ListNode{1, &ListNode{2, nil}}},
		{&ListNode{1, &ListNode{3, nil}}, &ListNode{2, nil}, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := MergeTwoLists(tt.list1, tt.list2)
			assert.Equal(t, tt.expected, got)
		})
	}
}

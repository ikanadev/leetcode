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

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 0}, 1},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := RemoveDuplicates(tt.input)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		input    []int
		val      int
		expected int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
		{[]int{2, 2, 2}, 2, 0},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := RemoveElement(tt.input, tt.val)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestFindIndexOcurrence(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		expected int
	}{
		{"sadbutsad", "sad", 0},
		{"leetcode", "et", 2},
		{"a", "a", 0},
		{"iseedeadpeople", "le", 12},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := FindIndexOcurrence(tt.haystack, tt.needle)
			assert.Equal(t, tt.expected, got)
		})
	}
}

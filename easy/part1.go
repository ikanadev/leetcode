package easy

// https://leetcode.com/problems/roman-to-integer/description/
func RomanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var sum int
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && m[s[i]] < m[s[i+1]] {
			sum -= m[s[i]]
		} else {
			sum += m[s[i]]
		}
	}
	return sum
}

// https://leetcode.com/problems/longest-common-prefix/description/
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// https://leetcode.com/problems/valid-parentheses/description/
func IsValidParentheses(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	m := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]rune, 0, len(s))
	for _, char := range s {
		openChar, isClosingChar := m[char]
		if isClosingChar {
			if len(stack) == 0 || stack[len(stack)-1] != openChar {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

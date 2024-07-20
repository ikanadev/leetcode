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

// https://leetcode.com/problems/merge-two-sorted-lists/description/
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var head *ListNode
	if list1.Val < list2.Val {
		head = list1
		head.Next = MergeTwoLists(list1.Next, list2)
	} else {
		head = list2
		head.Next = MergeTwoLists(list1, list2.Next)
	}
	return head
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[index] != nums[i] {
			index++
			nums[index] = nums[i]
		}
	}
	return index + 1
}

// https://leetcode.com/problems/remove-element/description/
func RemoveElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
func FindIndexOcurrence(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

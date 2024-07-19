package main

import (
	"fmt"
	"leetcode/easy"
)

func main() {
	fmt.Println(easy.IsValidParentheses("()[]{}"))
	fmt.Println(easy.IsValidParentheses("([]{}"))
	fmt.Println(easy.IsValidParentheses("()[{}"))
	fmt.Println(easy.IsValidParentheses("([])[{}]{(())[]}"))
}

package main

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.
    Every close bracket has a corresponding open bracket of the same type.

Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

*/
func validParenthesis(input string) bool {
	stack := []rune{}
	matching := map[rune]rune{'}': '{', ')': '(', ']': '['}
	for _, char := range input {
		switch char {
		case '{', '(', '[':
			stack = append(stack, char)
		case '}', ')', ']':
			if len(stack) == 0 || stack[len(stack)-1] != matching[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

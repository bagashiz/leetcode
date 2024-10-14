package stack

// https://leetcode.com/problems/valid-parentheses/
func IsValid(s string) bool {
	stack := make(runeStack, 0)

	for _, c := range s {
		if len(stack) == 0 {
			stack.push(c)
			continue
		}

		top := stack.peek()

		switch c {
		case '}':
			if top == '{' {
				stack.pop()
				continue
			}
		case ')':
			if top == '(' {
				stack.pop()
				continue
			}
		case ']':
			if top == '[' {
				stack.pop()
				continue
			}
		}

		stack.push(c)
	}

	return len(stack) == 0
}

// runeStack is a stack implementation for the rune data type.
type runeStack []rune

// Push adds a rune to the top of the stack.
func (rs *runeStack) push(char rune) {
	*rs = append(*rs, char)
}

// Pop removes and returns the rune at the top of the stack.
func (rs *runeStack) pop() rune {
	char := (*rs)[len(*rs)-1]
	*rs = (*rs)[:len(*rs)-1]
	return char
}

// Peek returns the rune at the top of the stack without removing it.
func (rs *runeStack) peek() rune {
	return (*rs)[len(*rs)-1]
}

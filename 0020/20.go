package main

func isValid(s string) bool {
	stk := []rune{}
	for _, c := range s {
		if c == '(' {
			stk = append(stk, ')')
		} else if c == '[' {
			stk = append(stk, ']')
		} else if c == '{' {
			stk = append(stk, '}')
		} else {
			if len(stk) == 0 || stk[len(stk)-1] != c {
				return false
			}
			stk = stk[:len(stk)-1]
		}
	}
	return len(stk) == 0
}

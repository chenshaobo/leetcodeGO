package validParentheses

//import "fmt"

//import "fmt"

func IsValid(s string) bool {

	length := len(s)
	if length % 2 != 0 {
		return false
	}
	m := map[byte]byte{
		'{':'}',
		'[':']',
		'(':')',
		'}':'{',
		')':'(',
		']':'[',
	}

	curIndex := -1
	stack := make([]byte, length)

	for i := 0; i < length; i++ {

		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			curIndex ++
			stack[curIndex] = s[i]
		} else if s[i] == '}' || s[i] == ']' || s[i] == ')' {
			if curIndex < 0{
				return false
			}
			v, _ := m[stack[curIndex]]
			if s[i] == v {
				curIndex--
			} else {
				return false
			}
		} else {
			return false
		}
	}

	if curIndex == -1 {
		return true
	}
	return false
}


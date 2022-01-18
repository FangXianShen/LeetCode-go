package main

import "fmt"

func checkStr(str string) bool {
	stack := util.NewStack()
	charMap := map[string]string{")": "(", "]": "[", "}": "{"} //先定义配对的map,右括号为键
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		char := string(str2[i])
		if value, ok := charMap[char]; !ok { //如果是左括号，压入栈
			stack.Push(char)
		} else if stack.Len == 0 || (stack.Pop() != value) { //如果是右括号，判断弹出的字符是否与之匹配
			return false
		}
	}
	if stack.Len == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	str := "[[]"
	res := checkStr(str)
	fmt.Println(res)
}

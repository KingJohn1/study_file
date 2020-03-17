// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _27_基本计算器2

import (
	"fmt"
	"strconv"
)

type Stack []string

func (s *Stack) pop() (r string) {
	if s == nil || len(*s) == 0 {
		return ""
	}
	*s, r = (*s)[:len(*s)-1], (*s)[len(*s)-1]
	return
}

func (s *Stack) push(str string) {
	if s == nil {
		return
	}
	*s = append(*s, str)
}

func (s Stack) isEmpty() bool {
	return len(s) == 0
}

func (s Stack) peek() string {
	if len(s) == 0 {
		return ""
	}
	return s[len(s)-1]
}

func comparePriority(a, b string) bool {
	if a == "" {
		return false
	}
	if b == ")" {
		return true
	}
	if a == "(" {
		return false
	}
	if a == "*" || a == "/" {
		return true
	}
	if b == "*" || b == "/" {
		return false
	}
	return true
}

var sign = []byte{'+', '-', '*', '/', '(', ')', ' '}
// 搞一个带括号的计算器

// 算法：以数字为基础点，如果数字右边的优先级高于栈中数字 先计算高优先级，如果数字右边优先级小于等于栈中优先级先计算栈中数字
// 如果遇到左括号直接入栈，认定左右括号优先级最低。如果栈中是（，右侧是） 则（）相互抵消
func calculate(s string) int {

	strs := splitString(s)

	var stack Stack

	// i的坐标只可能是数字和左括号
	for i := 0; i < len(strs); i++ {
		if strs[i] == "(" {
			stack = append(stack, strs[i])
			continue
		}

		calNum, _ := strconv.Atoi(strs[i])

		// 当前是数字
		if i == len(strs)-1 {
			return calStack(&stack, calNum)
		}
		if strs[i+1] == ")" {
			calNum = calStack(&stack, calNum)
			stack.pop()
			strs[i+1] = strconv.Itoa(calNum)
			continue
		}
		if stack.isEmpty() || stack.peek() == "(" {
			stack.push(strs[i])
			stack.push(strs[i+1])
			i++
			continue
		}

		if comparePriority(stack.peek(), strs[i+1]) {
			calNum = calStack(&stack, calNum)
			stack.push(strconv.Itoa(calNum))
			stack.push(strs[i+1])
			i++
		} else {
			rightNum, _ := strconv.Atoi(strs[i+2])
			calNum = calNumOp(calNum, rightNum, strs[i+1])
			strs[i+2] = strconv.Itoa(calNum)
			i++
		}
	}

	return 0
}

func calStack(stack *Stack, culNum int) int {
	for {
		if stack.peek() == "(" || stack.isEmpty() {
			return culNum
		}
		op := stack.pop()
		leftNum, _ := strconv.Atoi(stack.pop())
		culNum = calNumOp(leftNum, culNum, op)
	}
}

func calNumOp(a, b int, op string) int {
	if op == "+" {
		return a + b
	}
	if op == "-" {
		return a - b
	}
	if op == "*" {
		return a * b
	}
	if op == "/" {
		return a / b
	}
	fmt.Println("op invalid,op:", op)
	return a
}

func equal(b byte) bool {
	for _, v := range sign {
		if v == b {
			return true
		}
	}
	return false
}

func splitString(s string) (result []string) {
	lastIndex := -1
	for i := 0; i < len(s); i++ {
		if equal(s[i]) {
			if lastIndex != -1 {
				result = append(result, s[lastIndex:i])
				lastIndex = -1
			}
			if s[i] != ' ' {
				result = append(result, s[i:i+1])
			}
		} else {
			if lastIndex == -1 {
				lastIndex = i
			}
		}
	}
	if lastIndex != -1 {
		result = append(result, s[lastIndex:])
	}
	return
}

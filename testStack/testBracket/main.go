package main

import (
	"errors"
	"fmt"
	"log"
)

type Stack struct {
	len   int
	point int
	data  []byte
}

func (s *Stack) append(b byte) error {
	if s.point == s.len {
		return errors.New("stack is full")
	}
	s.data[s.point] = b
	s.point += 1
	return nil
}

func (s *Stack) pop() (byte, error) {
	if s.point > 0 {
		data := s.data[s.point-1]
		s.point -= 1
		return data, nil
	}
	return ' ', errors.New("stack is empty")
}

func TestBracketLegal(k string) (bool, string) {

	brackMap := map[byte]byte{')': '(', ']': '[', '}': '{'}

	sStack := Stack{
		len:   100,
		point: 0,
		data:  make([]byte, 10),
	}

	for _, i := range []byte(k) {
		if i == '(' || i == '[' || i == '{' {
			err := sStack.append(i)
			if err != nil {
				log.Panic(err.Error())
			}
		} else if i == ')' || i == ']' || i == '}' {
			v, err := sStack.pop()
			if err != nil {
				return false, err.Error()
			} else {
				if v != brackMap[i] {
					return false, "括号不匹配"
				}
			}
		}
	}

	if sStack.point == 0 {
		return true, ""
	} else {
		return false, "存在多余的左括号"
	}
}

func main() {
	fmt.Println(TestBracketLegal("((()))"))
	fmt.Println(TestBracketLegal("((())"))
	fmt.Println(TestBracketLegal("(()()()())"))
	fmt.Println(TestBracketLegal("[(])()"))
	fmt.Println(TestBracketLegal("{[()()]}"))
	fmt.Println(TestBracketLegal("{[()()]"))

}

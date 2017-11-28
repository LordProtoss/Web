package main

import (
	"unicode"
)

func RemoveEven(s []int) (a []int) {
	a = make([]int, 0)
	for _, v := range s {
		if v%2 != 0 {
			a = append(a, v)
		}
	}
	return a
}

func PowerGenerator(a int) func() int {
	x := a
	return func() (b int) {
		b = x
		x = x * a
		return
	}
}
func DifferentWordsCount(s string) (r int) {
	c := make(map[string]int)
	word := ""
	l := 0
	number := 0
	for _, v := range s {
		if unicode.IsLetter(v) {
			c := unicode.ToLower(v)
			word += string(c)
			l += 1
		}
		if !unicode.IsLetter(v) {
			if l != 0 {
				if c[word] == 0 {
					c[word] += 1
					number = number + 1
				}
				word = ""
				l = 0
			}
		}
	}
	return number
}

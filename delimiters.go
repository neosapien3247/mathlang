package main

import (
	//"fmt"
	"sort"
)

func findDelimiters(txt String) (delims []int) {
	delims = make([]int, 0)

	// first search for double delimiters : $$math$$
	for i := 0; i < len(txt)-2; i++ {
		if txt[i:i+2] == "$$" {
			delims = append(delims, i)
		}
	}

	// then find single delimiters : $math$
	for i := 1; i < len(txt)-1; i++ {
		if txt[i] == '$' && txt[i+1] != '$' && txt[i-1] != '$' {
			delims = append(delims, i)
		}
	}
	delims = sort.IntSlice(delims)
	if len(delims)%2 != 0 {
		panic("delimiter missing!")
	}
	sort.Ints(delims)
	return
}

func smartReplace(math String) (s String) {
	s = ""
	if len(math) < 2 {
		panic("string too short!")
	} else {
		if math[1] == '$' {
			s += math[0:2]
			//s += replace(math[2 : len(math)-2])
			s += replace(math[2:])
			//s += math[len(math)-2:]
			return
		} else {
			s += math[0:1]
			//s += replace(math[1 : len(math)-1])
			s += replace(math[1:])
			//s += math[len(math)-1:]
			return
		}
	}
}

func replaceBetweenDelimiters(math String) (s String) {
	delims := findDelimiters(math)

	s = math[0:delims[0]]
	for i := 0; i < len(delims)-1; i += 2 {
		s += smartReplace(math[delims[i]:delims[i+1]])
		if i+2 < len(delims) {
			s += math[delims[i+1]:delims[i+2]]
		} else {
			s += math[delims[i+1]:]
		}
	}
	if s[len(s)-1] != math[len(math)-1] {
		s += math[delims[len(delims)-1]:]
	}

	return
}

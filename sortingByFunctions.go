package main

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

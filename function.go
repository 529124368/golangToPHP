package main

import "C"

//export calcFib
func calcFib(i int) int {
	if i < 2 {
		return i
	}
	return calcFib(i-1) + calcFib(i-2)
}

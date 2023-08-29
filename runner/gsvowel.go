package runner

import (
	"container/list"
	"fmt"
	"strings"
)

func runVowel() {
	printStart()
	l := findVowels("Sam Anderson")
	fmt.Println("Vowels are:")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%c ", e.Value)
	}
	fmt.Printf("\n")
	printEnd()
}

func inArray(needle rune, hayStack []string) bool {
	for _, b := range hayStack {
		if b == string(needle) {
			return true
		}
	}
	return false
}

func findVowels(ms string) *list.List {
	a := []string{"a", "e", "i", "o", "u"}
	l := list.New()
	for _, letter := range ms {
		if inArray(letter, a) {
			l.PushFront(letter)
		}
	}

	a1 := []rune{'a', 'e', 'i', 'o', 'u'}
	l1 := list.New()
	for _, letter := range a1 {
		if strings.ContainsRune(ms, letter) {
			l1.PushBack(letter)
		}
	}

	return l1
}

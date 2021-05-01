package main

import (
	"container/list"
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	t := list.New()
	isReverse := false
	for _, v := range strings.Split(s, "") {
		if v == "R" {
			isReverse = !isReverse
			continue
		}

		if isReverse {
			if front := t.Front(); front != nil && front.Value == v {
				t.Remove(front)
			} else {
				t.PushFront(v)
			}
			continue
		}

		if back := t.Back(); back != nil && back.Value == v {
			t.Remove(back)
		} else {
			t.PushBack(v)
		}
	}
	if isReverse {
		for e := t.Back(); e != nil; e = e.Prev() {
			fmt.Print(e.Value)
		}
	} else {
		for e := t.Front(); e != nil; e = e.Next() {
			fmt.Print(e.Value)
		}
	}
	fmt.Println("")
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, q int
	var s string

	fmt.Scanf("%d", &n)
	fmt.Scanf("%s", &s)
	fmt.Scanf("%d", &q)

	r := []rune(s)

	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	isSwapped := false
	for i := 0; i < q; i++ {
		var t, a, b int

		sc.Scan()
		t, _ = strconv.Atoi(sc.Text())

		sc.Scan()
		a, _ = strconv.Atoi(sc.Text())

		sc.Scan()
		b, _ = strconv.Atoi(sc.Text())

		if t == 2 {
			isSwapped = !isSwapped
		}

		if t == 1 {
			if isSwapped {
				if a <= n {
					a = a + n
				} else {
					a = a - n
				}

				if b <= n {
					b = b + n
				} else {
					b = b - n
				}

				if a > b {
					tmp := a
					a = b
					b = tmp
				}
			}
			r = swap(r, a-1, b-1)
		}
	}

	if isSwapped {
		r = append(r[n:], r[:n]...)
	}
	fmt.Print(string(r))
}

func splitToInt(s string) (intArr []int) {
	arr := strings.Split(s, " ")

	for _, v := range arr {
		intV, _ := strconv.Atoi(v)
		intArr = append(intArr, intV)
	}
	return
}

func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func swap(r []rune, a, b int) []rune {
	sa := r[a]
	sb := r[b]

	r[a] = sb
	r[b] = sa

	return r
}

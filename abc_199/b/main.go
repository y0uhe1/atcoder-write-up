package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var n int
	var a, b []int
	sc.Scan()
	n, _ = strconv.Atoi(sc.Text())

	sc.Scan()
	a = splitToInt(sc.Text())

	sc.Scan()
	b = splitToInt(sc.Text())

	ans := []int{}

	for i := 0; i < n; i++ {
		tmp := []int{}
		ai := a[i]
		bi := b[i]

		for j := ai; j <= bi; j++ {
			tmp = append(tmp, j)
		}

		if i == 0 {
			ans = append(ans, tmp...)
			continue
		}

		for j := 0; j < len(ans); j++ {
			if !contains(tmp, ans[j]) {
				ans = append(ans[:j], ans[j+1:]...)
				j--
			}
		}
	}

	fmt.Print(len(ans))
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

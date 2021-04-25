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
	var a, b, c int
	if sc.Scan() {
		abc := splitToInt(sc.Text())
		a = abc[0]
		b = abc[1]
		c = abc[2]
	}

	if a*a+b*b < c*c {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}

}

func splitToInt(s string) (intArr []int) {
	arr := strings.Split(s, " ")

	for _, v := range arr {
		intV, _ := strconv.Atoi(v)
		intArr = append(intArr, intV)
	}
	return
}

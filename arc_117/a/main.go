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
	var a, b int
	if sc.Scan() {
		ab := splitToInt(sc.Text())
		a = ab[0]
		b = ab[1]
	}

	arrA := []string{}
	arrB := []string{}

	sumA := (1 + a) * a / 2
	sumB := (1 + b) * b / 2

	if a == b {
		for i := 1; i <= a; i++ {
			arrA = append(arrA, strconv.Itoa(i))
		}

		for i := 1; i <= b; i++ {
			arrB = append(arrB, strconv.Itoa(i*-1))
		}
	}

	if a > b {
		base := (sumA - sumB) / b
		rem := (sumA - sumB) % b

		for i := 1; i <= a; i++ {
			arrA = append(arrA, strconv.Itoa(i))
		}

		for i := 1; i <= b; i++ {
			if i == b {
				arrB = append(arrB, strconv.Itoa((base+i+rem)*-1))
			} else {
				arrB = append(arrB, strconv.Itoa((base+i)*-1))
			}
		}
	}

	if a < b {
		base := (sumB - sumA) / a
		rem := (sumB - sumA) % a

		for i := 1; i <= b; i++ {
			arrB = append(arrB, strconv.Itoa(i*-1))
		}

		for i := 1; i <= a; i++ {
			if i == a {
				arrA = append(arrA, strconv.Itoa(base+i+rem))
			} else {
				arrA = append(arrA, strconv.Itoa(base+i))
			}
		}
	}

	fmt.Println(strings.Join(arrA, " ") + " " + strings.Join(arrB, " "))
}

func splitToInt(s string) (intArr []int) {
	arr := strings.Split(s, " ")

	for _, v := range arr {
		intV, _ := strconv.Atoi(v)
		intArr = append(intArr, intV)
	}
	return
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N int
	var D, H float64
	fmt.Scan(&N, &D, &H)

	var sc = bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	var y, tmp float64
	for i := 0; i < N; i++ {
		var d, h float64

		sc.Scan()
		d, _ = strconv.ParseFloat(sc.Text(), 64)

		sc.Scan()
		h, _ = strconv.ParseFloat(sc.Text(), 64)

		tmp = (D*h - d*H) / (D - d)

		if tmp > y {
			y = tmp
		}
	}

	fmt.Print(y)
}

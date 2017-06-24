package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	i1, _ := strconv.ParseUint(scanner.Text(), 0, 64)
	scanner.Scan()
	d1, _ := strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	s1 := scanner.Text()

	fmt.Println(i + i1)
	fmt.Printf("%.1f\n",d + d1)
	fmt.Println(s + s1)

}


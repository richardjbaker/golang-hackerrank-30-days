package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	var size int32
	fmt.Scanf("%d", &size)
	scanner:=bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := strings.Split(scanner.Text(), " ")

	for i := len(s)-1; i >= 0; i-- {
		fmt.Printf("%s ", s[i])
	}
}

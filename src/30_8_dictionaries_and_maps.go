package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	var size int32
	fmt.Scanf("%d", &size)
	scanner := bufio.NewScanner(os.Stdin)
	m := make(map[string]string)
	for i := size - 1; i >= 0; i-- {
		scanner.Scan()
		s := strings.Split(scanner.Text(), " ")
		m[s[0]] = s[1]
	}
	loop:
	for {
		scanner.Scan()
		switch {
		default:
			fmt.Printf("%s=%s\n", scanner.Text(), m[scanner.Text()])
		case scanner.Text() == "":
			break loop
		case m[scanner.Text()] == "":
			fmt.Println("Not found")
		}
	}
}

package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {

	scanner:=bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("Hello, World.");
	fmt.Println(scanner.Text())
}

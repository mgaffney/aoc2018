package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var total int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, "converting number:", err)
		}
		total = total + a
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println(total)
}

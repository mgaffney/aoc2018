package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var cs []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		a, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, "converting number:", err)
			os.Exit(1)
		}
		cs = append(cs, a)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}

	var freq int
	freqs := make(map[int]bool)
	for {
		for _, a := range cs {
			freq = freq + a
			if _, ok := freqs[freq]; !ok {
				freqs[freq] = true
			} else {
				fmt.Println("second time: ", freq)
				os.Exit(0)
			}
		}
	}
}

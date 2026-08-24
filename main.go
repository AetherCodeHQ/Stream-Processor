package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	mode := "passthrough"
	if len(os.Args) > 1 {
		mode = os.Args[1]
	}
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1<<20), 1<<20)
	lines, fields, bytes := 0, 0, 0
	var seen map[string]int
	if mode == "unique" {
		seen = map[string]int{}
	}
	for sc.Scan() {
		line := sc.Text()
		lines++
		fields += len(strings.Fields(line))
		bytes += len(line)
		switch mode {
		case "passthrough":
			fmt.Println(line)
		case "lower":
			fmt.Println(strings.ToLower(line))
		case "upper":
			fmt.Println(strings.ToUpper(line))
		case "unique":
			if seen != nil {
				seen[line]++
				if seen[line] == 1 {
					fmt.Println(line)
				}
			}
		case "count":
			// just count, don't output
		}
	}
	fmt.Fprintf(os.Stderr, "stream: %d lines, %d fields, %d bytes\n", lines, fields, bytes)
}

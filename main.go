package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// This is a simple application that takes a terraform 'output' format and makes it into
// a valid terraform input file.

func main() {
	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		fmt.Fprint(os.Stdout, fmt.Sprintf("%s\n", func(eq, txt string) string {
			if n := strings.SplitN(txt, eq, 2); len(n) == 2 {
				return strings.Join([]string{n[0], strconv.Quote(n[1])}, eq)
			}
			return txt
		}(" = ", s.Text())))
	}
}

package main

import (
	"bufio"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := scanner.Text()
		println(ucl)
	}

}

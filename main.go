package main

import (
	"bufio"
	"nars_go/nars"
	"os"
)

func main() {
	n := nars.New()

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			n.SetInputLine(scanner.Text())
		}
	}()

	n.Start()
}

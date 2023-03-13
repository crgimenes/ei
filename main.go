package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	reader := bufio.NewReader(os.Stdin)

	// Read a single rune from stdin
	for {
		r, _, err := reader.ReadRune()

		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Printf("Read rune: %q\r\n", r)

		if r == 'q' {
			return
		}
	}

}

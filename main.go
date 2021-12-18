// Prints its command-line arguments using strings.Join
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	echo(os.Stdout, os.Args[1:])
}

func echo(w io.Writer, args []string) {
	fmt.Fprintln(w, strings.Join(args, " "))
}

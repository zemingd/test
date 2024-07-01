package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	if n >= 30 {
		Fprintln(out, "Yes")
	} else {
		Fprintln(out, "No")
	}
}

func main() { run(os.Stdin, os.Stdout) }

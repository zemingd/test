package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	sc := NewScanner(os.Stdin)
	s, _ := sc.NextString()
	if s == "ABC" {
		fmt.Println("ARC")
	} else {
		fmt.Println("ABC")
	}
}

// Scanner is a wrapper of bufio.Scanner which is customized for competitive programing.
type Scanner struct {
	bufio.Scanner
}

// NewScanner is a constructor for Scanner.
func NewScanner(r io.Reader) *Scanner {
	sc := Scanner{*bufio.NewScanner(r)}
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
	return &sc
}

// NextString reads a string from io stream
func (sc *Scanner) NextString() (string, error) {
	if !sc.Scan() {
		return "", errors.New("failed to scan")
	}
	return sc.Text(), nil
}

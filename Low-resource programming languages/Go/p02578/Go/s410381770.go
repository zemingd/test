package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	_, _ = strconv.ParseUint(scanner.Text(), 10, 64)
	scanner.Scan()
	aall := scanner.Text()

	a := strings.Split(aall, ` `)
	var mina uint64
	var fumidaisum uint64
	for _, an := range a {
		anint, _ := strconv.ParseUint(an, 10, 64)
		if mina < anint {
			mina = anint
		} else {
			fumidaisum += mina - anint
		}
	}
	fmt.Println(fumidaisum)
}

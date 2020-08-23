package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

/*
var alphabet = []rune{'A', 'B',
	'C', 'D',
	'E', 'F',
	'G', 'H',
	'I', 'J',
	'K', 'L',
	'M', 'N',
	'O', 'P',
	'Q', 'R',
	'S', 'T',
	'U', 'V',
	'W', 'X',
	'Y', 'Z', ' '}
*/
var alphabet = []rune{'A', 'a', 'B', 'b',
	'C', 'c', 'D', 'd',
	'E', 'e', 'F', 'f',
	'G', 'g', 'H', 'h',
	'I', 'i', 'J', 'j',
	'K', 'k', 'L', 'l',
	'M', 'm', 'N', 'n',
	'O', 'o', 'P', 'p',
	'Q', 'q', 'R', 'r',
	'S', 's', 'T', 't',
	'U', 'u', 'V', 'v',
	'W', 'w', 'X', 'x',
	'Y', 'y', 'Z', 'z', ' '}

func main() {
	table := make(map[rune]int)
	for _, r := range alphabet {
		table[r] = 0
	}

	//r := bufio.NewReader(strings.NewReader(os.Stdin))
	r := bufio.NewReader(os.Stdin)
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			table[c] += 1
		}
	}

	// unsorted
	for r, t := range table {
		// exclude 0
		if t != 0 {
			fmt.Printf("%s : %d\n", string(r), t)
		}
	}
}

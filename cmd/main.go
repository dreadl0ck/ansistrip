// ANSISTRIP - ANSI Sequence Removal Library
// Copyright (c) 2017 Philipp Mieden <dreadl0ck [at] protonmail [dot] ch>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/dreadl0ck/ansistrip"
)

func printHelp() {
	fmt.Println("usage: ansistrip <file> OR cat <file> | ansistrip")
}

func main() {

	if len(os.Args) > 1 {

		f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0600)
		if err != nil {
			log.Fatal(err)
		}

		filterANSI(f)
		return
	}

	filterANSI(os.Stdin)
	return
}

func filterANSI(w io.Reader) {

	// read data line by line, strip ANSI Sequences and print to stdout
	r := bufio.NewReader(w)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Fatal("failed to read line: ", err)
			} else {
				fmt.Println("done.")
				os.Exit(0)
			}
		}
		fmt.Println(string(ansistrip.StripAnsi(line)))
	}
}

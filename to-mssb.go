package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

// MSSBCA MATHEMATICAL SANS-SERIF BOLD CAPITAL A
const MSSBCA = 0x1D5D4
// MSSBSA MATHEMATICAL SANS-SERIF BOLD SMALL A
const MSSBSA = 0x1D5EE
// SSBONE MATHEMATICAL SANS-SERIF BOLD DIGIT ZERO
const MSSBONE = 0x1D7EC

func main() {
	if len(os.Args) > 1 {
		var out = toMSSBC(os.Args[1])
		for i := 2; i < len(os.Args); i++ {
			out += " " + toMSSBC(os.Args[i])
		}
		fmt.Println(out)
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(toMSSBC(scanner.Text()))
		}
	}
}

func oneASCII(r int32, line []byte) []byte {
	var utf8bytes = make([]byte, utf8.RuneLen(r))
	utf8.EncodeRune(utf8bytes, r)
	for _, u := range utf8bytes {
		line = append(line, u)
	}
	return line
}

func toMSSBC(in string) string {
	var line []byte
	for _, c := range []byte(in) {
		if c >= 'A' && c <= 'Z' {
			r := int32(int(c-'A') + MSSBCA)
			line = oneASCII(r, line)
		} else if (c >= 'a' && c <= 'z') {
			r := int32(int(c-'a') + MSSBSA)
			line = oneASCII(r, line)
		} else if (c >= '0' && c <= '9') {
			r := int32(int(c-'0') + MSSBONE)
			line = oneASCII(r, line)
		} else {
			line = append(line, c)
		}
	}
	return string(line)
}

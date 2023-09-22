package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var d int
	var astr, bstr string
	var err error

	if scanner.Scan() {
		d, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, "1行目の変換エラー:", err)
			return
		}
	}

	if scanner.Scan() {
		astr = scanner.Text()
	}

	if scanner.Scan() {
		bstr = scanner.Text()
	}

	fa, _, _ := big.NewFloat(0).SetPrec(366).Parse(astr, 10)
	fb, _, _ := big.NewFloat(0).SetPrec(366).Parse(bstr, 10)
	sum := new(big.Float).SetPrec(366).Add(fa, fb)
	format := fmt.Sprintf("%.*f", d, sum)
	fmt.Println(format)
}

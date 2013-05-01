package main

import (
	"noodle/account"
	"noodle/report"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var input *os.File

func init() {
	filename := ""
	flag.StringVar(&filename, "input", "", "path of file to process")
	flag.Parse()
	if filename == "" {
		input = os.Stdin
	} else {
		var err error
		if input, err = os.Open(filename); err != nil {
			panic(err)
		}
	}
}

func main() {
	balances := map[string]*account.Balance{}
	lineno := 0
	reader := bufio.NewReader(input)
	for {
		if line, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		} else {
			lineno += 1
			cols := strings.Split(strings.TrimRight(line, "\n"), " ")
			switch len(cols) {
			case 4:
				account.Add(balances, cols)
			case 3:
				account.Transaction(balances, cols)
			default:
				fmt.Fprintf(os.Stderr, "Unexpected input at line %d: '%s'", lineno, line)
			}
		}
	}
	fmt.Println(report.Summary(balances))
}

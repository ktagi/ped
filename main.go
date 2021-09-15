package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
)

func main() {

	var (
		e bool
	)
	flag.BoolVar(&e, "e", false, "encode mode")
	flag.Parse()

	sc := bufio.NewScanner(os.Stdin)
	if !sc.Scan() {
		os.Exit(0)
	}

	if e {
		fmt.Println(url.PathEscape(sc.Text()))
	} else {
		u, err := url.PathUnescape(sc.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(u)
	}
}

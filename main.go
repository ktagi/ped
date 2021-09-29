package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
)

func main() {
	cli := &CLI{inStream: os.Stdin, outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}

type CLI struct {
	inStream             io.Reader
	outStream, errStream io.Writer
}

func (c *CLI) Run(args []string) int {

	var (
		e bool
	)
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.BoolVar(&e, "e", false, "encode mode")
	flags.Parse(args[1:])

	sc := bufio.NewScanner(c.inStream)
	if !sc.Scan() {
		return 0
	}

	if e {
		fmt.Fprintln(c.outStream, url.PathEscape(sc.Text()))
	} else {
		u, err := url.PathUnescape(sc.Text())
		if err != nil {
			fmt.Fprintln(c.errStream, err)
			return 1
		}
		fmt.Fprintln(c.outStream, u)
	}

	return 0
}

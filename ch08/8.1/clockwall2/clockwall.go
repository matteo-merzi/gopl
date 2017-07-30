package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"text/tabwriter"
)

type clock struct {
	name, host string
}

func (c *clock) watch(r io.Reader) <-chan string {
	ch := make(chan string)
	s := bufio.NewScanner(r)
	go func() {
		for s.Scan() {
			ch <- fmt.Sprintf("%s:%s", c.name, s.Text())
		}
	}()
	return ch
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "usage: clockwall NAME=HOST ...")
		os.Exit(1)
	}
	clocks := make([]*clock, 0)
	for _, arg := range os.Args[1:] {
		fields := strings.Split(arg, "=")
		if len(fields) != 2 {
			fmt.Fprintf(os.Stderr, "bad arg: %s\n", arg)
			os.Exit(1)
		}
		clocks = append(clocks, &clock{name: fields[0], host: fields[1]})
	}
	dial(clocks)
}

func dial(clocks []*clock) {
	chx := make([]<-chan string, 0)
	for _, c := range clocks {
		net.Dial("tcp", c.host)
		conn, err := net.Dial("tcp", c.host)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		chx = append(chx, c.watch(conn))
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	for {
		sep := ""
		for _, ch := range chx {
			fmt.Fprintf(w, "%s%s", sep, <-ch)
			sep = "\t"
		}
		fmt.Fprintf(w, "\n")
		w.Flush()
	}
}

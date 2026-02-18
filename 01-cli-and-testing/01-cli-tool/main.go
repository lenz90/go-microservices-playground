package main

import (
	"flag"
	"fmt"
	"os"
)

func run(a []string) (string, error) {
	if len(a) < 1 {
		return "", fmt.Errorf("missing command")
	}
	switch a[0] {
	case "greet":
		fs := flag.NewFlagSet("g", flag.ContinueOnError)
		n := fs.String("name", "world", "")
		if err := fs.Parse(a[1:]); err != nil {
			return "", err
		}
		return "hello, " + *n, nil
	case "calc":
		if len(a) < 2 {
			return "", fmt.Errorf("missing op")
		}
		fs := flag.NewFlagSet("c", flag.ContinueOnError)
		x := fs.Float64("a", 0, "")
		y := fs.Float64("b", 0, "")
		if err := fs.Parse(a[2:]); err != nil {
			return "", err
		}
		if a[1] == "add" {
			return fmt.Sprintf("%g", *x+*y), nil
		}
		if a[1] == "sub" {
			return fmt.Sprintf("%g", *x-*y), nil
		}
		return "", fmt.Errorf("unknown op")
	}
	return "", fmt.Errorf("unknown")
}
func main() {
	o, e := run(os.Args[1:])
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
	fmt.Println(o)
}

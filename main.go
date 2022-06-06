package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if os.Args[1] == "-" {
		scanner := bufio.NewScanner(os.Stdin)
		math := ""
		for scanner.Scan() {
			math += scanner.Text() + "\n"
		}
		fmt.Print(translate(math))
	} else if os.Args[1] == "-d" {
		scanner := bufio.NewScanner(os.Stdin)
		math := ""
		for scanner.Scan() {
			math += scanner.Text() + "\n"
		}
		fmt.Print(translate(math, true))
	} else if os.Args[1] == "-f" || os.Args[1] == "-fd" || os.Args[1] == "-df" {
		file, err := os.Open(os.Args[2])
		check(err)
		scanner := bufio.NewScanner(file)
		math := ""
		for scanner.Scan() {
			math += scanner.Text() + "\n"
		}
		fmt.Println(math)
		if os.Args[1] == "-fd" || os.Args[1] == "-df" {
			fmt.Print(translate(math, true))
		} else {
			fmt.Print(translate(math))
		}
	} else if os.Args[1] == "-e" {
		math := strings.Join(os.Args[2:], " ")
		fmt.Print(translate(math))
	} else {
		printHelp()
	}
}

func check(err ...error) {
	for _, e := range err {
		if e != nil {
			panic(e)
		}
	}
}

// delim here refers to wether to only parse between delimiters or not
func translate(txt string, delim ...bool) String {
	math := String(txt)
	// delim is taken to be false by default
	if len(delim) > 0 {
		if delim[0] {
			return replaceBetweenDelimiters(math)
		} else {
			return replace(math)
		}
	} else {
		return replace(math)
	}
}

func printHelp() {
	fmt.Println("mathlang usage:")
	fmt.Println("mathlang -edfh")
	fmt.Printf(`
-h              : show this help menu
-e [expression]	: translates expression to LaTeX math
-               : reads from stdin
-d              : reads from stdin only in delimiters ($math$ or $$math$$)
-f [file]       : reads from file and translates to stdout
-df / -fd [file]: reads from file and translates to stdout only in delimiters
		`)
}

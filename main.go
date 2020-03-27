package main

import (
	"fmt"
	"os"

	"github.com/rblack42/go-literate/parser"
)

func main() {
	res := parser.Execute(os.Args[1:])

	if res.Err != nil {
		fmt.Println("Error processing file", res.Err)
		os.Exit(-1)
	}

	fmt.Println("go-literate exited normally")
}

package parser

import (
	"bytes"
	"io"
	"fmt"
	"os"
	"bufio"
)

type Result struct {
	Err error
}

// Read a whole file into the memory and store it as array of lines
func readLines(path string) (lines []string, err error) {
    var (
        file *os.File
        part []byte
        prefix bool
    )
    if file, err = os.Open(path); err != nil {
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    buffer := bytes.NewBuffer(make([]byte, 0))
    for {
        if part, prefix, err = reader.ReadLine(); err != nil {
            break
        }
        buffer.Write(part)
        if !prefix {
            lines = append(lines, buffer.String())
            buffer.Reset()
        }
    }
    if err == io.EOF {
        err = nil
    }
    return
}
func Execute(args [ ]string) Result {
	var res Result

	for _, path := range args {
		fmt.Println("processing",path)
		lines, err := readLines(path)
		if err != nil {
			fmt.Println("Error: %s\n", err)
			res.Err = err
			return res
		}
		for _, line := range lines {
			fmt.Println(line)
		}
	}
	res.Err = nil
	
	return res
}


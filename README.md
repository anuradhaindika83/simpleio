# simpleio - File handling simplified

## Information

This library extended the Go file handling APIs. This provides some useful functions which are missing in the Go file handling package

## Installation

    go get github.com/anuradhaindika83/simpleio

## Usage

```Go
    package main

    import (
    	"fmt"
    	"github.com/anuradhaindika83/simpleio"
    )

    func main() {
    	file := simpleio.FileHandler{}
    	file.OpenFile("test.txt")
    	line := file.ReadLine()
    	fmt.Println(line)
    }
```

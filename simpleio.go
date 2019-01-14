package simpleio

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//FileHandler represents a text file
type FileHandler struct {
	File *os.File
	Name string
	EOF  bool
}

//OpenFile will open a given file for reading
func (fh *FileHandler) OpenFile(file string) {
	fhdl, err := os.Open(file)
	check(err)
	fh.File = fhdl
	fh.Name = fhdl.Name()
}

//Close will close file
func (fh *FileHandler) Close() {
	fh.File.Close()
}

//ReadLine reads the next line
func (fh *FileHandler) ReadLine() string {
	line := ""
	buf := make([]byte, 1)
	for {
		_, err := fh.File.Read(buf)
		if err == io.EOF {
			fh.EOF = true
			return line
		}
		check(err)
		line = line + string(buf)

		if strings.Index(line, "\r\n") != -1 {
			result := strings.Trim(line, "\r\n")
			return result
		}
	}
}

//ReadLines reads all the next line
func (fh *FileHandler) ReadLines() []string {
	line := ""
	lines := make([]string, 1)
	buf := make([]byte, 1)
	fh.File.Seek(0, 0)
	for {

		_, err := fh.File.Read(buf)
		if err == io.EOF {
			fh.EOF = true
			return lines
		}
		check(err)
		line = line + string(buf)

		if strings.Index(line, "\r\n") != -1 {
			result := strings.Trim(line, "\r\n")
			lines = append(lines, result)
			line = ""
		}
	}
}

//ReadBlock reads a specified amount of characters
func (fh *FileHandler) ReadBlock(start int64, count int64) string {
	block := ""
	buf := make([]byte, count)
	_, err := fh.File.ReadAt(buf, start)
	block = string(buf)
	if err == io.EOF {
		fh.EOF = true
		return block
	}
	check(err)
	return block

}

//ReadToEnd reads entier file
func (fh *FileHandler) ReadToEnd() string {
	line := ""
	buf := make([]byte, 1)
	fh.File.Seek(0, 0)
	for {

		_, err := fh.File.Read(buf)
		if err == io.EOF {
			fh.EOF = true
			return line
		}
		check(err)
		line = line + string(buf)
	}

}

//GetLength gets the length of the content
func (fh *FileHandler) GetLength() int64 {
	var length int64
	length = 0
	buf := make([]byte, 1)
	fh.File.Seek(0, 0)
	for {

		_, err := fh.File.Read(buf)
		if err == io.EOF {
			fh.EOF = true
			return length
		}
		check(err)
		length = length + 1
	}

}

//Reset will set file reading start position to 0
func (fh *FileHandler) Reset() {
	fh.File.Seek(0, 0)
}

func check(e error) {
	if e != nil {
		fmt.Println(e.Error())
		panic(e)
	}
}

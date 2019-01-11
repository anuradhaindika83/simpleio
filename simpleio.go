package simpleio

import (
	"fmt"
	"io"
	"os"
	"strings"
)

//FileHandler represents a text file
type FileHandler struct {
	file *os.File
	EOF  bool
}

//OpenFile will open a given file for reading
func (fh *FileHandler) OpenFile(file string) {
	fhdl, err := os.Open(file)
	check(err)
	fh.file = fhdl
}

//Close will close file
func (fh *FileHandler) Close() {
	fh.file.Close()
}

//ReadLine reads the next line
func (fh *FileHandler) ReadLine() string {
	line := ""
	buf := make([]byte, 1)
	for {
		_, err := fh.file.Read(buf)
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
	fh.file.Seek(0, 0)
	for {

		_, err := fh.file.Read(buf)
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
	_, err := fh.file.ReadAt(buf, start)
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
	fh.file.Seek(0, 0)
	for {

		_, err := fh.file.Read(buf)
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
	fh.file.Seek(0, 0)
	for {

		_, err := fh.file.Read(buf)
		if err == io.EOF {
			fh.EOF = true
			return length
		}
		check(err)
		length = length + 1
	}

}

func check(e error) {
	if e != nil {
		fmt.Println(e.Error())
		panic(e)
	}
}

package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type InputStream struct {
	rd             *bufio.Reader
	pos, line, col int
	eof            bool
	rune_size      int // for testing the eof
}

/* Open the file and initialize the pos, col, line variables*/
func (is *InputStream) open_file(fimename string) {
	f, err := os.Open(fimename)
	check(err) // check whether the file is opened correctly
	is.rd = bufio.NewReader(f)
	is.line = 1
	is.rune_size = -1 // not eof at the first beginning,
}

func (is *InputStream) peek() rune {
	r, size, _ := is.rd.ReadRune() // don't check the eof, is_eof() will handle it
	is.rune_size = size
	is.rd.UnreadRune()
	return r
}

func (is *InputStream) next() rune {
	r, size, _ := is.rd.ReadRune() // don't check the eof, is_eof() will handle it
	is.rune_size = size
	is.pos++

	char := string(r)
	if char == "\n" {
		is.line++
		is.col = 0
	} else {
		is.col++
	}
	return r
}

func (is *InputStream) is_eof() bool {
	return is.rune_size == 0
}

func (i *InputStream) err_msg(msg string) {
	err := errors.New(msg + " at line: " + strconv.Itoa(i.line) + " col: " + strconv.Itoa(i.col))
	panic(err)
}

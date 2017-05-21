package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type InputStream struct {
	rd             *bufio.Reader
	pos, line, col int
	eof            bool
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

/* Open the file and initialize the pos, col, line variables*/
func (is *InputStream) open_file(fimename string) {
	f, err := os.Open(fimename)
	check(err)
	is.rd = bufio.NewReader(f)
	is.line = 1
}

func (is *InputStream) peek() rune {
	r, _, err := is.rd.ReadRune()
	check(err)
	is.rd.UnreadRune()
	return r
}

func (is *InputStream) next() rune {
	r, _, err := is.rd.ReadRune()
	check(err)
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
	return string(is.peek()) == ""
}

/*
	We should define the types of all the token the JSON files can have.
	1. type = string, "cat"
	2. type = number 45, -23, 102.332, -52.01e-355, and -475246256524654 are all valid JSON numbers.
	3. type = punctuation, {} [] : ,
	4. type = keywords. true, false and null.
*/

// var keywords []string  = {"true"; "false"; "null"}

type Token struct {
	t_type string
	value  interface{}
}

type TokenStream struct {
	input   InputStream
	current Token
}

func (ts TokenStream) read_next() Token {
	read_while(is_whitespace)
	if ts.input.is_eof() {
		return nil
	}
	char = ts.input.peek()
	switch {
	case ch == "\"":
		return ts.read_string()
	case is_digit(ch):
		return ts.read_number()
	case is_punc(ch):
		return ts.read_punc()
	case is_kw_start(ch):
		return ts.read_kw()
	}
}

func is_digit(r rune) bool {
	return unicode.IsDigit(r)
}

func is_whitespace(r rune) bool {
	return unicode.IsSpace(r)
}

func is_punc(r rune) bool {
	return unicode.IsPunct(r)
}

func is_kw_start(r rune) bool {
	ch := string(r)
	for _, kw := range keywords {
		if ch == kw[0] {
			return true
		}
	}
	return false
}

// The major components of JSON file:
// 	1. punctuations like : {} []
// 	2. strings, starting and ending with ""
// 	3. numbers, including integers, floating point numbers and exponent
// 	4. keywords like true, false and null.

// should handle the situation of reading escaped
func (t *TokenStream) read_string() Token {

}

func (t *TokenStream) read_number() Token {

}

func (t *TokenStream) read_punc() Token {

}

func (t *TokenStream) read_kw() Token {

}

func (t *TokenStream) peek() Token {

}

func (t *TokenStream) next() Token {

}

func (t *TokenStream) is_eof() bool {

}

// func main() {
// 	var is InputStream
// 	is.open_file("test.json")
// 	for i := 0; i < 10; i++ {
// 		s := is.peek()
// 		fmt.Println(s, is.pos, is.line, is.col)
// 		s = is.next()
// 		fmt.Println(s, is.pos, is.line, is.col)
// 	}
// }

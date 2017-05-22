package main

import (
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"
)

/*
	We should define the types of all the token the JSON files can have.
	1. type = string, "cat"
	2. type = number 45, -23, 102.332, -52.01e-355, and -475246256524654 are all valid JSON numbers.
	3. type = punctuation, {} [] : ,
	4. type = keywords. true, false and null.
*/

type Token struct {
	t_type string
	value  interface{}
}

func (t Token) String() string {
	return fmt.Sprintf("type: %v, value: %v", t.t_type, t.value)
}

type TokenStream struct {
	input    *InputStream
	current  Token
	keywords []string
	puncs    []string
}

func (t *TokenStream) set_up(filename string) {
	t.input = new(InputStream)
	t.input.open_file(filename)
	t.keywords = []string{"true", "false", "null"}
	t.puncs = []string{"[", "]", ",", ":", "{", "}"}
}

func (t *TokenStream) _is_valid(candidate string, lst []string) bool {
	for _, kw := range lst {
		if candidate == kw {
			return true
		}
	}
	return false
}

// the core function for the tokenizer
func (ts *TokenStream) read_next() Token {
	ts.read_while(is_whitespace) // skip(ignore) all whitespace
	if ts.input.is_eof() {
		return Token{t_type: "eof", value: 0}
	}
	r := ts.input.peek()
	switch {
	case string(r) == "\"":
		return ts.read_string()
	case is_num(r):
		return ts.read_number()
	case is_punc(r):
		return ts.read_punc()
	case ts.is_kw_start(r):
		return ts.read_kw()
	}
	ts.input.err_msg("invalid token ")
	return Token{}
}

func is_num(r rune) bool {
	return unicode.IsDigit(r) || string(r) == "+" || string(r) == "-" ||
		string(r) == "." || string(r) == "e"
}

func is_whitespace(r rune) bool {
	return unicode.IsSpace(r)
}

func is_punc(r rune) bool {
	return unicode.IsPunct(r)
}

func (ts *TokenStream) is_kw_start(r rune) bool {
	for _, kw := range ts.keywords {
		first_r, _ := utf8.DecodeRuneInString(kw)
		if r == first_r {
			return true
		}
	}
	return false
}

func is_letter(r rune) bool {
	return unicode.IsLetter(r)
}

// The major components of JSON file:
// 	1. punctuations like : {} []
// 	2. strings, starting and ending with ""
// 	3. numbers, including integers, floating point numbers and exponent
// 	4. keywords like true, false and null.

// should handle the situation of reading escaped
func (t *TokenStream) read_while(predicate func(rune) bool) string {
	str := ""
	for !t.input.is_eof() && predicate(t.input.peek()) {
		str += string(t.input.next())
	}
	return str
}

func (t *TokenStream) read_string() Token {
	s := t._read_escaped("\"")
	string_tk := Token{t_type: "str", value: s}
	return string_tk
}

func (t *TokenStream) read_number() Token {
	str := t.read_while(is_num)
	_, err := strconv.ParseFloat(str, 64)
	if err != nil {
		msg := "invalid number " + str
		t.input.err_msg(msg)
		return Token{}
	} else {
		number_tk := Token{t_type: "num", value: str}
		return number_tk
	}
}

func (t *TokenStream) _read_escaped(end string) string {
	escaped := false
	result := ""
	t.input.next() // skip the first "\""
	for !t.input.is_eof() {
		s := string(t.input.next()) // read the next rune, convert to utf-8 string
		if escaped {                // the previous s was "\\", the current rune should be escaped
			result += ("\\" + s)
			escaped = false
		} else if s == "\\" {
			escaped = true
		} else if s == end {
			break
		} else {
			result += s
		}
	}
	return result
}

func (t *TokenStream) read_punc() Token {
	punc := string(t.input.next())
	if t._is_valid(punc, t.puncs) {
		punc_tk := Token{t_type: "punc", value: punc}
		return punc_tk
	} else {
		t.input.err_msg("invalid punctuation!")
		return Token{}
	}
}

func (t *TokenStream) read_kw() Token {
	str := t.read_while(is_letter)
	if t._is_valid(str, t.keywords) {
		kw_tk := Token{t_type: "kw", value: str}
		return kw_tk
	} else {
		t.input.err_msg("invalid keywords!")
		return Token{}
	}
}

func (t *TokenStream) peek() Token {
	// if current is nil, read then next token as current and return it
	// if current is not nil, which means we use contiguous peek(), we shouldn't
	// call read_next() for multiple times but instead should simply return the current
	if (Token{}) == t.current { // t.current is empty
		t.current = t.read_next()
		return t.current
	} else {
		return t.current
	}
}

func (t *TokenStream) next() Token {
	// it's a little tricky here. Since the token might have been read by peek(),
	// we CAN'T always call read_next(). Instead, we should firstly check whether
	// t.current is nil, if it is (no peek() before), we call read_next(); if not,
	// we should simply return the t.current(set by peek()) and set it back to nil.
	tok := t.current
	t.current = Token{}
	if tok != (Token{}) {
		return tok
	} else {
		return t.read_next()
	}
}

func (t *TokenStream) is_eof() bool {
	return t.peek().t_type == "eof" // check eof through the type of the current token
}

// func main() {
// 	var ts TokenStream
// 	ts.set_up("test.json")
// 	for !ts.is_eof() {
// 		s := ts.peek()
// 		fmt.Println(s)
// 		t := ts.next()
// 		fmt.Println(t)
// 	}
// }

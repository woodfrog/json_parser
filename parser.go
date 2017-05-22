package main

import (
	"fmt"
)

type Parser struct {
	tokens     *TokenStream
	html       string
	num_indent int
}

// punc should be the desired punctuation
func (p *Parser) is_punc(punc string) bool {
	tok := p.tokens.peek()
	if tok == (Token{}) || tok.t_type != "punc" || tok.value.(string) != punc {
		return false
	} else {
		return true
	}

}

func (p *Parser) skip_punc(punc string) {
	if p.is_punc(punc) {
		p.html_punc(punc) // the consuming of token is deferred to the html function
	}
}

// return the string of the html output
func (p *Parser) parse_toplevel() string {
	p.html = "<span style=\"font-family:monospace; white-space:pre\">\n"
	p.parse_object()
	p.html += "\n</span>"
	return p.html
}

func (p *Parser) parse_object() {
	first := true
	p.num_indent += 1
	p.skip_punc("{")
	p.html_new_line()
	for !p.tokens.is_eof() {
		if p.is_punc("}") {
			p.html_new_line()
			p.num_indent -= 1
			p.html_tab()
			p.html_punc("}")
			break
		}
		if first {
			first = false
		} else {
			p.skip_punc(",")
			p.html_new_line()
		}
		if p.is_punc("}") {
			p.html_new_line()
			p.num_indent -= 1
			p.html_tab()
			p.html_punc("}")
			break
		}
		p.html_tab()
		p.html_string()
		p.skip_punc(":")
		p.parse_value()
	}
}

// parse a sequence of values, since the length of sequence can be arbitrarily long, parse_array()
// should be recursive by itself
func (p *Parser) parse_array() {
	first := true
	p.skip_punc("[")
	for !p.tokens.is_eof() {
		if p.is_punc("]") {
			p.html_punc("]")
			break
		}
		if first {
			first = false
		} else {
			p.skip_punc(",")
		}
		if p.is_punc("]") {
			p.html_punc("]")
			break
		}
		p.parse_value()
	}
}

// this is a quite core part, the logic can enter into different parts
// depending on the type of value
func (p *Parser) parse_value() {
	switch {
	case p.is_punc("{"):
		p.parse_object()
	case p.is_punc("["):
		p.parse_array()
	case p.tokens.peek().t_type == "str":
		p.html_string()
	case p.tokens.peek().t_type == "num":
		p.html_num()
	case p.tokens.peek().t_type == "kw":
		p.html_kw()
	}
}

func (p *Parser) html_string() {
	str_tok := p.tokens.next() // read and consume
	fmt_str := "&quot;" + str_tok.value.(string) + "&quot;"
	p.html += fmt_str
}

func (p *Parser) html_num() {
	num_tok := p.tokens.next() // read and consume
	p.html += num_tok.value.(string)
}

func (p *Parser) html_kw() {
	kw_tok := p.tokens.next()
	p.html += kw_tok.value.(string)
}

func (p *Parser) html_punc(punc string) {
	p.tokens.next() // consume the punc token
	p.html += punc
}

func (p *Parser) html_new_line() {
	p.html += "\n"
}

func (p *Parser) html_tab() {
	for i := 0; i < p.num_indent; i++ {
		p.html += "\t"
	}
}

func main() {
	var ts TokenStream
	ts.set_up("test.json")
	var p Parser
	p.tokens = &ts
	result := p.parse_toplevel()
	fmt.Println(result)
}

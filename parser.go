package main

type parser struct {
	tokens *TokenStream
	html   string
}

// punc should be the desired punctuation
func (p *parser) is_punc(punc string) bool {
	tok := p.tokens.peek()
	return tok && tok.t_type == "punc" && string(tok.value) == punc
}

func (p *parser) skip_punc(punc string) {
	if p.is_punc(punc) {
		p.tokens.next()
		p.html_punc(punc)
	}
}

// return the string of the html output
func (p *parser) parse_toplevel() string {
	p.html = ""
	p.parse_object()
	return html
}

func (p *parser) parse_object() {
	p.skip_punc("{")
	// there should be a string token
	p.skip_punc(":")
	p.parse_value()
	p.skip_punc("}")
}

// this is a quite core part, the logic can enter into different parts
// depending on the type of value
func (p *parser) parse_value() {
	switch {
	case p.is_punc("{"):
		p.parse_object()
	case p.is_punc("["):
		p.parse_array()
	case p.is_string():
		p.html_string()
	case p.is_num():
		p.html_num()
	case p.is_kw():
		p.html_kw()
	}
}

// parse a sequence of values, since the length of sequence can be arbitrarily long, parse_array()
// should be recursive by itself
func (p *parser) parse_array() {
	first := true
	p.skip_punc("[")
	for !p.tokens.is_eof() {
		if p.is_punc("]") {
			p.tokens.next()
			p.html_punc("]")
			break
		}
		if first {
			first = false
		} else {
			p.skip_punc(",")
		}
		if p.is_punc("]") {
			p.tokens.next()
			p.html_punc("]")
			break
		}
		p.parse_value()
	}
}

func (p *parser) html_string() {

}

func (p *parser) html_num() {

}

func (p *parser) html_kw() {

}

func (p *parser) html_punc(punc string) {

}

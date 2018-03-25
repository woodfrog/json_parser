# Simple JSON Parser for Generating Colorized HTML Output

## Introduction

This project contains a JSON parser which converts a JSON input to a colorized HTML file. The parser's structure is based on [How to Implement a Programming Language](http://lisperator.net/pltut/). This project is based on [Go language](https://golang.org), so you need to setup Go environment before running it.

## How To Run

You can either build the project and run the executable file or directly run the scripts.

### a). Build The Binary And Run

1. Build the project with Go.
~~~
go build
~~~
2. Run the generated binary and specify the input and output files.
~~~
./json_parser inputFileName.json > outputFileName.html
~~~

### b). Run Without Build
This project can also run without building it. You need to specify each script in the command and provide input and output file names.
~~~
go run parser.go input_stream.go tokenizer.go inputFileName.json > outputFileName.html
~~~

## Design Review
### Basic Idea
The basic idea comes from recursive descent parser. A recursive descent parser is top-down parser built on a set of mutually recursive procedures. Each procedure implements one rule of the grammar.[2] Since JSON has strict syntax rules and these rules are corresponding to HTML outputs, a recursive descent parser can be used to check the grammar and convert it to an HTML string.

### Structure
This projects contain the following three components. Input Streaming, Tokenizer and Parser. These components are also executed in this order. This section talks about what these components are doing.

#### 1. Input Streaming

Input Streaming provides two major functions `peek()` and `next()` to read the input stream one rune by another. `peek()` checks the next token without consuming it. It decides the beginning of the next token and helps the tokenizer to choose different tokenizing logic. `next()` moves the pointer to the next token.

#### 2. Tokenizer

Tokenizer seperates the whole input text into a list of valid tokens. If there is any illegal token, the tokenizer would fail with proper error message.

#### 3. Parser 

The idea of "parser" in this project comes from recursive descent parser[2], but it is not a real parser. Parser in this project does not generate the abstract syntax tree (AST). It only checks the correctness of the token list and convert the token list into an HTML string. If there is syntax error, the parsing phase would fail and print corresponding error messages on stdout.

## Other Notes
1. Any valid JSON value to be the top-level object. In this case, an array can also be at top-level.[3]

2. The color mappings for different components are defined in `htmlColorMap` variable in `parser.go`. You may change the color mapping as you like.

## Reference
[1] How to Implement a Programming Language [http://lisperator.net/pltut/](http://lisperator.net/pltut/)

[2] Recursive Descent Parser [Wikipedia](https://en.wikipedia.org/wiki/Recursive_descent_parser)

[3] JSON LD 1.1 [https://json-ld.org/spec/latest/json-ld/](https://json-ld.org/spec/latest/json-ld/)
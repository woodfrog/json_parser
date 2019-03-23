# Simple JSON Parser for Generating Colorized HTML Output

## Introduction

This project contains a JSON parser that transforms a JSON file to a formatted HTML file with colour highlights. The structure of the parser follows a top-down approach inspired by the tutorial [How to Implement a Programming Language](http://lisperator.net/pltut/). The project is implemented in [Go language](https://golang.org) and requires Go environment to be set up before running the program.

## How To Run

Two methods are provided to run the program. Either build from binary and run the executable file or directly run the scripts.

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
Please specify each script in the command and provide the input and output file names.
~~~
go run parser.go input_stream.go tokenizer.go inputFileName.json > outputFileName.html
~~~

## Design Review
### Basic Idea
The basic idea originates from the recursive descent parser, which is a top-down parser built on a set of mutually recursive procedures. Each procedure implements a single rule of the grammar.[2] Since JSON has strict syntax rules that corresponds to HTML outputs, a recursive descent parser can be applied to validate and convert JSON grammar to an HTML string.

### Structure
This project contains the following three components: Input Streaming, Tokenizer, and Parser. These components are executed in the listed order. This section belows presents the objectives of each component.

#### 1. Input Streaming

Input Streaming provides two major functions `peek()` and `next()` to read the input stream one rune by another. `peek()` checks the next token without consuming it and decides the beginning of the next token, which helps the tokenizer to choose different tokenizing logic. `next()` moves the pointer to the next token.

#### 2. Tokenizer

Tokenizer seperates the whole input text into a list of valid tokens. If there is any illegal token, the tokenizer would fail with proper error message.

#### 3. Parser 

The parser in this project borrows the idea from a recursive descent parser[2], but does not fully implememt it. The parser implemented only checks the correctness of the token list and converts the token list into an HTML string but does not generate the abstract syntax tree (AST). If there is a syntax error, the parsing phase would fail and print corresponding error messages on stdout.

## Other Notes
1. Any valid JSON value can be the top-level object. In this case, an array can also be at the top-level.[3]

2. The color mappings for different components are defined in `htmlColorMap` variable in `parser.go`. You may change the color mapping as you like.

## Reference
[1] How to Implement a Programming Language [http://lisperator.net/pltut/](http://lisperator.net/pltut/)

[2] Recursive Descent Parser [Wikipedia](https://en.wikipedia.org/wiki/Recursive_descent_parser)

[3] JSON LD 1.1 [https://json-ld.org/spec/latest/json-ld/](https://json-ld.org/spec/latest/json-ld/)

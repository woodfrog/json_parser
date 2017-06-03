## Simple JSON Parser

A JSON parser which completes the stage of tokenization and elementary
parsing for converting the JSON text to colorized HTML file.

The overall parsing structure refers to the great tutorial on
[How to Implement a Programming Language](http://lisperator.net/pltut/). This
tutorial introduces the method of **recursive descent** to implement parser
for programming languages with specified syntax rules.

###Basic Structure

1.input streaming. This part provides two major functionalities peek() and next() to
read the input stream one rune by another. The peek() is especially important
for the tokenizer to decide the beginning of certain token so that it can
enter corresponding logic for different tokens.

2. Tokenizer. This module seperate the whole input text into a list of valid
   tokens, if there is any illegal token the tokenization would fail with
   proper error message.

3. Parser. Actually it's not a real parser since it doesn't generate the parse
   tree (AST). It basically further checks the correctness of the token list and
   convert the token list into a long HTML string. If there is syntax error the
   parsing phase would fail and print corresponding error messages on stdout.

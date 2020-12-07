# Base64 Golang Encoder/Decoder

A small project that takes command line input text and output the base64 encoding/decoding of it. The implementation is build without the use of any libraries.

## Run

The whole application is built in one file main.go You can choose to run the base application either by building the application `go build main.go` then sebsequently calling `b64 ...` or you can run without complilation with `go run main.go ...` both solutions work.

### Commands

The application takes 2 commands either encode or decode followed by text. You do not need to wrap the subsequent arguments to `encode` in quotes if the text is separated by a single white space. If you need double spaces then you should wrap it.

e.g.

`b64 encode hello world` = `aGVsbG8gd29ybGQ=`

`b64 encode "hello world"` = `aGVsbG8gd29ybGQ=`

`b64 encode hello  world` = `aGVsbG8gd29ybGQ=`

`b64 encode "hello  world"` = `aGVsbG8gIHdvcmxk`

## Notes 

This is a simple mini-project to get familar with bitwise operators in golang. It is not designed to be used in production code. For that you should use the *encoding/base64* module native to golang.

After building the binary you can easily but the binary file in one of the folders in your path (or adjust your path to include the local working directory). Doing so may give you a quick command line utility to go in and out of base64.

The code takes the character encoding it is given, there is no option at the moment to define different encodings.
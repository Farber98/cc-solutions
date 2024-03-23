# Compress CLI Tool

This is a command-line tool written in Go that provides functionality for Huffman encoding and decoding of files. It allows you to compress files using Huffman coding and then decompress them back to their original form.

## Usage

To use the tool, run the following command:

```sh
$ go run main.go [command] [file]
````

Replace [command] with one of the following:

- compress: Compress a file.
- decompress: Decompress a file.


## Example

Compress a file: 
```sh
go run main.go -compress tests/simple_test.txt
```

Decompress a file:
```sh
go run main.go -decompress tests/simple_test.txt.compressed
```



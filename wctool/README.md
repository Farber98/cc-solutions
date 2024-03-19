# Word Counting Tool

This is a command-line tool written in Go that provides functionality similar to the `wc` command. It counts the number of bytes, words, lines, and characters in a given file. To check the requirements, please visit [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-wc) page.

## Usage

To use the tool, run the following command:

```
$ go run main.go [options] [file]
```

Replace `[options]` with one of the following flags:

- `-c`: Count the number of bytes in the file.
- `-w`: Count the number of words in the file.
- `-l`: Count the number of lines in the file.
- `-m`: Count the number of characters in the file.
- `-all`: All the commands at one in the file.

## Examples

Count the number of bytes in a file:
```
$ go run main.go -c filename.txt
```

Count the number of words in a file:
```
$ go run main.go -w filename.txt
```

Count the number of lines in a file:
```
$ go run main.go -l filename.txt
```

Count the number of characters in a file:
```
$ go run main.go -m filename.txt
```

Execute all the commands together in a file:
```
$ go run main.go -all filename.txt
```

## Notes

- This tool assumes that words are separated by whitespace characters (space, tab, newline, etc.).
- For counting characters, the tool decodes the file contents in UTF-8 encoding and counts the number of runes.
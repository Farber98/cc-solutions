# Simple JSON Parser

This is a command-line tool written in Go that parses simple JSON objects. It parses valid and invalid JSON files, reporting which is which. This project is part of a coding challenge. For the full challenge description, please visit [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-jsonparser) page.

## Usage

To use the tool, run the following command:

```
$ go run main.go [file]
```

Replace `[file]` with the path to the JSON file you want to parse.

## Output

The tool will output a message indicating whether the JSON file is valid or invalid, along with the exit code (0 for valid, 1 for invalid).

## Examples

Parse a valid JSON file:

```
$ go run main.go tests/step1/valid.json
```

Parse an invalid JSON file:
```
$ go run main.go tests/step1/invalid.json
```

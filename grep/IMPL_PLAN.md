# Implementation Plan

## How to use

* [x] CLI application
    - stdin: `cat <test> | ./grep <pattern>`
    - file: `./grep <pattern> <file>`

```bash
$ cat test.txt | grep apple # or ./grep apple test.txt

I have an apple.
An apple a day keeps the doctor away.
```

* Input
    * Pattern: via the CLI argument
        * [x] Supported search pattern:
            * [x] Plain text
            * [x] Regex regular expression (e.g., `^apple$`, `apple|banana`)
    * Text: via stdin, e.g., `cat <file> | grep <pattern>`
* Output
    * All lines from the input text that contain the search pattern.

## Functionality

* [x] Receive input from the CLI argument and stdin
* [x] Search the pattern in the text
* [x] Output the content of the matching lines

## Extra bonus
* [x] Support a file path for the text, e.g., `grep <pattern> <file>`

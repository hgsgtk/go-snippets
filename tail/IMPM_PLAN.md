# Implementation Plan

## How to use

* CLI tool

```bash
./tail -f <file_path>

# The CLI outputs the new lines of the file
```

## Features

### Log rotation

* Log rotation: the file is renamed or replaced
    * If the file is renamed/replaced
        * e.g., app.log -> app.log.1 and new app.log is created, app.log
        * Detect the rename or the file deletion, and exit with error
        * If the file content is deleted, the CLI should exit with error
        * When the file is truncated
        * If there is no app.log, the CLI should exit with error

### Keyword filtering (optional)

TBA

## Implementation steps

- [x] CLI boilerplate
    * Input: file path (required)
    * Block the user thread until the user inputs `Ctrl+C`
- [x] Start reading the file
    * Validate the file path
        * If not exist, exit with error
    * Open the file
    * Read the file from the last line
- [x] Output the new lines of the file
    * Stream the new lines of the file to the CLI
- [x] Log rotation handling

## TODO

- [x] Add unit tests
- [x] Clean up temporal commits

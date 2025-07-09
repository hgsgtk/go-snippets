# Password Generator CLI ￼

Generate secure passwords with options for length, symbols, numbers, etc.

## Specification

- Interface: CLI
- Input:
    - [x] Length: default 8, min 8, max 100
    - [ ] Include symbols, `--include-symbols`
    - Include numbers
- Output:
    - Password
- Password requirements:
    - At least 8 characters to 100 characters
    - Default: lowercase 8 alphabet randomly
    - Opt-in: can specify characters uppercase, lowercase, numbers, symbols

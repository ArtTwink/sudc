# sudc

A simple Unix date calculator CLI tool written in Go.

## Features
- Parse and calculate Unix timestamps and date expressions
- Supports expressions like `now-2d`, `now+3h`, or direct Unix timestamps
- Output in Unix timestamp or UTC format
- Calculate duration between two Unix timestamps
- Flexible flag placement (thanks to Cobra)

## Usage

```
sudc [--unix|--utc] <expression>
```

### Examples

- Output the current time as a Unix timestamp:
  ```
  sudc now --unix
  ```
- Output the time 2 days ago as a Unix timestamp:
  ```
  sudc --unix "now-2d"
  ```
- Output the duration between two Unix timestamps:
  ```
  sudc --utc 1750071305-1749898505
  ```
- Output a Unix timestamp as a formatted date:
  ```
  sudc 1750071305
  ```

## Expression Syntax
- `now` — current time
- `now-2d` — 2 days ago
- `now+3h` — 3 hours from now
- `<timestamp1>-<timestamp2>` — duration between two Unix timestamps
- `<timestamp>` — parse and format a Unix timestamp

## Flags
- `--unix` — Output as Unix timestamp
- `--utc` — Output as UTC RFC3339 string

## Build

```
go build -o bin/sudc main.go
```

## Requirements
- Go 1.20+
- [Cobra](https://github.com/spf13/cobra)

## License
MIT

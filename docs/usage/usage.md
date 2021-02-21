---
title: Usage
has_children: true
nav_order: 3
---

## Usage
```
▶ gee [flags] [file1] [file2] ...
```


## Flags
```
▶ ~/go/bin/gee -h (if you install from source)
▶ gee -h
```

```
Usage: ./gee [flags] [file1] [file2] ...
(If you do not specify a file, only stdout is output)

Flags:
  -append
    	Append mode for files
  -chunked int
    	Chuked files from line (e.g output / output_1 / output_2)
  -debug
    	Show debug message!
  -distribute
    	Distribution to files
  -find string
    	Find string in line (colorize red)
  -format string
    	Change output format (json, md-table, html-table) (default "line")
  -grep string
    	Greping with Regular Expression (like grep)
  -grepv string
    	Inverse greping with Regular Expression (like grep -v)
  -prefix string
    	Prefix string
  -replace string
    	Replace string in line with '-find' option
  -rmnl
    	Remove newline(\r\n)
  -split string
    	Split string within line. (to line , to table, to md-table)
  -suffix string
    	Suffix string
  -uncolor
    	Uncolorize stdout
  -version
    	Version of gee
  -with-line
    	With line number (colorize blue)
  -with-time
    	With timestamp (colorize green)
```

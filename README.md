<p align="center">
  <br>
  <img src="https://user-images.githubusercontent.com/13212227/107402243-17bc5280-6b47-11eb-9f49-d97684c452c3.png" alt="gee" width="160px;">
  <br>
  🏵 Standard input to each files and stdout. similar to tee, But there's a little bit more and Write in go
  <br><br>
  <!--
  <img src="https://img.shields.io/github/v/release/hahwul/dalfox?style=flat"> 
  <a href="https://snapcraft.io/dalfox"><img alt="dalfox" src="https://snapcraft.io/dalfox/badge.svg" /></a>
  <img src="https://img.shields.io/github/languages/top/hahwul/dalfox?style=flat"> 
  <img src="https://api.codacy.com/project/badge/Grade/17cac7b8d1e849a688577f2bbdd6ecd0"> 
  <a href="https://goreportcard.com/report/github.com/hahwul/dalfox"><img src="https://goreportcard.com/badge/github.com/hahwul/dalfox"></a> 
    -->
  <img src="https://github.com/hahwul/gee/workflows/Build/badge.svg">
  <img src="https://github.com/hahwul/gee/workflows/SAST/badge.svg">
  <a href="https://twitter.com/intent/follow?screen_name=hahwul"><img src="https://img.shields.io/twitter/follow/hahwul?style=flat&logo=twitter"></a>
  <a href="https://github.com/hahwul"><img src="https://img.shields.io/github/stars/hahwul?style=flat&logo=github"></a></a>
</p>

## Introduction
Gee is tools of standard input to each files and stdout. It is similar to the tee command, but there are more functions for convenience. In addition, it was written as go.

## Installation
from source
```
$ go get github.com/hahwul/gee
```

## Usage
```
Usage: ./gee [flags] [file1] [file2] ...
(If you do not specify a file, only stdout is output)

Flags
  -append
    	Append mode for files
  -chunked int
    	Chuked files from line (e.g output / output_1 / output_2)
  -distribute
    	Distribution to files
  -prefix string
    	Prefix string
  -regex string
    	Match with Regular Expression (like grep)
  -regexv string
    	Unmatch with Regular Expression (like grep -v)
  -rmnl
    	Remove newline(\r\n)
  -suffix string
    	Suffix string
  -version
    	Version of gee
  -with-line
    	With line number
  -with-time
    	With timestamp
```

## Case of Study
### gee with prefix and suffix
```
$ cat urls | gee -prefix "curl -i -k " -suffix " -H 'Auth: abcd'" curls.sh
```

## Contribute
[Contribute](/CONTRIBUTING.md)

## Contributors
![](/CONTRIBUTORS.svg)

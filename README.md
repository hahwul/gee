<p align="center">
  <br>
  <img src="https://user-images.githubusercontent.com/13212227/107406554-e72ae780-6b4b-11eb-8f66-2f7e0d925f78.png" alt="gee" width="160px;">
  <br>
  Standard input to each files and stdout. similar to tee, But there's a little bit more and Write in go
  <br><br>
  <!--
  <img src="https://img.shields.io/github/v/release/hahwul/dalfox?style=flat"> 
  <a href="https://snapcraft.io/dalfox"><img alt="dalfox" src="https://snapcraft.io/dalfox/badge.svg" /></a>
  <img src="https://img.shields.io/github/languages/top/hahwul/dalfox?style=flat"> 
  <img src="https://api.codacy.com/project/badge/Grade/17cac7b8d1e849a688577f2bbdd6ecd0"> 
  <a href="https://goreportcard.com/report/github.com/hahwul/dalfox"><img src="https://goreportcard.com/badge/github.com/hahwul/dalfox"></a> 
    -->
  <a href="https://goreportcard.com/report/github.com/hahwul/gee"><img src="https://goreportcard.com/badge/github.com/hahwul/gee"></a>
  <a href="https://app.codacy.com/gh/hahwul/gee?utm_source=github.com&utm_medium=referral&utm_content=hahwul/gee&utm_campaign=Badge_Grade"><img src="https://api.codacy.com/project/badge/Grade/fac8a4d4755a4fb481432f7ed14db3ca"></a>
  <img src="https://github.com/hahwul/gee/workflows/Build/badge.svg">
  <img src="https://github.com/hahwul/gee/workflows/SAST/badge.svg">
  <a href="https://twitter.com/intent/follow?screen_name=hahwul"><img src="https://img.shields.io/twitter/follow/hahwul?style=flat&logo=twitter"></a>
  <a href="https://github.com/hahwul"><img src="https://img.shields.io/github/stars/hahwul?style=flat&logo=github"></a></a>
</p>

## 🔖 Introduction
Gee is tools of standard input to each files and stdout. It is similar to the tee command, but there are more functions for convenience. In addition, it was written as go.

## 🚀 Installation
from source
```
$ go get github.com/hahwul/gee
```
with homebrew (coming soon 🙏🏼 haven't released it yet. can't use it.)
```
$ brew tap hahwul/gee
$ brew install gee
```

## ☄️ Usage
```
$ ~/go/bin/gee -h (if you install from source)
$ gee -h
```
```
Usage: ./gee [flags] [file1] [file2] ...
(If you do not specify a file, only stdout is output)

Flags:
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

## 📚 Case of Study
### gee with prefix and suffix
```
$ cat urls | gee -prefix "curl -i -k " -suffix " -H 'Auth: abcd'" curls.sh
```
```
curl -i -k https://www.hahwul.com/?q=123 -H 'Auth: abcd'
curl -i -k http://testphp.vulnweb.com/listproducts.php?cat=asdf&ff=1 -H 'Auth: abcd'
curl -i -k https://xss-game.appspot.com/level1/frame  -H 'Auth: abcd'
```

## 🌟 Contribute
[Contribute](/CONTRIBUTING.md)

## 🌸 Contributors
![](/CONTRIBUTORS.svg)

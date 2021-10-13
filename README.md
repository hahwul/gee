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
  <a href="https://codecov.io/gh/hahwul/gee"><img src="https://codecov.io/gh/hahwul/gee/branch/main/graph/badge.svg"/></a>
  <img src="https://github.com/hahwul/gee/workflows/Build/badge.svg">
  <img src="https://github.com/hahwul/gee/workflows/SAST/badge.svg">
  <a href="https://twitter.com/intent/follow?screen_name=hahwul"><img src="https://img.shields.io/twitter/follow/hahwul?style=flat&logo=twitter"></a>
</p>

## 🔖 Introduction
Gee is tool of stdin to each files and stdout. It is similar to the tee command, but there are more functions for convenience. In addition, it was written as go. which provides output to stdout and files. In this process, it has various processing functions for lines such as replace, prefix, and suffix, so it can be used as a pipeline configuration or as a utility for testing. For more information, see the usage and case of study below!

## 🚀 Installation
From source
Go1.17
```
go install -v github.com/hahwul/gee@latest
```

Go1.16
```
GO111MODULE=on go get -v github.com/hahwul/gee
```
With homebrew (only macos)
```
brew tap hahwul/gee
brew install gee
```
Download from [release](https://github.com/hahwul/gee/releases) page. (macos,linux,freebsd,windows)
```
wget https://github.com/hahwul/gee/releases/download/v1.0.4/gee_1.0.4_linux_amd64.tar.gz
tar -xvf gee_1.0.4_linux_amd64.tar.gz
cp ./gee /usr/bin
```

## ☄️ Usage
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

## 📚 Case of Study
### gee with prefix and suffix
```
▶ cat urls | gee -prefix "curl -i -k " -suffix " -H 'Auth: abcd'" curls.sh
```
```
curl -i -k https://www.hahwul.com/?q=123 -H 'Auth: abcd'
curl -i -k http://testphp.vulnweb.com/listproducts.php?cat=asdf&ff=1 -H 'Auth: abcd'
curl -i -k https://xss-game.appspot.com/level1/frame  -H 'Auth: abcd'
```
### Find and replace
```
▶ cat raw.txt | gee -find keep-alive
▶ cat raw.txt | gee -find keep-alive -replace close
```
https://twitter.com/hahwul/status/1360495560843689989

### Specify the maximum length of the file and save it in multiple files.
```
▶ wc -l http.txt
2278

▶ cat http.txt | gee -chunked 500 output
```
https://twitter.com/hahwul/status/1360495565633540097

### Distribute each line sequentially to multiple files.
```
▶ wc -l http.txt
2278

▶ cat http.txt | gee -distribute alice.txt bob.txt charlie.txt
```
https://twitter.com/hahwul/status/1360495570922704897

## 🌟 Contribute
Contributions are always welcome. Please write/modify the code and send the PR.<br> 
Please read [Contribute](/CONTRIBUTING.md) for more information.

## 🌸 Contributors
![](/CONTRIBUTORS.svg)

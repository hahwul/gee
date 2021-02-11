## How to contribute
- First, fork this repository
- Second, clone repository in your develop machine
- Finaly, writing code and push/PR to me(hahwul/gee)

## Writing code
I'm checking the quality of code through Codacy when PR/Merge/Push. If you want to consider code quality in advance, please check the link below (not perfect, but very helpful).

https://goreportcard.com/report/github.com/hahwul/gee

e.g: `https://goreportcard.com/report/github.com/{your github account}/gee`

## Add new flag
First, Add new flag in sturcture in `pkg/model/options.go`
```go
type Options struct {
	Files         []string
	Append        bool
  // snip
}
```
Second, Add flag in `main.go`
```go
// e.g
appendOption := flag.Bool("append", false, "Append mode for files")
...
options := model.Options{
		Files:         files,
		Append:        *appendOption,
    // snip 
}
```
Finaly, Add logic
```go
if options.Append {
  mode = os.O_APPEND | os.O_CREATE | os.O_WRONLY
}
```

## Build
```
go build
```

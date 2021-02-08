package model

// Options is struct of gee options
type Options struct {
	Files         []string
	Append        bool
	ChunkedLine   bool
	ChunkedSize   bool
	Prefix        string
	Suffix        string
	WithTimestamp bool
	WithLine      bool
	RemoveLine    bool
	TableMarkdown bool
	TableHTML     bool
	Regex         string
	RegexV        string
	Split         string
	Distribute    bool
}

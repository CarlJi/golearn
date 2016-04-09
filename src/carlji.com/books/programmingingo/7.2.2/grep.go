package main

import (
	"regexp"
	"runtime"
)

type Job struct {
	filename string
	results  chan<- Result
}

type Result struct {
	filename string
	index    int
	line     string
}

func main() {
	grep(nil, nil)
}

var workers = runtime.NumCPU()

func grep(lineRx *regexp.Regexp, filenames []string) {
	jobs := make(chan Job, workers)
	results := make(chan Result, 1000)
	done := make(chan struct{}, workers)


}

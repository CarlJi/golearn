package main

import (
	"fmt"
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
	fmt.Println("Num of CPU", workers)
}

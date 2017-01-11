package main

import (
	"bufio"
	"os"

	"fmt"
	"strings"

	"flag"

	"strconv"

	"qiniupkg.com/x/log.v7"
)

// Profile represents the profiling data for a specific file.
type Profile struct {
	FileName   string
	Mode       string
	Blocks     []ProfileBlock
	TotalLines int
}

// ProfileBlock represents a single block of profiling data.
type ProfileBlock struct {
	StartLine, StartCol int
	EndLine, EndCol     int
	NumStmt, Count      int
}

func main() {
	file := flag.String("f", "", "coverage profile to summart ")
	flag.Parse()

	totalLine := 0
	pm := calcFile(*file)
	for k, v := range pm {
		sl := int(v)
		fmt.Printf("file: %v, MaxLine: %d  \n", k, sl)
		totalLine += sl
	}
	fmt.Printf("Total Line: %d \n", totalLine)

}

func calcFile(coverfile string) map[string]float64 {
	f, err := os.Open(coverfile)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	profileMap := make(map[string]float64)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		oneLine := scanner.Text()
		ss := strings.FieldsFunc(oneLine, split)
		fmt.Println(len(ss))
		if len(ss) > 3 {
			endLine, err := strconv.ParseFloat(ss[2], 64)
			if err != nil {
				fmt.Printf("strconv.ParseFloat: %v failed", ss[2])
				continue
			}
			v, ok := profileMap[ss[0]]
			if !ok || (ok && endLine > v) {
				profileMap[ss[0]] = endLine
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("scanner.Err():", err)
	}

	return profileMap
}

func split(r rune) bool {
	return r == ':' || r == ',' || r == ' '

}

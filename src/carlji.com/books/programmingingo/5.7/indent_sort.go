package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("|   Original   |   Sorted     |")
	fmt.Println("|--------------|--------------|")

	original := []string{"abc", " bgd", "  ci", " bd", " bf", "  cbf", "   abc", "   cbg"}

	sorted := SortedIndentStrings(original)
	for i, _ := range sorted {
		fmt.Printf("|%-14s|%-14s|\n", original[i], sorted[i])
	}
}

type Entry struct {
	key      string
	value    string
	children Entries
}

type Entries []Entry

func (entries Entries) Len() int { return len(entries) }

func (entries Entries) Less(i, j int) bool {
	return entries[i].key < entries[j].key
}

func (entries Entries) Swap(i, j int) {
	entries[i], entries[j] = entries[j], entries[i]
}

func SortedIndentStrings(s []string) []string {
	entries := populateEntries(s)
	return sortEntries(entries)
}

func populateEntries(s []string) Entries {
	indent, indentSize := computeIndent(s)
	entries := make(Entries, 0)
	for _, item := range s {
		i, level := 0, 0
		for strings.HasPrefix(item[i:], indent) {
			i += indentSize
			level++
		}
		key := strings.ToLower(strings.TrimSpace(item))
		addEntry(level, key, item, &entries)
	}

	return entries
}

func computeIndent(s []string) (string, int) {
	for j, item := range s {
		if len(item) > 0 && (item[0] == ' ' || item[0] == '\t') {
			whitespace := rune(item[0])
			for i, char := range item[1:] {
				if char != whitespace {
					return strings.Repeat(string(whitespace), j+i), j + i
				}
			}
		}
	}
	return "", 0
}

func addEntry(level int, key, value string, entries *Entries) {
	if level == 0 {
		*entries = append(*entries, Entry{key, value, make(Entries, 0)})
	} else {
		addEntry(level-1, key, value, &((*entries)[entries.Len()-1].children))
	}
}

func sortEntries(entries Entries) []string {
	var indentSlice []string
	sort.Sort(entries)
	for _, entry := range entries {
		sortChildren(entry, &indentSlice)
	}
	return indentSlice
}

func sortChildren(entry Entry, s *[]string) {
	*s = append(*s, entry.value)
	sort.Sort(entry.children)
	for _, child := range entry.children {
		sortChildren(child, s)
	}
}

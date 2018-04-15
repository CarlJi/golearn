package main

import (
	"io"
	"os"
	"strings"
)

func main() {

	reader := io.TeeReader(strings.NewReader("Golearn Everybody! \n"), os.Stdout)
	reader.Read(make([]byte, 20))
}

package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

func main() {
	pipe()
}

func pipe() {
	r, w := io.Pipe()
	go func() {
		// close the writer, so the reader knows there's no more data
		defer w.Close()

		if err := json.NewEncoder(w).Encode("hello word"); err != nil {
			log.Fatal(err)
		}
	}()

	bs, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(bs))
}

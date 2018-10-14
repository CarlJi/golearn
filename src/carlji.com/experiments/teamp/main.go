package main

import "qiniupkg.com/x/log.v7"

var NOP = "dsfdsf"

func main() {
	a, b := tta()
	log.Println(a)
	log.Println(b)
}

func tta() (no string, err error) {
	return NOP, err
}

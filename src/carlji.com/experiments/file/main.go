package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	num := 10240000
	path := strconv.FormatInt(int64(num), 10) + "O_Sync"
	_, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	u := time.Now()
	createFileOne(num, path)
	fmt.Println(time.Now().Sub(u).Nanoseconds())

	err = os.Remove(path)
	if err != nil {
		panic(err)
	}

	path = path + "_2"
	_, err = os.Create(path)
	if err != nil {
		panic(err)
	}

	u = time.Now()
	createFileTwo(num, path)
	fmt.Println(time.Now().Sub(u).Nanoseconds())

	err = os.Remove(path)
	if err != nil {
		panic(err)
	}

}

func createFileOne(num int, path string) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_SYNC, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	by := make([]byte, num)
	rand.Read(by)
	off := int64(0)
	_, err = file.WriteAt(by, off)
	if err != nil {
		panic(err)
	}

}

func createFileTwo(num int, path string) {
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	by := make([]byte, num)
	rand.Read(by)
	off := int64(0)
	_, err = file.WriteAt(by, off)
	if err != nil {
		panic(err)
	}
	file.Sync()
}

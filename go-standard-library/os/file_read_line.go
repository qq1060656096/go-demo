package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	f, err := os.Open(dir + "/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(f)
	for ;; {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("read line err: %v", err)
		}
		fmt.Printf("read line: %s \n", string(b))
	}

}
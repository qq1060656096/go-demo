package main

import (
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	os.Mkdir(dir + "/mkdir",  os.ModePerm)
	os.MkdirAll(dir + "/mkdirall/sub/sub", os.ModePerm)
}

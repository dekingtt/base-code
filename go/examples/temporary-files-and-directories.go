package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkerr(err error) {
	if nil != err {
		panic(err)
	}
}
func main() {
	f, err := os.CreateTemp("", "sample")
	checkerr(err)
	defer os.Remove(f.Name())

	fmt.Println("Temp file name:", f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	checkerr(err)
	dname, err := os.MkdirTemp("", "sampledir")
	checkerr(err)
	defer os.RemoveAll(dname)
	fmt.Println("Temp dir name:", dname)

	fname := filepath.Join(dname, "file1.temp")
	err = os.WriteFile(fname, []byte{1, 2}, os.ModeAppend)
	checkerr(err)
	defer os.Remove(fname)
}

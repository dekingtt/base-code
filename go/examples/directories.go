package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	err := os.Mkdir("subdir", os.ModeDir)
	checkerr(err)

	defer os.RemoveAll("subdir")

	createEmptyFile := func(name string) {
		dat := []byte("")
		checkerr(os.WriteFile(name, dat, os.ModeAppend))
	}

	createEmptyFile("subdir/file1.temp")
	err = os.MkdirAll("subdir/parent/child", os.ModeDir)
	checkerr(err)

	createEmptyFile("subdir/parent/file2.temp")
	createEmptyFile("subdir/parent/file3.temp")
	createEmptyFile("subdir/parent/child/file4.temp")

	c, err := os.ReadDir("subdir/parent")
	checkerr(err)
	fmt.Println("listing subdir/parent")

	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	fmt.Println("listing subdir/parent/child")
	err = os.Chdir("subdir/parent/child")
	c, err = os.ReadDir(".")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	err = os.Chdir("../../..")
	checkerr(err)

	fmt.Println("visting dir")
	err = filepath.Walk("subdir", visit)
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}

package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func checkError(err error, exitOnError bool) {
	if err != nil {
		log.Println(err)
		if exitOnError {
			os.Exit(1)
		}
	}
}

func printDir(dir string) string {
	d, f := path.Dir(dir), path.Base(dir)
	fmt.Println(d, f)
	if d == "/" {
		fmt.Println(f)
		return fmt.Sprintf("%c ", 0x2514)
	} else {
		indent := printDir(d)
		fmt.Println(indent + f)
		return "  " + indent
	}
}

func main() {
	wd, err := os.Getwd()
	checkError(err, true)
	printDir(wd)
}

// godatetime.go (c) 2014-2015 David Rook - all rights reserved

package main

import (
	"flag"
	"fmt"
	"time"
)

/*
 for use in makefile with something like:

	# <Makefile> for program

	all:
		godatetime -pkg=main > compileDateTime.go
		go build

	# </Makefile>

	Then in program you can use: fmt.Printf("Compiled on %s\n", CompileDateTime)
*/
var pkgName string

func init() {
	flag.StringVar(&pkgName, "pkg", "main", "Name of package")
}

func main() {
	var t1 = `package `
	var t2 = `
var CompileDateTime = "`

	flag.Parse()
	fmt.Printf("%s%s%s%s\"\n", t1, pkgName, t2, time.Now())
}

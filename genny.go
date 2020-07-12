package main

import (
	"go/format"
	"io"
	"log"
	"os"

	"github.com/cheekybits/genny/parse"
)

func gennyGen(filename, pkgName string, typesets []map[string]string, out io.Writer) error {

	var output []byte
	var err error

	var file *os.File
	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	output, err = parse.Generics(filename, "outputFilename.go", pkgName, "", file, typesets)
	if err != nil {
		return err
	}

	fmtOutput, err := format.Source([]byte(output))
	if err != nil {
		log.Fatal(err)
	}
	out.Write(fmtOutput)
	return nil
}

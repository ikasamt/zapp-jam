package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	exitCode      = 0
	prefixAutoGen = "zzz-"
	buildTag      = `+jam`
)

var parseFlags = func() []string {
	flag.Parse()
	return flag.Args()
}

func rogerMain() {
	packagePaths := parseFlags()
	if len(packagePaths) == 0 {
		return
	}
	path := packagePaths[0]

	packageName, task := ParseDir(path)
	log.Println(fmt.Sprintf("# %s", packageName))
	for inFn, typesets := range task {
		_, fn := filepath.Split(inFn)
		outFn := filepath.Join(path, fmt.Sprintf("%s%s", prefixAutoGen, fn))
		log.Println(fmt.Sprintf("%s -> %s", inFn, outFn))
		out, err := os.Create(outFn)
		defer out.Close()
		if err != nil {
			log.Println(err)
			return
		}
		gennyGen(filepath.Join(path, inFn), packageName, typesets, out)
	}

}

func main() {
	rogerMain()
	os.Exit(exitCode)
}

package main

import (
	"bytes"
	"go/build"
	"go/format"
	"strings"

	// "bytes"
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
	validationSep = `,`
)

var GoPath string
var ZappJamSrcRoot string

const SemanticVersion = `v0.6.8`

func init() {
	GoPath = os.Getenv("GOPATH")
	if GoPath == "" {
		GoPath = build.Default.GOPATH
	}
	ZappJamSrcRoot = filepath.Join(GoPath, `/pkg/mod/github.com/ikasamt/zapp-jam@`+SemanticVersion)
	log.Println(ZappJamSrcRoot)
}

var parseFlags = func() []string {
	flag.Parse()
	return flag.Args()
}

func jamMain() {
	packagePaths := parseFlags()
	if len(packagePaths) == 0 {
		return
	}
	packagePath := packagePaths[0]

	packageName, task, validations := ParseDir(packagePath)
	log.Println(task)
	log.Println(fmt.Sprintf("# %s", packageName))
	for inFn, typesets := range task {
		_, fn := filepath.Split(inFn)
		outFn := filepath.Join(packagePath, fmt.Sprintf("%s%s", prefixAutoGen, fn))
		log.Println(fmt.Sprintf("%s -> %s", inFn, outFn))
		out, err := os.Create(outFn)
		defer out.Close()
		if err != nil {
			log.Println(err)
			return
		}
		gennyGen(inFn, packageName, typesets, out)
	}

	if len(validations) != 0 {
		for funcName, validation := range validations {
			switch funcName {
			case `Ngram`:
				var buf bytes.Buffer
				buf.WriteString("package " + packageName)
				buf.WriteString("\n\n// Auto-generated DO NOT EDIT!")
				buf.WriteString("\n")
				buf.WriteString("\n\nimport (")
				buf.WriteString("\"github.com/ikasamt/zapp/zapp\"")
				buf.WriteString("\n)")
				buf.WriteString("\n")
				buf.WriteString("\n")
				for _, value := range validation {
					for structName, fieldsStr := range value {
						fields := strings.Split(fieldsStr, validationSep)
						buf.WriteString("\n")
						fmt.Fprintf(&buf, "\nfunc (x %s) Ngrams() []string {", structName)
						fmt.Fprintf(&buf, "\nngrams := []string{}")
						for _, v := range fields {
							fmt.Fprintf(&buf, "\nngrams = append(ngrams, zapp.SplitNgramsRange(x.%s, 3)...)", v)
						}
						fmt.Fprintf(&buf, "\nreturn ngrams}")
						buf.WriteString("\n")
					}
				}
				buf.WriteString("\n")
				formatted, err := format.Source(buf.Bytes())
				if err != nil {
					log.Printf("%v", err)
					log.Printf("%s", buf.Bytes())
				}

				file := filepath.Join(packagePath, fmt.Sprintf("%s%s", prefixAutoGen, `ngram.go`))
				fh, err := os.Create(file)
				if err != nil {
					log.Printf(`failed to open file %s for writing`, file)
				}
				defer fh.Close()
				fh.Write(formatted)
			case `ValidatePresenceOf`:
				var buf bytes.Buffer
				buf.WriteString("package " + packageName)
				buf.WriteString("\n\n// Auto-generated DO NOT EDIT!")
				buf.WriteString("\n")
				buf.WriteString("\n\nimport (")
				buf.WriteString("  validation \"github.com/go-ozzo/ozzo-validation\"")
				buf.WriteString("\n)")
				buf.WriteString("\n")
				buf.WriteString("\n")
				for _, value := range validation {
					for structName, fieldsStr := range value {
						fields := strings.Split(fieldsStr, validationSep)
						buf.WriteString("\n")
						fmt.Fprintf(&buf, "\nfunc (x %s) Validations() error {", structName)
						fmt.Fprintf(&buf, "\nreturn validation.ValidateStruct(&x,")
						for _, v := range fields {
							fmt.Fprintf(&buf, "\nvalidation.Field(&x.%s, validation.Required),", v)
						}
						fmt.Fprintf(&buf, "\n)")
						fmt.Fprintf(&buf, "\n}")
						buf.WriteString("\n")
					}
				}
				buf.WriteString("\n")
				formatted, err := format.Source(buf.Bytes())
				if err != nil {
					log.Printf("%s", buf.Bytes())
				}

				file := filepath.Join(packagePath, fmt.Sprintf("%s%s", prefixAutoGen, `validate-presence-of.go`))
				fh, err := os.Create(file)
				if err != nil {
					log.Printf(`failed to open file %s for writing`, file)
				}
				defer fh.Close()
				fh.Write(formatted)
			}
		}

	}

}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	jamMain()
	os.Exit(exitCode)
}

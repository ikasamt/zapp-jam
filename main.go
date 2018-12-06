package main

import (
	"bytes"
	"github.com/iancoleman/strcase"
	"go/format"
	"strings"

	// "bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	// "strings"
)

const (
	exitCode      = 0
	prefixAutoGen = "zzz-"
	buildTag      = `+jam`
	validationSep = `,`
)

var parseFlags = func() []string {
	flag.Parse()
	return flag.Args()
}

func jamMain() {
	packagePaths := parseFlags()
	if len(packagePaths) == 0 {
		return
	}
	path := packagePaths[0]

	packageName, task, tasks2 := ParseDir(path)
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

	if len(tasks2) != 0 {
		for funcName, validation := range tasks2 {
			var buf bytes.Buffer
			buf.WriteString("package " + packageName)
			buf.WriteString("\n\n// Auto-generated DO NOT EDIT!")
			buf.WriteString("\n")
			baseName := ``
			if funcName == `Setter` {
				baseName = `setter.go`
				buf.WriteString("\n\nimport (")
				buf.WriteString("  \"github.com/gin-gonic/gin\"")
				buf.WriteString("  \"github.com/ikasamt/zapp/zapp\"")
				buf.WriteString("\n)")
				buf.WriteString("\n")
				buf.WriteString("\n")
				for _, value := range validation {
					for structName, fieldsStr := range value {
						fields := strings.Split(fieldsStr, validationSep)
						buf.WriteString("\n")
						fmt.Fprintf(&buf, "\nfunc (x *%s) Setter(c *gin.Context) {", structName)
						for _, fieldName := range fields {
							// fieldType := getFieldType("./"+path, packageName, structName, fieldName)
							snakeFieldName := strcase.ToSnake(fieldName)
							fmt.Fprintf(&buf, "\nx.%s = zapp.GetParams(c, \"%s\"),", fieldName, snakeFieldName)
						}
						fmt.Fprintf(&buf, "\n}")
						buf.WriteString("\n")
					}
				}
			} else if funcName == `ValidatePresenceOf` {
				baseName = `validate-presence-of.go`
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
			}
			buf.WriteString("\n")
			formatted, err := format.Source(buf.Bytes())
			if err != nil {
				log.Printf("%s", buf.Bytes())
			}

			file := filepath.Join(path, fmt.Sprintf("%s%s", prefixAutoGen, baseName))
			fh, err := os.Create(file)
			if err != nil {
				log.Printf(`failed to open file %s for writing`, file)
			}
			defer fh.Close()
			fh.Write(formatted)
		}

	}

}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	jamMain()
	os.Exit(exitCode)
}

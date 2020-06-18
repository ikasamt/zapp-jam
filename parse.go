package main

import (
	"go/doc"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"strings"
)

type Task map[string][]map[string]string
type Validation map[string][]map[string]string

func ParseDir(packagePath string) (packageName string, task Task, validation Validation) {
	task = Task{}
	validation = Validation{}

	fset := token.NewFileSet()
	d, err := parser.ParseDir(fset, packagePath, nil, parser.ParseComments)
	if err != nil {
		log.Println(err)
		return
	}

	for pn, f := range d {
		packageName = pn
		p := doc.New(f, "./", 0)
		for _, t := range p.Types {
			doc := t.Doc
			structName := t.Name
			lines := strings.Split(doc, "\n")
			for _, line := range lines {
				if strings.Contains(line, buildTag) {
					tmp := strings.Fields(line)
					switch len(tmp) {
					case 1:
						continue // needs filename
					case 2:
						filename := tmp[1]
						if strings.HasPrefix(filename, `/clefs/`) {
							filename = filepath.Join(ZappJamSrcRoot, filename)
						} else {
							filename = filepath.Join(packagePath, filename)
						}
						task[filename] = append(task[filename], map[string]string{`Anything`: structName})
					default:
						if strings.Contains(tmp[1], `.go`) {
							filename := tmp[1]
							if strings.HasPrefix(filename, `/clefs/`) {
								filename = filepath.Join(ZappJamSrcRoot, filename)
							} else {
								filename = filepath.Join(packagePath, filename)
							}
							args := tmp[2:]
							for _, arg := range args {
								task[filename] = append(task[filename], map[string]string{`Anything`: structName, `Something`: arg})
							}
						} else {
							key := tmp[1]
							validation[key] = append(validation[key], map[string]string{structName: tmp[2]})
						}
					}
				}
			}
		}
	}
	return
}

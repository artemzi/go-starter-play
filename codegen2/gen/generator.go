package main

import (
	"go/parser"
	"go/token"
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl = `package {{.Package}}

type {{.MyType}}Queue struct {
	q []{{.MyType}}
}

func New{{.MyType}}Queue() *{{.MyType}}Queue {
	return &{{.MyType}}Queue {
		q: []{{.MyType}}{},
	}
}

func (o *{{.MyType}}Queue) Insert(v {{.MyType}}) {
	o.q = append(o.q, v)
}

func (o *{{.MyType}}Queue) Remove() {{.MyType}} {
	if len(o.q) == 0 {
		panic("Oops.")
	}
	first := o.q[0]
	o.q = o.q[1:]
	return first
}
`

func main() {
	tt := template.Must(template.New("queue").Parse(tpl))
	// first argument is filename for getting package name so let start loop from 2
	for i := 2; i < len(os.Args); i++ {
		dest := strings.ToLower(os.Args[i]) + "_queue.go"
		file, err := os.Create(dest)
		if err != nil {
			log.Printf("Could not create %s: %s (skip)\n", dest, err)
			continue
		}

		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, os.Args[1], nil, parser.ParseComments)
		if err != nil {
			log.Fatal(err)
		}

		vals := map[string]string{
			"MyType":  os.Args[i],
			"Package": node.Name.Name,
		}
		tt.Execute(file, vals)

		file.Close()
	}
}

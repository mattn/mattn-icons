package main

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"text/template"
)

const tmpl = `# Icon Collections
{{range .}}
![{{.}}](https://raw.githubusercontent.com/mattn/mattn-icons/master/{{.}})
{{end}}

# License

NYSL: http://www.kmonos.net/nysl/index.en.html
`

type item struct {
	Title string
	URL   string
}

func main() {
	m, err := filepath.Glob("logo*.png")
	if err != nil {
		log.Fatal(err)
	}
	sort.Strings(m)
	f, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	template.Must(template.New("readme").Parse(tmpl)).Execute(f, m)
}

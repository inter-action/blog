package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

/*
{
    entries:[
        {
            contents:[{title: ,}],
            entries:[
                contents:{},
                entries:
            ]
        }
    ]
}
*/
type Root struct {
	entries []Entries
}

type Entries struct {
	contents []Content
	entries  []Entries
	path     string
}

type Content struct {
	title string
}

var ignores = []string{"tpl", ".git", ".gitignore", ".DS_STORE"}

const HOST_ROOT = "https://github.com/inter-action/blog/blob/master/"
const OUTPUT_FILE = "./README.md"

// this func may cause stackoverflow, refractor this if needed
func walk_file(path string) Entries {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	entries := Entries{}
	entries.path = path

	for _, fi := range files {
		if contains(ignores, fi.Name()) {
			continue
		}

		if fi.IsDir() {
			entries.entries = append(entries.entries, walk_file(filepath.Join(path, fi.Name())))
		} else {
			entries.contents = append(entries.contents, Content{title: fi.Name()})

		}
	}

	return entries
}

func contains(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func join_contents(contents []Content, path string) string {
	var r []string

	for _, c := range contents {
		//todo: find a way to encodingURI in golang
		// fullpath := filepath.Join(HOST_ROOT, path, c.title)
		// subpath := url.QueryEscape(fullpath[len(HOST_ROOT)-1:])
		r = append(r, fmt.Sprintf("[%s](%s)  ", c.title, filepath.Join(HOST_ROOT, path, url.QueryEscape(c.title))))
	}

	return strings.Join(r, "\n")
}

func walk(entries []Entries, depth int) string {
	var rs string
	for _, entry := range entries {
		rs += strings.Repeat("#", depth+3) + entry.path + "\n"
		rs += join_contents(entry.contents, entry.path) + "\n"
		rs += walk(entry.entries, depth+1)
	}

	return rs
}

func walk_root(root Root) string {
	depth := 1
	return walk(root.entries, depth)
}

func GenIndex() string {
	_filePath, err := os.Getwd()

	if err != nil {
		log.Fatal("error: failed to get file path", _filePath)
	}

	entries := walk_file("./")
	_root := Root{entries: []Entries{entries}}
	fmt.Println(_root)

	return walk_root(_root)
}

func test_struct() {
	c1 := Content{}
	c1.title = "content lvl 1"

	e1 := Entries{}
	e1.contents = append(e1.contents, c1)

	e2 := Entries{}
	c2 := Content{title: "content lvl 2"}
	e2.contents = append(e2.contents, c2)

	e1.entries = append(e1.entries, e2)

	root := Root{}
	root.entries = append(root.entries, e1)

	if out, err := json.Marshal(&root); err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(string(out), root, root.entries[0].contents[0].title)
	}
}

func main() {
	context := template.FuncMap{
		"GenIndex": GenIndex,
	}
	// t, err := template.New("readme").ParseFiles(filepath.Join(_filePath, "tpl/readme.md.tpl"))
	t, err := template.New("_readme_tpl.md").Funcs(context).ParseFiles("./tpl/_readme_tpl.md")
	if err != nil {
		log.Fatal(err)
	}

	outfile, err := os.Create(OUTPUT_FILE)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(outfile, nil)
	if err != nil {
		log.Fatal(err)
	}
}
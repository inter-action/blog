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

// 默认隐藏所有以 .开头的 dir or file
var ignores = []string{"tpl"}

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
		if strings.HasPrefix(fi.Name(), ".") || contains(ignores, fi.Name()) {
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
		e := encodeURI(c.title)
		r = append(r, fmt.Sprintf("[%s](%s)  ", c.title, "./"+filepath.Join(path, e)))
	}

	return strings.Join(r, "\n")
}

func encodeURI(filepath string) string {
	s := url.QueryEscape(filepath)
	s = strings.Replace(s, "+", "%20", -1)
	return s
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
	// fmt.Println(_root)

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
	//这个地方的文件名 new 和 parseFiles 这两个地方必须一致，涉及到golang 内部实现。太坑了！具体看下面链接
	//http://stackoverflow.com/questions/10199219/go-template-function
	t, err := template.New("_readme_tpl.md").Funcs(context).ParseFiles("./tpl/_readme_tpl.md")
	if err != nil {
		log.Fatal(err)
	}

	outfile, err := os.Create(OUTPUT_FILE)
	defer outfile.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = t.Execute(outfile, nil)
	if err != nil {
		log.Fatal(err)
	}
}

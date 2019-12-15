// +build ignore

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

const (
	dictFileLocation = "/usr/share/dict"
	ignoreFile       = "README"
	dictFileTemplate = `
	package dict
	
	
	var DictionaryBytes = []byte{ {{ range .DictBytes -}} 0x{{printf "%X" .}}, {{- end }} }
	var DictionaryString = []string{ {{ range .DictString -}} {{ printf "%q" .}}, {{- end }} }
	var DictionaryBytesWords = [][]byte{ {{ range .DictString -}} []byte({{ printf "%q" .}}), {{- end }} } 
	
	
	`
)

var dictTemplate = template.Must(template.New("").Parse(dictFileTemplate))

type Dictionary struct {
	DictBytes []byte
	DictString []string
}




func main() {
	dirList, err := ioutil.ReadDir(dictFileLocation)
	if err != nil {
		log.Fatalln("Unable to read dict file location", err)
	}
	dictWords := map[string]struct{}{}
	for _, file := range dirList {
		if file.Name() == ignoreFile {
			continue
		}
		fileName := filepath.Join(dictFileLocation, file.Name())
		fmt.Println("Reading ", fileName)
		dictFile, err  := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatalln("Unable to read dict file location", err)
		}

		words := bytes.Split(dictFile, []byte{'\n'})
		for _, b := range words {
			dictWords[string(b)] = struct{}{}
		}
	}
	d := Dictionary{}
	dictBytes := &bytes.Buffer{}
	for k := range dictWords {
		dictBytes.Write([]byte(k))
		d.DictString = append(d.DictString, k)
	}
	d.DictBytes = dictBytes.Bytes()
	f, err := os.Create("dict.go")
	if err != nil {
		log.Fatalln("Unable to create dict.go", err)
	}
	if err := dictTemplate.Execute(f, d); err != nil {
		log.Fatalln("Unable to execute template", err)
	}
}

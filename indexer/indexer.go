package indexer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"text/scanner"
)

//Mapping vocab terms to document names
type InvertedIndex struct {
	index map[string][]string
}

type Document struct {
	Title  string
	Author string
	Body   string
}

func Index(files []string) InvertedIndex {
	var index InvertedIndex
	index.index = make(map[string][]string)
	var doc Document
	for _, file := range files {
		raw, _ := ioutil.ReadFile(file)
		if err := json.Unmarshal(raw, &doc); err != nil {
			log.Fatal(err)
		}
		tokenize(index, doc.Body, file)
	}
	return index
}

func tokenize(index InvertedIndex, text string, docName string) {
	var s scanner.Scanner
	s.Init(strings.NewReader(text))
	for token := s.Scan(); token != scanner.EOF; token = s.Scan() {
		//fmt.Printf("%d: %s\n", s.Position.Offset, s.TokenText())
		index.index[s.TokenText()] = append(index.index[s.TokenText()], docName)
	}
}

func (index InvertedIndex) String() string {
	var buffer bytes.Buffer
	for key, value := range index.index {
		buffer.WriteString(fmt.Sprint(key, ": ", value, "\n"))
	}
	return buffer.String()
}

func (index InvertedIndex) Size() int {
	return len(index.index)
}

func (index InvertedIndex) Search(query string) []string {
	return index.index[query]
}

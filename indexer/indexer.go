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

type InvertedIndex struct {
	dict map[string]map[int][]int
	docs map[int]string
}

type Document struct {
	Title  string
	Author string
	Body   string
}

func Index(files []string) InvertedIndex {
	var index InvertedIndex
	index.dict = make(map[string]map[int][]int)
	index.docs = make(map[int]string)

	var doc Document
	for _, file := range files {
		raw, _ := ioutil.ReadFile(file)
		if err := json.Unmarshal(raw, &doc); err != nil {
			log.Fatal(err)
		}
		docId := len(index.docs)
		index.docs[docId] = file
		tokenize(index, docId, doc.Body)
	}
	return index
}

func tokenize(index InvertedIndex, docId int, text string) {
	var s scanner.Scanner
	s.Init(strings.NewReader(text))
	for token := s.Scan(); token != scanner.EOF; token = s.Scan() {
		if docMap, ok := index.dict[s.TokenText()]; ok {
			if posting, ok := docMap[docId]; ok {
				index.dict[s.TokenText()][docId] = append(posting, s.Position.Offset)
			} else {
				index.dict[s.TokenText()][docId] = []int{s.Position.Offset}
			}
		} else {
			index.dict[s.TokenText()] = map[int][]int{docId: {s.Position.Offset}}
		}
	}
}

func (index InvertedIndex) String() string {
	var buffer bytes.Buffer
	for key, value := range index.dict {
		buffer.WriteString(fmt.Sprint(key, ": ", value, "\n"))
	}
	return buffer.String()
}

func (index InvertedIndex) Size() int {
	return len(index.dict)
}

func (index InvertedIndex) Search(query string) map[int][]int {
	return index.dict[query]
}

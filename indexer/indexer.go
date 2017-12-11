package indexer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"unicode"
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
	for docId, file := range files {
		//go func(file string) {
		raw, _ := ioutil.ReadFile(file)
		if err := json.Unmarshal(raw, &doc); err != nil {
			log.Fatal(err)
		}
		index.docs[docId] = file
		tokenize(index, docId, doc.Body)
		//}(file)
	}
	return index
}

func tokenize(index InvertedIndex, docId int, text string) {
	var buffer bytes.Buffer
	for currPos, rune := range text {
		switch {
		case unicode.IsPunct(rune):
		case unicode.IsSpace(rune):
			if buffer.Len() != 0 {
				putToken(index, docId, buffer.String(), currPos-buffer.Len())
				buffer.Reset()
			}
		default:
			buffer.WriteRune(rune)
		}
	}
	if buffer.Len() != 0 {
		putToken(index, docId, buffer.String(), len(text)-buffer.Len())
	}
}

func putToken(index InvertedIndex, docId int, token string, offset int) {
	if _, ok := index.dict[token]; ok {
		index.dict[token][docId] = append(index.dict[token][docId], offset)
	} else {
		index.dict[token] = map[int][]int{docId: {offset}}
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

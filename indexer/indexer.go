package indexer

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"log"
	"bytes"
	"time"
)

type InvertedIndex struct {
	data map[string]interface{}
}

func Index(files []string) (InvertedIndex, time.Duration)  {
	start := time.Now();
	var index InvertedIndex
	for _, file := range files {
		fmt.Println(file)
		raw, _ := ioutil.ReadFile(file)
		if err := json.Unmarshal(raw, &index.data); err != nil {
			log.Fatal(err)
		}
	}
	return index, time.Now().Sub(start)
}

func (index InvertedIndex) String() string {
	var buffer bytes.Buffer
	for key, value := range index.data {
		buffer.WriteString(fmt.Sprint(key, ": ", value, "\n"))
	}
	return buffer.String()
}
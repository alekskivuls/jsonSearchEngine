package main

import (
	"fmt"
	"os"
	"time"

	"github.com/alekskivuls/jsonSearchEngine/common"
	"github.com/alekskivuls/jsonSearchEngine/indexer"
)

func main() {
	files := common.GetJsonFilesFromPath(os.Args[1:])
	start := time.Now()
	fmt.Println(files)
	index := indexer.Index(files)
	fmt.Println("Took", time.Now().Sub(start), "to index")
	fmt.Println(index.Size())

	fmt.Print("Enter search term: ")
	var input string
	fmt.Scanln(&input)
	fmt.Print(input)
	fmt.Println(index.Search(input))
}

package main

import (
	"fmt"
	"os"

	"github.com/alekskivuls/jsonSearchEngine/common"
	"github.com/alekskivuls/jsonSearchEngine/indexer"
)

func main() {
	files := common.GetJsonFilesFromPath(os.Args[1:])
	index, duration := indexer.Index(files)
	fmt.Println("Took", duration, "to index")
	fmt.Print(index)
}

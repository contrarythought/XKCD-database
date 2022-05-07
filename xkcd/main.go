package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"xkcd/xkcdJSON"
)

const DB_NAME = "xkcdDB.txt"

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	contents, err := os.ReadDir(currentDir)
	if err != nil {
		log.Fatal(err)
	}

	var containsDB bool = false
	for i := 0; i < len(contents); i++ {
		if strings.Compare(DB_NAME, contents[i].Name()) == 0 {
			containsDB = true
			break
		}
	}

	var database map[int]xkcdJSON.Comic
	fastLookup := make(map[string]string)
	if !containsDB {
		file, err := os.Create(DB_NAME)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		database, err = xkcdJSON.BuildXKCDDatabase()
		if err != nil {
			err = fmt.Errorf("failed to build database")
			if err != nil {
				log.Fatal(err)
			}
		}

		for i := 0; i < len(database); i++ {
			//fmt.Println(database[i].Title)
			file.WriteString(database[i].Title + "\n")
			file.WriteString(database[i].Img + "\n")
			fastLookup[database[i].Title] = database[i].Img
		}

		// grab user input

	} else {
		// grab most recent comic
		comic, err := xkcdJSON.GetNewComic()
		if err != nil {
			log.Fatal(err)
		}

		// obtain user input
	}

}

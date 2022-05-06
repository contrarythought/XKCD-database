package xkcdJSON

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

const (
	HOST    = "https://xkcd.com/"
	URL_END = "/info.0.json"
)

func BuildXKCDDatabase() (map[int]Comic, error) {
	database := make(map[int]Comic)
	issueNumber := 1

	for {
		query, err := http.Get(HOST + strconv.Itoa(issueNumber) + URL_END)
		if err != nil {
			return nil, err
		}

		//fmt.Println("searching: ", HOST+strconv.Itoa(issueNumber)+URL_END)

		if query.StatusCode != http.StatusOK {
			fmt.Println("finished parsing")
			query.Body.Close()
			break
		}

		// get a comic
		var comic Comic
		if err := json.NewDecoder(query.Body).Decode(&comic); err != nil {
			query.Body.Close()
			fmt.Println("Failed to decode: ", err.Error())
			return nil, err
		}

		database[issueNumber] = comic
		issueNumber++
	}
	return database, nil
}

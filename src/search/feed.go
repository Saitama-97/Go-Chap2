package search

/**
 * @Time    : 2023/6/25 16:33
 * @File    : feed.go
 * @Project : chapter2
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    :
 */

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// defer
	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}

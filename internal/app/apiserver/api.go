package apiserver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func printTime() {
	dt := time.Now()
	fmt.Println("time: ", dt.String())
}

func AddFile() error {
	var path2 = "https://docs.google.com/spreadsheets/d/e/2PACX-1vTQn40xmfHXLqWjLXBZ3UzfxAh6h9lB8z9_7sne33xfugzCwotWA5j1W80skZ7-iyIPizNYpcB9lVqR/pub?output=xlsx"

	// getting xlsx file from url
	resp, err := http.Get(path2)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// copying from response to byte slice
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	rowStuff(body)
	return nil
}

package apiserver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//"https://docs.google.com/spreadsheets/d/e/2PACX-1vTQn40xmfHXLqWjLXBZ3UzfxAh6h9lB8z9_7sne33xfugzCwotWA5j1W80skZ7-iyIPizNYpcB9lVqR/pub?output=xlsx"

func printTime() {
	dt := time.Now()
	fmt.Println("time: ", dt.String())
}

func AddFile(a *AddRequest) error {

	// getting xlsx file from url
	resp, err := http.Get(a.FileLink)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// copying from response to byte slice
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	ParseXlsx(body)
	return nil
}

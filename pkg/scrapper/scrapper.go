package scrapper

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func DownloadContent(url string, as string) {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic(errors.New(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status)))
	}

	log.Println("Downloaded content from: ", url)

	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file, err := os.Create(basePath + "/html/" + as + ".html")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	length, err := io.Copy(file, res.Body)

	if err != nil {
		panic(err)
	}

	log.Println("Downloaded content length: ", length)
}

func ParseFile(fileName string) string {
	var description string

	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := basePath + "/html/" + fileName + ".html"

	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		panic(err.Error())
	}

	text := doc.Find(".infobox ~ p").First().Text()

	reg, err := regexp.Compile(`\[\d+\]`)
	if err != nil {
		panic(err)
	}

	description = reg.ReplaceAllString(text, "")

	return description
}

package api

import (
	"bufio"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

// For now since pypi index is quite huge with over 500k entries
// Updating the indexes is optional
const PyPiIndex = "https://pypi.org/simple/"

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// This should be one time thing if user want to re-fetch index from Pypi
func GenerateIndex() {
	htmlIndex, e := os.Open("index/pypi.html")
	check(e)
	defer htmlIndex.Close()

	var doc *goquery.Document
	if doc, e = goquery.NewDocumentFromReader(htmlIndex); e != nil {
		log.Fatal(e)
	}

	var index []string

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		index = append(index, s.Text())
	})

	textIndex, err := os.Create("index/pypi.index")
	check(err)

	w := bufio.NewWriter(textIndex)

	for _, data := range index {
    log.Println("Writing>")
		_, _ = w.WriteString(data + "\n")
	}

	w.Flush()
	textIndex.Close()
}

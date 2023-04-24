package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

type Fact struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func main() {
	allFacts := make([]Fact, 0)

	collector := colly.NewCollector(colly.AllowedDomains("factretriever.com", "www.factretriever.com"))

	collector.OnHTML(".factsList li", func(element *colly.HTMLElement) {
		factId, err := strconv.Atoi(element.Attr("id"))
		if err != nil {
			log.Println("Failed to retrieve the ID")
		}

		factDesc := element.Text

		fact := Fact{
			ID:          factId,
			Description: factDesc,
		}

		allFacts = append(allFacts, fact)

	})
	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL.String())
	})
	collector.Visit("https://www.factretriever.com/dinosaur-facts")

	writeJson(allFacts)

}

func writeJson(data []Fact) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create the JSON File.")
		return
	}
	_ = ioutil.WriteFile("dinosaurFacts.json", file, 0644)
}

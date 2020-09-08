package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	)

func exampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://www.ah.nl/allerhande/recept/R-R1194290/makkelijke-risotto")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the ingredients, split them and print the right data
	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		attr, _ := s.Attr("itemprop")
		if attr == "ingredients" {
			list := strings.Split(strings.Replace(s.Text(), "\n", "", -1), " ")
			fmt.Println(list[len(list) - 1])
		}
	})
}

func getIngredientsFromRecipe(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the ingredients, split them and print the right data
	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		attr, _ := s.Attr("itemprop")
		if attr == "ingredients" {
			//list := strings.Split(strings.Replace(s.Text(), "\n", "", -1), " ")
			//fmt.Println(list[len(list) - 1])

			reg, err := regexp.Compile("[^a-zA-Z() ]+")
			if err != nil {
				log.Fatal(err)
			}

			str := reg.ReplaceAllString(strings.Replace(s.Text(), "\n", "", -1), "")

			fmt.Println(strings.Replace(str, " ", "", 1))
		}
	})
}


func main() {
	//exampleScrape()

	file, err := ioutil.ReadFile("./gerechten.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	for _, line := range lines {
		getIngredientsFromRecipe(line)
	}
}

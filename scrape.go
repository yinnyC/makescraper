package main

import (
	"encoding/json"
	"os"

	"github.com/gocolly/colly"
)

type foodRecipe struct {
	Title string `json:"title"`
	Img   string `json:"img"`
	Link  string `json:"href"`
}

func main() {
	c := colly.NewCollector()

	c.OnHTML("article.category-taiwanese", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		img := e.ChildAttr("img", "src")
		title := e.ChildText("h3")
		print(link, img, title)

		foodRecipe := foodRecipe{
			Title: title,
			Img:   img,
			Link:  link,
		}
		foodRecipeJson, _ := json.Marshal(foodRecipe)

		f, err := os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		foodRecipeStr := string(foodRecipeJson)

		if _, err = f.WriteString(foodRecipeStr); err != nil {
			panic(err)
		}

	})

	c.Visit("https://chejorge.com/")

}

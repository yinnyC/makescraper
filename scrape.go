package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type foodRecipe struct {
	Title string `json:"title"`
	Img   string `json:"img"`
	Href  string `json:"href"`
}

func main() {
	var recipes []foodRecipe
	c := colly.NewCollector()
	c.OnHTML("article.category-taiwanese", func(e *colly.HTMLElement) {
		href := e.ChildAttr("a", "href")
		img := e.ChildAttr("img", "src")
		title := e.ChildText("h3")

		recipes = append(recipes, foodRecipe{
			Title: title,
			Img:   img,
			Href:  href,
		})
	})
	c.Visit("https://chejorge.com/")

	foodRecipeJson, _ := json.MarshalIndent(recipes, "", " ")

	print(string(foodRecipeJson))

	// writing json to file
	_ = ioutil.WriteFile("output.json", foodRecipeJson, 0644)

	// to append to a file
	// create the file if it doesn't exists with O_CREATE, Set the file up for read write, add the append flag and set the permission
	f, err := os.OpenFile("./debug-web.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Fatal(err)
	}
	// write to file, f.Write()
	f.Write(foodRecipeJson)

}

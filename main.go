package main

import "github.com/gin-gonic/gin"

type person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func main() {

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", index)
	// router.POST("/", postIndex)
	router.GET("/json", json)
	router.GET("/htmlpage", htmlPage)
	// router.POST("/postjson", postJson)
	router.GET("/blog/:author/:article", blog)
	router.GET("/query", query)

	router.Run(":9000")

}

func index(c *gin.Context) {
	c.String(200, "Hello world!")
}

func json(c *gin.Context) {
	c.String(200, "Sade!\n")

	c.JSON(200, gin.H{
		"name":    "Furkan Sade",
		"surname": "Uçkun",
		"age":     21,
	})
}

func htmlPage(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Home Page!",
	})
}

func blog(c *gin.Context) {
	author := c.Param("author")
	article := c.Param("article")
	c.String(200, "Yazar: "+author+"\nMakale: "+article)
}

// http://localhost:9000/query?tur=bilim-kurgu&siralama=imdb
func query(c *gin.Context) {
	tur := c.Query("tur")
	siralama := c.Query("siralama")
	c.String(200, tur+" türünden filmler "+siralama+" olarak sıralanıyor.")

}

// func postIndex(c *gin.Context) {
// 	c.String(200, "Connect with post method!")
// }

// func postJson(c *gin.Context) {
// 	c.String(200, "Connect to post method!\n")

// 	var postPerson person

// 	c.BindJSON(&postPerson)

// 	c.String(200, "Json Data:")
// 	c.JSON(200, postPerson)

// }

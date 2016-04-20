package main

import (
	"log"
	"net/http"
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
)

func main() {
	// githubKey := os.Getenv("GITHUB_KEY")
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	githubClient := github.NewClient(nil)

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		page := 1

		opts := &github.SearchOptions{
			Sort: "stars",
      ListOptions: github.ListOptions{
		    PerPage: 100,
				Page: page,
      },
    }
		repos, _, err := githubClient.Search.Repositories("language:go", opts)

		if err != nil {
			fmt.Println("Github API Error: ", err)
			c.HTML(http.StatusOK, "index.tmpl.html", nil)
			return;
		}
		obj := gin.H{"Repos": repos.Repositories, "TotalRepos": repos.Total}
		c.HTML(http.StatusOK, "index.tmpl.html", obj)
	})

	router.Run(":" + port)
}

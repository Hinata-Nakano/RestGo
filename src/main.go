package main

import (
	"example.com/RestCRUD/presentation/http"
)

type post struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

var posts = []post{
	{ID: "1", Title: "Title 1", Content: "Content 1", CreatedAt: "2025-01-01"},
	{ID: "2", Title: "Title 2", Content: "Content 2", CreatedAt: "2025-01-02"},
	{ID: "3", Title: "Title 3", Content: "Content 3", CreatedAt: "2025-01-03"},
}

func main() {
	dependencies := InjectDependencies()
	tweetHandler := http.NewTweetHandler(dependencies.TweetDependencies.CreateTweetUseCase)
	router := http.SetupRouter(tweetHandler)

	router.Run("localhost:3000")
}

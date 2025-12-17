package main

import (
	"example.com/RestCRUD/infrastructure/clock"
	"example.com/RestCRUD/infrastructure/repository"
	tweetusecase "example.com/RestCRUD/usecase/tweet"
)

type TweetDependencies struct {
	CreateTweetUseCase *tweetusecase.CreateTweetUseCase
}

type Dependencies struct {
	TweetDependencies *TweetDependencies
}

func InjectDependencies() *Dependencies {
	return &Dependencies{
		TweetDependencies: &TweetDependencies{
			CreateTweetUseCase: tweetusecase.NewCreateTweetUseCase(repository.NewTweetRepository(), clock.RealClock{}),
		},
	}
}

package repository

import (
	"example.com/RestCRUD/domain/tweet"
)

// このファイルはすべて仮の内容です。実際のデータベースとは関係ありません。
type mockDB struct {
	createFunc func(tweet tweet.Tweet) (tweet.Tweet, error)
}

type TweetRepositoryImpl struct {
	db mockDB
}

func (m *TweetRepositoryImpl) Create(t tweet.Tweet) (tweet.Tweet, error) {
	return m.db.createFunc(t)
}

func NewTweetRepository() tweet.TweetRepository {
	return &TweetRepositoryImpl{db: mockDB{createFunc: func(tweet tweet.Tweet) (tweet.Tweet, error) {
		return tweet, nil
	}}}
}

package tweet

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type createTweetRequest struct {
	UserID  string `json:"user_id"`
	Content string `json:"content"`
}

// errorじゃなくてresult型を定義してもいいかも？
func CreateTweet(request createTweetRequest) (TweetEntity, error) {
	if request.UserID == "" {
		return TweetEntity{}, errors.New("user_id is required")
	}
	if request.Content == "" {
		return TweetEntity{}, errors.New("content is required")
	}
	tweetEntity := TweetEntity{
		ID:        uuid.New().String(),
		UserID:    request.UserID,
		Content:   request.Content,
		CreatedAt: time.Now().Format(time.RFC3339),
	}
	return tweetEntity, nil
}

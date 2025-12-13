// domain/tweet/tweet.go
package tweet

import (
	"example.com/RestCRUD/domain"
	"github.com/google/uuid"
	"time"
)

type Tweet struct {
	ID        string
	UserID    string
	Content   string
	CreatedAt string
}

type CreateTweetRequest struct {
	UserID  string
	Content string
}

// NewTweet は新しいTweetエンティティを作成する
// バリデーションを通過した有効なエンティティのみを返す
func NewTweet(createTweetRequest CreateTweetRequest, clock domain.Clock) (Tweet, error) {
	if createTweetRequest.UserID == "" {
		return Tweet{}, domain.ErrUserIDRequired
	}
	if createTweetRequest.Content == "" {
		return Tweet{}, domain.ErrContentRequired
	}

	return Tweet{
		ID:        uuid.New().String(),
		UserID:    createTweetRequest.UserID,
		Content:   createTweetRequest.Content,
		CreatedAt: clock.Now().Format(time.RFC3339),
	}, nil
}

// domain/tweet/tweet.go
package tweet

import (
	"example.com/RestCRUD/domain"
	"github.com/google/uuid"
	"time"
)

type TweetID string
type TweetContent string

func newTweetContent(content string) TweetContent {
	return TweetContent(content)
}

func newTweetID() TweetID {
	return TweetID(uuid.New().String())
}

type Tweet struct {
	ID        TweetID
	UserID    string //ここもいずれ変更したい
	Content   TweetContent
	CreatedAt string //ここもいずれ変更したい。domainで形式を定義して、インフラ層でそのようにパースするような依存関係にしたい。
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
		ID:        newTweetID(),
		UserID:    createTweetRequest.UserID,
		Content:   newTweetContent(createTweetRequest.Content),
		CreatedAt: clock.Now().Format(time.RFC3339),
	}, nil
}

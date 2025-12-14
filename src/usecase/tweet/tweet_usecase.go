package tweet

import (
	"example.com/RestCRUD/domain"
	"example.com/RestCRUD/domain/tweet"
)

// CreateTweetUseCase はツイート作成のユースケース
type CreateTweetUseCase struct {
	repo  tweet.TweetRepository
	clock domain.Clock
}

// NewCreateTweetUseCase はCreateTweetUseCaseのコンストラクタ
func NewCreateTweetUseCase(repo tweet.TweetRepository, clock domain.Clock) *CreateTweetUseCase {
	return &CreateTweetUseCase{
		repo:  repo,
		clock: clock,
	}
}

// Execute はツイート作成ユースケースを実行する
// 1. ドメインサービスを呼び出してエンティティを作成
// 2. リポジトリを使って永続化
func (u *CreateTweetUseCase) Execute(userID, content string) (*tweet.Tweet, error) {
	// ドメインサービスを呼び出してエンティティを作成
	request := tweet.CreateTweetRequest{
		UserID:  userID,
		Content: content,
	}

	tweetEntity, err := tweet.NewTweet(request, u.clock)
	if err != nil {
		return nil, err
	}

	// リポジトリを使って永続化
	savedTweet, err := u.repo.Create(tweetEntity)
	if err != nil {
		return nil, err
	}

	return &savedTweet, nil
}

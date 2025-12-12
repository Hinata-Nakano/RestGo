package tweet

type ITweetRepository interface {
	CreateTweet(tweetEntity TweetEntity) (TweetEntity, error)
}

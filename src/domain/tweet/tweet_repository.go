package tweet

type TweetRepository interface {
	Create(tweet Tweet) (Tweet, error)
}

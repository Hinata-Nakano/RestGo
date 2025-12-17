package dto

// CreateTweetRequest はツイート作成リクエストのDTO
type CreateTweetRequest struct {
	UserID  string `json:"user_id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// TweetResponse はツイートレスポンスのDTO
type TweetResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

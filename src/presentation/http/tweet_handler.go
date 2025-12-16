package http

import (
	"net/http"

	"example.com/RestCRUD/domain"
	tweetdomain "example.com/RestCRUD/domain/tweet"
	"example.com/RestCRUD/presentation/dto"
	tweetusecase "example.com/RestCRUD/usecase/tweet"
	"github.com/gin-gonic/gin"
)

// TweetHandler はツイート関連のHTTPハンドラー
type TweetHandler struct {
	createUseCase *tweetusecase.CreateTweetUseCase
}

// NewTweetHandler はTweetHandlerのコンストラクタ
func NewTweetHandler(createUseCase *tweetusecase.CreateTweetUseCase) *TweetHandler {
	return &TweetHandler{
		createUseCase: createUseCase,
	}
}

// CreateTweet はツイート作成のHTTPハンドラー
func (h *TweetHandler) CreateTweet(c *gin.Context) {
	// リクエストをDTOにバインド
	var req dto.CreateTweetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
		})
		return
	}

	// ユースケースを実行
	createdTweet, err := h.createUseCase.Execute(req.UserID, req.Content)
	if err != nil {
		// ドメインエラーのハンドリング
		switch err {
		case domain.ErrUserIDRequired, domain.ErrContentRequired:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		default:
			// その他のエラー（リポジトリエラーなど）
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create tweet",
			})
			return
		}
	}

	// エンティティをDTOに変換して返す
	response := h.toResponse(createdTweet)
	c.JSON(http.StatusCreated, response)
}

// toResponse はTweetエンティティをTweetResponse DTOに変換する
func (h *TweetHandler) toResponse(t *tweetdomain.Tweet) dto.TweetResponse {
	return dto.TweetResponse{
		ID:        t.ID,
		UserID:    t.UserID,
		Content:   t.Content,
		CreatedAt: t.CreatedAt,
	}
}

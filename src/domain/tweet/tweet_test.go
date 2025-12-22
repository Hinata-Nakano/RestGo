package tweet

import (
	"testing"
	"time"

	"example.com/RestCRUD/domain"
)

// FixedClock はテスト用の固定時刻を返すClock実装
type FixedClock struct {
	FixedTime time.Time
}

func (f FixedClock) Now() time.Time {
	return f.FixedTime
}

func TestNewTweet(t *testing.T) {
	// 固定時刻を設定
	fixedTime := time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)
	clock := FixedClock{FixedTime: fixedTime}

	tests := []struct {
		name        string
		userID      string
		content     string
		clock       domain.Clock
		wantErr     error
		wantUserID  string
		wantContent TweetContent
		wantTime    string
	}{
		{
			name:        "正常系: 有効なUserIDとContentでTweetを作成",
			userID:      "user123",
			content:     "Hello, World!",
			clock:       clock,
			wantErr:     nil,
			wantUserID:  "user123",
			wantContent: TweetContent("Hello, World!"),
			wantTime:    "2025-01-01T12:00:00Z",
		},
		{
			name:        "異常系: UserIDが空文字",
			userID:      "",
			content:     "Hello, World!",
			clock:       clock,
			wantErr:     domain.ErrUserIDRequired,
			wantUserID:  "",
			wantContent: TweetContent(""),
			wantTime:    "",
		},
		{
			name:        "異常系: Contentが空文字",
			userID:      "user123",
			content:     "",
			clock:       clock,
			wantErr:     domain.ErrContentRequired,
			wantUserID:  "",
			wantContent: TweetContent(""),
			wantTime:    "",
		},
		{
			name:        "異常系: UserIDとContentが両方空文字",
			userID:      "",
			content:     "",
			clock:       clock,
			wantErr:     domain.ErrUserIDRequired, // 最初のバリデーションエラーが返される
			wantUserID:  "",
			wantContent: TweetContent(""),
			wantTime:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTweet(CreateTweetRequest{UserID: tt.userID, Content: tt.content}, tt.clock)

			// エラーチェック
			if err != tt.wantErr {
				t.Errorf("NewTweet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 正常系の場合のみ、作成されたTweetの値をチェック
			if tt.wantErr == nil {
				if got.ID == TweetID("") {
					t.Error("NewTweet() ID should not be empty")
				}
				if got.UserID != tt.wantUserID {
					t.Errorf("NewTweet() UserID = %v, want %v", got.UserID, tt.wantUserID)
				}
				if got.Content != tt.wantContent {
					t.Errorf("NewTweet() Content = %v, want %v", got.Content, tt.wantContent)
				}
				if got.CreatedAt != tt.wantTime {
					t.Errorf("NewTweet() CreatedAt = %v, want %v", got.CreatedAt, tt.wantTime)
				}
			} else {
				// 異常系の場合、Tweetは空であるべき
				if got.ID != TweetID("") || got.UserID != "" || got.Content != TweetContent("") || got.CreatedAt != "" {
					t.Errorf("NewTweet() should return empty Tweet on error, got %+v", got)
				}
			}
		})
	}
}

func TestNewTweet_UniqueID(t *testing.T) {
	// IDがユニークに生成されることを確認
	fixedTime := time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)
	clock := FixedClock{FixedTime: fixedTime}

	userID := "user123"
	content := "Hello, World!"

	tweet1, err1 := NewTweet(CreateTweetRequest{UserID: userID, Content: content}, clock)
	if err1 != nil {
		t.Fatalf("NewTweet() unexpected error: %v", err1)
	}

	tweet2, err2 := NewTweet(CreateTweetRequest{UserID: userID, Content: content}, clock)
	if err2 != nil {
		t.Fatalf("NewTweet() unexpected error: %v", err2)
	}

	if tweet1.ID == tweet2.ID {
		t.Errorf("NewTweet() should generate unique IDs, got same ID: %v", tweet1.ID)
	}
}

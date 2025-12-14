package tweet

import (
	"errors"
	"testing"
	"time"

	"example.com/RestCRUD/domain"
	"example.com/RestCRUD/domain/tweet"
)

// MockClock はテスト用のモックClock
type MockClock struct {
	FixedTime time.Time
}

func (m MockClock) Now() time.Time {
	return m.FixedTime
}

// MockTweetRepository はテスト用のモックリポジトリ
type MockTweetRepository struct {
	CreateFunc func(tweet tweet.Tweet) (tweet.Tweet, error)
}

func (m *MockTweetRepository) Create(t tweet.Tweet) (tweet.Tweet, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(t)
	}
	return t, nil
}

func TestCreateTweetUseCase_Execute(t *testing.T) {
	fixedTime := time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)
	clock := MockClock{FixedTime: fixedTime}

	tests := []struct {
		name           string
		userID         string
		content        string
		repo           *MockTweetRepository
		clock          domain.Clock
		wantErr        error
		wantSavedTweet bool
		wantUserID     string
		wantContent    string
	}{
		{
			name:    "正常系: 有効なUserIDとContentでツイートを作成",
			userID:  "user123",
			content: "Hello, World!",
			repo: &MockTweetRepository{
				CreateFunc: func(tw tweet.Tweet) (tweet.Tweet, error) {
					// リポジトリが正しいエンティティを受け取ったことを確認
					if tw.UserID != "user123" {
						return tweet.Tweet{}, errors.New("repository received wrong UserID")
					}
					if tw.Content != "Hello, World!" {
						return tweet.Tweet{}, errors.New("repository received wrong Content")
					}
					return tw, nil
				},
			},
			clock:          clock,
			wantErr:        nil,
			wantSavedTweet: true,
			wantUserID:     "user123",
			wantContent:    "Hello, World!",
		},
		{
			name:    "異常系: UserIDが空文字の場合、ドメインエラーが返される",
			userID:  "",
			content: "Hello, World!",
			repo: &MockTweetRepository{
				CreateFunc: func(tw tweet.Tweet) (tweet.Tweet, error) {
					return tweet.Tweet{}, errors.New("Repository.Create should not be called when domain validation fails")
				},
			},
			clock:          clock,
			wantErr:        domain.ErrUserIDRequired,
			wantSavedTweet: false,
		},
		{
			name:    "異常系: Contentが空文字の場合、ドメインエラーが返される",
			userID:  "user123",
			content: "",
			repo: &MockTweetRepository{
				CreateFunc: func(tw tweet.Tweet) (tweet.Tweet, error) {
					return tweet.Tweet{}, errors.New("Repository.Create should not be called when domain validation fails")
				},
			},
			clock:          clock,
			wantErr:        domain.ErrContentRequired,
			wantSavedTweet: false,
		},
		{
			name:    "異常系: リポジトリがエラーを返した場合、エラーが伝播される",
			userID:  "user123",
			content: "Hello, World!",
			repo: &MockTweetRepository{
				CreateFunc: func(t tweet.Tweet) (tweet.Tweet, error) {
					return tweet.Tweet{}, errors.New("repository error")
				},
			},
			clock:          clock,
			wantErr:        errors.New("repository error"),
			wantSavedTweet: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := NewCreateTweetUseCase(tt.repo, tt.clock)
			got, err := usecase.Execute(tt.userID, tt.content)

			// エラーチェック
			if (err != nil) != (tt.wantErr != nil) {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 正常系の場合のみ、作成されたTweetの値をチェック
			if tt.wantErr == nil && tt.wantSavedTweet {
				if got == nil {
					t.Error("Execute() should return non-nil tweet on success")
					return
				}
				if got.UserID != tt.wantUserID {
					t.Errorf("Execute() UserID = %v, want %v", got.UserID, tt.wantUserID)
				}
				if got.Content != tt.wantContent {
					t.Errorf("Execute() Content = %v, want %v", got.Content, tt.wantContent)
				}
				if got.ID == "" {
					t.Error("Execute() ID should not be empty")
				}
				if got.CreatedAt == "" {
					t.Error("Execute() CreatedAt should not be empty")
				}
			} else {
				// 異常系の場合、nilが返されるべき
				if got != nil {
					t.Errorf("Execute() should return nil tweet on error, got %+v", got)
				}
			}
		})
	}
}

func TestCreateTweetUseCase_Execute_RepositoryCalledWithCorrectEntity(t *testing.T) {
	fixedTime := time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)
	clock := MockClock{FixedTime: fixedTime}

	var receivedTweet tweet.Tweet
	repo := &MockTweetRepository{
		CreateFunc: func(tw tweet.Tweet) (tweet.Tweet, error) {
			receivedTweet = tw
			return tw, nil
		},
	}

	usecase := NewCreateTweetUseCase(repo, clock)
	_, err := usecase.Execute("user123", "Test content")
	if err != nil {
		t.Fatalf("Execute() unexpected error: %v", err)
	}

	// リポジトリが正しいエンティティを受け取ったことを確認
	if receivedTweet.UserID != "user123" {
		t.Errorf("Repository received wrong UserID: got %v, want user123", receivedTweet.UserID)
	}
	if receivedTweet.Content != "Test content" {
		t.Errorf("Repository received wrong Content: got %v, want Test content", receivedTweet.Content)
	}
	if receivedTweet.ID == "" {
		t.Error("Repository received Tweet with empty ID")
	}
	if receivedTweet.CreatedAt != "2025-01-01T12:00:00Z" {
		t.Errorf("Repository received wrong CreatedAt: got %v, want 2025-01-01T12:00:00Z", receivedTweet.CreatedAt)
	}
}

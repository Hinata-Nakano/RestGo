package tweet

import (
	"example.com/RestCRUD/domain/tweet"
	"testing"
	"time"
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

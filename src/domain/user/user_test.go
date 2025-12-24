package user

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

func TestNewUser(t *testing.T) {
	// 固定時刻を設定
	fixedTime := time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)
	clock := FixedClock{FixedTime: fixedTime}

	tests := []struct {
		name      string
		nameValue string
		email     string
		password  string
		clock     domain.Clock
		wantErr   error
		wantName  UserName
		wantEmail UserEmail
		wantTime  string
	}{
		{
			name:      "正常系: 有効なNameとEmailでUserを作成",
			nameValue: "John Doe",
			email:     "john@example.com",
			password:  "password123",
			clock:     clock,
			wantErr:   nil,
			wantName:  UserName("John Doe"),
			wantEmail: UserEmail("john@example.com"),
			wantTime:  "2025-01-01T12:00:00Z",
		},
		{
			name:      "異常系: Nameが空文字",
			nameValue: "",
			email:     "john@example.com",
			password:  "password123",
			clock:     clock,
			wantErr:   domain.ErrNameRequired,
			wantName:  UserName(""),
			wantEmail: UserEmail(""),
			wantTime:  "",
		},
		{
			name:      "異常系: Emailが空文字",
			nameValue: "John Doe",
			email:     "",
			password:  "password123",
			clock:     clock,
			wantErr:   domain.ErrEmailRequired,
			wantName:  UserName(""),
			wantEmail: UserEmail(""),
			wantTime:  "",
		},
		{
			name:      "異常系: NameとEmailが両方空文字",
			nameValue: "",
			email:     "",
			password:  "password123",
			clock:     clock,
			wantErr:   domain.ErrNameRequired, // 最初のバリデーションエラーが返される
			wantName:  UserName(""),
			wantEmail: UserEmail(""),
			wantTime:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(CreateUserRequest{
				Name:     tt.nameValue,
				Email:    tt.email,
				Password: tt.password,
			}, tt.clock)

			// エラーチェック
			if err != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 正常系の場合のみ、作成されたUserの値をチェック
			if tt.wantErr == nil {
				if got == nil {
					t.Error("NewUser() should return non-nil user on success")
					return
				}
				if got.ID == UserID("") {
					t.Error("NewUser() ID should not be empty")
				}
				if got.Name != tt.wantName {
					t.Errorf("NewUser() Name = %v, want %v", got.Name, tt.wantName)
				}
				if got.Email != tt.wantEmail {
					t.Errorf("NewUser() Email = %v, want %v", got.Email, tt.wantEmail)
				}
				if got.Password != UserPassword(tt.password) {
					t.Errorf("NewUser() Password = %v, want %v", got.Password, tt.password)
				}
				if got.CreatedAt != tt.wantTime {
					t.Errorf("NewUser() CreatedAt = %v, want %v", got.CreatedAt, tt.wantTime)
				}
				if got.UpdatedAt != tt.wantTime {
					t.Errorf("NewUser() UpdatedAt = %v, want %v", got.UpdatedAt, tt.wantTime)
				}
			} else {
				// 異常系の場合、nilが返されるべき
				if got != nil {
					t.Errorf("NewUser() should return nil user on error, got %+v", got)
				}
			}
		})
	}
}

func TestNewUser_UniqueID(t *testing.T) {
	// IDがユニークに生成されることを確認
	fixedTime := time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)
	clock := FixedClock{FixedTime: fixedTime}

	name := "John Doe"
	email := "john@example.com"
	password := "password123"

	user1, err1 := NewUser(CreateUserRequest{
		Name:     name,
		Email:    email,
		Password: password,
	}, clock)
	if err1 != nil {
		t.Fatalf("NewUser() unexpected error: %v", err1)
	}

	user2, err2 := NewUser(CreateUserRequest{
		Name:     name,
		Email:    email,
		Password: password,
	}, clock)
	if err2 != nil {
		t.Fatalf("NewUser() unexpected error: %v", err2)
	}

	if user1.ID == user2.ID {
		t.Errorf("NewUser() should generate unique IDs, got same ID: %v", user1.ID)
	}
}

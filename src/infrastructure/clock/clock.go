package clock

import (
	"time"

	"example.com/RestCRUD/domain"
)

// RealClock は実際の現在時刻を返すClockの実装
type RealClock struct{}

// Now は現在時刻を返す
func (r RealClock) Now() time.Time {
	return time.Now()
}

// 型チェック: RealClockがdomain.Clockインターフェースを実装していることを確認
var _ domain.Clock = RealClock{}

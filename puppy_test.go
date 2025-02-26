// puppy_test.go 正確結構
package puppy // 必須在第一行宣告所屬包
// puppy_test.go
import (
	"strings"
	"testing"
)

func TestBidBark(t *testing.T) {
	t.Run("debug-mode", func(t *testing.T) {
		result := BidBark() // 在此设断点
		if !strings.Contains(result, "WOOF") {
			t.Fatal("Unexpected output")
		}
	})
}

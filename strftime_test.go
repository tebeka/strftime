package strftime

import (
	"testing"
	"time"
)

var testTime = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

func TestBasic(t *testing.T) {
	s := strftime("%a", testTime)
	if s != "Tue" {
		t.Fatalf("Bad day for %s, got %s - expected Tue", testTime, s)
	}
}

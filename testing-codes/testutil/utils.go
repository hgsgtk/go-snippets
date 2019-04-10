package testutil

import (
	"testing"
	"time"
)

func GetJstLocation(t *testing.T) *time.Location {
	t.Helper()

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		t.Fatalf("Failed to load JST time Location")
	}
	return jst
}

package date

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatDate(t *testing.T) {
	testCases := []struct {
		name string
		t    time.Time
		want string
	}{
		{"Christmas date", time.Date(2022, 12, 25, 0, 0, 0, 0, time.UTC), "2022-12-25"},
		{"Datetime with timezone", time.Date(2025, 1, 2, 3, 4, 5, 6, time.FixedZone("UTC-3", -3)), "2025-01-02"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := FormatDate(tc.t)
			assert.Equal(t, tc.want, got)
		})
	}
}

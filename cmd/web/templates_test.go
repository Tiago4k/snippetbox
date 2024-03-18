package main

import (
	"testing"
	"time"

	"github.com/Tiago4k/snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {

	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2024, 3, 18, 15, 30, 0, 0, time.UTC),
			want: "18 Mar 2024 at 15:30",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2024, 3, 18, 15, 30, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "18 Mar 2024 at 14:30",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			assert.Equal(t, hd, tt.want)
		})
	}
}

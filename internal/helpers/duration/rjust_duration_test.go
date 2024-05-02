package duration_helpers

import (
	"testing"
	"time"
)

func TestRjustDuration(t *testing.T) {
	type args struct {
		duration time.Duration
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1 hour 1 minute 1 second",
			args: args{duration: time.Hour + time.Minute + time.Second},
			want: "01:01:01",
		},
		{
			name: "666 hours 17 minutes 59 seconds",
			args: args{duration: time.Hour*666 + time.Minute*17 + time.Second*59},
			want: "666:17:59",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RjustDuration(tt.args.duration); got != tt.want {
				t.Errorf("RjustDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

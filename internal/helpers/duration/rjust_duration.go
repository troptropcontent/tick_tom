package duration_helpers

import (
	"fmt"
	"math"
	"time"
)

func RjustDuration(duration time.Duration) string {
	hours := math.Floor(duration.Hours())
	minutes := math.Floor(duration.Minutes() - (hours * 60))
	seconds := math.Floor(duration.Seconds() - (minutes * 60) - (hours * 3600))
	return fmt.Sprintf("%02d:%02d:%02d", int(hours), int(minutes), int(seconds))
}

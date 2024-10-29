package helper

import (
	"fmt"
	"time"
)

func TimeDiffCalculator(initTime time.Time) string {
	currentTime := time.Now()
	timeDiff := currentTime.Sub(initTime)

	switch {
	case timeDiff.Seconds() < 60:
		return "A few seconds ago"
	case timeDiff.Minutes() < 60:
		return fmt.Sprintf("%v minutes ago", int(timeDiff.Minutes()))
	case timeDiff.Hours() < 24:
		return fmt.Sprintf("%v hours ago", int(timeDiff.Hours()))
	default:
		return fmt.Sprintf("%v days ago", int(timeDiff.Hours()/24))
	}
}

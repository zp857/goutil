package progress

import (
	"fmt"
	"github.com/zp857/goutil/mathutil"
	"time"
)

func Calculate(total, finished int, start time.Time) (doneString string, doneFloat, remainingFloat float64) {
	done := float64(finished) / float64(total)
	doneString = fmt.Sprintf("%.2f", done)
	doneFloat = done
	if done == 1 {
		remainingFloat = 0
		doneString = "1.0"
	} else {
		elapsed := time.Since(start).Seconds()
		remainingFloat = elapsed / done
		remainingFloat = mathutil.Decimal(remainingFloat)
	}
	return
}

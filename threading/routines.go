package threading

import (
	"github.com/zp857/goutil/rescue"
	"go.uber.org/zap"
	"runtime"
	"time"
)

func GoSafe(fn func()) {
	go RunSafe(fn)
}

func RunSafe(fn func()) {
	defer rescue.Recover()
	fn()
}

func WatchGoroutines(ticker int) {
	tk := time.NewTicker(time.Duration(ticker) * time.Second)
	for {
		zap.L().Sugar().Debug("Current goroutine number: %v", runtime.NumGoroutine())
		<-tk.C
	}
}

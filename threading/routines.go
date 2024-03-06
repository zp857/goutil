package threading

import (
	"github.com/zp857/goutil/rescue"
)

func GoSafe(fn func()) {
	go RunSafe(fn)
}

func RunSafe(fn func()) {
	defer rescue.Recover()
	fn()
}

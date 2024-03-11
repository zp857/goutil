package headless

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/launcher/flags"
	"time"
)

type RodOptions struct {
	Headless bool   `json:"headless"`
	Proxy    string `json:"proxy"`
	Timeout  int    `json:"timeout"`
	Trace    bool   `json:"trace"`
}

func NewRod(options *RodOptions) (browser *rod.Browser) {
	l := launcher.New().
		Set("ignore-certificate-errors", "true").
		Set("ignore-certificate-errors", "1").
		Set("mute-audio", "true").
		Set("incognito", "true").
		Set("disable-blink-features", "AutomationControlled").
		Leakless(true).
		Headless(options.Headless).
		Devtools(false)
	if options.Proxy != "" {
		l.Set(flags.ProxyServer, options.Proxy)
	}
	browser = rod.New().ControlURL(l.MustLaunch())
	if options.Timeout > 0 {
		browser = browser.Timeout(time.Duration(options.Timeout) * time.Second)
	}
	if options.Trace {
		browser = browser.Trace(options.Trace).SlowMotion(2 * time.Second)
	}
	return
}

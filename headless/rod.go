package headless

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/launcher/flags"
	"time"
)

type RodOptions struct {
	Headless   bool   `json:"headless"`
	Proxy      string `json:"proxy"`
	MaxRuntime int    `json:"maxRuntime"`
	Timeout    int    `json:"timeout"`
	Trace      bool   `json:"trace"`
}

func NewRod(options *RodOptions) (browser *rod.Browser) {
	l := launcher.New().
		Headless(options.Headless).
		Set("ignore-certificate-errors", "true").
		Set("no-sandbox").
		Set("disable-gpu").
		Set("no-first-run").
		Set("no-default-browser-check").
		Set("enable-automation", "false"). // 防止监测 webdriver
		Set("disable-blink-features", "AutomationControlled"). // 禁用 blink 特征，绕过了加速乐检测
		Leakless(true).
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
	browser.NoDefaultDevice()
	return
}

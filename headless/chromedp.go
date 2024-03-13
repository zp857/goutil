package headless

import (
	"context"
	"github.com/chromedp/chromedp"
	"time"
)

type ChromedpOptions struct {
	Headless   bool   `json:"headless"`
	Proxy      string `json:"proxy"`
	ChromePath string `json:"chromePath"`
	Timeout    int    `json:"timeout"`
}

func NewChromedp(options *ChromedpOptions) (ctx context.Context, cancel context.CancelFunc) {
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", options.Headless),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.NoSandbox,
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		chromedp.NoFirstRun,
		chromedp.Flag("enable-automation", false),                       // 防止监测 webdriver
		chromedp.Flag("disable-blink-features", "AutomationControlled"), // 禁用 blink 特征，绕过了加速乐检测
	)
	if options.Proxy != "" {
		opts = append(opts, chromedp.ProxyServer(options.Proxy))
	}
	if options.ChromePath != "" {
		opts = append(opts, chromedp.ExecPath(options.ChromePath))
	}
	ctx, _ = chromedp.NewExecAllocator(
		context.Background(),
		opts...,
	)
	ctx, cancel = chromedp.NewContext(
		ctx,
	)
	ctx, cancel = context.WithTimeout(ctx, time.Duration(options.Timeout)*time.Second)
	return
}

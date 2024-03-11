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
		// 以默认配置的数组为基础，覆写 headless 参数
		// 当然也可以根据自己的需要进行修改，这个 flag 是浏览器的设置
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoSandbox,
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		chromedp.NoFirstRun,
		chromedp.Flag("headless", options.Headless), // 显示界面
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("ignore-certificate-errors", true),
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

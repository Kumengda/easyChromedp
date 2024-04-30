package chrome

import (
	"context"
	. "github.com/Kumengda/easyChromedp/runtime"
	"github.com/Kumengda/easyChromedp/utils"
	"github.com/chromedp/chromedp"
	"os"
	"path/filepath"
	"time"
)

type Chrome struct {
	ctx         context.Context
	cancels     []context.CancelFunc
	debugModule bool
	tmpPath     string
	timeout     int
}

func (c *Chrome) EnableDebug() {
	c.debugModule = true
}
func (c *Chrome) GetTimeout() int {
	return c.timeout
}
func NewChromeWithTimout(timeout int, option ...chromedp.ExecAllocatorOption) (*Chrome, error) {
	var cancels []context.CancelFunc
	absPwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	tmpPath := absPwd + string(filepath.Separator) + utils.GenerateRandomString(8)
	option = append(option,
		chromedp.UserDataDir(tmpPath),
		chromedp.Flag("disk-cache-dir", tmpPath),
	)
	ctx, cancel := chromedp.NewContext(context.Background())
	cancels = append(cancels, cancel)
	ctx, _ = context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		option...,
	)
	allocCtx, cancel2 := chromedp.NewExecAllocator(ctx, opts...)
	cancels = append(cancels, cancel2)
	ctx, _ = chromedp.NewContext(allocCtx)
	chromedp.Run(ctx)
	return &Chrome{ctx: ctx, cancels: cancels, tmpPath: tmpPath, timeout: timeout}, nil
}
func (c *Chrome) RunWithListen(ctx context.Context, listenFun func(ev interface{}), action ...chromedp.Action) error {
	chromedp.ListenTarget(ctx, listenFun)
	err := chromedp.Run(ctx,
		action...,
	)
	if err != nil {
		return err
	}

	return nil
}
func (c *Chrome) GetContext() context.Context {
	return c.ctx
}
func (c *Chrome) RunWithOutListen(ctx context.Context, action ...chromedp.Action) error {
	err := chromedp.Run(ctx,
		action...,
	)
	if err != nil {
		return err
	}
	return nil
}
func (c *Chrome) Close() {
	for _, v := range c.cancels {
		v()
	}
	for {
		time.Sleep(1 * time.Second)
		if c.debugModule {
			MainInsp.Print(LEVEL_DEBUG, Text("尝试删除文件"+c.tmpPath))
		}
		err := os.RemoveAll(c.tmpPath)
		if err != nil {
			if c.debugModule {
				MainInsp.Print(LEVEL_DEBUG, Text("删除文件失败,即将重试"+c.tmpPath+err.Error()))
			}
		} else {
			if c.debugModule {
				MainInsp.Print(LEVEL_DEBUG, Text("删除文件成功"+c.tmpPath))
			}
			break
		}
	}

}
func init() {
	InitLogger()
}

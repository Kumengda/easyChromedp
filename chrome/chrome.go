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
	ctx     context.Context
	cancels []context.CancelFunc
	tmpPath string
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
	return &Chrome{ctx: ctx, cancels: cancels, tmpPath: tmpPath}, nil
}
func (c *Chrome) RunWithListen(listenFun func(ev interface{}), action ...chromedp.Action) error {
	chromedp.ListenTarget(c.ctx, listenFun)
	err := chromedp.Run(c.ctx,
		action...,
	)
	if err != nil {
		return err
	}

	return nil
}
func (c *Chrome) RunWithOutListen(action ...chromedp.Action) error {
	err := chromedp.Run(c.ctx,
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
	os.RemoveAll(c.tmpPath)

}
func init() {
	InitLogger()
}

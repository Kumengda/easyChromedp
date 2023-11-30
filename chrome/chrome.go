package chrome

import (
	"context"
	. "github.com/Kumengda/easyChromedp/runtime"
	"github.com/chromedp/chromedp"
	"time"
)

type Chrome struct {
	ctx     context.Context
	cancels []context.CancelFunc
}

func NewChromeWithTimout(timeout int, option ...chromedp.ExecAllocatorOption) (*Chrome, error) {
	var cancels []context.CancelFunc
	ctx, cancel := chromedp.NewContext(context.Background())
	cancels = append(cancels, cancel)
	ctx, _ = context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		option...,
	)
	allocCtx, cancel2 := chromedp.NewExecAllocator(ctx, opts...)
	cancels = append(cancels, cancel2)
	ctx, _ = chromedp.NewContext(allocCtx)
	return &Chrome{ctx: ctx, cancels: cancels}, nil
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
}
func init() {
	InitLogger()
}

package helpers

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/metrics"
	"time"
)

// IntrumentingMiddleware is a struct representing middleware
type IntrumentingMiddleware struct {
	RequestCount metrics.Counter
	RequestLatency metrics.Histogram
	Next EncryptService
}

func (mw IntrumentingMiddleware) Encrypt(ctx context.Context, key, text string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "encrypt", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	output, err = mw.Next.Encrypt(ctx, key, text)
	return
}

func (mw IntrumentingMiddleware) Decrypt(ctx context.Context, key, text string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "decrypt", "error", "false"}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())
	output, err = mw.Next.Decrypt(ctx, key, text)
	return
}

package helpers

import (
	"context"
	log "github.com/go-kit/kit/log"
	"time"
)

// LoggingMiddleware wraps the logs for incoming request
type LoggingMiddleware struct {
	Logger log.Logger
	Next EncryptService
}

// Encrypt logs the Encryption requests
func (mw LoggingMiddleware) Encrypt(ctx context.Context, key, text string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "encrypt",
			"key", key,
			"text", text,
			"output", output,
			"error", err,
			"took", time.Since(begin),
			)
	}(time.Now())
	output, err = mw.Next.Encrypt(ctx, key, text)
	return
}

// Decrypt logs the Decryption requests
func (mw LoggingMiddleware) Decrypt(ctx context.Context, key, text string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "decrypt",
			"key", key,
			"text", text,
			"output", output,
			"error", err,
			"took", time.Since(begin),
			)
	}(time.Now())
	output, err = mw.Next.Decrypt(ctx,key, text)
	return
}
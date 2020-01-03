package log

import (
	"cloud.google.com/go/logging"
	"context"
	"net/http"

	"GoStuff/gcp/ae/log/internal"
)

func Flush(ctx context.Context) {
	internal.FlushLog(ctx)
}

func NewContext(r *http.Request) context.Context {
	return internal.WithLogContentValue(r)
}

// Debugf formats its arguments according to the format
func Debugf(ctx context.Context, format string, args ...interface{}) {
	internal.Logf(ctx, logging.Debug, format, args...)
}

// Infof is like Debugf, but at Info level.
func Infof(ctx context.Context, format string, args ...interface{}) {
	internal.Logf(ctx, logging.Info, format, args...)
}

// Warningf is like Debugf, but at Warning level.
func Warningf(ctx context.Context, format string, args ...interface{}) {
	internal.Logf(ctx, logging.Warning, format, args...)
}

// Errorf is like Debugf, but at Error level.
func Errorf(ctx context.Context, format string, args ...interface{}) {
	internal.Logf(ctx, logging.Error, format, args...)
}

// Criticalf is like Debugf, but at Critical level.
func Criticalf(ctx context.Context, format string, args ...interface{}) {
	internal.Logf(ctx, logging.Critical, format, args...)
}

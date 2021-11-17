package logging

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
)

const loggerCtx = "loggerCtx"

type zapLogger struct {
	l *zap.SugaredLogger
}

func NewLogger() Logger {
	zapLog, err := zap.NewProduction()
	errCheck(err)
	log := NewZapLogger(zapLog.Sugar())

	return log
}

func NewZapLogger(z *zap.SugaredLogger) *zapLogger {
	return &zapLogger{l: z}
}

func (l *zapLogger) Info(args ...interface{}) {
	l.l.Info(args...)
}

func (l *zapLogger) Error(args ...interface{}) {
	l.l.Error(args...)
}

func (l *zapLogger) Warn(args ...interface{}) {
	l.l.Warn(args...)
}

func (l *zapLogger) Fatal(args ...interface{}) {
	l.l.Fatal(args...)
}

func (l *zapLogger) With(args ...interface{}) Logger {
	return NewZapLogger(l.l.With(args...))
}

func (l *zapLogger) Sync() error {
	return l.l.Sync()
}

func InjectIntoHandler(log Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, loggerCtx, log.With("requestId", uuid.New().String())) // TODO: request id managing, add to traces

			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}

func InjectIntoContext(ctx context.Context, log Logger) context.Context {
	return context.WithValue(ctx, loggerCtx, log)
}

func GetFromContext(ctx context.Context) (Logger, bool) {
	log, ok := ctx.Value(loggerCtx).(Logger)
	if ok {
		return log, true
	}
	return nil, false
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

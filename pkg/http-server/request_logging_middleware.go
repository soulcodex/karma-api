package httpserver

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/soulcodex/karma-api/pkg/logger"
)

type RequestLoggingMiddleware struct {
	logger logger.Logger
}

func NewRequestLoggingMiddleware(logger logger.Logger) *RequestLoggingMiddleware {
	return &RequestLoggingMiddleware{logger: logger}
}

func (rlm *RequestLoggingMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		recorder, ipAddress, startedAt := NewRequestStatusRecorder(w), ClientIp(req), time.Now()

		defer func() {
			totalTime := time.Since(startedAt).Milliseconds()
			rlm.logger.Info(
				req.Context(),
				fmt.Sprintf("%s %s %d %s %s %d %s", req.Method, req.RequestURI, recorder.Status, time.Now().Format(time.RFC3339), ipAddress, totalTime, req.Referer()),
				slog.String("remote_addr_ip", ipAddress),
				slog.Int64("request_time_d", totalTime),
				slog.Int("status", recorder.Status),
				slog.String("request", req.RequestURI),
				slog.String("request_method", req.Method),
				slog.String("http_referrer", req.Referer()),
				slog.String("http_user_agent", req.Header.Get("User-Agent")),
				slog.String("response_content_type", w.Header().Get("Content-Type")),
				slog.String("correlation_id", w.Header().Get(HeaderRequestIdentifier)),
				slog.String("message", fmt.Sprintf("%s %s %d %s %s %d %s", req.Method, req.RequestURI, recorder.Status, time.Now().Format(time.RFC3339), ipAddress, totalTime, req.Referer())),
			)
		}()

		next.ServeHTTP(recorder, req)
	})
}

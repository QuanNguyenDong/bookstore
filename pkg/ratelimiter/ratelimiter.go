package ratelimiter

import (
	"net/http"
	"time"
)

type Limiter interface {
	AllowRequest(ip string) (bool, time.Duration)
	Middleware(http.Handler) http.Handler
}

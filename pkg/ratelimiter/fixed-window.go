package ratelimiter

import (
	"net/http"
	"sync"
	"time"
)

type FixedWindowRateLimiter struct {
	sync.RWMutex
	clients map[string]int
	limit int
	window time.Duration
}

func NewFixedWindowLimiter(limit int, window time.Duration) *FixedWindowRateLimiter {
	return &FixedWindowRateLimiter{
		clients: map[string]int{},
		limit: limit,
		window: window,
	}
}

func (limiter *FixedWindowRateLimiter) AllowRequest(ip string) (bool, time.Duration) {
	limiter.RLock()
	count, exists := limiter.clients[ip]
	limiter.RUnlock()

	if !exists || count < limiter.limit {
		limiter.Lock()
		if !exists {
			go limiter.resetCount(ip)
		}

		limiter.clients[ip] += 1
		limiter.Unlock()
		return true, 0
	}


	return false, limiter.window
}

func (limiter *FixedWindowRateLimiter) resetCount(ip string) {
	time.Sleep(limiter.window)
	limiter.Lock()
	delete(limiter.clients, ip)
	limiter.Unlock()
}

func (limiter *FixedWindowRateLimiter) Middleware(next http.Handler) http.Handler {
	
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if allow, retryAfter := limiter.AllowRequest(request.RemoteAddr); !allow {
			writer.Header().Set("Retry-After", retryAfter.String())
			writer.WriteHeader(http.StatusTooManyRequests)
			writer.Write([]byte("To many request"))
			return
		}

		next.ServeHTTP(writer, request)
	})
}

package rate

/*
Based on https://github.com/julianshen/gin-limiter
Copyright 2016 Julian Shen
Apache License 2.0
*/
import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

type KeyFunc func(ctx *gin.Context) (string, error)

type LimiterMiddleware struct {
	fillInterval time.Duration
	capacity     int64
	rateKeyGen   KeyFunc
	limiters     map[string]*ratelimit.Bucket
}

func (r *LimiterMiddleware) get(ctx *gin.Context) (*ratelimit.Bucket, error) {
	key, err := r.rateKeyGen(ctx)

	if err != nil {
		return nil, err
	}

	if limiter, existed := r.limiters[key]; existed {
		return limiter, nil
	}

	limiter := ratelimit.NewBucketWithQuantum(r.fillInterval, r.capacity, r.capacity)
	r.limiters[key] = limiter
	return limiter, nil
}

func (r *LimiterMiddleware) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limiter, err := r.get(ctx)
		if err != nil || limiter.TakeAvailable(1) == 0 {
			if err == nil {
				err = errors.New("too many requests")
			}
			ctx.AbortWithError(429, err)
		} else {
			ctx.Next()
		}
	}
}

func NewRateLimiter(interval time.Duration, capacity int64, keyGen KeyFunc) *LimiterMiddleware {
	limiters := make(map[string]*ratelimit.Bucket)
	return &LimiterMiddleware{
		interval,
		capacity,
		keyGen,
		limiters,
	}
}

package limiter

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// 限流器
type MethodLimiter struct {
	*Limiter
}

func NewMethodLimiter() LimiterIface {
	return MethodLimiter{
		Limiter: &Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)},
	}
}

// 获取对应限流器键值对名称, 路由地址
/**
 * @description: 实现LimiterIface Key 方法
 * @param {*gin.Context} c
 * @return {*}
 */
func (l MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	// Index返回s中substr的第一个实例的索引，如果s中没有substr，则返回-1
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index]
}

// 获取令牌桶
/**
 * @description: 实现LimiterIface GetBucket 方法
 * @param {string} key
 * @return {*}
 */
func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.limiterBuckets[key]
	return bucket, ok
}

// 新增令牌桶
/**
 * @description: 实现LimiterIface AddBuckets 方法
 * @param {...LimiterBucketRule} rules
 * @return {*}
 */
func (l MethodLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterIface {
	for _, rule := range rules {
		if _, ok := l.limiterBuckets[rule.Key]; !ok {
			// 创建令牌桶，
			l.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quantum)
		}
	}
	return l
}

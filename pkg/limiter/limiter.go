/*
 * @Author: GG
 * @Date: 2023-02-08 16:03:07
 * @LastEditTime: 2023-02-08 16:24:32
 * @LastEditors: GG
 * @Description: 限流控制
 * @FilePath: \oms\pkg\limiter\limiter.go
 *
 */
package limiter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

// 定义限流器接口
type LimiterIface interface {
	Key(c *gin.Context) string                          // 获取对应限流器键值对名称
	GetBucket(key string) (*ratelimit.Bucket, bool)     // 获取令牌桶
	AddBuckets(rules ...LimiterBucketRule) LimiterIface // 新增多个令牌桶
}

// 存储令牌桶与键值对名称的映射关系
type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

// 令牌桶规则
type LimiterBucketRule struct {
	Key          string        // 键值对名称
	FillInterval time.Duration // 间隔多久放N个令牌
	Capacity     int64         // 令牌桶容量
	Quantum      int64         // 每次到达间隔时间后所放的具体令牌数量
}

/*
 * @Author: GG
 * @Date: 2023-03-14 15:41:56
 * @LastEditTime: 2023-03-14 16:41:21
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\pkg\cache\cache.go
 *
 */
package cache

import "oms/pkg/setting"

type CacheInterface interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
	Delete(key string) error
	Clear() error
	GetType() string
}

type CacheStore struct {
	Engine CacheInterface
}

func NewCacheStore(CacheSetting *setting.CacheSettingS) (*CacheStore, error) {

	var cacheStore *CacheStore
	switch CacheSetting.CacheStore {
	case "bigCache", "bigcache":
		cacheStore = &CacheStore{
			Engine: NewBigCache(),
		}
		return cacheStore, nil
	default:
		return nil, nil
	}
	return nil, nil
}

func GetStore() {}

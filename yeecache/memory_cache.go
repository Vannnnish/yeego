/**
 * Created by angelina-zf on 17/2/28.
 */

// yeecache
// 内存缓存类
package yeecache

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	DefaultExpiration time.Duration = 0
	NoExpiration      time.Duration = -1
)

// MemoryCache 内存缓存
type MemoryCache struct {
	*cache
}

// yeecache
type cache struct {
	// 默认的过期时间，如果设置为0，则永不过期
	defaultExpiration time.Duration
	// 缓存的数据
	items map[string]Item
	// 读写锁
	mu sync.RWMutex
	// 在即将消失的时候执行的函数，在替换的时候不执行
	onDisappeared func(string, interface{})
	// 新开一个go-routine去监控那些过期的数据
	monitor *monitor
}

// newCache
// 初始化
func newCache(defaultExpiration time.Duration, m map[string]Item) *cache {
	if defaultExpiration == 0 {
		defaultExpiration = -1
	}
	c := &cache{
		defaultExpiration: defaultExpiration,
		items:             m,
	}
	return c
}

// Set
// 存储数据，如果存在则替换
func (c *cache) Set(k string, v interface{}, d time.Duration) {
	var e int64
	if d == DefaultExpiration {
		d = c.defaultExpiration
	}
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.mu.Lock()
	c.items[k] = Item{
		Object:     v,
		Expiration: e,
	}
	c.mu.Unlock()
}

// SetDefault
// 默认过期时间
func (c *cache) SetDefault(k string, v interface{}) {
	c.Set(k, v, c.defaultExpiration)
}

func (c *cache) set(k string, v interface{}, d time.Duration) {
	var e int64
	if d == DefaultExpiration {
		d = c.defaultExpiration
	}
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.items[k] = Item{
		Object:     v,
		Expiration: e,
	}
}

// Add
// 添加数据，如果已经存在，则报错
func (c *cache) Add(k string, v interface{}, d time.Duration) error {
	c.mu.Lock()
	_, found := c.get(k)
	if found {
		c.mu.Unlock()
		return fmt.Errorf("Item %s already exists", k)
	}
	c.set(k, v, d)
	c.mu.Unlock()
	return nil
}

// Replace
// 替换
func (c *cache) Replace(k string, v interface{}, d time.Duration) error {
	c.mu.Lock()
	_, found := c.get(k)
	if !found {
		c.mu.Unlock()
		return fmt.Errorf("Item %s doesn't exist", k)
	}
	c.set(k, v, d)
	c.mu.Unlock()
	return nil
}

// Get
// 获取数据，如果不存在，返回nil,false
func (c *cache) Get(k string) (interface{}, bool) {
	c.mu.RLock()
	item, found := c.items[k]
	if !found {
		c.mu.RUnlock()
		return nil, false
	}
	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			c.mu.RUnlock()
			return nil, false
		}
	}
	c.mu.RUnlock()
	return item.Object, true
}

func (c *cache) get(k string) (interface{}, bool) {
	item, found := c.items[k]
	if !found {
		return nil, false
	}
	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			return nil, false
		}
	}
	return item.Object, true
}

// Delete
// 删除某个key
func (c *cache) Delete(k string) {
	c.mu.Lock()
	v, found := c.delete(k)
	c.mu.Unlock()
	if found {
		c.onDisappeared(k, v)
	}
}

func (c *cache) delete(k string) (interface{}, bool) {
	if c.onDisappeared != nil {
		if v, found := c.items[k]; found {
			delete(c.items, k)
			return v.Object, true
		}
	}
	delete(c.items, k)
	return nil, false
}

// Items
// 获取全部未过期的数据
func (c *cache) Items() map[string]Item {
	c.mu.Lock()
	defer c.mu.Unlock()
	m := make(map[string]Item, len(c.items))
	now := time.Now().UnixNano()
	for k, v := range c.items {
		if v.Expiration > 0 {
			if now > v.Expiration {
				continue
			}
		}
		m[k] = v
	}
	return m
}

// ItemCount
// 获取已缓存的数量
// 其中包含了那些已经过期但是没有被清理的
func (c *cache) ItemCount() int {
	c.mu.RLock()
	n := len(c.items)
	c.mu.Unlock()
	return n
}

// Flush
// 清空全部的缓存
func (c *cache) Flush() {
	c.mu.Lock()
	c.items = map[string]Item{}
	c.mu.Unlock()
}

type KV struct {
	key   string
	value interface{}
}

// DeleteExpired
// 删除过期的数据
func (c *cache) DeleteExpired() {
	var items []KV
	now := time.Now().UnixNano()
	c.mu.Lock()
	for k, v := range c.items {
		if v.Expiration > 0 && now > v.Expiration {
			data, found := c.delete(k)
			if found {
				items = append(items, KV{key: k, value: data})
			}
		}
	}
	c.mu.Unlock()
	for _, v := range items {
		c.onDisappeared(v.key, v.value)
	}
}

// Item
// 缓存的数据的结构
type Item struct {
	Object     interface{}
	Expiration int64
}

// Expired
// 判断数据是否过期
func (item *Item) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > item.Expiration
}

// monitor
// 循环监控过期数据
type monitor struct {
	Interval time.Duration
	stop     chan bool
}

// run
// 不断循环
func (m *monitor) run(c *cache) {
	m.stop = make(chan bool)
	ticker := time.NewTicker(m.Interval)
	for {
		select {
		case <-ticker.C:
			c.DeleteExpired()
		case <-m.stop:
			ticker.Stop()
			return
		}
	}
}

// stopMonitor
// 停止
func stopMonitor(c *MemoryCache) {
	c.monitor.stop <- true
}

// runMonitor
// 开启监控
func runMonitor(c *cache, ci time.Duration) {
	m := &monitor{
		Interval: ci,
	}
	c.monitor = m
	go m.run(c)
}

// newMemoryCacheWithMonitor
// 初始化缓存，同时开启监控routine
func newMemoryCacheWithMonitor(de time.Duration, ci time.Duration, m map[string]Item) *MemoryCache {
	c := newCache(de, m)
	MC := &MemoryCache{c}
	if ci > 0 {
		runMonitor(c, ci)
		runtime.SetFinalizer(MC, stopMonitor)
	}
	return MC
}

// NewMemoryCache
// 初始化内存缓存
func NewMemoryCache(defaultExpiration, cleanupInterval time.Duration) *MemoryCache {
	items := make(map[string]Item)
	return newMemoryCacheWithMonitor(defaultExpiration, cleanupInterval, items)
}

// NewMemoryCache
// 初始化内存缓存
// 可以传入map数据
func NewMemoryCacheWithMap(defaultExpiration, cleanupInterval time.Duration, items map[string]Item) *MemoryCache {
	return newMemoryCacheWithMonitor(defaultExpiration, cleanupInterval, items)
}

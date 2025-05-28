package options_memory_cache

import (
	"context"
	"sync"
	"time"

	exchanges_repo "Arbitrax/pkg/repositories/exchanges"
	strategy_repo "Arbitrax/pkg/repositories/strategies"
)

type Cache struct {
	exchanges  []*exchanges_repo.Model
	strategies []*strategy_repo.Model

	expiresAt time.Time
	ttl       time.Duration
	mutex     sync.RWMutex
}

func New(ttl time.Duration) *Cache {
	return &Cache{
		ttl: ttl,
	}
}

func (c *Cache) GetExchanges(getter func(ctx context.Context) ([]*exchanges_repo.Model, error), ctx context.Context) ([]*exchanges_repo.Model, error) {
	c.mutex.RLock()
	if time.Now().Before(c.expiresAt) && c.exchanges != nil {
		defer c.mutex.RUnlock()
		return c.exchanges, nil
	}
	c.mutex.RUnlock()

	// Upgrade to write lock
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// Double check in case another goroutine already refreshed
	if time.Now().Before(c.expiresAt) && c.exchanges != nil {
		return c.exchanges, nil
	}

	data, err := getter(ctx)
	if err != nil {
		return nil, err
	}
	c.exchanges = data
	c.expiresAt = time.Now().Add(c.ttl)
	return data, nil
}

func (c *Cache) GetStrategies(getter func(ctx context.Context) ([]*strategy_repo.Model, error), ctx context.Context) ([]*strategy_repo.Model, error) {
	c.mutex.RLock()
	if time.Now().Before(c.expiresAt) && c.strategies != nil {
		defer c.mutex.RUnlock()
		return c.strategies, nil
	}
	c.mutex.RUnlock()

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if time.Now().Before(c.expiresAt) && c.strategies != nil {
		return c.strategies, nil
	}

	data, err := getter(ctx)
	if err != nil {
		return nil, err
	}
	c.strategies = data
	c.expiresAt = time.Now().Add(c.ttl)
	return data, nil
}

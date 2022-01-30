package main

import (
	"sync"
	"time"
)

// Cache offers a means to cache the result of an expensive operation with a timed expiry.
type Cache struct {
	fetcher Fetcher
	in      FetcherInput
	out     FetcherOutput

	mu            sync.Mutex
	cond          *sync.Cond
	expiry        time.Time
	fetchEnqueued bool
}

// FetcherInput is the input to the fetch() function of the Fetcher interface supplied and used to compute the expensive operation and cache its result.
type FetcherInput interface {
}

// FetcherOutput is the output of the expensive operation from the fetch() function of the Fetcher interface supplied to the cache.
type FetcherOutput interface {
}

// Fetcher interface used by the cache to fetch the result of the expensive operation.
type Fetcher interface {
	// Fetch is the expensive operation that will be cached.
	Fetch(FetcherInput) (FetcherOutput, time.Time)
}

// NewCache returns a Cache.
func NewCache(f Fetcher, in FetcherInput) *Cache {
	c := Cache{
		fetcher: f,
		in:      in,
		out:     nil,
	}
	c.cond = sync.NewCond(&c.mu)
	return &c
}

// update updates the value in the cache along with the specified expiry timestamp.
func (c *Cache) update(newVal FetcherOutput, expiry time.Time) {
	c.mu.Lock()
	c.fetchEnqueued = false
	c.out = newVal
	c.expiry = expiry
	c.cond.Broadcast()
	c.mu.Unlock()
}

// getCachedLocked returns whether an unexpired valid cache value is available
// along with the value.
// NOTE: This function should only be called with the mutex c.mu locked.
func (c *Cache) getCachedLocked() (bool, FetcherOutput) {
	// Check if a cached value is available and has not expired.
	if time.Now().After(c.expiry) {
		return false, nil
	}
	return true, c.out
}

// enqueueFetchLocked queues up a new request to invoke the Fetch() function of the
// fetcher if one has not already been enqueued.
// NOTE: This function should only be called with the mutex c.mu locked.
func (c *Cache) enqueueFetchLocked() {
	if c.fetchEnqueued {
		// A fetch has already been enqueued do not initiate a new one.
		return
	}

	// Clear existing cached state.
	c.out = nil

	// Launch goroutine to fetch the latest output.
	go c.fetch()
	c.fetchEnqueued = true
}

// fetch invokes the Fetch() function of the fetcher.
func (c *Cache) fetch() {
	c.update(c.fetcher.Fetch(c.in))
}

// Get returns the cached value if it has not expired, otherwise invokes the
// Fetch() function of the Fetcher interface that was supplied while the
// Cache was constructed to compute the result of the expensive operation.
// This function is thread-safe and guarantees that at most only one Fetch()
// function call will be invoked as a side-effect (from a separate go-routine)
// and only if the cached value has expired.
func (c *Cache) Get() FetcherOutput {
	logger := buildLogger()
	defer logger.Sync() // nolint - flushes buffer, if any
	log := logger.Sugar()

	cacheHitOnFirstTry := true

	var valid bool
	var out FetcherOutput
	c.mu.Lock()
	for {
		valid, out = c.getCachedLocked()
		if valid {
			if cacheHitOnFirstTry {
				log.Debugf("Found cached status on first try, using it.")
			} else {
				log.Debugf("Found cached status after waiting.")
			}
			break
		}
		cacheHitOnFirstTry = false
		c.enqueueFetchLocked()
		log.Debugf("No valid status found in the cache, waiting for a cache hit.")
		c.cond.Wait()
	}
	c.mu.Unlock()
	return out
}

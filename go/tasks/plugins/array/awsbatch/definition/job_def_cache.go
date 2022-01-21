/*
 * Copyright (c) 2018 Lyft. All rights reserved.
 */

package definition

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/batch"
	"github.com/coocood/freecache"
	"github.com/shomali11/util/xhashes"
)

type JobDefinitionArn = string

// An JobDefinition cache interface
type Cache interface {
	// Gets a JobDefinition if one is in memory. Otherwise found is set to false.
	Get(key CacheKey) (jobDefinition JobDefinitionArn, found bool)

	// Stores a definition in cache.
	Put(key CacheKey, definition JobDefinitionArn) error
}

// A generic CacheKey interface
type CacheKey interface {
	fmt.Stringer
}

type cacheKey struct {
	role               string
	image              string
	jobDefinitionInput batch.RegisterJobDefinitionInput
}

func (k cacheKey) String() string {
	return fmt.Sprintf("%v-%v-%v", k.image, k.role, xhashes.FNV64a(k.jobDefinitionInput.String()))
}

type cache struct {
	raw *freecache.Cache
}

func (c cache) Get(key CacheKey) (jobDefinition JobDefinitionArn, found bool) {
	if raw, err := c.raw.Get([]byte(key.String())); err == nil {
		return string(raw), true
	}

	return "", false
}

func (c cache) Put(key CacheKey, definition JobDefinitionArn) error {
	return c.raw.Set([]byte(key.String()), []byte(definition), 0)
}

// Creates a new deterministic cache key.
func NewCacheKey(role, image string, jobDefinitionInput batch.RegisterJobDefinitionInput) CacheKey {
	return cacheKey{
		role:               role,
		image:              image,
		jobDefinitionInput: jobDefinitionInput,
	}
}

// Creates a new cache using cache size from aws config.
func NewCache(jobDefinitionCacheSize int) Cache {
	return cache{
		raw: freecache.NewCache(jobDefinitionCacheSize),
	}
}

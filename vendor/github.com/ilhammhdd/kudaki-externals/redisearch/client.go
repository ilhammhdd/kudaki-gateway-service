package redisearch

import (
	"github.com/RediSearch/redisearch-go/redisearch"
)

type RedisClient interface {
	Name() string
	Schema() *redisearch.Schema
}

type Client int

const (
	Item Client = iota
	Storefront
	User
)

func (c Client) Name() string {
	return []string{
		"item",
		"storefront",
		"user",
	}[c]
}

func (c Client) Schema() *redisearch.Schema {
	return []*redisearch.Schema{
		new(itemSchema).makeSchema(),
		new(storefrontSchema).makeSchema(),
		new(userSchema).makeSchema(),
	}[c]
}

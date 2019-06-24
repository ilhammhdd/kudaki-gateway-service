package redisearch

import (
	"github.com/RediSearch/redisearch-go/redisearch"
)

type itemSchema struct{}

func (it *itemSchema) makeSchema() *redisearch.Schema {
	schema := redisearch.NewSchema(redisearch.DefaultOptions)
	schema.Fields = append(schema.Fields, itemFields...)
	schema.Fields = append(schema.Fields, storefrontFields...)
	schema.Fields = append(schema.Fields, userFields...)

	return schema
}

type storefrontSchema struct{}

func (ss *storefrontSchema) makeSchema() *redisearch.Schema {
	schema := redisearch.NewSchema(redisearch.DefaultOptions)
	schema.Fields = append(schema.Fields, storefrontFields...)
	schema.Fields = append(schema.Fields, userFields...)

	return schema
}

type userSchema struct{}

func (us *userSchema) makeSchema() *redisearch.Schema {
	schema := redisearch.NewSchema(redisearch.DefaultOptions)
	schema.Fields = append(schema.Fields, userFields...)

	return schema
}

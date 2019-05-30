package kudakiredisearch

import (
	"github.com/RediSearch/redisearch-go/redisearch"
)

type RentalSchema int

const (
	CartsSchema RentalSchema = iota
	CheckoutsSchema
	CartItemsSchema
)

func (rs RentalSchema) Schema() *redisearch.Schema {
	return []*redisearch.Schema{
		redisearch.NewSchema(redisearch.DefaultOptions).
			AddField(redisearch.NewSortableNumericField("id")).
			AddField(redisearch.NewTextField("uuid")).
			AddField(redisearch.NewTextField("user_uuid")).
			AddField(redisearch.NewNumericField("total_price")).
			AddField(redisearch.NewNumericField("open")),
		redisearch.NewSchema(redisearch.DefaultOptions).
			AddField(redisearch.NewSortableNumericField("id")).
			AddField(redisearch.NewTextField("uuid")).
			AddField(redisearch.NewTextField("cart_uuid")).
			AddField(redisearch.NewTextField("issued_at")),
		redisearch.NewSchema(redisearch.DefaultOptions).
			AddField(redisearch.NewSortableNumericField("id")).
			AddField(redisearch.NewTextField("uuid")).
			AddField(redisearch.NewTextField("cart_uuid")).
			AddField(redisearch.NewTextField("item_uuid")).
			AddField(redisearch.NewSortableNumericField("total_item")).
			AddField(redisearch.NewSortableNumericField("total_price")).
			AddField(redisearch.NewSortableNumericField("unit_price")),
	}[rs]
}

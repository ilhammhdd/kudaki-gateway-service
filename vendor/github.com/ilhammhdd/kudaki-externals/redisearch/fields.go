package redisearch

import "github.com/RediSearch/redisearch-go/redisearch"

var itemFields = []redisearch.Field{
	redisearch.NewSortableNumericField("item_id"),
	redisearch.NewTextField("item_uuid"),
	redisearch.NewTextField("item_name"),
	redisearch.NewSortableNumericField("item_amount"),
	redisearch.NewTextField("item_unit"),
	redisearch.NewSortableNumericField("item_price"),
	redisearch.NewTextField("item_description"),
	redisearch.NewTextField("item_photo"),
	redisearch.NewSortableNumericField("item_rating"),
	redisearch.NewSortableNumericField("item_created_at"),
}

var storefrontFields = []redisearch.Field{
	redisearch.NewSortableNumericField("storefront_id"),
	redisearch.NewTextField("storefront_uuid"),
	redisearch.NewSortableNumericField("storefront_total_item"),
	redisearch.NewSortableNumericField("storefront_rating"),
	redisearch.NewSortableNumericField("storefront_created_at"),
}

var userFields = []redisearch.Field{
	redisearch.NewTextField("user_uuid"),
	redisearch.NewTextField("user_email"),
	redisearch.NewTextField("user_password"),
	redisearch.NewTextField("user_token"),
	redisearch.NewTagField("user_role"),
	redisearch.NewTextField("user_phone_number"),
	redisearch.NewTagField("user_account_type"),
}

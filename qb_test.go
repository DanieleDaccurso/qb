package qb

import (
	"testing"
)

func TestQueryBuilder_Select(t *testing.T) {
	new(QueryBuilder).Select("v.title", "v.duration", "w.name").
		From("videos", "v").
		Join("websites", "w", "ON", "w.id = v.website_id").
		OrderBy("v.title", "ASC")

}

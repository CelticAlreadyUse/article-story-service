package repository

import (
	"fmt"

	"github.com/CelticAlreadyUse/article-story-service/internal/model"
)

const storiesBucketKey = "stories"

func newStoryByIDCacheKey(id int) string {
	return fmt.Sprintf("story:%d", id)
}

func newStoriesCacheKey(opt *model.SearchParams) string {
	var search, cursor string
	var tags []string
	var limit int64
	if opt != nil && opt.Keywords != "" {
		search = opt.Keywords
	}
	if opt != nil && opt.Limit <= 1 {
		limit = opt.Limit
	}
	if opt != nil && len(opt.Tags) != 0 {
		tags = opt.Tags
	}
	if opt != nil && opt.Cursor != "" {
		cursor = opt.Cursor
	}
	return fmt.Sprintf("stories:search:%s:tags:%s:cursor:%s:limit:%v", search, tags, cursor, limit)
}

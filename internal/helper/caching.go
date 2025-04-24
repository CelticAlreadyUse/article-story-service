package helper

import (
	"fmt"

	"github.com/CelticAlreadyUse/article-story-service/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
const StoriesBucketKey = "stories"

func NewStoryByIDCacheKey(id primitive.ObjectID) string {
	return fmt.Sprintf("story:%d", id)
}
func NewStoryByUseridKey(id int64) string {
	return fmt.Sprintf("user_id:%d", id)
}
func NewStoriesCacheKey(opt *model.SearchParams) string {
	var keywords, cursor string
	var tags []string
	var limit int64
	if opt != nil && opt.Keywords != "" {
		keywords = opt.Keywords
	}
	if opt != nil && opt.Limit > 0 {
		limit = opt.Limit
	}
	if opt != nil && len(opt.Tags) > 0 {
		tags = opt.Tags
	}
	if opt != nil && opt.Cursor != "" {
		cursor = opt.Cursor
	}
	return fmt.Sprintf("stories:search:%s:tags:%s:cursor:%s:limit:%v", keywords, tags, cursor, limit)
}

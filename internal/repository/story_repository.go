package repository

import (
	"context"
	"time"

	"github.com/CelticAlreadyUse/article-story-service/internal/helper"
	"github.com/CelticAlreadyUse/article-story-service/internal/model"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StoryRepository struct {
	db    *mongo.Database
	redis model.RedisClient
}

func InitStoryStruct(collection *mongo.Database, redis model.RedisClient) model.StoryRepository {
	return &StoryRepository{db: collection, redis: redis}
}

func (r *StoryRepository) Create(ctx context.Context, story model.Story) (*model.Story, error) {
	story.Created_at = time.Now()
	res, err := r.db.Collection("story_service").InsertOne(ctx, story)
	if err != nil {
		return nil, err
	}
	logrus.Info(res.InsertedID)
	story.ID = res.InsertedID.(primitive.ObjectID)
	go func() {
		err := r.redis.HDelByBucketKey(context.Background(), helper.StoriesBucketKey)
		if err != nil {
			log.Errorf("failed to delete data from redis %v", err)
		}
	}()
	return &story, nil
}
func (u *StoryRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	opts := bson.D{primitive.E{Key: "_id", Value: id}}
	_, err := u.db.Collection("story_service").DeleteOne(ctx, opts)

	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	go func() {
		err = u.redis.Del(context.Background(), helper.NewStoryByIDCacheKey(id))
		if err != nil {
			log.Errorf("failed when delete data from redis, error: %v", err)
		}
		err = u.redis.HDelByBucketKey(context.Background(), helper.StoriesBucketKey)
		if err != nil {
			log.Errorf("failed when delete data from redis, error: %v", err)
		}
	}()
	return nil
}
func (u *StoryRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*model.Story, error) {
	var row *model.Story
	err := u.db.Collection("story_service").FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&row)
	if err != nil {
		return nil, err
	}
	return row, nil
}
func (u *StoryRepository) GetAll(ctx context.Context, params *model.SearchParams) ([]model.Story, string, error) {
	var stories []model.Story
	var nextCursor string
	filter := bson.M{}
	if params.Keywords != "" {
		filter["title"] = bson.M{
			"$regex":   params.Keywords,
			"$options": "i",
		}
	}
	if len(params.Tags) > 0 {
		filter["tags_id"] = bson.M{
			"tags_id": bson.M{"$in": params.Tags},
		}
	}
	if params.Cursor != "" {
		cursor, err := helper.DecodeCursor(params.Cursor)
		if err != nil {
			logrus.Error("decode cursor error: ", err)
			return nil, "", err
		}
		filter["$or"] = []bson.M{
			{"created_at": bson.M{"$lt": cursor.Time}},
			{
				"created_at": cursor.Time,
				"_id":        bson.M{"$lt": cursor.ID},
			},
		}
	}
	queryLimit := params.Limit + 1
	opts := options.Find().
		SetSort(bson.D{
			{Key: "created_at", Value: -1},
			{Key: "_id", Value: -1},
		}).
		SetLimit(int64(queryLimit))
	rows, err := u.db.Collection("story_service").Find(ctx, filter, opts)
	if err != nil {
		logrus.Error("find error: ", err)
		return nil, "", err
	}
	defer rows.Close(ctx)

	if err := rows.All(ctx, &stories); err != nil {
		logrus.Error("cursor decode error: ", err)
		return nil, "", err
	}
	logrus.Infof("Len Stories: %d,limit:%d", len(stories), params.Limit)
	if len(stories) > int(params.Limit) {
		last := stories[params.Limit-1]
		nextCursor = helper.EncodeCursor(last.Created_at, last.ID)
		stories = stories[:params.Limit]
	}
	return stories, nextCursor, nil
}
func (u *StoryRepository) Update(ctx context.Context, id primitive.ObjectID, body model.Story) (*model.Story, int64, error) {
	body.Updated_at = time.Now()
	newStory := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: body.Title},
			{Key: "tags_id", Value: body.TagsID},
			{Key: "content", Value: body.Content},
			{Key: "updated_at", Value: body.Updated_at},
		}},
	}
	filter := bson.D{{Key: "_id", Value: id}}
	results, err := u.db.Collection("story_service").UpdateOne(ctx, filter, newStory)
	if err != nil {
		logrus.Error(err)
		return nil, 0, err
	}
	go func() {
		err = u.redis.Del(context.Background(), helper.NewStoryByIDCacheKey(id))
		if err != nil {
			log.Errorf("failed when delete data from redis, error: %v", err)
		}
		err = u.redis.HDelByBucketKey(context.Background(), helper.StoriesBucketKey)
		if err != nil {
			log.Errorf("failed when delete data from redis, error: %v", err)
		}
	}()

	return &body, results.ModifiedCount, nil
}

func (u *StoryRepository) GetStoriesByUserID(ctx context.Context, author_id int64) ([]*model.Story, error) {
	var stories []*model.Story
	filter := bson.M{"author_id": author_id}
	opts := options.Find().
		SetSort(bson.D{
			{Key: "created_at", Value: -1},
			{Key: "_id", Value: -1},
		})
	rows, err := u.db.Collection("story_service").Find(ctx, filter, opts)
	if err != nil {
		logrus.Error("find error: ", err)
		return nil, err
	}
	defer rows.Close(ctx)
	if err := rows.All(ctx, &stories); err != nil {
		logrus.Error("cursor decode error: ", err)
		return nil, err
	}
	return stories, nil
}

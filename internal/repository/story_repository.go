package repository

import (
	"context"
	"time"

	"github.com/CelticAlreadyUse/article-story-service/internal/helper"
	"github.com/CelticAlreadyUse/article-story-service/internal/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StoryRepository struct {
	db *mongo.Database
}

func InitStoryStruct(collection *mongo.Database) model.StoryRepository {
	return &StoryRepository{db: collection}
}

func (r *StoryRepository) Create(ctx context.Context, story model.Story) error {
	story.Created_at = time.Now()
	insertResults, err := r.db.Collection("story_service").InsertOne(ctx, story)
	if err != nil {
		return err
	}
	if _, ok := insertResults.InsertedID.(primitive.ObjectID); ok {
		return nil
	} else {
		return err
	}
}
func (u *StoryRepository) Delete(ctx context.Context, id string) {
	panic("Implmenet me")
}
func (u *StoryRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*model.Story, error) {
	var row *model.Story
	err := u.db.Collection("story_service").FindOne(ctx, bson.D{primitive.E{Key: "_id", Value: id}}).Decode(&row)
	if err != nil {
		return nil, err
	}
	return row, nil
}
func (u *StoryRepository) GetAll(ctx context.Context, params model.SearchParams) ([]model.Story, string, error) {
	filter := bson.M{}
	if params.Keywords != "" {
		filter["title"] = bson.M{
			"$regex":   params.Keywords,
			"$options": "i",
		}
	}
	if len(params.Tags) > 0 {
		filter["tags"] = bson.M{"$in": params.Tags}
	}
	if params.Cursor != "" {
		cursor, err := helper.DecodeCursor(params.Cursor)
		if err != nil {
			logrus.Error(err)
			return nil, "", err
		}
		filter["$or"] = []bson.M{
			{"created_at": bson.M{"$gt": cursor.Time}},
			{
				"created_at": cursor.Time,
				"_id":        bson.M{"$gt": cursor.ID},
			},
		}
	}

	// Validate and set default limit
	if params.Limit <= 0 || params.Limit > 100 {
		params.Limit = 10
	}

	opts := options.Find().
		SetSort(bson.D{{"created_at", 1}, {"_id", 1}}).
		SetLimit(int64(params.Limit))

	rows, err := u.db.Collection("story_service").Find(ctx, filter, opts)
	if err != nil {
		logrus.Error(err)
		return nil, "", err
	}
	defer rows.Close(ctx)

	var stories []model.Story
	if err := rows.All(ctx, &stories); err != nil {
		logrus.Error(err)
		return nil, "", err
	}
	var nextCursor string
	if len(stories) == int(params.Limit) {
		last := stories[len(stories)-1]
		nextCursor = helper.EncodeCursor(last.Created_at, last.ID)
	}

	return stories, nextCursor, nil
}

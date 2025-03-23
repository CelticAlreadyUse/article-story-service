package repository

import (
	"context"
	"time"

	"github.com/CelticAlreadyUse/article-story-service/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

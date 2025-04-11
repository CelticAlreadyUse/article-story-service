package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StoryUsecase interface {
	Create(ctx context.Context, Story Story) (*Story,error)
	Delete(ctx context.Context, id primitive.ObjectID) (error)
	Update(ctx context.Context, id primitive.ObjectID, story Story) (*Story,int64, error)
	GetAll(ctx context.Context, params SearchParams) ([]Story,string, error)
	GetStoryByID(ctx context.Context, userID string) (*Story, error)

	//Users and account
	GetStoriesByUserID(ctx context.Context,id int64)([]*Story,error)

}
type ParamsShowStories struct {
	Limit  int
	Offset int
}
type StoryRepository interface {
	Create(ctx context.Context, story Story) (*Story,error)
	GetAll(ctx context.Context, params SearchParams) ([]Story,string, error)
	Delete(ctx context.Context, id primitive.ObjectID) error
	Update(ctx context.Context, id primitive.ObjectID, story Story) (*Story,int64, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*Story, error)

	//Story and Account 
	GetStoriesByUserID(ctx context.Context,id int64)([]*Story,error)
}
type Story struct {
	ID         primitive.ObjectID        `json:"id,omitempty" bson:"_id,omitempty"`
	AuthorID   int64          `json:"author_id" bson:"author_id" validate:"required"`
	Title      string         `json:"title" bson:"title" validate:"required"`
	Tags       []string       `json:"tags" bson:"tags" validate:"required"`
	Created_at time.Time      `json:"created_at" bson:"created_at"`
	Updated_at time.Time     `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Content    []*StoryElement `json:"content" bson:"content" validate:"required"`
}
type Type string

const (
	PARAGRAPH Type = "paragraph"
	IMAGE     Type = "image"
)

type StoryElement struct {
	Type            Type             `json:"type,omitempty" bson:"type"`
	Text            *string          `json:"text,omitempty" bson:"text,omitempty"`
	ParagraphStyles *ParagraphStyles `json:"paragraph_styles,omitempty" bson:"paragraph_styles,omitempty"`
	Url             *string          `json:"url,omitempty" bson:"url,omitempty"`
	ImageStyles     *ImageStyles     `json:"image_styles,omitempty" bson:"image_styles,omitempty"`
	Caption         *string          `json:"caption,omitempty" bson:"caption,omitempty"`
	AltText         *string          `json:"alt_text,omitempty" bson:"alt_text,omitempty"`
}
type FontStyle string
type FontSize string
type FontStyleQuote string

const (
	STANDART FontSize = "default"
	HEADING  FontSize = "heading"
	SUBTITLE FontSize = "subtitle"
)
const (
	NORMAL     FontStyleQuote = "normal"
	BLACKQUOTE FontStyleQuote = "blackquote"
	PULLQUOTE  FontStyleQuote = "pullquote"
)
const (
	MARKUP FontStyle = "markup"
	STRONG FontStyle = "strong"
	BASIC  FontStyle = "basic"
)

type ParagraphStyles struct {
	FontSize   FontSize       `json:"font-size"`
	FontStyle  FontStyle      `json:"font_style"`
	FontFamily FontStyleQuote `json:"font_family"`
}
type ImageSize string
type ImageStyles struct {
	ImageSize ImageSize
}

const (
	DEFAULT        ImageSize = "Default"
	OUTSIDE_COLUMN ImageSize = "outside_column"
	FULL_SIZE      ImageSize = "full_size"
)

type Image struct {
	Url         string      `json:"url"`
	ImageStyles ImageStyles `json:"styles"`
	Caption     string      `json:"caption"`
	AltText     string      `json:"alt_text"`
}

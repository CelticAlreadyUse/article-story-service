package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SearchParams struct{
	Tags []string `params:"tags"`
	Keywords string `params:"keywords"`
	Limit int64 `params:"limit" validate:"number"`
	Cursor string `params:"cursor"`
}
type CategoryParams struct{
	Alph string `params:"alph"`
	Limit int64 `params:"limit"`
	Keyword string `params:"keyword"`
}
type Pagination struct{
	Before string
	HasBefore bool
	After string
	HasAfter string
}

type Cursor struct{
	Time time.Time `json:"t"`
	ID primitive.ObjectID `json:"i"`
}
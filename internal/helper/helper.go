package helper

import (
	"encoding/base64"
	"encoding/json"
	"regexp"
	"strings"
	"time"

	"github.com/CelticAlreadyUse/article-story-service/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EncodeCursor(t time.Time, id primitive.ObjectID) string {
	data, _ := json.Marshal(model.Cursor{Time: t, ID: id})

	return base64.URLEncoding.EncodeToString(data)
}
func DecodeCursor(s string) (model.Cursor, error) {
	data, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return model.Cursor{}, err
	}
	var c model.Cursor
	if err := json.Unmarshal(data, &c); err != nil {
		return model.Cursor{}, err
	}
	return c, nil
}
func Slugify(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, "_", "-")
	slug = strings.ReplaceAll(slug, " ", "-")
	reg := regexp.MustCompile(`[^a-z0-9\-]`)
	slug = reg.ReplaceAllString(slug, "")
	slug = regexp.MustCompile(`-+`).ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	return slug
}
func NotBlank(s string) bool {
	return strings.TrimSpace(s) != ""
}
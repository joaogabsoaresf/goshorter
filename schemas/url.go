package schemas

import "go.mongodb.org/mongo-driver/bson/primitive"

type Url struct {
	OriginalPath string
	Domain       string
	ShorterID    string
}

type UrlResponse struct {
	ID           primitive.ObjectID `json:"_id"`
	OriginalPath string             `json:"original_path"`
	Domain       string             `json:"domain"`
	ShorterID    string             `json:"shorter_id"`
}

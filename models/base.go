package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Base 结果 基础 字段
type Base struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	CreateAt time.Time          `json:"create_at" bson:"create_at"`
	Name     string             `json:"name" bson:"name"`
	URL      string             `json:"url" bson:"url"`
	Adapter  string             `json:"adapter" bson:"adapter"`
	Args     []string           `json:"args" bson:"args"`
	Regex    [][]string         `json:"regex" bson:"regex"`
	Error    string             `json:"error" bson:"error"`
}

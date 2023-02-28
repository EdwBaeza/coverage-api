package domain

import (
	"time"
)

type BaseModel struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

package EntityTypes

import (
	"time"
)

type Document struct {
	UpdatedAt time.Time `json:"updatedat"`
	CreatedAt time.Time `json:"createdat"`
}

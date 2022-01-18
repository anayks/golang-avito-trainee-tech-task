package ChatUser

import (
	"time"
)

type ChatUser struct {
	ID         int64  `json:"user"`
	Username   string `json:"username"`
	Created_at time.Time
}

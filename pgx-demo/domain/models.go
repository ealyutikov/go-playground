package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	IsActive  bool
	CreatedAt time.Time
}

type Outbox struct {
	ID        uint64
	Data      []byte
	CreatedAt time.Time
}

package token

import (
	"time"

	"github.com/google/uuid"
)

// Payload contains thee payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"uuid"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issed_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

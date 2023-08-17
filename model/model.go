package model

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"index"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if base.ID == uuid.Nil {
		base.ID = uuid.New()
	} else {
		err = errors.New("Can't save UUID to DB")
	}
	return
}

type User struct {
	Base
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Perishable struct {
	Base
	UserID         uuid.UUID `json:"userId"`
	Name           string    `json:"name"`
	Quantity       int       `json:"quantity"`
	IsConsumed     bool      `json:"isConsumed"`
	ExpirationDate time.Time `json:"expirationDate"`
}

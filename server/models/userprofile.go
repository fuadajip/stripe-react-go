package models

import "time"

type UserProfile struct {
	ID        uint       `json:"id" gorm:"column:id;primary_key" validate:"-"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at" validate:"-"`
	CreatedBy *int64     `json:"created_by" gorm:"column:created_by" validate:"-"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at" validate:"-"`
	UpdatedBy *int64     `json:"updated_by" gorm:"column:updated_by" validate:"-"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at" validate:"-"`
	UserID    *int64     `json:"user_id" gorm:"column:user_id" validate:"required"`
	Fullname  *string    `json:"fullname" gorm:"column:fullname" validate:"-"`
	Gender    *string    `json:"gender" gorm:"golumn:gender" validate:"-"`
	Avatar    *string    `json:"avatar" gorm:"column:avatar" validate:"-"`
	Bio       *string    `json:"bio" gorm:"column:bio" validate:"-"`
}

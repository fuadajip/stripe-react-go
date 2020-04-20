package models

import "time"

type User struct {
	ID        *int64     `json:"id" gorm:"column:id;primary_key" validate:"-"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at" validate:"-"`
	CreatedBy *int64     `json:"created_by" gorm:"column:created_by" validate:"-"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at" validate:"-"`
	UpdatedBy *int64     `json:"updated_by" gorm:"column:updated_by" validate:"-"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at" validate:"-"`
	Username  *string    `json:"username" gorm:"column:username" validate:"required"`
	Email     *string    `json:"email"  gorm:"column:email" validate:"required"`
	Phone     *string    `json:"phone"  gorm:"column:phone" validate:"required"`
	Password  *string    `json:"password"  gorm:"column:password" validate:"required"`
	IsActive  *bool      `json:"is_active" gorm:"column:is_active" validate:"-"`
	Type      *string    `json:"type" gorm:"column:type" validate:"-"`
}

type UserRegistrationRequest struct {
	Fullname *string `json:"fullname" validate:"required"`
	Username *string `json:"username" validate:"required"`
	Email    *string `json:"email" validate:"required"`
	Phone    *string `json:"phone" validate:"required"`
	Password *string `json:"password" validate:"required"`
}

type UserRegistrationResponse struct {
	Fullname *string `json:"fullname"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
}

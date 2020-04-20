package models

import "time"

type UserAddress struct {
	ID            uint       `json:"id" gorm:"column:id;primary_key" validate:"-"`
	CreatedAt     time.Time  `json:"created_at" gorm:"column:created_at" validate:"-"`
	CreatedBy     *int64     `json:"created_by" gorm:"column:created_by" validate:"-"`
	UpdatedAt     time.Time  `json:"updated_at" gorm:"column:updated_at" validate:"-"`
	UpdatedBy     *int64     `json:"updated_by" gorm:"column:updated_by" validate:"-"`
	DeletedAt     *time.Time `json:"deleted_at" gorm:"column:deleted_at" validate:"-"`
	UserProfileID *int64     `json:"user_profile_id" gorm:"column:user_profile_id" validate:"required"`
	Address       *string    `json:"address" gorm:"column:address" validate:"-"`
	PostalCode    *string    `json:"postal_code" gorm:"column:postal_code" validate:"-"`
	CityCode      *string    `json:"city_code" gorm:"column:city_code" validate:"-"`
	StateCode     *string    `json:"state_code" gorm:"golumn:state_code" validate:"-"`
	CountryCode   *string    `json:"country_code" gorm:"column:country_code" validate:"-"`
}

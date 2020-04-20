package models

import "time"

type Transaction struct {
	ID         uint       `json:"id" gorm:"column:id;primary_key" validate:"-"`
	CreatedAt  time.Time  `json:"created_at" gorm:"column:created_at" validate:"-"`
	CreatedBy  *int64     `json:"created_by" gorm:"column:created_by" validate:"-"`
	UpdatedAt  time.Time  `json:"updated_at" gorm:"column:updated_at" validate:"-"`
	UpdatedBy  *int64     `json:"updated_by" gorm:"column:updated_by" validate:"-"`
	DeletedAt  *time.Time `json:"deleted_at" gorm:"column:deleted_at" validate:"-"`
	UserID     *int64     `json:"user_id" gorm:"column:user_id" validate:"required"`
	Invoice    *string    `json:"invoice" gorm:"column:invoice" validate:"-"`
	Currency   *string    `json:"currency" gorm:"column:currency" validate:"-"`
	Total      *float64   `json:"total" gorm:"column:total" validate:"-"`
	AdminFee   *float64   `json:"admin_fee" gorm:"column:admin_fee" validate:"-"`
	Discount   *float64   `json:"discount" gorm:"golumn:discount" validate:"-"`
	Shipping   *float64   `json:"shipping" gorm:"golumn:shipping" validate:"-"`
	Tax        *float64   `json:"tax" gorm:"golumn:tax" validate:"-"`
	GrandTotal *float64   `json:"grand_total" gorm:"golumn:grand_total" validate:"-"`
}

package models

import "time"

type Order struct {
	ID            uint       `json:"id" gorm:"column:id;primary_key" validate:"-"`
	CreatedAt     time.Time  `json:"created_at" gorm:"column:created_at" validate:"-"`
	CreatedBy     *int64     `json:"created_by" gorm:"column:created_by" validate:"-"`
	UpdatedAt     time.Time  `json:"updated_at" gorm:"column:updated_at" validate:"-"`
	UpdatedBy     *int64     `json:"updated_by" gorm:"column:updated_by" validate:"-"`
	DeletedAt     *time.Time `json:"deleted_at" gorm:"column:deleted_at" validate:"-"`
	TransactionID *int64     `json:"transaction_id" gorm:"column:transaction_id" validate:"required"`
	ProductCode   *string    `json:"product_code" gorm:"column:product_code" validate:"-"`
	Qty           *float64   `json:"qty" gorm:"column:qty" validate:"-"`
	Price         *float64   `json:"price" gorm:"column:price" validate:"-"`
	Note          *string    `json:"note" gorm:"golumn:note" validate:"-"`
}

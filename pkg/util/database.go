package util

import (
	"gorm.io/gorm"
)

type DatabaseTx struct {
	*gorm.DB
}

func NewDatabaseTx(db *gorm.DB) *DatabaseTx {
	return &DatabaseTx{db}
}

func (d *DatabaseTx) Tx(fn func(q *gorm.DB) error) error {
	tx := d.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	fn(tx)

	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	return tx.Commit().Error
}

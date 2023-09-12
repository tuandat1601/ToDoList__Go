package storage

import "gorm.io/gorm"

type sqlStrore struct {
	db *gorm.DB
}
func NewSQLStrore(db *gorm.DB) *sqlStrore{
	return &sqlStrore{db:db}
}
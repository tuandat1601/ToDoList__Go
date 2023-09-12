package model

import (
	"errors"
	"todolistgo/common"
)

var (
	ErrTitleIsBlank = errors.New("title cann't is blank")
	ErrItemDeleted = errors.New("Item is deleted")
)

type TodoItem struct {
	common.SQLModel
	Title  string      `json:"title" gorm:"column:title"`
	Status *ItemStatus `json:"status" gorm:"column:status"`
}

type TodoItemCreation struct {
	Id     int         `json:"id" gorm:"column:id;"`
	Title  string      `json:"title" gorm:"column:title;"`
	Status *ItemStatus `json:"status" gorm:"column:status;"`
}

func (TodoItem) TableName() string         { return "todo_items" }
func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

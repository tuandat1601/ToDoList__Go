package storage

import (
	"context"
	"todolistgo/modules/item/model"
)

func (s *sqlStrore) CreateItem(ctx context.Context, data *model.TodoItemCreation) error{
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
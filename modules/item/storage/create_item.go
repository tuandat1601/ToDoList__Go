package storage

import (
	"context"
	"todolistgo/common"
	"todolistgo/modules/item/model"
)

func (s *sqlStrore) CreateItem(ctx context.Context, data *model.TodoItemCreation) error{
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
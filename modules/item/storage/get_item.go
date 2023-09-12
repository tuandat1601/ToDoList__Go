package storage

import (
	"context"
	
	"todolistgo/modules/item/model"
)

func (s *sqlStrore) GetItem(ctx context.Context,cond map[string]interface{}) (*model.TodoItem,error){
	var data model.TodoItem
	
	if err:= s.db.Where(cond).First(&data).Error;err!=nil{
		return nil, err
	}
	return &data, nil
}
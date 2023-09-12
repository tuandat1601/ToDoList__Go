package storage

import (
	"context"
	
	"todolistgo/modules/item/model"
)

func (s *sqlStrore) UpdateItem(ctx context.Context,cond map[string]interface{},dataupdate *model.TodoItemCreation) (error){

	
	if err:= s.db.Where(cond).Updates(dataupdate).Error;err!=nil{
		return err
	}
	return  nil
}
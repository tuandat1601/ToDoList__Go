package storage

import (
	"context"
	"todolistgo/modules/item/model"
)

func (s *sqlStrore) DeleteItem(ctx context.Context,cond map[string]interface{}) (error){

	
	if err:= s.db.Table(model.TodoItem{}.TableName()).Where(cond).Updates(map[string]interface{}{
		"status":model.ItemStatusDeleted,
	}).Error;err!=nil{
		return err
	}
	return  nil
}
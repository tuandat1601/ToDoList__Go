package biz

import (
	"context"
	"todolistgo/modules/item/model"
)

type UpdateItemStorage interface {
	GetItem(ctx context.Context,cond map[string]interface{})(*model.TodoItem,error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, dataupdate *model.TodoItemCreation) error
}
type updateItemBiz struct {
	store UpdateItemStorage
}
func UpdateItem(store UpdateItemStorage) *updateItemBiz{
	return &updateItemBiz{store: store}
}
func (biz *updateItemBiz) UpdateItemById(ctx context.Context, id int, dataupdate *model.TodoItemCreation) (error){
	
	
	data, err:=biz.store.GetItem(ctx,map[string]interface{}{"id":id})
	if err!=nil{
		return  err
	}	
	if *data.Status==model.ItemStatusDeleted && data.Status!=nil{
		return model.ErrItemDeleted
	}
	if err :=biz.store.UpdateItem(ctx, map[string]interface{}{"id":id},dataupdate);err!=nil{
		return err
	}
	return nil
}
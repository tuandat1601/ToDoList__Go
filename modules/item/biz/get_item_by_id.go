package biz

import (
	"context"
	"todolistgo/modules/item/model"
)

type GetItemStorage interface {
	GetItem(ctx context.Context,cond map[string]interface{})(*model.TodoItem,error)
}
type getItemBiz struct {
	store GetItemStorage
}
func NewGetItem(store GetItemStorage) *getItemBiz{
	return &getItemBiz{store: store}
}
func (biz *getItemBiz) GetItemById(ctx context.Context, id int) (*model.TodoItem,error){
	data, err:=biz.store.GetItem(ctx,map[string]interface{}{"id":id})
	if err!=nil{
		return nil, err
	}	
	// if data.Status!=nil&& *data.Status == model.ItemStatusDeleted{
	// 	return nil, model.ErrItemDeleted
	// }
	return data, nil
}
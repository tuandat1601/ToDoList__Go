package biz

import (
	"context"
	"todolistgo/common"
	"todolistgo/modules/item/model"
)

type ListItemStorage interface {
	ListItem(ctx context.Context,filter *model.Filter,
		paging *common.Paging, moreKey ...string)([]model.TodoItem,error);
}
type listItemBiz struct{
	store ListItemStorage
}
func NewListItemBiz(store ListItemStorage) *listItemBiz{
	return &listItemBiz{store: store}
}
func (biz * listItemBiz) ListItemById(ctx context.Context,
	filter *model.Filter,paging *common.Paging, moreKey ...string)([]model.TodoItem,error){
	data,err :=biz.store.ListItem(ctx,filter,paging)
	if err!=nil{
		return nil, err
	}
	return data,nil
}
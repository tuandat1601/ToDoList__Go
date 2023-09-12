package biz

import (
	"context"
	"strings"
	"todolistgo/modules/item/model"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreation) error
}
type createItemBiz struct {
	store CreateItemStorage
}
func NewCreateItem(store CreateItemStorage) *createItemBiz{
	return &createItemBiz{store: store}
}
func (biz *createItemBiz) CreateNewItem(ctx context.Context, data *model.TodoItemCreation) error{
	title:=strings.TrimSpace(data.Title)
	if title ==""{
		return model.ErrTitleIsBlank
	}
	if err:=biz.store.CreateItem(ctx,data);err!=nil{
		return err
	}
	return nil
}
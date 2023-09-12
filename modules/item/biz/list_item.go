package biz

import (
	"context"
	"todolistgo/common"
	"todolistgo/modules/item/model"
)

type ListItemStorage interface {
	ListItem(ctx context.Context,cond map[string]interface{},paging *common.Paging)([]model.TodoItem,error);
}
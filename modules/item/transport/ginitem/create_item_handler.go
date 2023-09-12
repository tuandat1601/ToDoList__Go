package ginitem

import (
	"net/http"
	"todolistgo/common"
	"todolistgo/modules/item/biz"
	"todolistgo/modules/item/model"
	"todolistgo/modules/item/storage"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var data model.TodoItemCreation
		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store:=storage.NewSQLStrore(db)
		business:=biz.NewCreateItem(store)
		if err := business.CreateNewItem(ctx.Request.Context(), &data) ; err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}

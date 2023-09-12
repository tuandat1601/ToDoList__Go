package ginitem

import (
	"net/http"
	"strconv"
	"todolistgo/common"
	"todolistgo/modules/item/biz"
	"todolistgo/modules/item/model"

	"todolistgo/modules/item/storage"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func UpdateItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var data model.TodoItemCreation
		if err := ctx.ShouldBind(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store:=storage.NewSQLStrore(db)
		business:=biz.UpdateItem(store)
		// if *data.Status == model.ItemStatusDeleted{
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": model.ErrItemDeleted})
		// 	return
		// }
		if err := business.UpdateItemById(ctx.Request.Context(),id, &data);err!=nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		
		
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

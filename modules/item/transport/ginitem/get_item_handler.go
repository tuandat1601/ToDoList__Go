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

func GetItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store:=storage.NewSQLStrore(db)
		business:=biz.NewGetItem(store)
		data, err := business.GetItemById(ctx.Request.Context(),id)
		if *data.Status == model.ItemStatusDeleted{
			ctx.JSON(http.StatusBadRequest, gin.H{"error": model.ErrItemDeleted})
			return
		}
		
		if err!=nil{
			
			return
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

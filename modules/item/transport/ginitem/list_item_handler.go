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

func ListItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		
		var paging common.Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		paging.Process()
		var filter model.Filter
		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		store:=storage.NewSQLStrore(db)
		business:=biz.NewListItemBiz(store)
		result,err := business.ListItemById(ctx.Request.Context(),&filter, &paging);
		if err!=nil{
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		
		
		ctx.JSON(http.StatusOK, common.NewSuccessResponse(result,paging,filter))
	}
}

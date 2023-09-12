package ginitem

import (
	"net/http"
	"strconv"
	"todolistgo/common"
	"todolistgo/modules/item/biz"

	"todolistgo/modules/item/storage"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func DeleteItem(db *gorm.DB) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSQLStrore(db)
		business := biz.NewDeleteItem(store)
		// if err := db.Table(model.TodoItem{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
		// 	"status": "Deleted",
		// }).Error; err != nil {
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }
		if err := business.DeleteItemById(ctx.Request.Context(), id); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse("Delete successfully"))
	}
}

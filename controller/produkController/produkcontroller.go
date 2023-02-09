package produkcontroller

import (
	"net/http"
	"strconv"

	"github.com/Billy278/belajar-go-restapi-gin/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var produk []model.Product
	model.DB.Find(&produk)
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": "Ok",
		"data":   produk,
	})
}

func Show(c *gin.Context) {
	var produk []model.Product
	id := c.Param("id")
	if err := model.DB.First(&produk, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"code":   http.StatusNotFound,
				"status": "Not Found",
				"data":   err,
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":   http.StatusInternalServerError,
				"status": "Internal Server Error",
				"data":   err,
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusNotFound,
		"status": "Not Found",
		"data":   produk,
	})
}

func Create(c *gin.Context) {
	var produk model.Product

	if err := c.ShouldBindJSON(&produk); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"status": "Bad Request",
			"data":   err,
		})
		return
	}
	model.DB.Create(&produk)
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": "Ok",
		"data":   produk,
	})
}

func Update(c *gin.Context) {
	var produk model.Product

	id := c.Param("id")
	ID, _ := strconv.Atoi(id)
	produk.Id = int64(ID)
	if err := c.ShouldBindJSON(&produk); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"status": "Bad Request",
			"data":   err,
		})
		return
	}
	if model.DB.Model(&produk).Where("id=?", id).Updates(&produk).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"status": "Bad Request",
			"data":   nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": "Ok",
		"data":   produk,
	})

}

func Delete(c *gin.Context) {
	var produk model.Product
	inp := map[string]int{"id": 0}
	// var input struct {
	// 	Id json.Number
	// }
	if err := c.ShouldBindJSON(&inp); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"status": "Bad Request",
			"data":   err,
		})
		return
	}
	//id, _ := input.Id.Int64()
	id := inp["id"]
	//fmt.Println("tes id", id)
	if model.DB.Delete(&produk, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"status": "Bad Request",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"status": "Ok",
	})
}

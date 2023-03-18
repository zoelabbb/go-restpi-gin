package productcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zoelabbb/go-restpi-gin/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var pasien []models.Pasien

	models.DB.Find(&pasien)
	c.JSON(http.StatusOK, gin.H{"pasien": pasien})

}
func Show(c *gin.Context) {
	var pasien models.Pasien
	id := c.Param("id")

	if err := models.DB.First(&pasien, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"pasien": pasien})
}
func Create(c *gin.Context) {

	var pasien models.Pasien

	if err := c.ShouldBindJSON(&pasien); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&pasien)
	c.JSON(http.StatusOK, gin.H{"pasien": pasien})
}
func Update(c *gin.Context) {
	var pasien models.Pasien
	id := c.Param("id")

	if err := c.ShouldBindJSON(&pasien); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&pasien).Where("id = ?", id).Updates(&pasien).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate pasien"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}
func Delete(c *gin.Context) {

	var pasien models.Pasien

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&pasien, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus pasien"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}

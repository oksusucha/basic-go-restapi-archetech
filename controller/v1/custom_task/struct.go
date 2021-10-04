package customtask

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CustomTaskAPIEvent struct {
	DB *gorm.DB
}

type customtaskInterface interface {
	PleaseChangeToYourAPIServiceName(c *gin.Context)
}

package customtask

import "github.com/gin-gonic/gin"

func (api *CustomTaskAPIEvent) PleaseChangeToYourAPIServiceName(c *gin.Context) {
	db := api.DB

	// Start Transaction
	db.Begin()

	// End Transaction
	db.Commit()
}

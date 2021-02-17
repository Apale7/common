package utils

import (
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func RetData(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"msg":         "success",
		"status_code": 0,
		"data":        data,
	})
}

func RetErr(c *gin.Context, err error) {
	if err != nil {
		log.Error("Error: " + err.Error())
		c.JSON(200, gin.H{
			"msg":         err.Error(),
			"status_code": -1,
		})
	} else {
		c.JSON(200, gin.H{
			"msg":         "success",
			"status_code": -1,
		})
	}
}

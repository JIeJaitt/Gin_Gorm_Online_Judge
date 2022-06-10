package service

import (
	"gin_gorm_oj/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProblemList(c *gin.Context) {
	models.GetProblemList()
	c.String(http.StatusOK, "Get Problem List")
}

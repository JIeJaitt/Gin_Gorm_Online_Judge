package router

import (
	_ "gin_gorm_oj/docs"
	"gin_gorm_oj/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// GetProblemList
// @Tags 公共方法
// @Summary 问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Param category_identity query string false "category_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problem-list [get]
func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// 配置路由规则
	r.GET("/ping", service.Ping)
	r.GET(" ", service.GetProblemList)

	return r
}

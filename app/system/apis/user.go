package apis

import (
	"sky/pkg/tools/response"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	response.OK(c, "", "ok")
}

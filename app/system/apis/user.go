package apis

import (
	"sky/app/system/models"
	"sky/pkg/conn"
	"sky/pkg/tools/response"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	response.OK(c, "", "ok")
}

func CreateUser(c *gin.Context) {
	var (
		err       error
		user      models.User
		userCount int64
		password  []byte
	)

	err = c.ShouldBind(&user)
	if err != nil {
		response.Error(c, err, response.InvalidParameterError)
		return
	}

	// 判断用户是否存在
	err = conn.Orm.Model(&models.User{}).Where("username = ?", user.Username).Count(&userCount).Error
	if err != nil {
		response.Error(c, err, response.QueryUserError)
		return
	}
	if userCount > 0 {
		response.Error(c, err, response.UserExistError)
		return
	}

	// 加密密码
	password, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Error(c, err, response.EncryptPasswordError)
		return
	}

	user.Password = string(password)

	// 创建用户
	err = conn.Orm.Create(&user).Error
	if err != nil {
		response.Error(c, err, response.CreateUserError)
		return
	}

	response.OK(c, "", "")
}

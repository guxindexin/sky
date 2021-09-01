package apis

import (
	"sky/app/system/models"
	"sky/pkg/conn"
	"sky/pkg/pagination"
	"sky/pkg/tools/response"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	var (
		err      error
		userList []*models.User
		result   interface{}
	)

	SearchParams := map[string]map[string]interface{}{
		"like": pagination.RequestParams(c),
	}

	result, err = pagination.Paging(&pagination.Param{
		C:  c,
		DB: conn.Orm.Model(&models.User{}),
	}, &userList, SearchParams)
	if err != nil {
		response.Error(c, err, response.UserListError)
		return
	}

	response.OK(c, result, "")
}

func UserInfo(c *gin.Context) {
	var (
		err  error
		user struct {
			models.User
			Password string `json:"-"`
		}
	)

	err = conn.Orm.Model(&models.User{}).Where("username = ?", c.GetString("username")).Find(&user).Error
	if err != nil {
		response.Error(c, err, response.GetUserInfoError)
		return
	}

	response.OK(c, user, "")
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

func UpdateUser(c *gin.Context) {
	var (
		err    error
		userId string
		user   models.User
	)

	userId = c.Param("id")

	err = c.ShouldBind(&user)
	if err != nil {
		response.Error(c, err, response.InvalidParameterError)
		return
	}

	// 更新用户
	err = conn.Orm.Model(&models.User{}).Omit("password").Where("id = ?", userId).Save(&user).Error
	if err != nil {
		response.Error(c, err, response.UpdateUserError)
		return
	}

	response.OK(c, "", "")
}

func DeleteUser(c *gin.Context) {
	var (
		err    error
		userId string
	)

	userId = c.Param("id")

	err = conn.Orm.Delete(&models.User{}, userId).Error
	if err != nil {
		response.Error(c, err, response.DeleteUserError)
		return
	}

	response.OK(c, "", "")
}

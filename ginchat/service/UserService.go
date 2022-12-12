package service

import (
	"ginchat/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Tags 用户
// @Summary 获取用户列表
// @Description 获取用户列表
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()

	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Tags 用户
// @Summary 新增用户
// @Description 创建新用户
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [post]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	if password != repassword {
		c.JSON(-1, gin.H{
			"message": "different password",
		})
		return
	}
	user.Password = password
	models.CreateUser(&user)
	c.JSON(200, gin.H{
		"message": "add successfully",
	})
}

// DeleteUser
// @Tags 用户
// @Summary 删除用户
// @Description 删除用户
// @param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [delete]
func DeleteUser(c *gin.Context) {
	user := &models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"code":    0, //0成功 -1失败
		"message": "delete successfully",
		"data":    user,
	})
}

// UpdateUser
// @Tags 用户
// @Summary 更新用户
// @Description 更新用户
// @param id query string false "id"
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [put]
func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	user := &models.UserBasic{}
	user.ID = uint(id)
	user.Name = c.Query("name")
	user.Password = c.Query("password")
	models.UpdateUser(user)
	c.JSON(200, gin.H{
		"message": "update successfully",
	})
}

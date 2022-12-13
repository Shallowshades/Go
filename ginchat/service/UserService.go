package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"math/rand"
	"strconv"

	"github.com/asaskevich/govalidator"
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
	//user.Password = password

	data := models.FindUserByName(user.Name)
	if data.Name != "" {
		c.JSON(-1, gin.H{
			"message": "user has existed",
		})
		return
	}

	//密码加密
	user.Salt = fmt.Sprintf("%06d", rand.Int31())
	user.Password = utils.MakePassword(password, user.Salt)

	models.CreateUser(&user)
	c.JSON(200, gin.H{
		"message": "add successfully",
	})
}

// login
// @Tags 用户
// @Summary 通过用户名和密码查找用户
// @Description 通过用户名和密码查找用户
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/login [get]
func Login(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"message": "user doesn't exist",
		})
		return
	}
	if !utils.ValidPassword(password, user.Salt, user.Password) {
		c.JSON(200, gin.H{
			"message": "password error",
		})
		return
	}
	password = utils.Md5Encode(password + user.Salt)
	data := models.FindUserByNameAndPwd(name, password)
	if data.Name != "" && data.Password != "" {
		c.JSON(200, gin.H{
			"message": "find successfully",
		})
	}
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
// @param phone query string false "手机号"
// @param email query string false "邮箱"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [put]
func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	user := &models.UserBasic{}
	user.ID = uint(id)
	user.Name = c.Query("name")
	user.Password = c.Query("password")
	user.Phone = c.Query("phone")
	user.Email = c.Query("email")
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": "update failed",
		})
		return
	}
	models.UpdateUser(user)
	c.JSON(200, gin.H{
		"message": "update successfully",
	})
}

//防止跨域站点伪造请求
// var upGrader = websocket.Upgrader(
// 	CheckOrigin : func(r *http.Request)bool{
// 		return true
// 	},
// )

// func RedisMsg(c *gin.Context) {
// 	userIdA, _ := strconv.Atoi(c.PostForm("userIdA"))
// 	userIdB, _ := strconv.Atoi(c.PostForm("userIdB"))
// 	start, _ := strconv.Atoi(c.PostForm("start"))
// 	end, _ := strconv.Atoi(c.PostForm("end"))
// 	isRev, _ := strconv.ParseBool(c.PostForm("isRev"))
// 	res := models.RedisMsg(int64(userIdA), int64(userIdB), int64(start), int64(end), isRev)
// 	utils.RespOKList(c.Writer, "ok", res)
// }

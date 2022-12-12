package models

import (
	"ginchat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Avatar        string //头像
	Identity      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     time.Time
	HeartBeatTime time.Time
	LogoutTime    time.Time //`gorm:"column:login_out_time" json:"login_out_time"`
	IsLogout      bool
	DeviceInfo    string
}

// 表名
func (tb *UserBasic) TableName() string {
	return "user_basic"
}

// 获取用户列表
func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	return data
}

// 按名字查找用户
func FindUserByName(name string) *UserBasic {
	user := &UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}

// 按手机号查找用户
func FindUserByPhone(phone string) *gorm.DB {
	user := &UserBasic{}
	return utils.DB.Where("phone = ?", phone).First(&user)
}

// 按邮箱查找用户
func FindUserByEmail(email string) *gorm.DB {
	user := &UserBasic{}
	return utils.DB.Where("email = ?", email).First(&user)
}

// 创建新用户
func CreateUser(user *UserBasic) *gorm.DB {
	return utils.DB.Create(user)
}

// 删除用户
func DeleteUser(user *UserBasic) *gorm.DB {
	return utils.DB.Delete(user)
}

// 按id删除用户
func DeleteUserById(id uint) *gorm.DB {
	return utils.DB.Delete(id)
}

// 更新用户信息
func UpdateUser(user *UserBasic) *gorm.DB {
	return utils.DB.Model(user).Updates(UserBasic{
		Name:     user.Name,
		Password: user.Password,
	})
}

// 按id查找用户
func FindByID(id uint) *UserBasic {
	user := &UserBasic{}
	utils.DB.Where("id = ?", id).First(user)
	return user
}

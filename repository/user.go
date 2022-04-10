package repository

import (
	"github.com/Moonlight-Zhao/go-project-example/util"
	"gorm.io/gorm"
	"sync"
	"time"
)

type User struct {
	Id         int64     `gorm:"column:id"`
	Name       string    `gorm:"column:name"`
	Avatar     string    `gorm:"column:avatar"`
	Level      int       `gorm:"column:level"`
	CreateTime time.Time `gorm:"column:create_time"`
	ModifyTime time.Time `gorm:"column:modify_time"`
}

func (User) TableName() string {
	return "user"
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		})
	return userDao
}

func (*UserDao) QueryUserById(id int64) (*User, error) {
	var user User
	err := db.Where("id = ?", id).Find(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		util.Logger.Error("find user by id err:" + err.Error())
		return nil, err
	}
	return &user, nil
}

func (*UserDao) MQueryUserById(ids []int64) (map[int64]*User, error) {
	var users []*User
	err := db.Where("id in (?)", ids).Find(&users).Error
	if err != nil {
		util.Logger.Error("batch find user by id err:" + err.Error())
		return nil, err
	}
	userMap := make(map[int64]*User)
	for _, user := range users {
		userMap[user.Id] = user
	}
	return userMap, nil
}

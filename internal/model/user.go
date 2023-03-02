package model

import (
	"oms/global"
	"oms/pkg/enum"
	"oms/pkg/util"
)

type User struct {
	*Model
	Username    string `gorm:"type:varchar(100);not null;" json:"username"`
	Password    string `gorm:"not null;" json:"password"`
	Level       uint8  `gorm:"default:2" json:"level"`
	State       uint8  `gorm:"default:1" json:"state"`
	GroupID     uint32 `gorm:"default:0" json:"group_id"`
	GroupLeader uint8  `gorm:"default:0" json:"group_leader"`
	Salt        string `gorm:"type:varchar(10);not null"`
}

func init() {
	global.ModelAutoMigrate = append(global.ModelAutoMigrate, &User{})
}

func NewUser() *User {
	return &User{
		Level:       enum.DEFAULT_USER_LEVEL,
		State:       enum.DEFAULT_STATE,
		GroupID:     0,
		GroupLeader: 0,
		Model:       &Model{},
	}
}

func (u *User) TableName() string {
	return "oms_user"
}

// 获取并设置salt
func (u *User) SetSalt() {
	if u.Salt == "" {
		u.Salt = util.CreateSalt()
	}
}

// 设置加密密码
func (u *User) SetPassword() error {
	hashPassword, err := util.HashPassword(u.Password + u.Salt)
	u.Password = hashPassword
	return err
}

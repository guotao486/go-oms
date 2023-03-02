/*
 * @Author: GG
 * @Date: 2023-02-28 08:57:41
 * @LastEditTime: 2023-02-28 10:20:50
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\model\demo\auth.go
 *
 */
package demo

import (
	"oms/global"
	"oms/internal/model"

	"github.com/jinzhu/gorm"
)

type Auth struct {
	*model.Model
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

func (a Auth) TableName() string {
	return "blog_auth"
}

// ====== gorm =======
func (a Auth) Get(db *gorm.DB) (Auth, error) {
	var auth Auth
	db = db.Where("app_key = ? and app_secret = ? and is_del = ?", a.AppKey, a.AppSecret, global.IS_DEL_DISABLE)
	err := db.First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return auth, err
	}
	return auth, nil
}

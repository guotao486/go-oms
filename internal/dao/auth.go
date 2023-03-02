/*
 * @Author: GG
 * @Date: 2023-02-28 08:57:46
 * @LastEditTime: 2023-02-28 11:17:47
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\dao\auth.go
 *
 */
package dao

import "oms/internal/model/demo"

func (d *Dao) GetAuth(appKey, appSecret string) (demo.Auth, error) {
	auth := demo.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}

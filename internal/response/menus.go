/*
 * @Author: GG
 * @Date: 2023-03-30 16:45:03
 * @LastEditTime: 2023-03-31 16:11:28
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\response\menus.go
 *
 */
package response

import "oms/internal/model"

type MenusResponse struct {
	ID        uint32 `json:"id"`
	Title     string `json:"title"`
	Router    string `json:"router"`
	Sort      uint8  `json:"sort"`
	Role      string `json:"role"`
	ParentID  uint32 `json:"parent_id"`
	ChildNode []*MenusResponse
}

func (m *MenusResponse) TableName() string {
	return model.NewMenus().TableName()
}

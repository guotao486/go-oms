/*
 * @Author: GG
 * @Date: 2023-03-02 16:52:24
 * @LastEditTime: 2023-03-03 14:08:29
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\response\user.go
 *
 */
package response

import (
	"encoding/json"
	"html/template"
	"oms/pkg/enum"
	"oms/pkg/util"
)

type UserResponse struct {
	ID          uint32 `json:"id"`
	Username    string `json:"username"`
	Level       uint8  `json:"level"`
	LevelText   string `json:"level_text"`
	State       uint8  `json:"state"`
	GroupID     uint32 `json:"group_id"`
	GroupLeader uint8  `json:"group_leader"`
	CreatedOn   uint32 `json:"created_on"`
}

func (u *UserResponse) CustomMarshal() {
	u.getLevelText()
	u.getUsername()
}

func (u *UserResponse) getUsername() string {
	return template.HTMLEscapeString(u.Username)
}

// levelText
func (u *UserResponse) getLevelText() string {
	if u.Level == enum.USER_LEVEL_ADMIN {
		return enum.USER_LEVEL_ADMIN_TEXT
	}

	if u.Level == enum.USER_LEVEL_STAFF {
		return enum.USER_LEVEL_STAFF_TEXT
	}
	return ""
}

func (u *UserResponse) MarshalJSON() ([]byte, error) {
	type Alias UserResponse
	return json.Marshal(&struct {
		CreatedOn string `json:"created_on"`
		LevelText string `json:"level_text"`
		Username  string `json:"username"`
		*Alias
	}{
		Username:  u.getUsername(),
		CreatedOn: util.UnixToString(int64(u.CreatedOn)),
		LevelText: u.getLevelText(),
		Alias:     (*Alias)(u),
	})
}

package response

import (
	"encoding/json"
	"oms/internal/model"
	"oms/pkg/util"
)

type UserGroupResponse struct {
	ID        uint32          `json:"id"`
	Title     string          `json:"title"`
	Leader    uint8           `json:"leader"`
	Members   string          `json:"members"`
	UserList  []*UserResponse `gorm:"foreignKey:GroupID" json:"userList"`
	CreatedOn uint32          `json:"created_on"`
}

func (u *UserGroupResponse) TableName() string {
	return model.NewUserGroup().TableName()
}

func (u *UserGroupResponse) GetUserMember() string {
	member := "组员："
	leader := ""
	if len(u.UserList) > 0 {
		for _, user := range u.UserList {
			if uint8(user.ID) == u.Leader {
				leader += "组长：" + user.getUsername() + "<br/>"
			} else {
				member += user.getUsername() + ","
			}
		}
	}
	return leader + member
}

func (u *UserGroupResponse) MarshalJSON() ([]byte, error) {
	type Alias UserGroupResponse
	return json.Marshal(&struct {
		Member    string `json:"members"`
		CreatedOn string `json:"created_on"`
		*Alias
	}{
		Member:    u.GetUserMember(),
		CreatedOn: util.UnixToString(int64(u.CreatedOn)),
		Alias:     (*Alias)(u),
	})
}

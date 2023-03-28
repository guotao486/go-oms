/*
 * @Author: GG
 * @Date: 2023-03-27 14:54:45
 * @LastEditTime: 2023-03-27 15:14:30
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\pkg\sessions\sessions.go
 *
 */
package sessions

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Sessions struct {
	Engine sessions.Session
}

func NewSession(c *gin.Context) *Sessions {
	Engine := sessions.Default(c)
	return &Sessions{
		Engine: Engine,
	}
}

func (s *Sessions) Set(name interface{}, value interface{}) {
	s.Engine.Set(name, value)
}

func (s *Sessions) Delete(name string) {
	s.Engine.Delete(name)
}

func (s *Sessions) Clear() {
	s.Engine.Clear()
}

func (s *Sessions) Get(name string) interface{} {
	return s.Engine.Get(name)
}

// 调用 session 方法： Set()、 Delete()、 Clear()、方法后，必须调用一次 Save() 方法。否则session数据不会更新
func (s *Sessions) Save() error {
	return s.Engine.Save()
}

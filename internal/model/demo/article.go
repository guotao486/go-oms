/*
 * @Author: GG
 * @Date: 2023-01-28 11:04:28
 * @LastEditTime: 2023-02-28 09:44:30
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\model\demo\article.go
 *
 */
package demo

import (
	"oms/internal/model"
	"oms/pkg/app"
)

type Article struct {
	*model.Model
	Title         string `json:"title"`           // 文章标题
	Desc          string `json:"desc"`            // 文章简述
	Content       string `json:"content"`         // 封面图片地址
	CoverImageUrl string `json:"cover_image_url"` // 文章内容
	State         uint8  `json:"state"`           // 状态 0 禁用 1 启用
}

func (a Article) TableName() string {
	return "blog_article"
}

// 定义一个针对swagger 的对象
type ArticleSwagger struct {
	List  []*Tag
	Pager *app.Pager
}

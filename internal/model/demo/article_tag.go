/*
 * @Author: GG
 * @Date: 2023-02-28 08:57:43
 * @LastEditTime: 2023-02-28 10:19:51
 * @LastEditors: GG
 * @Description:
 * @FilePath: \oms\internal\model\demo\article_tag.go
 *
 */
package demo

import "oms/internal/model"

type ArticleTag struct {
	*model.Model
	TagID     uint32 `json:"tag_id"`     // 标签id
	ArticleID uint32 `json:"article_id"` // 文章id
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}

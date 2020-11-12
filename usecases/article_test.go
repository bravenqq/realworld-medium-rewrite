// Package usecases provides ...
package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"iohttps.com/live/realworld-medium-rewrite/domain"
)

var itor ArticleInteractor

//TestSaveDraft 测试保存草稿
func TestSaveDraft(t *testing.T) {
	a := assert.New(t)
	//1. 创建草稿态文章
	//1.1 保存空内容失败
	article := domain.Article{}
	_, err := itor.SaveDraft(GenerateUUID, article)
	a.NotNil(err)

	//1.2 创建草稿文章成功
	article.Title = "test"
	ID, err := itor.SaveDraft(GenerateUUID, article)
	a.True(assert.Nil(err) && a.IsType(new(domain.NUUID), &ID, nil))

	//2. 更新草稿文章
	article.Title = "testja"
	article.Content = "test"
	ID, err = itor.SaveDraft(GenerateUUID, article)
	a.True(a.Nil(err) && a.IsType(new(domain.NUUID), &ID, nil))
}

func TestPublish(t *testing.T) {
	a := assert.New(t)
	//1.测试发布失败
	//1.1发布的文章不存在
	art := domain.Article{}
	err := itor.Publish(art.ID, 0)
	a.NotNil(err)
	//1.2 作者没有权限
	//1.3 文章的标题或者内容为空
	//测试发布成功

}

package service

import (
	"go-blob/config"
	"go-blob/dao"
	"go-blob/models"
	"html/template"
)

func GetAllIndexInfo(slug string, page, pageSize int) (*models.HomeResponse, error) {
	//var categorys = []models.Category{
	//	{Cid: 0, Name: "123"},
	//}
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	//var posts = []models.PostMore{
	//	{
	//		Pid:          1,
	//		Title:        "go博客",
	//		Content:      "测试内容",
	//		UserName:     "郑海强",
	//		ViewCount:    123,
	//		CreateAt:     "2022-02-20",
	//		CategoryId:   1,
	//		CategoryName: "go",
	//		Type:         0,
	//	},
	//}

	var posts []models.Post
	var total int

	if slug == "" {
		posts, err = dao.GetPostPage(page, pageSize)

		total = dao.CountGetAllPost()

	} else {
		posts, err = dao.GetPostPageBySlug(slug, page, pageSize)

		total = dao.CountGetAllPostBySlug(slug)
	}
	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.CategoryId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}
		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			models.DateDay(post.CreateAt),
			models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}
	//11
	//(11-1)/10 + 1 = 2

	pagesCount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}

	var hr = &models.HomeResponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     postMores,
		Total:     total,
		Page:      page,
		Pages:     pages,
		PageEnd:   page != pagesCount,
	}
	return hr, nil
}

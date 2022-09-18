package service

import (
	"go-blob/config"
	"go-blob/dao"
	"go-blob/models"
	"html/template"
)

func GetPostsByCategoryId(cId, page, pageSize int) (*models.CategoryResponse, error) {
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
	posts, err := dao.GetPostPageByCategoryId(cId, page, pageSize)
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
	total := dao.CountGetAllPostByCategoryId(cId)

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
	categoryName := dao.GetCategoryNameById(cId)
	categoryResponse := &models.CategoryResponse{
		hr,
		categoryName,
	}
	return categoryResponse, nil
}

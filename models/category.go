package models

//从数据库中获取的
type Category struct {
	Cid      int
	Name     string
	CreateAt string
	UpdateAt string
}

type CategoryResponse struct {
	*HomeResponse
	CategoryName string
}

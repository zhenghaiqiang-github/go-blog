package models

import (
	"go-blob/config"
	"html/template"
	"time"
)

type Post struct {
	Pid        int       `json:"pid"`        //文章ID
	Title      string    `json:"title"`      //文章名
	Slug       string    `json:"slug"`       //自定义页面Path
	Content    string    `json:"content"`    //文章的html
	Markdown   string    `json:"markdown"`   //文章的markdown
	CategoryId int       `json:"categoryId"` //分类id
	UserId     int       `json:"userId"`     //用户id
	ViewCount  int       `json:"viewCount"`  //阅览数量
	Type       int       `json:"type"`       //文章类型
	CreateAt   time.Time `json:"createAt"`   //创建时间
	UpdateAt   time.Time `json:"updateAt"`   //修改时间
}

type PostMore struct {
	Pid          int           `json:"pid"`          //文章ID
	Title        string        `json:"title"`        //文章名
	Slug         string        `json:"slug"`         //自定义页面Path
	Content      template.HTML `json:"content"`      //文章的html
	CategoryId   int           `json:"categoryId"`   //分类ID
	CategoryName string        `json:"categoryName"` //分类名
	UserId       int           `json:"userId"`       //用户ID
	UserName     string        `json:"userName"`     //用户名
	ViewCount    int           `json:"viewCount"`
	Type         int           `json:"type"`
	CreateAt     string        `json:"createAt"`
	UpdateAt     string        `json:"updateAt"`
}
type PostReq struct {
	Pid        int               `json:"pid"`     //文章ID
	Title      string            `json:"title"`   //文章名
	Slug       string            `json:"slug"`    //自定义页面Path
	Content    template.Template `json:"content"` //文章的html
	Markdown   string            `json:"markdown"`
	CategoryId int               `json:"categoryId"`
	UserId     int               `json:"userId"`
	Type       int               `json:"type"`
}

type SearchResp struct {
	Pid   int    `json:"pid"` //文章ID
	Title string `json:"title"`
}
type PostRes struct {
	config.Viewer
	config.SystemConfig
	Article PostMore
}
type WritingRes struct {
	Title     string
	CdnURL    string
	Categorys []Category
}

type PigeonholeRes struct {
	config.Viewer
	config.SystemConfig
	Categorys []Category
	Lines     map[string][]Post
}

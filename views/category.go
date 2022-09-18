package views

import (
	"errors"
	"fmt"
	"go-blob/common"
	"go-blob/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	categoryTemplate := common.Template.Category
	//http://localhost:8080/c/1 1参数 分类id
	path := r.URL.Path
	fmt.Println(path)
	cIdStr := strings.TrimPrefix(path, "/c/")
	cId, err := strconv.Atoi(cIdStr)
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("不识别此请求路径"))
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败:", err)
		categoryTemplate.WriteError(w, errors.New("系统错误,请联系管理员！"))
		return
	}
	pageStr := r.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	//每页显示的数量
	pageSize := 10
	categoryResponse, err := service.GetPostsByCategoryId(cId, page, pageSize)
	if err != nil {
		categoryTemplate.WriteError(w, err)
		return
	}
	categoryTemplate.WriteData(w, categoryResponse)
}

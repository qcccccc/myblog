package views

import (
	"errors"
	"gocode/common"
	"gocode/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取数据出错", err)
		index.WriteError(w, errors.New("系统错误"))
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSize := 10
	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/")
	//fmt.Println(slug)

	hr, err := service.GetAllIndexform(slug, page, pageSize)
	if err != nil {
		log.Println("首页获取数据出错", err)
		index.WriteError(w, errors.New("系统错误"))
	}
	index.WriteData(w, hr)
}

package api

import (
	"errors"
	"gocode/common"
	"gocode/dao"
	"gocode/service"
	"net/http"
	"strconv"
	"strings"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	//log.Println(userName)
	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, loginRes)
}

func (*Api) Register(w http.ResponseWriter, r *http.Request) {

	///
	f := errors.New("功能维护中，暂不支持注册")
	common.Error(w, f)
	return
	////

	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)
	//log.Println(userName)
	err := service.Register(userName, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}

	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, loginRes)
}

func (*Api) Delete(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/delete/id=")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		common.Error(w, errors.New("不识别"))
		return
	}
	dao.DeleteByPid(pId)
	common.Success(w, pId)
}

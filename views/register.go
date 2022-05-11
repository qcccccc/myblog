package views

import (
	"gocode/common"
	"gocode/config"
	"net/http"
)

func (*HTMLApi) Register(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("register")
	register := common.Template.Rigister
	register.WriteData(w, config.Cfg.Viewer)
}

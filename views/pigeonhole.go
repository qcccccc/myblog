package views

import (
	"gocode/common"
	"gocode/service"
	"net/http"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pig := common.Template.Pigeonhole

	pigeonholeRes := service.FindPostPigeonhole()
	pig.WriteData(w, pigeonholeRes)

}

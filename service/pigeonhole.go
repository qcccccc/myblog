package service

import (
	"gocode/config"
	"gocode/dao"
	"gocode/models"
)

func FindPostPigeonhole() models.PigeonholeRes {
	categorys, _ := dao.GetAllCategory()
	posts, _ := dao.GetPostAll()
	pigeonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		at := post.CreateAt
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}
	return models.PigeonholeRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		categorys,
		pigeonholeMap,
	}
}

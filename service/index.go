package service

import (
	"gocode/config"
	"gocode/dao"
	"gocode/models"
	"html/template"
)

func GetAllIndexform(slug string, page, pageSize int) (*models.HomeResponse, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	var total int
	if slug == "golang" {
		//	fmt.Println(slug)
		posts, err = dao.GetPostPageBySlug(slug, page, pageSize)
		total = dao.CountGetAllPostBySlug(slug)
	} else {
		//fmt.Println("asdas")
		posts, err = dao.GetPostPage(page, pageSize)
		total = dao.CountGetAllPost()
	}

	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 0 {
			content = content[0:0]
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
	pages := (total + 9) / 10
	var p []int
	for i := 1; i <= pages; i++ {
		p = append(p, i)
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		postMores,
		total,
		page,
		p,
		page != pages,
	}

	return hr, nil
}

package dao

import (
	"fmt"
	"gocode/models"
	"log"
)

func DeleteByPid(pId int) error {
	res, err := DB.Exec("delete from blog_post where pid=?", pId)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		fmt.Println("rows faoiled")
		return err
	}
	return nil
}

func GetCategoryNameById(cId int) string {
	row := DB.QueryRow("select  name from blog_category where cid=?", cId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var categoryName string
	_ = row.Scan(&categoryName)
	return categoryName
}

func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select cid,name from blog_category order by cid")
	if err != nil {
		log.Println("GetAllCategory 查询错误：", err)
		return nil, err
	}
	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name)
		if err != nil {
			log.Println("GetAllCategory 取值出错：", err)
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}

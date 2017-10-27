package main

import (
	"fmt"
	"net/http"

	. "wx-gin/app/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/media"
)

// Create another GORM-backend model

type Result struct {
	Ads []Ad
}

func main() {
	//	DB, _ := gorm.Open("sqlite3", "demo.db")
	db, _ := gorm.Open("mysql", "root:654654@tcp(db:3306)/db3?charset=utf8&parseTime=True&loc=Local")
	media.RegisterCallbacks(db)
	//	db.AutoMigrate(&Ad{})
	//=============================
	var (
		ads []Ad
	)

	//"id","name", "list_pic_url", "retail_price"
	/*
		db.Where("parent_id = ?", 0).Where("name <> ?", "推荐").Find(&categorys)
		db.Limit(3).Find(&topics)
		db.Where("is_new = ?", 1).Order("new_sort_order").Limit(4).Find(&brands)
		db.Select([]string{"id", "name", "list_pic_url", "retail_price"}).Where("is_new = ?", 1).Limit(3).Find(&newgoods)
		db.Select([]string{"id", "name", "list_pic_url", "retail_price"}).Where("is_hot = ?", 1).Limit(3).Find(&hotgoods)
		db.Find(&channels)
		db.Where("ad_position_id = ?", 1).Find(&ads)
	*/
	db.Find(&ads)
	//=============================
	//=============================
	result := Result{Ads: ads}
	//============
	Admin := admin.New(&admin.AdminConfig{DB: db})
	Admin.AddResource(&Ad{})

	//serves
	mux := http.NewServeMux()
	r := gin.New()
	//static dir
	for _, path := range []string{"system", "javascripts", "stylesheets", "images"} {
		r.Static(fmt.Sprintf("/%s", path), fmt.Sprintf("public/%s", path))
	}
	//=========
	Admin.MountTo("/admin", mux)
	r.Any("/admin/*w", gin.WrapH(mux))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, result)
	})
	r.Run(":80")

}

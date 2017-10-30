package main

import (
	"fmt"
	"log"
	"net/http"

	. "wx-gin/app/models"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/media"
)

// Create another GORM-backend model

type Result struct {
	Ads      []Ad
	Channels []Channel
}

func main() {
	//	DB, _ := gorm.Open("sqlite3", "demo.db")
	db, _ := gorm.Open("mysql", "root:654654@tcp(db:3306)/db3?charset=utf8&parseTime=True&loc=Local")
	media.RegisterCallbacks(db)
	db.AutoMigrate(&Ad{}, &Channel{})
	//=============================
	var (
		ads      []Ad
		channels []Channel
	)

	Admin := admin.New(&admin.AdminConfig{DB: db})
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
	db.Find(&channels)
	//=============================
	//=============================
	result := Result{Ads: ads, Channels: channels}
	//============
	Admin.AddResource(&Ad{})
	Admin.AddResource(&Channel{})

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
	//	r.Run(":8080")
	log.Fatal(autotls.Run(r, "wcqt.site"))

}

package main

import (
	"fmt"
	"net/http"

	. "wx-gin/app/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
)

// index/index
type IndexJson struct {
	Ads   []NideshopAd      `json:"adad"`
	Goods []NideshopChannel `json:"channelssss"`
}

// Create a GORM-backend model
type User struct {
	gorm.Model
	Name string
}

// Create another GORM-backend model
type Product struct {
	gorm.Model
	Name        string
	Description string
}

func main() {
	//	DB, _ := gorm.Open("sqlite3", "demo.db")
	db, _ := gorm.Open("mysql", "root:654654@tcp(db:3306)/nideshop?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	//	DB.AutoMigrate(&User{}, &Product{})

	Admin := admin.New(&admin.AdminConfig{DB: db})
	// Allow to use Admin to manage User, Product
	Admin.AddResource(&NideshopGoods{})
	Admin.AddResource(&NideshopAd{})

	ads := []NideshopAd{}
	channels := []NideshopChannel{}

	db.Where("ad_position_id = ?", 1).Find(&ads)
	db.Find(&channels)
	indexjson := &IndexJson{ads, channels}

	//serves
	mux := http.NewServeMux()

	r := gin.Default()
	//static dir
	for _, path := range []string{"system", "javascripts", "stylesheets", "images"} {
		r.Static(fmt.Sprintf("/%s", path), fmt.Sprintf("public/%s", path))
	}
	Admin.MountTo("/admin", mux)
	r.Any("/admin/*w", gin.WrapH(mux))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, indexjson)
	})
	r.Run(":8081")

}

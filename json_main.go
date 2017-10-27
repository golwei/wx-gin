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

type Index struct {
	Channels  []NideshopChannel
	Ads       []NideshopAd
	Newgoods  []NideshopGoods
	Hotgoods  []NideshopGoods
	Brands    []NideshopBrand
	Topics    []NideshopTopic
	Categorys []NideshopCategory
	//	Categorygoods  []NideshopGoods
	Categorygoodss map[int][]NideshopGoods
}

func main() {
	//	DB, _ := gorm.Open("sqlite3", "demo.db")
	db, _ := gorm.Open("mysql", "root:654654@tcp(db:3306)/nideshop?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	media.RegisterCallbacks(db)
	db.AutoMigrate(&Ad{})
	//=============================
	var (
		channels  []NideshopChannel
		ads       []NideshopAd
		newgoods  []NideshopGoods
		hotgoods  []NideshopGoods
		brands    []NideshopBrand
		topics    []NideshopTopic
		categorys []NideshopCategory
	)
	//	var categoryids []int
	//"id","name", "list_pic_url", "retail_price"

	db.Where("parent_id = ?", 0).Where("name <> ?", "推荐").Find(&categorys)
	db.Limit(3).Find(&topics)
	db.Where("is_new = ?", 1).Order("new_sort_order").Limit(4).Find(&brands)
	db.Select([]string{"id", "name", "list_pic_url", "retail_price"}).Where("is_new = ?", 1).Limit(3).Find(&newgoods)
	db.Select([]string{"id", "name", "list_pic_url", "retail_price"}).Where("is_hot = ?", 1).Limit(3).Find(&hotgoods)
	db.Find(&channels)
	db.Where("ad_position_id = ?", 1).Find(&ads)

	//=============================
	categorygoodss := make(map[int][]NideshopGoods)
	for _, category := range categorys {
		//categoryids = append(categoryids, category.Id)
		var categorygoods []NideshopGoods
		var ids []int
		newcategory := make(map[int][]int)

		db.Table("nideshop_category").Where("parent_id = ?", category.Id).Limit(100).Pluck("id", &ids)
		newcategory[category.Id] = ids
		//
		db.Select([]string{"id", "name", "list_pic_url", "retail_price"}).Where("category_id in (?)", ids).Limit(7).Find(&categorygoods)
		categorygoodss[category.Id] = categorygoods
		//		childs := []NideshopCategory{}
		//		db.Select("id").Where("parent_id = ?", category.Id).Limit(100).Find(&childs)
		//		newcategory[category.Id] = childs
	}
	//=============================
	/*
		fmt.Printf("%#v\n", ads)
		fmt.Printf("%#v\n", channels)
		fmt.Printf("%#v\n", newgoods)
		fmt.Printf("%#v\n", hotgoods)
		fmt.Printf("%#v\n", brands)
		fmt.Printf("%#v\n", topics)
		fmt.Printf("%#v\n", categorys)
		fmt.Printf("%#v\n", categorygoodss)
	*/
	//=============================
	index := Index{Ads: ads, Channels: channels, Newgoods: newgoods, Hotgoods: hotgoods, Brands: brands, Topics: topics, Categorys: categorys, Categorygoodss: categorygoodss}
	//============
	Admin := admin.New(&admin.AdminConfig{DB: db})
	// Allow to use Admin to manage User, Product
	Admin.AddResource(&NideshopGoods{})
	Admin.AddResource(&NideshopAd{})
	Admin.AddResource(&Ad{})

	//serves
	mux := http.NewServeMux()
	r := gin.New()
	//static dir
	for _, path := range []string{"system", "javascripts", "stylesheets", "images"} {
		r.Static(fmt.Sprintf("/%s", path), fmt.Sprintf("public/%s", path))
	}
	Admin.MountTo("/admin", mux)
	r.Any("/admin/*w", gin.WrapH(mux))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, index)
	})
	r.Run(":8081")

}

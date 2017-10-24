package main // simulate some private data
import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qor/admin"
	"github.com/qor/media"
	"github.com/qor/media/oss"
	"github.com/qor/qor"
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
	Image       oss.OSS
}

func main() {
	DB, _ := gorm.Open("sqlite3", "demo.db")
	DB.AutoMigrate(&User{}, &Product{})
	media.RegisterCallbacks(DB)

	// Initalize
	Admin := admin.New(&qor.Config{DB: DB})

	// Create resources from GORM-backend model
	Admin.AddResource(&User{})
	Admin.AddResource(&Product{})
	// Register route
	//http service

	mux := http.NewServeMux()
	for _, path := range []string{"system", "javascripts", "stylesheets", "images"} {
		beego.SetStaticPath(fmt.Sprintf("/%s/", path), fmt.Sprintf("public/%s", path))
	}
	Admin.MountTo("/admin", mux)
	beego.Handler("/admin/*", mux)

	beego.Run()
}

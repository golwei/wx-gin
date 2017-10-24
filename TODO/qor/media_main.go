package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qor/admin"
	"github.com/qor/media"
	"github.com/qor/media/oss"
	"github.com/qor/qor"
	"github.com/qor/qor/utils"
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
	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)
	//http service
	for _, path := range []string{"system", "javascripts", "stylesheets", "images"} {
		mux.Handle(fmt.Sprintf("/%s/", path), utils.FileServer(http.Dir("public")))
	}
	// amount to /admin, so visit `/admin` to view the admin interface

	fmt.Println("Listening on: 8081")
	http.ListenAndServe(":8081", mux)
}

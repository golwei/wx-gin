package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qor/admin"
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

func main() {
	DB, _ := gorm.Open("sqlite3", "demo.db")
	DB.AutoMigrate(&User{}, &Product{})

	// Initalize
	Admin := admin.New(&admin.AdminConfig{DB: DB})
	//redir
	Admin.GetRouter().Get("/", func(c *admin.Context) {
		fmt.Fprint(c.Writer, "hello world~")
		//http.Redirect(c.Writer, c.Request, "/admin/products", http.StatusSeeOther)
	})

	// Allow to use Admin to manage User, Product
	Admin.AddResource(&User{})
	Admin.AddResource(&Product{})

	// initalize an HTTP request multiplexer
	mux := http.NewServeMux()
	mux.HandleFunc("/aa", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello~")

	})

	// Mount admin interface to mux
	Admin.MountTo("/admin", mux)

	fmt.Println("Listening on: 8081")
	http.ListenAndServe(":8081", mux)
}

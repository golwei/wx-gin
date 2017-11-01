package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/media/oss"
)

type Ad struct {
	gorm.Model
	AdPositionId uint16 `orm:"column(ad_position_id)"`
	MediaType    uint8  `orm:"column(media_type)"`
	Name         string `orm:"column(name);size(60)"`
	Link         string `orm:"column(link);size(255)"`
	ImageUrl     string `orm:"column(image_url)"`
	Content      string `orm:"column(content);size(255)"`
	Image        oss.OSS
}
type Channel struct {
	gorm.Model
	Name      string `orm:"column(name);size(45)"`
	Url       string `orm:"column(url);size(255)"`
	IconUrl   string `orm:"column(icon_url);size(255)"`
	SortOrder uint   `orm:"column(sort_order)"`
	Image     oss.OSS
}
type Brand struct {
	gorm.Model
	Name          string `orm:"column(name);size(255)"`
	ListPicUrl    string `orm:"column(list_pic_url);size(255)"`
	SimpleDesc    string `orm:"column(simple_desc);size(255)"`
	Pic           oss.OSS
	SortOrder     uint8   `orm:"column(sort_order)"`
	IsShow        uint8   `orm:"column(is_show)"`
	FloorPrice    float64 `orm:"column(floor_price);digits(10);decimals(2)"`
	AppListPicUrl string  `orm:"column(app_list_pic_url);size(255)"`
	IsNew         uint8   `orm:"column(is_new)"`
	NewPic        oss.OSS
	NewSortOrder  uint8 `orm:"column(new_sort_order)"`
}

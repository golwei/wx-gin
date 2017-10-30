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

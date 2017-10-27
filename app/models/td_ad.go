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

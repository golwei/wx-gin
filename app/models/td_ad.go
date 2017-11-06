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
type Goods struct {
	gorm.Model
	Image             oss.OSS
	CategoryId        uint    `orm:"column(category_id)"`
	GoodsSn           string  `orm:"column(goods_sn);size(60)"`
	Name              string  `orm:"column(name);size(120)"`
	BrandId           uint    `orm:"column(brand_id)"`
	GoodsNumber       uint32  `orm:"column(goods_number)"`
	Keywords          string  `orm:"column(keywords);size(255)"`
	GoodsBrief        string  `orm:"column(goods_brief);size(255)"`
	GoodsDesc         string  `orm:"column(goods_desc);null"`
	IsOnSale          uint8   `orm:"column(is_on_sale)"`
	AddTime           uint    `orm:"column(add_time)"`
	SortOrder         uint16  `orm:"column(sort_order)"`
	IsDelete          uint8   `orm:"column(is_delete)"`
	AttributeCategory uint    `orm:"column(attribute_category)"`
	CounterPrice      float64 `orm:"column(counter_price);digits(10);decimals(2)" description:"专柜价格"`
	ExtraPrice        float64 `orm:"column(extra_price);digits(10);decimals(2)" description:"附加价格"`
	IsNew             uint8   `orm:"column(is_new)"`
	GoodsUnit         string  `orm:"column(goods_unit);size(45)" description:"商品单位"`
	PrimaryPicUrl     string  `orm:"column(primary_pic_url);size(255)" description:"商品主图"`
	ListPicUrl        string  `orm:"column(list_pic_url);size(255)" description:"商品列表图"`
	RetailPrice       float64 `orm:"column(retail_price);digits(10);decimals(2)" description:"零售价格"`
	SellVolume        uint    `orm:"column(sell_volume)" description:"销售量"`
	PrimaryProductId  uint    `orm:"column(primary_product_id)" description:"主sku　product_id"`
	UnitPrice         float64 `orm:"column(unit_price);digits(10);decimals(2)" description:"单位价格，单价"`
	PromotionDesc     string  `orm:"column(promotion_desc);size(255)"`
	PromotionTag      string  `orm:"column(promotion_tag);size(45)"`
	AppExclusivePrice float64 `orm:"column(app_exclusive_price);digits(10);decimals(2)" description:"APP专享价"`
	IsAppExclusive    uint8   `orm:"column(is_app_exclusive)" description:"是否是APP专属"`
	IsLimited         uint8   `orm:"column(is_limited)"`
	IsHot             uint8   `orm:"column(is_hot)"`
}

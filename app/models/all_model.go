package models

type NideshopAd struct {
	Id           int    `orm:"column(id);auto"`
	AdPositionId uint16 `orm:"column(ad_position_id)"`
	MediaType    uint8  `orm:"column(media_type)"`
	Name         string `orm:"column(name);size(60)"`
	Link         string `orm:"column(link);size(255)"`
	ImageUrl     string `orm:"column(image_url)"`
	Content      string `orm:"column(content);size(255)"`
	EndTime      int    `orm:"column(end_time)"`
	Enabled      uint8  `orm:"column(enabled)"`
}
type NideshopAdPosition struct {
	Id     int    `orm:"column(id);auto"`
	Name   string `orm:"column(name);size(60)"`
	Width  uint16 `orm:"column(width)"`
	Height uint16 `orm:"column(height)"`
	Desc   string `orm:"column(desc);size(255)"`
}
type NideshopAddress struct {
	Id         int    `orm:"column(id);auto"`
	Name       string `orm:"column(name);size(50)"`
	UserId     uint32 `orm:"column(user_id)"`
	CountryId  int16  `orm:"column(country_id)"`
	ProvinceId int16  `orm:"column(province_id)"`
	CityId     int16  `orm:"column(city_id)"`
	DistrictId int16  `orm:"column(district_id)"`
	Address    string `orm:"column(address);size(120)"`
	Mobile     string `orm:"column(mobile);size(60)"`
	IsDefault  uint8  `orm:"column(is_default)"`
}
type NideshopAttribute struct {
	Id                  int    `orm:"column(id);auto"`
	AttributeCategoryId uint   `orm:"column(attribute_category_id)"`
	Name                string `orm:"column(name);size(60)"`
	InputType           uint8  `orm:"column(input_type)"`
	Values              string `orm:"column(values)"`
	SortOrder           uint8  `orm:"column(sort_order)"`
}
type NideshopAttributeCategory struct {
	Id      int    `orm:"column(id);pk"`
	Name    string `orm:"column(name);size(60)"`
	Enabled uint8  `orm:"column(enabled)"`
}
type NideshopBrand struct {
	Id            int     `orm:"column(id);auto"`
	Name          string  `orm:"column(name);size(255)"`
	ListPicUrl    string  `orm:"column(list_pic_url);size(255)"`
	SimpleDesc    string  `orm:"column(simple_desc);size(255)"`
	PicUrl        string  `orm:"column(pic_url);size(255)"`
	SortOrder     uint8   `orm:"column(sort_order)"`
	IsShow        uint8   `orm:"column(is_show)"`
	FloorPrice    float64 `orm:"column(floor_price);digits(10);decimals(2)"`
	AppListPicUrl string  `orm:"column(app_list_pic_url);size(255)"`
	IsNew         uint8   `orm:"column(is_new)"`
	NewPicUrl     string  `orm:"column(new_pic_url);size(255)"`
	NewSortOrder  uint8   `orm:"column(new_sort_order)"`
}
type NideshopCart struct {
	Id                        int     `orm:"column(id);auto"`
	UserId                    uint32  `orm:"column(user_id)"`
	SessionId                 string  `orm:"column(session_id);size(32)"`
	GoodsId                   uint32  `orm:"column(goods_id)"`
	GoodsSn                   string  `orm:"column(goods_sn);size(60)"`
	ProductId                 uint32  `orm:"column(product_id)"`
	GoodsName                 string  `orm:"column(goods_name);size(120)"`
	MarketPrice               float64 `orm:"column(market_price);digits(10);decimals(2)"`
	RetailPrice               float64 `orm:"column(retail_price);digits(10);decimals(2)"`
	Number                    uint16  `orm:"column(number)"`
	GoodsSpecifitionNameValue string  `orm:"column(goods_specifition_name_value)" description:"规格属性组成的字符串，用来显示用"`
	GoodsSpecifitionIds       string  `orm:"column(goods_specifition_ids);size(60)" description:"product表对应的goods_specifition_ids"`
	Checked                   uint8   `orm:"column(checked)"`
	ListPicUrl                string  `orm:"column(list_pic_url);size(255)"`
}
type NideshopCategory struct {
	Id           int    `orm:"column(id);pk"`
	Name         string `orm:"column(name);size(90)"`
	Keywords     string `orm:"column(keywords);size(255)"`
	FrontDesc    string `orm:"column(front_desc);size(255)"`
	ParentId     uint   `orm:"column(parent_id)"`
	SortOrder    uint8  `orm:"column(sort_order)"`
	ShowIndex    int8   `orm:"column(show_index)"`
	IsShow       uint8  `orm:"column(is_show)"`
	BannerUrl    string `orm:"column(banner_url);size(255)"`
	IconUrl      string `orm:"column(icon_url);size(255)"`
	ImgUrl       string `orm:"column(img_url);size(255)"`
	WapBannerUrl string `orm:"column(wap_banner_url);size(255)"`
	Level        string `orm:"column(level);size(255)"`
	Type         int    `orm:"column(type)"`
	FrontName    string `orm:"column(front_name);size(255)"`
}
type NideshopChannel struct {
	Id        int    `orm:"column(id);auto"`
	Name      string `orm:"column(name);size(45)"`
	Url       string `orm:"column(url);size(255)"`
	IconUrl   string `orm:"column(icon_url);size(255)"`
	SortOrder uint   `orm:"column(sort_order)"`
}
type NideshopCollect struct {
	Id          int    `orm:"column(id);auto"`
	UserId      uint32 `orm:"column(user_id)"`
	ValueId     uint32 `orm:"column(value_id)"`
	AddTime     uint   `orm:"column(add_time)"`
	IsAttention int8   `orm:"column(is_attention)" description:"是否是关注"`
	TypeId      uint   `orm:"column(type_id)"`
}
type NideshopComment struct {
	Id      int    `orm:"column(id);auto"`
	TypeId  uint8  `orm:"column(type_id)"`
	ValueId uint   `orm:"column(value_id)"`
	Content string `orm:"column(content);size(6550)" description:"储存为base64编码"`
	AddTime uint64 `orm:"column(add_time)"`
	Status  uint8  `orm:"column(status)"`
	UserId  uint   `orm:"column(user_id)"`
}
type NideshopCommentPicture struct {
	Id        int    `orm:"column(id);auto"`
	CommentId uint   `orm:"column(comment_id)"`
	PicUrl    string `orm:"column(pic_url);size(255)"`
	SortOrder uint8  `orm:"column(sort_order)"`
}
type NideshopCoupon struct {
	Id             int     `orm:"column(id);auto"`
	Name           string  `orm:"column(name);size(60)"`
	TypeMoney      float64 `orm:"column(type_money);digits(10);decimals(2)"`
	SendType       uint8   `orm:"column(send_type)"`
	MinAmount      float64 `orm:"column(min_amount);digits(10);decimals(2)"`
	MaxAmount      float64 `orm:"column(max_amount);digits(10);decimals(2)"`
	SendStartDate  int     `orm:"column(send_start_date)"`
	SendEndDate    int     `orm:"column(send_end_date)"`
	UseStartDate   int     `orm:"column(use_start_date)"`
	UseEndDate     int     `orm:"column(use_end_date)"`
	MinGoodsAmount float64 `orm:"column(min_goods_amount);digits(10);decimals(2)"`
}
type NideshopFeedback struct {
	Id         int    `orm:"column(msg_id);auto"`
	ParentId   uint32 `orm:"column(parent_id)"`
	UserId     uint32 `orm:"column(user_id)"`
	UserName   string `orm:"column(user_name);size(60)"`
	UserEmail  string `orm:"column(user_email);size(60)"`
	MsgTitle   string `orm:"column(msg_title);size(200)"`
	MsgType    uint8  `orm:"column(msg_type)"`
	MsgStatus  uint8  `orm:"column(msg_status)"`
	MsgContent string `orm:"column(msg_content)"`
	MsgTime    uint   `orm:"column(msg_time)"`
	MessageImg string `orm:"column(message_img);size(255)"`
	OrderId    uint   `orm:"column(order_id)"`
	MsgArea    uint8  `orm:"column(msg_area)"`
}
type NideshopFootprint struct {
	Id      int `orm:"column(id);auto"`
	UserId  int `orm:"column(user_id)"`
	GoodsId int `orm:"column(goods_id)"`
	AddTime int `orm:"column(add_time)"`
}
type NideshopGoods struct {
	Id                int     `orm:"column(id);pk"`
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
type NideshopGoodsAttribute struct {
	Id          int    `orm:"column(id);auto"`
	GoodsId     uint   `orm:"column(goods_id)"`
	AttributeId uint   `orm:"column(attribute_id)"`
	Value       string `orm:"column(value)"`
}
type NideshopGoodsGallery struct {
	Id        int    `orm:"column(id);auto"`
	GoodsId   uint   `orm:"column(goods_id)"`
	ImgUrl    string `orm:"column(img_url);size(255)"`
	ImgDesc   string `orm:"column(img_desc);size(255)"`
	SortOrder uint   `orm:"column(sort_order)"`
}
type NideshopGoodsIssue struct {
	Id       int    `orm:"column(id);auto"`
	GoodsId  string `orm:"column(goods_id);null"`
	Question string `orm:"column(question);size(255);null"`
	Answer   string `orm:"column(answer);size(45);null"`
}
type NideshopGoodsSpecification struct {
	Id              int    `orm:"column(id);auto"`
	GoodsId         uint   `orm:"column(goods_id)"`
	SpecificationId uint   `orm:"column(specification_id)"`
	Value           string `orm:"column(value);size(50)"`
	PicUrl          string `orm:"column(pic_url);size(255)"`
}
type NideshopKeywords struct {
	Keyword   string `orm:"column(keyword);size(90)"`
	IsHot     uint8  `orm:"column(is_hot)"`
	IsDefault uint8  `orm:"column(is_default)"`
	IsShow    uint8  `orm:"column(is_show)"`
	SortOrder uint   `orm:"column(sort_order)"`
	SchemeUrl string `orm:"column(scheme _url);size(255)" description:"关键词的跳转链接"`
	Id_RENAME int    `orm:"column(id)"`
	Type      uint   `orm:"column(type)"`
}
type NideshopOrder struct {
	Id             int     `orm:"column(id);auto"`
	OrderSn        string  `orm:"column(order_sn);size(20)"`
	UserId         uint32  `orm:"column(user_id)"`
	OrderStatus    uint8   `orm:"column(order_status)"`
	ShippingStatus uint8   `orm:"column(shipping_status)"`
	PayStatus      uint8   `orm:"column(pay_status)"`
	Consignee      string  `orm:"column(consignee);size(60)"`
	Country        uint16  `orm:"column(country)"`
	Province       uint16  `orm:"column(province)"`
	City           uint16  `orm:"column(city)"`
	District       uint16  `orm:"column(district)"`
	Address        string  `orm:"column(address);size(255)"`
	Mobile         string  `orm:"column(mobile);size(60)"`
	Postscript     string  `orm:"column(postscript);size(255)"`
	ShippingId     int8    `orm:"column(shipping_id)"`
	ShippingName   string  `orm:"column(shipping_name);size(120)"`
	PayId          int8    `orm:"column(pay_id)"`
	PayName        string  `orm:"column(pay_name);size(120)"`
	ShippingFee    float64 `orm:"column(shipping_fee);digits(10);decimals(2)"`
	ActualPrice    float64 `orm:"column(actual_price);digits(10);decimals(2)" description:"实际需要支付的金额"`
	Integral       uint    `orm:"column(integral)"`
	IntegralMoney  float64 `orm:"column(integral_money);digits(10);decimals(2)"`
	OrderPrice     float64 `orm:"column(order_price);digits(10);decimals(2)" description:"订单总价"`
	GoodsPrice     float64 `orm:"column(goods_price);digits(10);decimals(2)" description:"商品总价"`
	AddTime        uint    `orm:"column(add_time)"`
	ConfirmTime    uint    `orm:"column(confirm_time)"`
	PayTime        uint    `orm:"column(pay_time)"`
	FreightPrice   uint    `orm:"column(freight_price)" description:"配送费用"`
	CouponId       uint32  `orm:"column(coupon_id)" description:"使用的优惠券id"`
	ParentId       uint32  `orm:"column(parent_id)"`
	CouponPrice    float64 `orm:"column(coupon_price);digits(10);decimals(2)"`
	CallbackStatus string  `orm:"column(callback_status);null"`
}
type NideshopOrderGoods struct {
	Id                        int     `orm:"column(id);auto"`
	OrderId                   uint32  `orm:"column(order_id)"`
	GoodsId                   uint32  `orm:"column(goods_id)"`
	GoodsName                 string  `orm:"column(goods_name);size(120)"`
	GoodsSn                   string  `orm:"column(goods_sn);size(60)"`
	ProductId                 uint32  `orm:"column(product_id)"`
	Number                    uint16  `orm:"column(number)"`
	MarketPrice               float64 `orm:"column(market_price);digits(10);decimals(2)"`
	RetailPrice               float64 `orm:"column(retail_price);digits(10);decimals(2)"`
	GoodsSpecifitionNameValue string  `orm:"column(goods_specifition_name_value)"`
	IsReal                    uint8   `orm:"column(is_real)"`
	GoodsSpecifitionIds       string  `orm:"column(goods_specifition_ids);size(255)"`
	ListPicUrl                string  `orm:"column(list_pic_url);size(255)"`
}
type NideshopProduct struct {
	Id                    int     `orm:"column(id);auto"`
	GoodsId               uint32  `orm:"column(goods_id)"`
	GoodsSpecificationIds string  `orm:"column(goods_specification_ids);size(50)"`
	GoodsSn               string  `orm:"column(goods_sn);size(60)"`
	GoodsNumber           uint32  `orm:"column(goods_number)"`
	RetailPrice           float64 `orm:"column(retail_price);digits(10);decimals(2)"`
}
type NideshopRegion struct {
	Id       int    `orm:"column(id);auto"`
	ParentId uint16 `orm:"column(parent_id)"`
	Name     string `orm:"column(name);size(120)"`
	Type     int8   `orm:"column(type)"`
	AgencyId uint16 `orm:"column(agency_id)"`
}
type NideshopRelatedGoods struct {
	Id             int  `orm:"column(id);auto"`
	GoodsId        uint `orm:"column(goods_id)"`
	RelatedGoodsId uint `orm:"column(related_goods_id)"`
}
type NideshopSearchHistory struct {
	Id      int    `orm:"column(id);auto"`
	Keyword string `orm:"column(keyword);size(50)"`
	From    string `orm:"column(from);size(45)" description:"搜索来源，如PC、小程序、APP等"`
	AddTime int    `orm:"column(add_time)" description:"搜索时间"`
	UserId  string `orm:"column(user_id);size(45);null"`
}
type NideshopSpecification struct {
	Id        int    `orm:"column(id);auto"`
	Name      string `orm:"column(name);size(60)"`
	SortOrder uint8  `orm:"column(sort_order)"`
}
type NideshopTopic struct {
	Id_RENAME       int     `orm:"column(id)"`
	Title           string  `orm:"column(title);size(255)"`
	Content         string  `orm:"column(content);null"`
	Avatar          string  `orm:"column(avatar);size(255)"`
	ItemPicUrl      string  `orm:"column(item_pic_url);size(255)"`
	Subtitle        string  `orm:"column(subtitle);size(255)"`
	TopicCategoryId uint    `orm:"column(topic_category_id)"`
	PriceInfo       float64 `orm:"column(price_info);digits(10);decimals(2)"`
	ReadCount       string  `orm:"column(read_count);size(255)"`
	ScenePicUrl     string  `orm:"column(scene_pic_url);size(255)"`
	TopicTemplateId uint    `orm:"column(topic_template_id)"`
	TopicTagId      uint    `orm:"column(topic_tag_id)"`
}
type NideshopTopicCategory struct {
	Id     int    `orm:"column(id);auto"`
	Title  string `orm:"column(title);size(255)"`
	PicUrl string `orm:"column(pic_url);size(255)"`
}
type NideshopUser struct {
	Id            int    `orm:"column(id);auto"`
	Username      string `orm:"column(username);size(60)"`
	Password      string `orm:"column(password);size(32)"`
	Gender        uint8  `orm:"column(gender)"`
	Birthday      uint   `orm:"column(birthday)"`
	RegisterTime  uint   `orm:"column(register_time)"`
	LastLoginTime uint   `orm:"column(last_login_time)"`
	LastLoginIp   string `orm:"column(last_login_ip);size(15)"`
	UserLevelId   uint8  `orm:"column(user_level_id)"`
	Nickname      string `orm:"column(nickname);size(60)"`
	Mobile        string `orm:"column(mobile);size(20)"`
	RegisterIp    string `orm:"column(register_ip);size(45)"`
	Avatar        string `orm:"column(avatar);size(255)"`
	WeixinOpenid  string `orm:"column(weixin_openid);size(50)"`
}
type NideshopUserCoupon struct {
	Id           int    `orm:"column(id);auto"`
	CouponId     uint8  `orm:"column(coupon_id)"`
	CouponNumber string `orm:"column(coupon_number);size(20)"`
	UserId       uint   `orm:"column(user_id)"`
	UsedTime     uint   `orm:"column(used_time)"`
	OrderId      uint32 `orm:"column(order_id)"`
}
type NideshopUserLevel struct {
	Id          int    `orm:"column(id);auto"`
	Name        string `orm:"column(name);size(30)"`
	Description string `orm:"column(description);size(255)"`
}

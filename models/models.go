package models

// 输出表名;
const (
	TableNameAdminUserExtend           = "zby_admin_user_extend"            // 员工扩展表
	TableNameComponentAppid            = "zby_component_appid"              //
	TableNameCoupon                    = "zby_coupon"                       //
	TableNameCustomer                  = "zby_customer"                     //
	TableNameCustomerBpConfig          = "zby_customer_bp_config"           //
	TableNameCustomerBpDetail          = "zby_customer_bp_detail"           //
	TableNameCustomerBpLog             = "zby_customer_bp_log"              //
	TableNameCustomerLevelConfig       = "zby_customer_level_config"        //
	TableNameCustomerLog               = "zby_customer_log"                 //
	TableNameCustomerModifyTrace       = "zby_customer_modify_trace"        // 会员重要信息修改记录表, 记录会员的姓名、手机号、生日、结婚纪念日的修改，用来限制修改次数
	TableNameCustomerRule              = "zby_customer_rule"                // 客户规则配置表
	TableNameCustomerSearchLog         = "zby_customer_search_log"          // 客户搜索日志
	TableNameCustomerTag               = "zby_customer_tag"                 // 客户标签
	TableNameCustomerTagRel            = "zby_customer_tag_rel"             // 客户关联的标签
	TableNameCustomerUserJoin          = "zby_customer_user_join"           //
	TableNameExclusiveServerChatRecord = "zby_exclusive_server_chat_record" //
	TableNameExclusiveServerLog        = "zby_exclusive_server_log"         //
	TableNameExclusiveServerRelation   = "zby_exclusive_server_relation"    //
	TableNameGoldSalePriceConfig       = "zby_gold_sale_price_config"       // 金价配制信息表
	TableNameMakeTagLog                = "zby_make_tag_log"                 // 记录打的客户标签日志
	TableNameMerchantSetting           = "zby_merchant_setting"             //
	TableNameMerchantShop              = "zby_merchant_shop"                //
	TableNameMerchantStaff             = "zby_merchant_staff"               //
	TableNameOnlineGoldPrice           = "zby_online_gold_price"            //
	TableNameOnlineShowcaseSetting     = "zby_online_showcase_setting"      // 在线橱窗配置表
	TableNameOrderCheck                = "zby_order_check"                  // 单据审核流程表
	TableNameQyCustomer                = "zby_qy_customer"                  // 企业微信,外部联系人
	TableNameQyTag                     = "zby_qy_tag"                       // 企业微信标签
	TableNameRefundOrderGoods          = "zby_refund_order_goods"           // 退货订单商品表
	TableNameRoleTemplateField         = "zby_role_template_field"          // 角色用户功能字段配置表
	TableNameSaleOrder                 = "zby_sale_order"                   // 销售单据表
	TableNameSaleOrderGoods            = "zby_sale_order_goods"             // 销售单据商品明细表
	TableNameSaleOrderHistory          = "zby_sale_order_history"           // 历史销售单据表
	TableNameSaleOrderNewExchange      = "zby_sale_order_new_exchange"      // 新品换购订单信息
	TableNameSaleOrderOldExchange      = "zby_sale_order_old_exchange"      // 销售订单旧换新商品表
	TableNameSaleOrderPrize            = "zby_sale_order_prize"             // 销售单据赠品明细表
)

// 员工扩展表
type AdminUserExtend struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/* 商户ID */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 门店ID */
	MerchantShopId int32 `gorm:"column:merchant_shop_id;type:int(11);" json:"merchant_shop_id"`
	/* 员工ID */
	AdminUserId int32 `gorm:"column:admin_user_id;type:int(11);" json:"admin_user_id"`
	/* 在线橱窗名称 */
	OnlineShowcaseName string `gorm:"column:online_showcase_name;type:varchar(64);" json:"online_showcase_name"`
	/* 创建时间 */
	CreateDatetime string `gorm:"column:create_datetime;type:datetime;" json:"create_datetime"`
	/* 更新时间 */
	UpdateDatetime string `gorm:"column:update_datetime;type:datetime;" json:"update_datetime"`
}

type ComponentAppid struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/* 商户ID */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 品牌ID */
	BrandId int32 `gorm:"column:brand_id;type:int(11);" json:"brand_id"`
	/* 门店ID */
	MerchantShopId int32 `gorm:"column:merchant_shop_id;type:int(11);" json:"merchant_shop_id"`
	/*  */
	AuthorizerAppid string `gorm:"column:authorizer_appid;type:varchar(200);" json:"authorizer_appid"`
	/*  */
	AuthorizerAccessToken string `gorm:"column:authorizer_access_token;type:text;" json:"authorizer_access_token"`
	/*  */
	ExpiresIn int32 `gorm:"column:expires_in;type:int(11);" json:"expires_in"`
	/*  */
	AuthorizerRefreshToken string `gorm:"column:authorizer_refresh_token;type:text;" json:"authorizer_refresh_token"`
	/*  */
	TicketJs string `gorm:"column:ticket_js;type:text;" json:"ticket_js"`
	/*  */
	TicketJsExpiresIn int32 `gorm:"column:ticket_js_expires_in;type:int(11);" json:"ticket_js_expires_in"`
	/*  */
	TicketCard string `gorm:"column:ticket_card;type:text;" json:"ticket_card"`
	/*  */
	TicketCardExpiresIn int32 `gorm:"column:ticket_card_expires_in;type:int(11);" json:"ticket_card_expires_in"`
	/*  */
	FuncInfo string `gorm:"column:func_info;type:text;" json:"func_info"`
	/*  */
	NickName string `gorm:"column:nick_name;type:varchar(200);" json:"nick_name"`
	/* 头像 */
	HeadImg string `gorm:"column:head_img;type:text;" json:"head_img"`
	/*  */
	ServiceTypeInfo string `gorm:"column:service_type_info;type:text;" json:"service_type_info"`
	/*  */
	VerifyTypeInfo string `gorm:"column:verify_type_info;type:text;" json:"verify_type_info"`
	/*  */
	UserName string `gorm:"column:user_name;type:varchar(200);" json:"user_name"`
	/*  */
	Alias string `gorm:"column:alias;type:varchar(200);" json:"alias"`
	/*  */
	Idc int32 `gorm:"column:idc;type:tinyint(4);" json:"idc"`
	/*  */
	BusinessInfo string `gorm:"column:business_info;type:text;" json:"business_info"`
	/*  */
	QrcodeUrl string `gorm:"column:qrcode_url;type:text;" json:"qrcode_url"`
	/*  */
	QrcodeLocalUrl string `gorm:"column:qrcode_local_url;type:text;" json:"qrcode_local_url"`
	/*  */
	QrcodeLocalUrlBak string `gorm:"column:qrcode_local_url_bak;type:varchar(512);" json:"qrcode_local_url_bak"`
	/*  */
	CardImg string `gorm:"column:card_img;type:text;" json:"card_img"`
	/* 是否授权 */
	Status int32 `gorm:"column:status;type:tinyint(4);" json:"status"`
	/*  */
	Statuschangetime int32 `gorm:"column:statuschangetime;type:int(11);" json:"statuschangetime"`
	/* json格式 涉及到的收费项目 */
	AuthInfo string `gorm:"column:auth_info;type:text;" json:"auth_info"`
	/* 授权类型，1:公众号，2:小程序 */
	AuthType int32 `gorm:"column:auth_type;type:tinyint(1);" json:"auth_type"`
	/*  */
	CreateTime int32 `gorm:"column:create_time;type:int(11);" json:"create_time"`
	/*  */
	UpdateTime int32 `gorm:"column:update_time;type:int(11);" json:"update_time"`
	/*  */
	DeleteTime int32 `gorm:"column:delete_time;type:int(11);" json:"delete_time"`
	/* 过期时间 */
	ExpireTime int32 `gorm:"column:expire_time;type:int(11);" json:"expire_time"`
	/* 小程序审核ID */
	Auditid int64 `gorm:"column:auditid;type:bigint(20);" json:"auditid"`
	/* 小程序当前的模版ID */
	TemplateId int32 `gorm:"column:template_id;type:int(11);" json:"template_id"`
}

type Coupon struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(10) unsigned;" json:"id"`
	/* 商户id */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 品牌id */
	BrandId int32 `gorm:"column:brand_id;type:int(11);" json:"brand_id"`
	/* 品牌名 */
	BrandName string `gorm:"column:brand_name;type:varchar(255);" json:"brand_name"`
	/* 门店id（如果为空则代表品牌下全部门店） */
	MerchantShopId string `gorm:"column:merchant_shop_id;type:varchar(50);" json:"merchant_shop_id"`
	/* 之前删除过的门店记录 */
	DelMerchantShopId string `gorm:"column:del_merchant_shop_id;type:varchar(50);" json:"del_merchant_shop_id"`
	/* 券名 */
	Name string `gorm:"column:name;type:varchar(50);" json:"name"`
	/* 优惠券券类型 1-通用券 2-人情券 3-生日券 4-连单券 */
	Type int32 `gorm:"column:type;type:tinyint(3) unsigned;" json:"type"`
	/* 发放节点 1-消费赠送 2-生日赠送 3-精准投放 4-会员卡 5-在线导购 6-扫码赠送  7-活动送券 */
	RuleType int32 `gorm:"column:rule_type;type:tinyint(1) unsigned zerofill;" json:"rule_type"`
	/* 发放日期开始时间 */
	StartTime string `gorm:"column:start_time;type:timestamp;" json:"start_time"`
	/* 发放日期结束时间 */
	EndTime string `gorm:"column:end_time;type:timestamp;" json:"end_time"`
	/* 优惠券数量 */
	TotalNum int32 `gorm:"column:total_num;type:int(10) unsigned;" json:"total_num"`
	/* 券发送数量 */
	SendNum int32 `gorm:"column:send_num;type:int(10) unsigned;" json:"send_num"`
	/* 参与商品分类的id串 */
	GoodsTypeIdString string `gorm:"column:goods_type_id_string;type:varchar(254);" json:"goods_type_id_string"`
	/* 扫码券的码 */
	QrcodeUrl string `gorm:"column:qrcode_url;type:varchar(255);" json:"qrcode_url"`
	/* 生日券提前几天赠送 */
	AfterDay int32 `gorm:"column:after_day;type:int(11);" json:"after_day"`
	/* 优惠方式 1-固定金额 2-随机金额 3-满减阶梯 4-累计 */
	DiscountType int32 `gorm:"column:discount_type;type:tinyint(4);" json:"discount_type"`
	/* 计算类型 1:标价券,2:克重券 */
	DiscountCalType int32 `gorm:"column:discount_cal_type;type:tinyint(4);" json:"discount_cal_type"`
	/* 优惠内容 */
	DiscountContent string `gorm:"column:discount_content;type:varchar(255);" json:"discount_content"`
	/* 会员等级权益ID */
	CustomerWelfareId int32 `gorm:"column:customer_welfare_id;type:int(11);" json:"customer_welfare_id"`
	/* 优惠券JSON */
	Content string `gorm:"column:content;type:json;" json:"content"`
	/* 创建人id */
	AdminUserId int32 `gorm:"column:admin_user_id;type:int(10) unsigned;" json:"admin_user_id"`
	/* 创建人名 */
	AdminUserName string `gorm:"column:admin_user_name;type:varchar(50);" json:"admin_user_name"`
	/* 状态 1-正常 2-作废(逻辑删除) 3-停用 4-已过期 */
	Status int32 `gorm:"column:status;type:tinyint(3) unsigned;" json:"status"`
	/* 说明 */
	Description string `gorm:"column:description;type:text;" json:"description"`
	/* 作废人id */
	CancelAdminUserId int32 `gorm:"column:cancel_admin_user_id;type:int(11);" json:"cancel_admin_user_id"`
	/* 作废人名 */
	CancelAdminUserName string `gorm:"column:cancel_admin_user_name;type:varchar(50);" json:"cancel_admin_user_name"`
	/* 作废原因 */
	Remark string `gorm:"column:remark;type:varchar(255);" json:"remark"`
	/* 分享图片 */
	ShareImg string `gorm:"column:share_img;type:varchar(255);" json:"share_img"`
	/* 分享描述 */
	ShareMark string `gorm:"column:share_mark;type:varchar(255);" json:"share_mark"`
	/* 分享标题 */
	ShareTitle string `gorm:"column:share_title;type:varchar(255);" json:"share_title"`
	/* 创建时间 */
	CreateTime string `gorm:"column:create_time;type:timestamp;" json:"create_time"`
	/* 更新时间 */
	UpdateTime string `gorm:"column:update_time;type:timestamp;" json:"update_time"`
	/* 删除时间 */
	DeleteTime string `gorm:"column:delete_time;type:timestamp;" json:"delete_time"`
	/* 卡券来源 0:珠宝益 1:老庙crm */
	CouponSource int32 `gorm:"column:coupon_source;type:int(11);" json:"coupon_source"`
	/* 卡券编码 */
	CouponCode string `gorm:"column:coupon_code;type:varchar(50);" json:"coupon_code"`
}

type Customer struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/*  */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 品牌id */
	BrandId int32 `gorm:"column:brand_id;type:int(10) unsigned;" json:"brand_id"`
	/* 商户下门店id */
	MerchantShopId int32 `gorm:"column:merchant_shop_id;type:int(11);" json:"merchant_shop_id"`
	/* 公众号user_id */
	UserId int32 `gorm:"column:user_id;type:int(11);" json:"user_id"`
	/* 小程序user_id */
	MiniProgramUserId int32 `gorm:"column:mini_program_user_id;type:int(11);" json:"mini_program_user_id"`
	/* 推荐人id */
	UnionCustomerId int32 `gorm:"column:union_customer_id;type:int(11);" json:"union_customer_id"`
	/* 联客姓名 */
	UnionCustomerName string `gorm:"column:union_customer_name;type:varchar(50);" json:"union_customer_name"`
	/* 推荐人电话 */
	UnionCustomerTelephone string `gorm:"column:union_customer_telephone;type:varchar(15);" json:"union_customer_telephone"`
	/* 微信Openid */
	Openid string `gorm:"column:openid;type:varchar(100);" json:"openid"`
	/* 开放平台unionid */
	Unionid string `gorm:"column:unionid;type:varchar(100);" json:"unionid"`
	/* 客户姓名 */
	Name string `gorm:"column:name;type:varchar(100);" json:"name"`
	/* 卡券编号 */
	CardCode string `gorm:"column:card_code;type:varchar(48);" json:"card_code"`
	/* 电话 */
	Telephone string `gorm:"column:telephone;type:varchar(15);" json:"telephone"`
	/* 手机后缀4位 */
	SuffixTelephone string `gorm:"column:suffix_telephone;type:varchar(4);" json:"suffix_telephone"`
	/*  */
	BirthdayDt string `gorm:"column:birthday_dt;type:date;" json:"birthday_dt"`
	/*  */
	Birthday int32 `gorm:"column:birthday;type:int(11);" json:"birthday"`
	/*  */
	Password string `gorm:"column:password;type:varchar(255);" json:"password"`
	/*  */
	WechatId string `gorm:"column:wechat_id;type:varchar(200);" json:"wechat_id"`
	/* 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知 */
	Sex int32 `gorm:"column:sex;type:tinyint(4);" json:"sex"`
	/* 枚举: 1-微信 2-手工 5-企业微信同步 */
	Source int32 `gorm:"column:source;type:tinyint(4);" json:"source"`
	/* 状态 （-1禁用，1正常）2-逻辑删除 */
	Status int32 `gorm:"column:status;type:tinyint(4);" json:"status"`
	/* 会员真实头像 */
	Headimgurl string `gorm:"column:headimgurl;type:varchar(255);" json:"headimgurl"`
	/* 是否订阅公众号 1: 关注 2: 已取关 0: 未关注 */
	Subscribe int32 `gorm:"column:subscribe;type:tinyint(1);" json:"subscribe"`
	/* 微信昵称 */
	Nickname string `gorm:"column:nickname;type:varchar(200);" json:"nickname"`
	/* 最后消费时间 */
	LastBuyTime string `gorm:"column:last_buy_time;type:datetime;" json:"last_buy_time"`
	/* 客单价 单位分 */
	UnitOrderAmount int32 `gorm:"column:unit_order_amount;type:int(11);" json:"unit_order_amount"`
	/* 消费次数 */
	TotalOrderCnt int32 `gorm:"column:total_order_cnt;type:int(11);" json:"total_order_cnt"`
	/* 累计消费金额 单位分 */
	TotalOrderAmount int32 `gorm:"column:total_order_amount;type:int(11);" json:"total_order_amount"`
	/* 累计会员等级积分总额 */
	TotalBonus float64 `gorm:"column:total_bonus;type:decimal(20,2);" json:"total_bonus"`
	/* 积分 */
	PayBonus float64 `gorm:"column:pay_bonus;type:decimal(20,2);" json:"pay_bonus"`
	/* 迁移之前的积分备份 */
	PayBonusOld float64 `gorm:"column:pay_bonus_old;type:decimal(11,2);" json:"pay_bonus_old"`
	/* 省 */
	Province string `gorm:"column:province;type:varchar(100);" json:"province"`
	/* 市 */
	City string `gorm:"column:city;type:varchar(100);" json:"city"`
	/* 区/县 */
	District string `gorm:"column:district;type:varchar(100);" json:"district"`
	/* 详细地址 */
	Address string `gorm:"column:address;type:varchar(255);" json:"address"`
	/* 备注 */
	Comment string `gorm:"column:comment;type:text;" json:"comment"`
	/* 1-lv1,2-lv2 以此类推 */
	Level int32 `gorm:"column:level;type:tinyint(4);" json:"level"`
	/* 结婚纪念日 */
	MarryDate string `gorm:"column:marry_date;type:date;" json:"marry_date"`
	/* 注册时间 */
	AddDatetime string `gorm:"column:add_datetime;type:datetime;" json:"add_datetime"`
	/* 创建时间 */
	CreateTime int32 `gorm:"column:create_time;type:int(11);" json:"create_time"`
	/* 修改时间 */
	UpdateTime int32 `gorm:"column:update_time;type:int(11);" json:"update_time"`
	/* 删除时间 */
	DeleteTime int32 `gorm:"column:delete_time;type:int(11);" json:"delete_time"`
	/*  */
	ReferrerId int32 `gorm:"column:referrer_id;type:int(11);" json:"referrer_id"`
	/*  */
	ReferrerName string `gorm:"column:referrer_name;type:varchar(255);" json:"referrer_name"`
	/*  */
	ReferrerPhone string `gorm:"column:referrer_phone;type:varchar(255);" json:"referrer_phone"`
	/*  */
	AdminIdOld int32 `gorm:"column:admin_id_old;type:int(11);" json:"admin_id_old"`
	/* 会员归属员工id */
	AdminId int32 `gorm:"column:admin_id;type:int(11);" json:"admin_id"`
	/* 会员归属员工姓名 */
	AdminName string `gorm:"column:admin_name;type:varchar(25);" json:"admin_name"`
	/* 收益金总金额 单位分 */
	EarningAmount int32 `gorm:"column:earning_amount;type:int(11);" json:"earning_amount"`
	/* 最后消费门店 */
	LastSaleMerchantShopId int32 `gorm:"column:last_sale_merchant_shop_id;type:int(11);" json:"last_sale_merchant_shop_id"`
	/* 绑卡时间 */
	BindCardTime int32 `gorm:"column:bind_card_time;type:int(11);" json:"bind_card_time"`
	/* 首次消费时间 */
	FirstConsumptionTime int32 `gorm:"column:first_consumption_time;type:int(11);" json:"first_consumption_time"`
	/* 会员相册 */
	Albums string `gorm:"column:albums;type:text;" json:"albums"`
	/* 是否农历生日 */
	IsBirthdayLunar int32 `gorm:"column:is_birthday_lunar;type:tinyint(4);" json:"is_birthday_lunar"`
	/* 是否农历结婚纪念日 */
	IsMarryDateLunar int32 `gorm:"column:is_marry_date_lunar;type:tinyint(4);" json:"is_marry_date_lunar"`
	/* 生日(农历) */
	BirthdayLunar string `gorm:"column:birthday_lunar;type:varchar(20);" json:"birthday_lunar"`
	/* 结婚纪念日(农历) */
	MarryDateLunar string `gorm:"column:marry_date_lunar;type:varchar(20);" json:"marry_date_lunar"`
	/* 来源平台 1抖音  2小红书  3微信视频号 4其他 */
	OriginPlatform int32 `gorm:"column:origin_platform;type:tinyint(2);" json:"origin_platform"`
	/* 来源渠道 1直播间  2视频 3线下  4其他 */
	OriginChannel int32 `gorm:"column:origin_channel;type:tinyint(2);" json:"origin_channel"`
}

type CustomerBpConfig struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/*  */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 品牌id */
	BrandId int32 `gorm:"column:brand_id;type:int(11);" json:"brand_id"`
	/* 门店id */
	MerchantShopId int32 `gorm:"column:merchant_shop_id;type:int(11);" json:"merchant_shop_id"`
	/* 0-不计分 1-所有相同积分 2-不同品类分开积分 */
	RuleType int32 `gorm:"column:rule_type;type:tinyint(1);" json:"rule_type"`
	/* 所有品类积分规则下  每消费一次积分数 */
	PointPerOrder float64 `gorm:"column:point_per_order;type:float;" json:"point_per_order"`
	/* 所有品类积分规则下  每消费一元积分数 */
	PointPerYuan float64 `gorm:"column:point_per_yuan;type:float;" json:"point_per_yuan"`
	/* [["商品分类,多个用逗号分割 id", "商品品牌,多个逗号分割 id", "商品标记,多个逗号分割 name", "每消费一元计分", "每消费一件计分", "参与促销活动不发积分", "参与折扣活动不发积分", "低于此折扣率不发积分"]]. */
	RuleJson string `gorm:"column:rule_json;type:json;" json:"rule_json"`
	/*  */
	RuleJsonOld string `gorm:"column:rule_json_old;type:json;" json:"rule_json_old"`
	/* 推荐人获得积分开关 1-开 0-关 */
	ReferrerPointIsOpen int32 `gorm:"column:referrer_point_is_open;type:tinyint(3) unsigned;" json:"referrer_point_is_open"`
	/* 推荐人获得积分倍数 */
	ReferrerMultiple float64 `gorm:"column:referrer_multiple;type:float unsigned;" json:"referrer_multiple"`
	/* 积分清零类型：  0：永不清零  1：年底清零  2：定义时间 */
	ClearMode int32 `gorm:"column:clear_mode;type:tinyint(4);" json:"clear_mode"`
	/* 积分清零的值 */
	ClearValue int32 `gorm:"column:clear_value;type:int(11);" json:"clear_value"`
	/* 积分清零时间 */
	ClearDatetime string `gorm:"column:clear_datetime;type:datetime;" json:"clear_datetime"`
	/* 积分抵现设置 */
	ExchangeCashConfig string `gorm:"column:exchange_cash_config;type:json;" json:"exchange_cash_config"`
	/*  */
	CreateTime int32 `gorm:"column:create_time;type:int(11);" json:"create_time"`
	/*  */
	UpdateTime int32 `gorm:"column:update_time;type:int(11);" json:"update_time"`
	/* 规则名称 */
	Name string `gorm:"column:name;type:varchar(255);" json:"name"`
	/* 状态 1 启用 0 禁止 3删除 */
	Status int32 `gorm:"column:status;type:tinyint(1);" json:"status"`
	/* 操作人id */
	OperatorId int32 `gorm:"column:operator_id;type:int(11);" json:"operator_id"`
	/* 操作人姓名 */
	OperatorName string `gorm:"column:operator_name;type:varchar(30);" json:"operator_name"`
	/* 多门店ids */
	MerchantShopIds string `gorm:"column:merchant_shop_ids;type:varchar(255);" json:"merchant_shop_ids"`
	/* 订单互斥策略 */
	OrderDetailPolicy string `gorm:"column:order_detail_policy;type:varchar(255);" json:"order_detail_policy"`
	/*  */
	OrderValue float64 `gorm:"column:order_value;type:float;" json:"order_value"`
}

type CustomerBpDetail struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/* 商户id */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 品牌id */
	BrandId int32 `gorm:"column:brand_id;type:int(11);" json:"brand_id"`
	/* 门店id */
	MerchantShopId int32 `gorm:"column:merchant_shop_id;type:int(11);" json:"merchant_shop_id"`
	/* 会员id */
	CustomerId int32 `gorm:"column:customer_id;type:int(11);" json:"customer_id"`
	/*  */
	CustomerName string `gorm:"column:customer_name;type:varchar(255);" json:"customer_name"`
	/*  */
	Content string `gorm:"column:content;type:json;" json:"content"`
	/* 发放积分 */
	Point float64 `gorm:"column:point;type:decimal(11,2);" json:"point"`
	/* 可用积分 */
	PayBonus float64 `gorm:"column:pay_bonus;type:decimal(11,2);" json:"pay_bonus"`
	/* 获得类型或者来源
	1-销售单 （收银端消费）+
	2-手动录入（手动调整）+
	3-连客积分 （转介绍顾客消费）+
	4-积分兑换 -
	5-激活微信会员卡+
	6-会员导入+
	7-积分抵现-
	8-取消兑换+
	9-退货-赠送积分收回-
	10-删单-赠送积分回收-
	11-新品换购-赠送积分收回
	12-积分清零
	13-退货-抵现积分退回
	14-删单-抵现积分退回
	15-新品换购-抵现积分退回
	16-盲盒抽奖 21.老庙crm差值 */
	Type int32 `gorm:"column:type;type:tinyint(4);" json:"type"`
	/* 0 永不清零 1 年底清零 2 自定义清零 3 周期清零 */
	ClearModel int32 `gorm:"column:clear_model;type:tinyint(4);" json:"clear_model"`
	/* 状态 0-未清零 1已清零 */
	Status int32 `gorm:"column:status;type:tinyint(4);" json:"status"`
	/* 过期时间 */
	ExpireTime int32 `gorm:"column:expire_time;type:int(11);" json:"expire_time"`
	/* 更新时间 */
	UpdateTime int32 `gorm:"column:update_time;type:int(11);" json:"update_time"`
	/* 创建时间 */
	CreateTime int32 `gorm:"column:create_time;type:int(11);" json:"create_time"`
}

type CustomerBpLog struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/*  */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 品牌id */
	BrandId int32 `gorm:"column:brand_id;type:int(11);" json:"brand_id"`
	/*  */
	MerchantShopId int32 `gorm:"column:merchant_shop_id;type:int(11);" json:"merchant_shop_id"`
	/*  */
	CustomerId int32 `gorm:"column:customer_id;type:int(11);" json:"customer_id"`
	/*  */
	CustomerName string `gorm:"column:customer_name;type:varchar(128);" json:"customer_name"`
	/* 本地获得积分 */
	Point float64 `gorm:"column:point;type:decimal(12,2);" json:"point"`
	/* 获得总积分 */
	PointTotal float64 `gorm:"column:point_total;type:decimal(12,2);" json:"point_total"`
	/* 获得类型或者来源
	1-销售单 （收银端消费）+
	2-手动录入（手动调整）+
	3-连客积分 （转介绍顾客消费）+
	4-积分兑换 -
	5-激活微信会员卡+
	6-会员导入+
	7-积分抵现-
	8-取消兑换+
	9-退货-赠送积分收回-
	10-删单-赠送积分回收-
	11-新品换购-赠送积分收回
	12-积分清零
	13-退货-抵现积分退回
	14-删单-抵现积分退回
	15-新品换购-抵现积分退回
	16-盲盒抽奖 */
	Type int32 `gorm:"column:type;type:int(11);" json:"type"`
	/* 表示 该记录中的会员是否变更 （变更需要作废  更改订单信息 会导致出现多条数据导致消息多发）0 正常 1 已变更 */
	IsDel int32 `gorm:"column:is_del;type:tinyint(1);" json:"is_del"`
	/* 备注 */
	Comment string `gorm:"column:comment;type:varchar(256);" json:"comment"`
	/* 订单类型 */
	OrderType int32 `gorm:"column:order_type;type:tinyint(4);" json:"order_type"`
	/* 订单ID */
	OrderId int32 `gorm:"column:order_id;type:int(11);" json:"order_id"`
	/* 订单SN */
	OrderSn string `gorm:"column:order_sn;type:varchar(255);" json:"order_sn"`
	/* 订单商品ID */
	OrderGoodsId int32 `gorm:"column:order_goods_id;type:int(11);" json:"order_goods_id"`
	/* 条码 */
	StockCode string `gorm:"column:stock_code;type:varchar(128);" json:"stock_code"`
	/* 商品名称 */
	OrderGoodsName string `gorm:"column:order_goods_name;type:varchar(255);" json:"order_goods_name"`
	/* 活动类型 1-自定义 2-生日 3-纪念日 4会员日 */
	ActivityType int32 `gorm:"column:activity_type;type:tinyint(4);" json:"activity_type"`
	/* 活动id */
	ActivityId int32 `gorm:"column:activity_id;type:int(11);" json:"activity_id"`
	/* 活动名称 */
	ActivityName string `gorm:"column:activity_name;type:varchar(255);" json:"activity_name"`
	/* 过期时间 */
	ExpireTime int32 `gorm:"column:expire_time;type:int(11);" json:"expire_time"`
	/* 操作人id */
	AdminUserId int32 `gorm:"column:admin_user_id;type:int(11);" json:"admin_user_id"`
	/* 操作人name */
	AdminUserName string `gorm:"column:admin_user_name;type:varchar(255);" json:"admin_user_name"`
	/*  */
	CreateTime int32 `gorm:"column:create_time;type:int(11);" json:"create_time"`
	/*  */
	UpdateTime int32 `gorm:"column:update_time;type:int(11);" json:"update_time"`
}

type CustomerLevelConfig struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/*  */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 品牌id */
	BrandId int32 `gorm:"column:brand_id;type:int(11);" json:"brand_id"`
	/* 商户门店ID */
	MerchantShopId int32 `gorm:"column:merchant_shop_id;type:int(11);" json:"merchant_shop_id"`
	/* 等级编号  1-n */
	Level int32 `gorm:"column:level;type:int(11);" json:"level"`
	/* 等级名称 例如：黄金会员，白银会员 */
	Name string `gorm:"column:name;type:varchar(256);" json:"name"`
	/* 等级分 */
	Point int32 `gorm:"column:point;type:int(11);" json:"point"`
	/* 该等级会员数 */
	CustomerCount int32 `gorm:"column:customer_count;type:int(11);" json:"customer_count"`
	/*  */
	CreateTime int32 `gorm:"column:create_time;type:int(11);" json:"create_time"`
	/*  */
	UpdateTime int32 `gorm:"column:update_time;type:int(11);" json:"update_time"`
}

type CustomerLog struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/*  */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/*  */
	MerchantShopId int32 `gorm:"column:merchant_shop_id;type:int(11);" json:"merchant_shop_id"`
	/*  */
	AdminId int32 `gorm:"column:admin_id;type:int(11);" json:"admin_id"`
	/*  */
	AdminName string `gorm:"column:admin_name;type:varchar(128);" json:"admin_name"`
	/*  */
	CustomerId int32 `gorm:"column:customer_id;type:int(11);" json:"customer_id"`
	/*  */
	CustomerName string `gorm:"column:customer_name;type:varchar(128);" json:"customer_name"`
	/* 1-自己修改 2-管理员备注,3-系统记录 */
	Type int32 `gorm:"column:type;type:tinyint(4);" json:"type"`
	/* 修改或备注内容 */
	Comment string `gorm:"column:comment;type:varchar(256);" json:"comment"`
	/*  */
	CreateTime string `gorm:"column:create_time;type:datetime;" json:"create_time"`
}

// 会员重要信息修改记录表, 记录会员的姓名、手机号、生日、结婚纪念日的修改，用来限制修改次数
type CustomerModifyTrace struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/*  */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/*  */
	CustomerId int32 `gorm:"column:customer_id;type:int(11);" json:"customer_id"`
	/* 操作人id */
	AdminUserId int32 `gorm:"column:admin_user_id;type:int(11);" json:"admin_user_id"`
	/* 1 姓名 2 手机号 3 生日 4 结婚纪念日 */
	Types int32 `gorm:"column:types;type:tinyint(1);" json:"types"`
	/*  */
	CreateTime int32 `gorm:"column:create_time;type:int(11);" json:"create_time"`
	/*  */
	UpdateTime int32 `gorm:"column:update_time;type:int(11);" json:"update_time"`
}

// 客户规则配置表
type CustomerRule struct {
	/* 主键 */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/* 商户id */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 品牌id */
	BrandId int32 `gorm:"column:brand_id;type:int(11);" json:"brand_id"`
	/* 生日提醒配置 */
	BdayRemind string `gorm:"column:bday_remind;type:json;" json:"bday_remind"`
	/* 结婚纪恋日提醒配置 */
	WeddingDayRemind string `gorm:"column:wedding_day_remind;type:json;" json:"wedding_day_remind"`
	/* 星级客户定义 */
	StarDefine string `gorm:"column:star_define;type:json;" json:"star_define"`
	/* 沉默客户定义 */
	SilenceDefine string `gorm:"column:silence_define;type:json;" json:"silence_define"`
	/* 活跃客户定义 */
	ActiveDefine string `gorm:"column:active_define;type:json;" json:"active_define"`
	/* 0333跟进提醒 */
	FollowRemind string `gorm:"column:follow_remind;type:json;" json:"follow_remind"`
	/* 创建时间 */
	CreateDatetime string `gorm:"column:create_datetime;type:datetime;" json:"create_datetime"`
	/* 更新时间 */
	UpdateDatetime string `gorm:"column:update_datetime;type:datetime;" json:"update_datetime"`
	/*  */
	UpdateTime int32 `gorm:"column:update_time;type:int(11);" json:"update_time"`
}

// 客户搜索日志
type CustomerSearchLog struct {
	/* 主键id */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/* 商户id */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 用户id */
	AdminUserId int32 `gorm:"column:admin_user_id;type:int(11);" json:"admin_user_id"`
	/* 搜索内容 */
	Content string `gorm:"column:content;type:varchar(255);" json:"content"`
	/* 状态, 1: 正常 0: 删除 */
	Status int32 `gorm:"column:status;type:tinyint(4);" json:"status"`
	/* 保存时间(时间戳) */
	SaveTime int32 `gorm:"column:save_time;type:int(11);" json:"save_time"`
}

// 客户标签
type CustomerTag struct {
	/* 主键id */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/* 商户id */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 标签|标签组名称 */
	Name string `gorm:"column:name;type:varchar(50);" json:"name"`
	/* 父id,0表示为group */
	Pid int32 `gorm:"column:pid;type:int(11);" json:"pid"`
	/* 状态(枚举, 1:启用 0: 删除) */
	Status int32 `gorm:"column:status;type:tinyint(4);" json:"status"`
	/* 排序号(数字越小,排序在前) */
	Order int32 `gorm:"column:order;type:int(11);" json:"order"`
	/* 是否为系统标签:1系统标签 其他:非系统标签 */
	IsSys int32 `gorm:"column:is_sys;type:tinyint(4);" json:"is_sys"`
	/* 操作人 */
	OpUser string `gorm:"column:op_user;type:varchar(255);" json:"op_user"`
	/* 创建时间戳 */
	CreateTime int64 `gorm:"column:create_time;type:bigint(20);" json:"create_time"`
	/* 更新时间戳 */
	UpdateTime int64 `gorm:"column:update_time;type:bigint(20);" json:"update_time"`
}

// 客户关联的标签
type CustomerTagRel struct {
	/* 主键id */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/* 商户id */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* scrm的客户id */
	CustomerId int32 `gorm:"column:customer_id;type:int(11);" json:"customer_id"`
	/* 标签id */
	TagId int32 `gorm:"column:tag_id;type:int(11);" json:"tag_id"`
	/* 企业微信员工id */
	QyUserId string `gorm:"column:qy_user_id;type:varchar(50);" json:"qy_user_id"`
	/* 员工id */
	AdminUserId int32 `gorm:"column:admin_user_id;type:int(11);" json:"admin_user_id"`
	/* 具体打标签时间 */
	MakeTime int64 `gorm:"column:make_time;type:bigint(20);" json:"make_time"`
	/* 创建时间 */
	CreateTime int64 `gorm:"column:create_time;type:bigint(20);" json:"create_time"`
}

type CustomerUserJoin struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/* 商户id */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 品牌id */
	BrandId int32 `gorm:"column:brand_id;type:int(11);" json:"brand_id"`
	/* 员工id */
	AdminUserId int32 `gorm:"column:admin_user_id;type:int(11);" json:"admin_user_id"`
	/* 粉丝id */
	UserId int32 `gorm:"column:user_id;type:int(11);" json:"user_id"`
	/* 会员id */
	CustomerId int32 `gorm:"column:customer_id;type:int(11);" json:"customer_id"`
	/* 状态 1：绑定 2：解绑 */
	Status int32 `gorm:"column:status;type:tinyint(4);" json:"status"`
	/* 建立绑定时间(员工与会员，员工与粉丝) */
	BindDatetime string `gorm:"column:bind_datetime;type:datetime;" json:"bind_datetime"`
	/* 解绑时间 */
	UnbindDatetime string `gorm:"column:unbind_datetime;type:datetime;" json:"unbind_datetime"`
	/* 解绑备注 */
	UnbindRemarks string `gorm:"column:unbind_remarks;type:varchar(128);" json:"unbind_remarks"`
	/* 在线导购最后聊天时间 */
	LastWechatDatetime string `gorm:"column:last_wechat_datetime;type:datetime;" json:"last_wechat_datetime"`
	/* a、新增客户，建立绑定关系的时间
	b、最后1条聊天信息的发送时间
	c、最后1次拨打电话记录的时间 */
	UpdateDatetime string `gorm:"column:update_datetime;type:datetime;" json:"update_datetime"`
	/* 5.未知（之前旧数据）   1.销售开单   2.扫名片   3.手动修改   4.公众号  6.企微同步 */
	Source int32 `gorm:"column:source;type:tinyint(1);" json:"source"`
}

type ExclusiveServerChatRecord struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(10) unsigned;" json:"id"`
	/* 商户id */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 品牌id */
	BrandId int32 `gorm:"column:brand_id;type:int(10) unsigned;" json:"brand_id"`
	/* 门店id */
	MerchantShopId int32 `gorm:"column:merchant_shop_id;type:int(10) unsigned;" json:"merchant_shop_id"`
	/* 员工id */
	StaffId int32 `gorm:"column:staff_id;type:int(10) unsigned;" json:"staff_id"`
	/* 员工名 */
	StaffName string `gorm:"column:staff_name;type:varchar(255);" json:"staff_name"`
	/* 粉丝id */
	UserId int32 `gorm:"column:user_id;type:int(10) unsigned;" json:"user_id"`
	/* 昵称 */
	NickName string `gorm:"column:nick_name;type:varchar(255);" json:"nick_name"`
	/* 用户id */
	CustomerId int32 `gorm:"column:customer_id;type:int(10) unsigned;" json:"customer_id"`
	/* 名字 */
	Name string `gorm:"column:name;type:varchar(255);" json:"name"`
	/* 聊天内容类型 1文本 2图片 3卡券 4模板 5自动回复 */
	Type int32 `gorm:"column:type;type:tinyint(3) unsigned;" json:"type"`
	/* 聊天内容 */
	Content string `gorm:"column:content;type:text;" json:"content"`
	/* 附加的信息，json格式 */
	AdditContent string `gorm:"column:addit_content;type:varchar(255);" json:"addit_content"`
	/* 发送来源	1-员工发出 2-客户发出 */
	Status int32 `gorm:"column:status;type:tinyint(3) unsigned;" json:"status"`
	/*  */
	CreateDatetime string `gorm:"column:create_datetime;type:timestamp;" json:"create_datetime"`
	/*  */
	UpdateDatetime string `gorm:"column:update_datetime;type:timestamp;" json:"update_datetime"`
	/*  */
	DeleteDatetime string `gorm:"column:delete_datetime;type:timestamp;" json:"delete_datetime"`
	/* 消息状态 1-成功 0-失败 2-服务消息 */
	MsgStatus int32 `gorm:"column:msg_status;type:tinyint(1);" json:"msg_status"`
}

type ExclusiveServerLog struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(10) unsigned;" json:"id"`
	/* 商户id */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 品牌id */
	BrandId int32 `gorm:"column:brand_id;type:int(10) unsigned;" json:"brand_id"`
	/* 门店id */
	MerchantShopId int32 `gorm:"column:merchant_shop_id;type:int(10) unsigned;" json:"merchant_shop_id"`
	/* 员工id */
	StaffId int32 `gorm:"column:staff_id;type:int(10) unsigned;" json:"staff_id"`
	/* 员工名 */
	StaffName string `gorm:"column:staff_name;type:varchar(255);" json:"staff_name"`
	/* 客户类型1粉丝 2客户 */
	Type int32 `gorm:"column:type;type:tinyint(1);" json:"type"`
	/* 粉丝表 */
	UserId int32 `gorm:"column:user_id;type:int(10) unsigned;" json:"user_id"`
	/* 头像 */
	HeadImg string `gorm:"column:head_img;type:varchar(255);" json:"head_img"`
	/* 昵称 */
	NickName string `gorm:"column:nick_name;type:varchar(255);" json:"nick_name"`
	/* 性别 默认0未知 1男 2女 */
	Sex int32 `gorm:"column:sex;type:tinyint(3) unsigned;" json:"sex"`
	/* 用户表 */
	CustomerId int32 `gorm:"column:customer_id;type:int(10) unsigned;" json:"customer_id"`
	/* 名字 */
	Name string `gorm:"column:name;type:varchar(255);" json:"name"`
	/* 电话 */
	Telephone string `gorm:"column:telephone;type:varchar(255);" json:"telephone"`
	/* 状态 1绑定 2解绑 */
	Status int32 `gorm:"column:status;type:tinyint(1);" json:"status"`
	/* 备注 */
	Remark string `gorm:"column:remark;type:varchar(255);" json:"remark"`
	/* 操作人 */
	AdminUserId int32 `gorm:"column:admin_user_id;type:int(11);" json:"admin_user_id"`
	/* 绑定人名 */
	AdminUserName string `gorm:"column:admin_user_name;type:varchar(255);" json:"admin_user_name"`
	/*  */
	CreateDatetime string `gorm:"column:create_datetime;type:timestamp;" json:"create_datetime"`
	/*  */
	UpdateDatetime string `gorm:"column:update_datetime;type:timestamp;" json:"update_datetime"`
	/*  */
	DeleteDatetime string `gorm:"column:delete_datetime;type:timestamp;" json:"delete_datetime"`
}

type ExclusiveServerRelation struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(10) unsigned;" json:"id"`
	/* 商户id */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 品牌id */
	BrandId int32 `gorm:"column:brand_id;type:int(10) unsigned;" json:"brand_id"`
	/* 门店id */
	MerchantShopId int32 `gorm:"column:merchant_shop_id;type:int(10) unsigned;" json:"merchant_shop_id"`
	/* 员工id */
	StaffId int32 `gorm:"column:staff_id;type:int(10) unsigned;" json:"staff_id"`
	/* 员工名 */
	StaffName string `gorm:"column:staff_name;type:varchar(255);" json:"staff_name"`
	/* 粉丝表 */
	UserId int32 `gorm:"column:user_id;type:int(10) unsigned;" json:"user_id"`
	/* 客户状态 默认1-粉丝 2-用户 */
	Type int32 `gorm:"column:type;type:tinyint(3) unsigned;" json:"type"`
	/* 头像 */
	HeadImg string `gorm:"column:head_img;type:varchar(255);" json:"head_img"`
	/* 昵称 */
	NickName string `gorm:"column:nick_name;type:varchar(255);" json:"nick_name"`
	/* 性别 默认0未知 1男 2女 */
	Sex int32 `gorm:"column:sex;type:tinyint(3) unsigned;" json:"sex"`
	/* 用户表 */
	CustomerId int32 `gorm:"column:customer_id;type:int(10) unsigned;" json:"customer_id"`
	/* 名字 */
	Name string `gorm:"column:name;type:varchar(255);" json:"name"`
	/* 电话 */
	Telephone string `gorm:"column:telephone;type:varchar(255);" json:"telephone"`
	/* 绑定时间 */
	BindTime string `gorm:"column:bind_time;type:timestamp;" json:"bind_time"`
	/* 操作人 */
	AdminUserId int32 `gorm:"column:admin_user_id;type:int(11);" json:"admin_user_id"`
	/* 绑定人名 */
	AdminUserName string `gorm:"column:admin_user_name;type:varchar(255);" json:"admin_user_name"`
	/* 状态 1绑定 2逻辑删除(解绑) */
	Status int32 `gorm:"column:status;type:tinyint(3) unsigned;" json:"status"`
	/*  */
	CreateDatetime string `gorm:"column:create_datetime;type:timestamp;" json:"create_datetime"`
	/*  */
	UpdateDatetime string `gorm:"column:update_datetime;type:timestamp;" json:"update_datetime"`
	/*  */
	DeleteDatetime string `gorm:"column:delete_datetime;type:timestamp;" json:"delete_datetime"`
}

// 金价配制信息表
type GoldSalePriceConfig struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/*  */
	ParentId int32 `gorm:"column:parent_id;type:int(11);" json:"parent_id"`
	/* 商户ID */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 门店ID */
	MerchantShopId int32 `gorm:"column:merchant_shop_id;type:int(11);" json:"merchant_shop_id"`
	/* 金价类型1-新品金价 2-旧料金价 */
	PriceType int32 `gorm:"column:price_type;type:tinyint(1);" json:"price_type"`
	/* 品牌id */
	BrandId int32 `gorm:"column:brand_id;type:int(11);" json:"brand_id"`
	/* 金价名称 */
	Name string `gorm:"column:name;type:varchar(50);" json:"name"`
	/* 商品分类 */
	GoodsTypeId int32 `gorm:"column:goods_type_id;type:int(11);" json:"goods_type_id"`
	/*  */
	GoodsTypeIdOld int32 `gorm:"column:goods_type_id_old;type:int(11);" json:"goods_type_id_old"`
	/* 商品分类名称（金价名称） */
	GoodsTypeName string `gorm:"column:goods_type_name;type:varchar(255);" json:"goods_type_name"`
	/* 销售金价（每克）单位：分 */
	SalePrice int32 `gorm:"column:sale_price;type:int(11);" json:"sale_price"`
	/* 回收金价(每克) 单位：分 */
	RecyclePrice int32 `gorm:"column:recycle_price;type:int(11);" json:"recycle_price"`
	/* 旧料换购单价（每克）单位：分 */
	ChangePrice int32 `gorm:"column:change_price;type:int(11);" json:"change_price"`
	/* 旧料换购工费单价（每克）单位：分 */
	ChangeLabourPrice int32 `gorm:"column:change_labour_price;type:int(11);" json:"change_labour_price"`
	/* 旧料回收单价（每克）单位：分 */
	RecyclePriceNew int32 `gorm:"column:recycle_price_new;type:int(11);" json:"recycle_price_new"`
	/* 旧料回收工费单价（每克）单位：分 */
	RecycleLabourPrice int32 `gorm:"column:recycle_labour_price;type:int(11);" json:"recycle_labour_price"`
	/* 操作人 */
	AdminUserId int32 `gorm:"column:admin_user_id;type:int(11);" json:"admin_user_id"`
	/* 操作人名称 */
	AdminUserName string `gorm:"column:admin_user_name;type:varchar(50);" json:"admin_user_name"`
	/* 是否锁定 1-锁定 2-非锁定 */
	Islock int32 `gorm:"column:islock;type:tinyint(1);" json:"islock"`
	/* 状态 1-正常 2-删除 */
	Status int32 `gorm:"column:status;type:tinyint(1);" json:"status"`
	/* 创建时间 */
	CreateTime int32 `gorm:"column:create_time;type:int(11);" json:"create_time"`
	/*  */
	UpdateTime int32 `gorm:"column:update_time;type:int(11);" json:"update_time"`
	/*  */
	DeleteTime int32 `gorm:"column:delete_time;type:int(11);" json:"delete_time"`
	/* 入库匹配字段集 */
	ImportMatching string `gorm:"column:import_matching;type:varchar(2550);" json:"import_matching"`
}

// 记录打的客户标签日志
type MakeTagLog struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/* 商户id */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 打标签的人 */
	AdminUserId int32 `gorm:"column:admin_user_id;type:int(11);" json:"admin_user_id"`
	/* 标签id */
	TagId int32 `gorm:"column:tag_id;type:int(11);" json:"tag_id"`
	/* 企微中的标签id */
	QyTagId string `gorm:"column:qy_tag_id;type:varchar(50);" json:"qy_tag_id"`
	/* 客户id */
	CustomerId int32 `gorm:"column:customer_id;type:int(11);" json:"customer_id"`
	/* 创建时间戳 */
	CreateTime int64 `gorm:"column:create_time;type:bigint(20);" json:"create_time"`
}

type MerchantSetting struct {
	/* 配置表id */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/* 金额取位数    1.四舍五入到小数点后两位 2.四舍五入到个位 */
	PriceToFixed int32 `gorm:"column:price_to_fixed;type:int(11);" json:"price_to_fixed"`
	/* 标签价规则 1.四舍五入到小数点后两位 2.四舍五入到个位 */
	PriceTagToFixed int32 `gorm:"column:price_tag_to_fixed;type:int(11);" json:"price_tag_to_fixed"`
	/* 退货管控设置 1=无需验证,2=密码验证,3=审核验证 */
	OldReturnsToFixed int32 `gorm:"column:old_returns_to_fixed;type:tinyint(4);" json:"old_returns_to_fixed"`
	/* 撤单管控设置 1=无需验证,2=密码验证,3=审核验证 */
	CancelOrderToFixed int32 `gorm:"column:cancel_order_to_fixed;type:tinyint(4);" json:"cancel_order_to_fixed"`
	/* 折扣超限设置 1.自动修改为最低折扣 2.按员工  3.按商品  4.按单项优惠 */
	DiscountToFixed int32 `gorm:"column:discount_to_fixed;type:tinyint(4);" json:"discount_to_fixed"`
	/* 商户id */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 1:升级0：不升级 */
	IsNewPolicy int32 `gorm:"column:is_new_policy;type:tinyint(4);" json:"is_new_policy"`
	/* 更新时间 */
	UpdateTime int32 `gorm:"column:update_time;type:int(11);" json:"update_time"`
	/* 创建时间 */
	CreateTime int32 `gorm:"column:create_time;type:int(11);" json:"create_time"`
	/* 修改人 */
	AdminUserName int32 `gorm:"column:admin_user_name;type:int(11);" json:"admin_user_name"`
	/* 收银班次制度 1=启用,2=不启用 */
	ShiftToFixed int32 `gorm:"column:shift_to_fixed;type:tinyint(4);" json:"shift_to_fixed"`
	/* 班次(默认A,B班) [{"id":"1","name":"A班"},{"id":"2","name":"B班"},{"id":"3","name":"C班"},{"id":"4","name":"A班和B班"},{"id":"5","name":"A班和C班"},{"id":"6","name":"B班和C班"}] */
	ShiftSetting string `gorm:"column:shift_setting;type:varchar(128);" json:"shift_setting"`
	/* 柜台是否允许销售 */
	IsSaleInWarehouse int32 `gorm:"column:is_sale_in_warehouse;type:tinyint(4);" json:"is_sale_in_warehouse"`
	/* 活动折扣管控  1:允许随意修改  2:不允许修改优惠 3:允许减少优惠幅度 */
	ActDiscount int32 `gorm:"column:act_discount;type:tinyint(4);" json:"act_discount"`
	/* 线下优惠券管控 1=开启不需要审核,2=关闭 ,3=开启需要审核 */
	OfflineCoupon int32 `gorm:"column:offline_coupon;type:tinyint(4);" json:"offline_coupon"`
	/* 是否允许非素旧料换购素金商品 1=是,2=否 */
	ChangeOld int32 `gorm:"column:change_old;type:tinyint(1);" json:"change_old"`
	/* 是否允许非素新品换购素金商品 */
	ChangeNew int32 `gorm:"column:change_new;type:tinyint(1);" json:"change_new"`
	/* 旧料抵扣方式设置 1-按商品权重分配 2-手动指定换购商品 */
	OldDeductType int32 `gorm:"column:old_deduct_type;type:tinyint(4);" json:"old_deduct_type"`
	/* 升级旧料按克管理1:升级0：不升级 */
	OldStockWeight int32 `gorm:"column:old_stock_weight;type:tinyint(1);" json:"old_stock_weight"`
	/* 旧料换购/回收时确定素非素的方式，1-跟新品一样依据分类，2-按照回收计价方式（按克-素金，按件-非素） */
	OldStockGoldTypeCalcType int32 `gorm:"column:old_stock_gold_type_calc_type;type:tinyint(1);" json:"old_stock_gold_type_calc_type"`
	/* 业绩设置 */
	PerformanceSetting string `gorm:"column:performance_setting;type:text;" json:"performance_setting"`
	/* 是否需要修改业绩比例  1需要  2不需要 */
	PerformanceRatio int32 `gorm:"column:performance_ratio;type:tinyint(1);" json:"performance_ratio"`
	/* 业绩提成设置  {
	    "is_open":"1", //开关
	    "setting":[
	        {
	            "id":"8453",  //类型id
	            "list":[
	                {
	                    "money":"",             //设置金额
	                    "discount_max":"100",   //最大金额
	                    "discount_min":"0"      //最小金额
	                }
	            ],
	            "name":"3D硬金", //类型名称
	            "type":"1"      //1.按克 2.按件
	        }
	    ],
	    "commission_type":  //1.实收金额 2.应收金额
	} */
	CommissionSetting string `gorm:"column:commission_setting;type:json;" json:"commission_setting"`
	/* 素转非是否计入门店收款 */
	IsExchangeTypeGoldToNogold int32 `gorm:"column:is_exchange_type_gold_to_nogold;type:tinyint(1);" json:"is_exchange_type_gold_to_nogold"`
	/* 旧料回收/换购 是否显示损耗 1 是 2 否 */
	OldExchangeOrRecycleViewLoss int32 `gorm:"column:old_exchange_or_recycle_view_loss;type:tinyint(4);" json:"old_exchange_or_recycle_view_loss"`
	/* 旧料回收/换购时修改金额影响 1.影响回收金价 2. 影响工费金额 */
	RecyclingPriceConfig int32 `gorm:"column:recycling_price_config;type:tinyint(1);" json:"recycling_price_config"`
	/* 是否允许收银端开单修改时间 */
	AllowSaleTime int32 `gorm:"column:allow_sale_time;type:tinyint(1);" json:"allow_sale_time"`
	/* 0:全部  1：自收 2：商收 */
	SaleType int32 `gorm:"column:sale_type;type:tinyint(1);" json:"sale_type"`
	/* 仅查总销售额的比例值 */
	SaleTypeValue int32 `gorm:"column:sale_type_value;type:int(11);" json:"sale_type_value"`
	/* 收银列表 */
	SaleTypeList string `gorm:"column:sale_type_list;type:json;" json:"sale_type_list"`
	/* 开单是否允许选择销售时间  1:允许  2：不允许 */
	IsSaleTime int32 `gorm:"column:is_sale_time;type:tinyint(1);" json:"is_sale_time"`
	/* 角色折扣权限配置的角色 */
	RoleIds string `gorm:"column:role_ids;type:varchar(200);" json:"role_ids"`
	/* 发送预警时间 */
	SendTime string `gorm:"column:send_time;type:varchar(256);" json:"send_time"`
	/* 库存预警筛选项最后配置 */
	WarningConfigSelected string `gorm:"column:warning_config_selected;type:varchar(256);" json:"warning_config_selected"`
	/* 盘点是否可用 1 可用 2 不可用 */
	IsInventory int32 `gorm:"column:is_inventory;type:tinyint(1);" json:"is_inventory"`
	/* 标签价尾数 个位 */
	TagPriceMantissaRulesOnes string `gorm:"column:tag_price_mantissa_rules_ones;type:json;" json:"tag_price_mantissa_rules_ones"`
	/* 标签价尾数 十位 */
	TagPriceMantissaRulesTens string `gorm:"column:tag_price_mantissa_rules_tens;type:json;" json:"tag_price_mantissa_rules_tens"`
	/* 1允许2不允许 */
	AllowEditNewGoldPrice int32 `gorm:"column:allow_edit_new_gold_price;type:tinyint(1);" json:"allow_edit_new_gold_price"`
	/* 1 允许 2不允许 */
	AllowEditOldGoldPrice int32 `gorm:"column:allow_edit_old_gold_price;type:tinyint(1);" json:"allow_edit_old_gold_price"`
	/* 1 允许 2 不允许 */
	AllowEditOldRecycleGoldPrice int32 `gorm:"column:allow_edit_old_recycle_gold_price;type:tinyint(1);" json:"allow_edit_old_recycle_gold_price"`
	/* 1 允许 2 不允许 */
	AllowEditNewLabourPrice int32 `gorm:"column:allow_edit_new_labour_price;type:tinyint(1);" json:"allow_edit_new_labour_price"`
	/* 1 允许 2 不允许 */
	AllowEditOldLabourPrice int32 `gorm:"column:allow_edit_old_labour_price;type:tinyint(1);" json:"allow_edit_old_labour_price"`
	/* 1 允许 2 不允许 */
	AllowEditOldRecycleLabourPrice int32 `gorm:"column:allow_edit_old_recycle_labour_price;type:tinyint(4);" json:"allow_edit_old_recycle_labour_price"`
	/* 1 开启 2 关闭 */
	Handover int32 `gorm:"column:handover;type:tinyint(1);" json:"handover"`
	/* 盘点是否记录盘盈数据   //默认0   0:是  1:否 */
	StockCheckType int32 `gorm:"column:stock_check_type;type:tinyint(1);" json:"stock_check_type"`
	/* 移动端是否允许查询无权限门店的库存 0否 1是 */
	IsShopStock int32 `gorm:"column:is_shop_stock;type:tinyint(3) unsigned;" json:"is_shop_stock"`
	/* 1允许 2不允许 修改业绩分配比列 */
	AllowEditPerformance int32 `gorm:"column:allow_edit_performance;type:tinyint(1);" json:"allow_edit_performance"`
	/* 手机开单是否运行自动打印 0:不允许 1允许 */
	AutoPrint int32 `gorm:"column:auto_print;type:tinyint(1);" json:"auto_print"`
	/* 是否允许隔天补录客户信息，1：允许，2：不允许 */
	BackTracingSetting int32 `gorm:"column:back_tracing_setting;type:tinyint(1);" json:"back_tracing_setting"`
	/* 收银台按日期筛选历史销售单，1：开启，2：关闭 */
	DateFilterSetting int32 `gorm:"column:date_filter_setting;type:tinyint(1);" json:"date_filter_setting"`
	/* 是否允许结算时使用截金功能 1 允许 2 不允许 */
	UseCutBySettlement int32 `gorm:"column:use_cut_by_settlement;type:tinyint(1);" json:"use_cut_by_settlement"`
	/* 是否可以查看历史导入销售单据 1 允许 2不允许 */
	ViewHistoryImportSaleOrderType int32 `gorm:"column:view_history_import_sale_order_type;type:tinyint(1);" json:"view_history_import_sale_order_type"`
	/* 导入历史销售单据是否生成电子保单，1：是，2：否 */
	IsCreatePolicyImportHistoryOrder int32 `gorm:"column:is_create_policy_import_history_order;type:tinyint(1);" json:"is_create_policy_import_history_order"`
	/* 库存状态 逗号隔开 1.正常 2.已售出 3.已报损  4.已转旧料 5.已预订 6.已锁定 8.已出库 */
	GoodsStockSearchStatus []byte `gorm:"column:goods_stock_search_status;type:varbinary(128);" json:"goods_stock_search_status"`
	/* 0-默认显示logo 1-不显示logo */
	NotShowLogo int32 `gorm:"column:not_show_logo;type:tinyint(1);" json:"not_show_logo"`
	/* 1 保持为领取时券内容 2更新为编辑后的券内容 */
	CouponSetting int32 `gorm:"column:coupon_setting;type:tinyint(1);" json:"coupon_setting"`
	/* 客户姓名、手机号修改模式 0 修改次数不限制 1 仅允许修改1次 2不允许修改 */
	CustomerNameModifyMode int32 `gorm:"column:customer_name_modify_mode;type:tinyint(1);" json:"customer_name_modify_mode"`
	/* 客户生日、结婚纪念日修改模式 0 修改次数不限制 1 仅允许修改1次 2不允许修改 */
	CustomerBirthdayModifyMode int32 `gorm:"column:customer_birthday_modify_mode;type:tinyint(1);" json:"customer_birthday_modify_mode"`
	/* 平均成本金价是否生效 */
	AvgCostPriceEnable int32 `gorm:"column:avg_cost_price_enable;type:tinyint(1);" json:"avg_cost_price_enable"`
	/* 员工在企业微信好友详情里是否可见会员信息的设置:1 都可见 2非归属人不可见 */
	CustomerAffiliationView int32 `gorm:"column:customer_affiliation_view;type:tinyint(1);" json:"customer_affiliation_view"`
	/* 历史销售单据类型   0：默认  1.全部门店 */
	HistoryOrderType int32 `gorm:"column:history_order_type;type:tinyint(1);" json:"history_order_type"`
	/* 是否运行跨度选择销售员  1：允许  2：不允许 */
	CanCrossShop int32 `gorm:"column:can_cross_shop;type:tinyint(1);" json:"can_cross_shop"`
	/* 销售单据日期筛选是否允许使用时、分、秒查询 */
	IsSaleDateAccurate int32 `gorm:"column:is_sale_date_accurate;type:tinyint(1);" json:"is_sale_date_accurate"`
	/* 员工名片码显示：0企微1公众号 */
	SalesmanQrcodeType int32 `gorm:"column:salesman_qrcode_type;type:tinyint(1);" json:"salesman_qrcode_type"`
	/* 老板端审核单据显示类型  1：被选人才能看到 2： 老板端就能看到 */
	AppAuditType int32 `gorm:"column:app_audit_type;type:tinyint(1);" json:"app_audit_type"`
}

type MerchantShop struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/*  */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 商户号 */
	MchId string `gorm:"column:mch_id;type:varchar(10);" json:"mch_id"`
	/* 是否有开通退款默认0-未开通 1已开通 */
	IsHaveRefund int32 `gorm:"column:is_have_refund;type:tinyint(4);" json:"is_have_refund"`
	/* 名称 */
	Name string `gorm:"column:name;type:varchar(100);" json:"name"`
	/*  */
	IsLeader int32 `gorm:"column:is_leader;type:tinyint(4);" json:"is_leader"`
	/* 是否显示在对应品牌的公众号附近门店中1是2否 */
	IsDisplay int32 `gorm:"column:is_display;type:tinyint(3) unsigned;" json:"is_display"`
	/* 是否开通收银台权限 */
	IsAllowedCasher int32 `gorm:"column:is_allowed_casher;type:tinyint(4);" json:"is_allowed_casher"`
	/* 门店电话 */
	ShopPhone string `gorm:"column:shop_phone;type:varchar(30);" json:"shop_phone"`
	/*  */
	Logo int32 `gorm:"column:logo;type:int(11);" json:"logo"`
	/*  */
	Photos string `gorm:"column:photos;type:text;" json:"photos"`
	/* 联系人姓名 */
	ContactName string `gorm:"column:contact_name;type:varchar(50);" json:"contact_name"`
	/* 联系人电话 */
	ContactPhone string `gorm:"column:contact_phone;type:varchar(50);" json:"contact_phone"`
	/* 证件照 */
	Certificate int32 `gorm:"column:certificate;type:int(11);" json:"certificate"`
	/* 省 */
	Province string `gorm:"column:province;type:varchar(50);" json:"province"`
	/* 城市 */
	City string `gorm:"column:city;type:varchar(50);" json:"city"`
	/* 区 */
	District string `gorm:"column:district;type:varchar(50);" json:"district"`
	/* 镇 */
	Town string `gorm:"column:town;type:varchar(100);" json:"town"`
	/* 详细地址 */
	Address string `gorm:"column:address;type:varchar(256);" json:"address"`
	/* 缩略图 */
	Thumb int32 `gorm:"column:thumb;type:int(11);" json:"thumb"`
	/* 纬度 */
	Lat string `gorm:"column:lat;type:varchar(100);" json:"lat"`
	/* 经度 */
	Lng string `gorm:"column:lng;type:varchar(100);" json:"lng"`
	/* 操作人ID */
	AdminUserId int32 `gorm:"column:admin_user_id;type:int(11);" json:"admin_user_id"`
	/* 修改人ID */
	UpdateAdminUserId int32 `gorm:"column:update_admin_user_id;type:int(11);" json:"update_admin_user_id"`
	/* 第一次编辑时间 */
	FirstChangeTime int32 `gorm:"column:first_change_time;type:int(11);" json:"first_change_time"`
	/* 退货密码 */
	BackCode string `gorm:"column:back_code;type:varchar(255);" json:"back_code"`
	/* 充值金额(单位：分) */
	SurplusMoney int64 `gorm:"column:surplus_money;type:bigint(20);" json:"surplus_money"`
	/* 短信收费配置 每条短信多少钱（单位：分） */
	SmsConfig int32 `gorm:"column:sms_config;type:int(11);" json:"sms_config"`
	/* 赠送金额配置(单位：分) */
	LargessMoney int32 `gorm:"column:largess_money;type:int(11);" json:"largess_money"`
	/* 剩余赠送金额(单位：分) */
	SurplusLargessMoney int32 `gorm:"column:surplus_largess_money;type:int(11);" json:"surplus_largess_money"`
	/* 门店铺编号 */
	ShopCode int32 `gorm:"column:shop_code;type:int(11);" json:"shop_code"`
	/* 商户+门店铺编码(前8位是商户后两位是门店铺) */
	ShopCodeFull string `gorm:"column:shop_code_full;type:varchar(20);" json:"shop_code_full"`
	/* 门店编码 */
	ShopSn string `gorm:"column:shop_sn;type:varchar(60);" json:"shop_sn"`
	/* 小程序商城 0:未开启，1:开启 */
	WechatMallEnabled int32 `gorm:"column:wechat_mall_enabled;type:tinyint(1);" json:"wechat_mall_enabled"`
	/* 小程序APPID */
	WechatMallAppid string `gorm:"column:wechat_mall_appid;type:varchar(64);" json:"wechat_mall_appid"`
	/* 品牌ID */
	BrandId int32 `gorm:"column:brand_id;type:int(11);" json:"brand_id"`
	/*  */
	BrandIdOld int32 `gorm:"column:brand_id_old;type:int(11);" json:"brand_id_old"`
	/* 1:打开 0：关闭 */
	IsPolicy int32 `gorm:"column:is_policy;type:tinyint(4);" json:"is_policy"`
	/* 签约类型 1-按年 2-按月 */
	SignType int32 `gorm:"column:sign_type;type:tinyint(4);" json:"sign_type"`
	/* 续费充值配置id */
	RenewRechargeConfigId int32 `gorm:"column:renew_recharge_config_id;type:int(11);" json:"renew_recharge_config_id"`
	/* 收银端默认支付方式  1.现金支付  2.微信支付  3.支付宝   4银行卡/信用卡 */
	DefaultPayMode int32 `gorm:"column:default_pay_mode;type:tinyint(4);" json:"default_pay_mode"`
	/* 录入时间
	是指门店信息被商家首次在商户管理页面编辑保存门店时的时间，后续更新门店信息不能再影响这里 */
	InsertDateTime string `gorm:"column:insert_date_time;type:timestamp;" json:"insert_date_time"`
	/* 启用时间 */
	StartDateTime string `gorm:"column:start_date_time;type:timestamp;" json:"start_date_time"`
	/* 到期时间 */
	DueDateTime string `gorm:"column:due_date_time;type:timestamp;" json:"due_date_time"`
	/* 内部企业id,外键 */
	QyCorpId int32 `gorm:"column:qy_corp_id;type:int(11);" json:"qy_corp_id"`
	/* 超级金店对应的门店id */
	CjjdShopId int32 `gorm:"column:cjjd_shop_id;type:int(11);" json:"cjjd_shop_id"`
	/* 门店分组 */
	GroupId int32 `gorm:"column:group_id;type:int(11);" json:"group_id"`
	/* 默认支付方式的ID */
	SaleBankId int32 `gorm:"column:sale_bank_id;type:int(11);" json:"sale_bank_id"`
	/* 默认启用支付方式json */
	SaleBankList string `gorm:"column:sale_bank_list;type:varchar(1024);" json:"sale_bank_list"`
	/* 金额取位数    1.四舍五入到小数点后两位 2.四舍五入到个位 */
	PriceToFixed int32 `gorm:"column:price_to_fixed;type:int(11);" json:"price_to_fixed"`
	/* 1-正常 -1禁用 2-删除 3-未启用 4-已到期 9-仅适用于超级金店的正常门店 */
	Status int32 `gorm:"column:status;type:tinyint(4);" json:"status"`
	/*  */
	CreateTime int32 `gorm:"column:create_time;type:int(11);" json:"create_time"`
	/*  */
	UpdateTime int32 `gorm:"column:update_time;type:int(11);" json:"update_time"`
	/*  */
	DeleteTime int32 `gorm:"column:delete_time;type:int(11);" json:"delete_time"`
	/* 归属于类型  0：zby ； 1: 超店 */
	BelongType int32 `gorm:"column:belong_type;type:tinyint(4);" json:"belong_type"`
}

type MerchantStaff struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/* 商户ID */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 头像 */
	Headimg string `gorm:"column:headimg;type:varchar(255);" json:"headimg"`
	/* 姓名 */
	Name string `gorm:"column:name;type:varchar(255);" json:"name"`
	/* 性别 1-男 2-女 */
	Sex int32 `gorm:"column:sex;type:tinyint(4);" json:"sex"`
	/* 身份证号码 */
	Idnumber string `gorm:"column:idnumber;type:varchar(20);" json:"idnumber"`
	/* 手机号码 */
	Cellphone string `gorm:"column:cellphone;type:varchar(15);" json:"cellphone"`
	/* 生日 */
	Birthday string `gorm:"column:birthday;type:date;" json:"birthday"`
	/* 学历 */
	Education int32 `gorm:"column:education;type:tinyint(1);" json:"education"`
	/* 婚姻状态 */
	Marriage int32 `gorm:"column:marriage;type:tinyint(1);" json:"marriage"`
	/* 紧急联系人1 */
	EmergencyName1 string `gorm:"column:emergency_name1;type:varchar(50);" json:"emergency_name1"`
	/* 紧急联系人1电话 */
	EmergencyCellphone1 string `gorm:"column:emergency_cellphone1;type:varchar(20);" json:"emergency_cellphone1"`
	/* 紧急联系人2 */
	EmergencyName2 string `gorm:"column:emergency_name2;type:varchar(50);" json:"emergency_name2"`
	/* 紧急联系人2电话 */
	EmergencyCellphone2 string `gorm:"column:emergency_cellphone2;type:varchar(20);" json:"emergency_cellphone2"`
	/* 底薪 */
	Salary int32 `gorm:"column:salary;type:int(11);" json:"salary"`
	/* 岗位补贴 */
	Subsidy int32 `gorm:"column:subsidy;type:int(11);" json:"subsidy"`
	/* 员工状态1-正式工2-试用期3-离职 */
	StaffStatus int32 `gorm:"column:staff_status;type:tinyint(1);" json:"staff_status"`
	/* 入职日期 */
	EntryDate string `gorm:"column:entry_date;type:date;" json:"entry_date"`
	/* 离职日期 */
	QuitDate string `gorm:"column:quit_date;type:varchar(20);" json:"quit_date"`
	/* 省 */
	Province string `gorm:"column:province;type:varchar(20);" json:"province"`
	/* 市 */
	City string `gorm:"column:city;type:varchar(20);" json:"city"`
	/* 区 */
	District string `gorm:"column:district;type:varchar(20);" json:"district"`
	/* 街道 */
	Town string `gorm:"column:town;type:varchar(20);" json:"town"`
	/* 详细地址 */
	Address string `gorm:"column:address;type:varchar(500);" json:"address"`
	/* 商户代码 */
	MerchantCode string `gorm:"column:merchant_code;type:varchar(255);" json:"merchant_code"`
	/* 所持账号ID  (已废弃) */
	AdminUserIds string `gorm:"column:admin_user_ids;type:varchar(1024);" json:"admin_user_ids"`
	/* 所属门店ID */
	MerchantShopIds string `gorm:"column:merchant_shop_ids;type:varchar(1024);" json:"merchant_shop_ids"`
	/* 所属角色 */
	AdminRoleIds string `gorm:"column:admin_role_ids;type:varchar(1024);" json:"admin_role_ids"`
	/* 证件信息，url_list，用 ,隔开 */
	Photos string `gorm:"column:photos;type:text;" json:"photos"`
	/* 是否分配账号 1-是 2-否 */
	Isdistribution int32 `gorm:"column:isdistribution;type:tinyint(4);" json:"isdistribution"`
	/* 是否需要首次编辑（APP） 默认1，表示APP需要编辑，编辑保存后变为0 */
	FirstEdit int32 `gorm:"column:first_edit;type:tinyint(1);" json:"first_edit"`
	/* 状态 1-正常 2-禁用 4-删除 */
	Status int32 `gorm:"column:status;type:tinyint(4);" json:"status"`
	/* 所持账号ID,一个员工对应一个帐号ID */
	AdminUserId int32 `gorm:"column:admin_user_id;type:int(11);" json:"admin_user_id"`
	/* 是否是销售员，0:不是，1:是 */
	IsSalePerson int32 `gorm:"column:is_sale_person;type:tinyint(4);" json:"is_sale_person"`
	/* 创建时间 */
	CreateTime int32 `gorm:"column:create_time;type:int(11);" json:"create_time"`
	/* 修改时间 */
	UpdateTime int32 `gorm:"column:update_time;type:int(11);" json:"update_time"`
	/* 删除时间 */
	DeleteTime int32 `gorm:"column:delete_time;type:int(11);" json:"delete_time"`
	/* 员工所在门店的二维码,如果有多个门店已最后生成的那个门店 */
	QrcodeUrl string `gorm:"column:qrcode_url;type:varchar(200);" json:"qrcode_url"`
	/* 员工编号 */
	EmployeeNumber string `gorm:"column:employee_number;type:varchar(50);" json:"employee_number"`
	/* 是否需要执行企微任务  0:否， 1：是 */
	IsCorpTask int32 `gorm:"column:is_corp_task;type:tinyint(1);" json:"is_corp_task"`
}

type OnlineGoldPrice struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/*  */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/*  */
	BrandId int32 `gorm:"column:brand_id;type:int(11);" json:"brand_id"`
	/* 类型是：1：新品金价，2：旧料价格，关联到gold_sale_price_config表 */
	GoldSalePriceConfigId int32 `gorm:"column:gold_sale_price_config_id;type:int(11);" json:"gold_sale_price_config_id"`
	/* 关联金价类型：1：新品金价，2：旧料价格，3：自定义 */
	Type int32 `gorm:"column:type;type:tinyint(4);" json:"type"`
	/* 金价名称 */
	Name string `gorm:"column:name;type:varchar(50);" json:"name"`
	/* 自定义关联金价价格：（每克）单位：分 */
	SalePrice int32 `gorm:"column:sale_price;type:int(11);" json:"sale_price"`
	/*  */
	UpdateTime int64 `gorm:"column:update_time;type:bigint(20);" json:"update_time"`
	/*  */
	CreateTime int64 `gorm:"column:create_time;type:bigint(20);" json:"create_time"`
	/*  */
	DeleteTime int64 `gorm:"column:delete_time;type:bigint(20);" json:"delete_time"`
}

// 在线橱窗配置表
type OnlineShowcaseSetting struct {
	/* 活动主键 */
	Id int32 `gorm:"primary_key;column:id;type:int(10) unsigned;" json:"id"`
	/* 商户id */
	MerchantId int32 `gorm:"column:merchant_id;type:int(10) unsigned;" json:"merchant_id"`
	/* 品牌ID */
	BrandId int32 `gorm:"column:brand_id;type:int(11);" json:"brand_id"`
	/* 是否打开 */
	IsOpen int32 `gorm:"column:is_open;type:tinyint(4);" json:"is_open"`
	/* 是否显示实时金价 1是 2否 */
	RealGoldPrice int32 `gorm:"column:real_gold_price;type:tinyint(4);" json:"real_gold_price"`
	/* 按克销售（元/克） */
	GramDiscount int32 `gorm:"column:gram_discount;type:int(11);" json:"gram_discount"`
	/* 按件销售商品是否支持支持定金：1是2否 */
	EnablePieceDiscount int32 `gorm:"column:enable_piece_discount;type:tinyint(1);" json:"enable_piece_discount"`
	/* 按件销售折扣 */
	PieceDiscount int32 `gorm:"column:piece_discount;type:int(11);" json:"piece_discount"`
	/* 启用预付0-不启用 1-启用预付 */
	EnablePrepay int32 `gorm:"column:enable_prepay;type:tinyint(1);" json:"enable_prepay"`
	/* 预付百分比如1505-表示预付15.05% */
	PrepayRate int32 `gorm:"column:prepay_rate;type:int(11);" json:"prepay_rate"`
	/* 提货有效天数 */
	ExpireDay int32 `gorm:"column:expire_day;type:int(11);" json:"expire_day"`
	/* 商品配送方式：JSON数组，如：[1,2]表示开启1-到店自提，2-快递邮寄 */
	DistributionWay string `gorm:"column:distribution_way;type:json;" json:"distribution_way"`
	/* 参与人数 */
	UserNum int32 `gorm:"column:user_num;type:int(11);" json:"user_num"`
	/* 活动背景色 */
	BgColor string `gorm:"column:bg_color;type:varchar(32);" json:"bg_color"`
	/* 活动背景图片 */
	BgImage string `gorm:"column:bg_image;type:varchar(512);" json:"bg_image"`
	/* 海报背景色 */
	PicBgColor string `gorm:"column:pic_bg_color;type:varchar(32);" json:"pic_bg_color"`
	/* 海报背景图片 */
	PicBgImage string `gorm:"column:pic_bg_image;type:varchar(512);" json:"pic_bg_image"`
	/* 活动规则说明 */
	Rules string `gorm:"column:rules;type:varchar(1000);" json:"rules"`
	/* 推广信息，富文本 */
	PromotionInfo string `gorm:"column:promotion_info;type:text;" json:"promotion_info"`
	/* 背景音乐播放路径 */
	MusicPath string `gorm:"column:music_path;type:varchar(512);" json:"music_path"`
	/* 推广海报path */
	PosterPath string `gorm:"column:poster_path;type:varchar(255);" json:"poster_path"`
	/* 微信分享title */
	WxShareTitle string `gorm:"column:wx_share_title;type:varchar(255);" json:"wx_share_title"`
	/* 微信分享description */
	WxShareDesc string `gorm:"column:wx_share_desc;type:varchar(512);" json:"wx_share_desc"`
	/* 微信分享img path url */
	WxShareImg string `gorm:"column:wx_share_img;type:varchar(255);" json:"wx_share_img"`
	/* 公司名称 */
	CompanyName string `gorm:"column:company_name;type:varchar(255);" json:"company_name"`
	/* 客服电话 */
	KefuTel string `gorm:"column:kefu_tel;type:varchar(64);" json:"kefu_tel"`
	/* 客服微信号 */
	KefuWx string `gorm:"column:kefu_wx;type:varchar(64);" json:"kefu_wx"`
	/* 客服QQ */
	KefuQq string `gorm:"column:kefu_qq;type:varchar(32);" json:"kefu_qq"`
	/* 操作人id */
	OperatorId int32 `gorm:"column:operator_id;type:int(11);" json:"operator_id"`
	/* 操作人姓名 */
	OperatorName string `gorm:"column:operator_name;type:varchar(255);" json:"operator_name"`
	/* 1：预览状态 2：线上状态 */
	Status int32 `gorm:"column:status;type:tinyint(3) unsigned;" json:"status"`
	/* 活动创建时间 */
	CreateDatetime string `gorm:"column:create_datetime;type:timestamp;" json:"create_datetime"`
	/* 活动更新时间 */
	UpdateDatetime string `gorm:"column:update_datetime;type:timestamp;" json:"update_datetime"`
	/* 是否打开在线金价功能 1：是 0：否 */
	OnlineGoldPrice int32 `gorm:"column:online_gold_price;type:tinyint(4);" json:"online_gold_price"`
}

// 单据审核流程表
type OrderCheck struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/* 销售订单id */
	OrderId int32 `gorm:"column:order_id;type:int(11);" json:"order_id"`
	/* 销售单号 */
	OrderSn string `gorm:"column:order_sn;type:varchar(50);" json:"order_sn"`
	/* 商户id */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 门店id */
	MerchantShopId int32 `gorm:"column:merchant_shop_id;type:int(11);" json:"merchant_shop_id"`
	/* 订单商品数量 */
	GoodsNum int32 `gorm:"column:goods_num;type:int(11);" json:"goods_num"`
	/* 订单总金额 单位(分) */
	PayAmount int32 `gorm:"column:pay_amount;type:int(11);" json:"pay_amount"`
	/* 单据类型 1=退货单据,2=撤单单据,3=回收单据删除,4=折扣超限设置 */
	Type int32 `gorm:"column:type;type:tinyint(4);" json:"type"`
	/* 审核状态: 1=审核通过 2=取消申请/废弃 3=待审核 4=审核退回/拒绝 */
	Status int32 `gorm:"column:status;type:tinyint(4);" json:"status"`
	/* 单据申请备注信息 */
	Remark string `gorm:"column:remark;type:varchar(200);" json:"remark"`
	/* 允许审核的人，多个用，隔开 */
	AuditId string `gorm:"column:audit_id;type:varchar(2048);" json:"audit_id"`
	/* 允许审核的人，多个用，隔开 */
	AuditName string `gorm:"column:audit_name;type:varchar(2048);" json:"audit_name"`
	/* 审核人id */
	OperatorId int32 `gorm:"column:operator_id;type:int(11);" json:"operator_id"`
	/* 审核人姓名 */
	OperatorName string `gorm:"column:operator_name;type:varchar(50);" json:"operator_name"`
	/* 订单其他相关信息 */
	Content string `gorm:"column:content;type:json;" json:"content"`
	/* 订单信息 */
	OrderContent string `gorm:"column:order_content;type:json;" json:"order_content"`
	/* 创建时间 */
	CreateTime string `gorm:"column:create_time;type:datetime;" json:"create_time"`
	/* 审核时间 */
	UpdateTime string `gorm:"column:update_time;type:datetime;" json:"update_time"`
	/* 申请人 */
	ApplyId int32 `gorm:"column:apply_id;type:int(11);" json:"apply_id"`
	/* 申请人姓名 */
	ApplyName string `gorm:"column:apply_name;type:varchar(50);" json:"apply_name"`
	/* 审核退回/拒绝备注信息 */
	CheckRemark string `gorm:"column:check_remark;type:varchar(200);" json:"check_remark"`
}

// 企业微信,外部联系人
type QyCustomer struct {
	/* 主键id */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/* 商户id */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 品牌id */
	BrandId int32 `gorm:"column:brand_id;type:int(11);" json:"brand_id"`
	/* 扫码时，确认的门店id */
	MerchantShopId int32 `gorm:"column:merchant_shop_id;type:int(11);" json:"merchant_shop_id"`
	/* 系统用户id(已废弃，查询数据不要通不过这个字段) */
	AdminUserId int32 `gorm:"column:admin_user_id;type:int(11);" json:"admin_user_id"`
	/* 企业id */
	Corpid string `gorm:"column:corpid;type:varchar(50);" json:"corpid"`
	/* 跟进成员的企业微信id */
	QyUserId string `gorm:"column:qy_user_id;type:varchar(50);" json:"qy_user_id"`
	/* 企业微信外部联系人id */
	ExternalUserid string `gorm:"column:external_userid;type:varchar(100);" json:"external_userid"`
	/* 开放平台的unionid */
	Unionid string `gorm:"column:unionid;type:varchar(100);" json:"unionid"`
	/* crm系统中的会员id */
	CustomerId int32 `gorm:"column:customer_id;type:int(11);" json:"customer_id"`
	/* 姓名 */
	Name string `gorm:"column:name;type:varchar(100);" json:"name"`
	/* 头像 */
	Avatar string `gorm:"column:avatar;type:varchar(500);" json:"avatar"`
	/* 类型 1:企业微信用户 2:微信用户 */
	Type int32 `gorm:"column:type;type:tinyint(4);" json:"type"`
	/* 性别: 0-未知 1-男性 2-女性 */
	Gender int32 `gorm:"column:gender;type:tinyint(4);" json:"gender"`
	/* 手机号(如果有多个,号隔开) */
	Mobiles string `gorm:"column:mobiles;type:varchar(255);" json:"mobiles"`
	/* 备注(50字符) */
	Remark string `gorm:"column:remark;type:varchar(100);" json:"remark"`
	/* 描述(最多100字符) */
	Description string `gorm:"column:description;type:varchar(255);" json:"description"`
	/* 存储企业微信标签id */
	Tags string `gorm:"column:tags;type:json;" json:"tags"`
	/* 客户管理 1:双向好友 3:成员删除客户  5:客户删除跟进人员 7:双向删除 */
	Relation int32 `gorm:"column:relation;type:tinyint(4);" json:"relation"`
	/* 企业微信的创建时间(实际内容为时间戳) */
	QyCreatetime int32 `gorm:"column:qy_createtime;type:int(11);" json:"qy_createtime"`
	/* 状态1: 启用 0:删除 (成员主动删除客户) */
	Status int32 `gorm:"column:status;type:tinyint(4);" json:"status"`
	/* 系统创建时间 */
	CreateTime string `gorm:"column:create_time;type:timestamp;" json:"create_time"`
	/* 系统更新时间 */
	UpdateTime string `gorm:"column:update_time;type:timestamp;" json:"update_time"`
	/* 客户删除跟进人时间 */
	DelFollowTime int32 `gorm:"column:del_follow_time;type:int(11);" json:"del_follow_time"`
}

// 企业微信标签
type QyTag struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/* 商户id */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 标签id,外键 */
	TagId int32 `gorm:"column:tag_id;type:int(11);" json:"tag_id"`
	/* 企业id */
	Corpid string `gorm:"column:corpid;type:varchar(50);" json:"corpid"`
	/* 企业微信中的标签id */
	QyTagId string `gorm:"column:qy_tag_id;type:varchar(50);" json:"qy_tag_id"`
	/* 企业微信中的标签id(使用客户联系秘钥获取的) */
	ContractTagId string `gorm:"column:contract_tag_id;type:varchar(50);" json:"contract_tag_id"`
	/* 企业微信的标签组id */
	QyGroupId string `gorm:"column:qy_group_id;type:varchar(50);" json:"qy_group_id"`
	/* 企业微信中的标签组id(使用客户联系秘钥获取的) */
	ContractGroupId string `gorm:"column:contract_group_id;type:varchar(50);" json:"contract_group_id"`
	/* 标签|组名称 */
	Name string `gorm:"column:name;type:varchar(50);" json:"name"`
	/* 类型: 1:标签组 2: 标签  */
	Type int32 `gorm:"column:type;type:tinyint(4);" json:"type"`
	/* 企业标签创建时间 */
	CreateTime int64 `gorm:"column:create_time;type:bigint(20);" json:"create_time"`
}

// 退货订单商品表
type RefundOrderGoods struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/* 退货订单 */
	OrderId int32 `gorm:"column:order_id;type:int(11);" json:"order_id"`
	/* 退货订单编号 */
	OrderSn string `gorm:"column:order_sn;type:varchar(50);" json:"order_sn"`
	/* 商户ID */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 门店ID */
	MerchantShopId int32 `gorm:"column:merchant_shop_id;type:int(11);" json:"merchant_shop_id"`
	/* 门店名称 */
	MerchantShopName string `gorm:"column:merchant_shop_name;type:varchar(128);" json:"merchant_shop_name"`
	/* 联客id */
	UnionCustomerId int32 `gorm:"column:union_customer_id;type:int(11);" json:"union_customer_id"`
	/* 联客姓名 */
	UnionCustomerName string `gorm:"column:union_customer_name;type:varchar(255);" json:"union_customer_name"`
	/* 联客手机号 */
	UnionCustomerTelephone string `gorm:"column:union_customer_telephone;type:varchar(15);" json:"union_customer_telephone"`
	/* 原销售订单order_goods 的id */
	OrderGoodsId int32 `gorm:"column:order_goods_id;type:int(11);" json:"order_goods_id"`
	/* 原销售订单zby_sale_order_old_exchange 的id */
	OrderOldExchangeId int32 `gorm:"column:order_old_exchange_id;type:int(11);" json:"order_old_exchange_id"`
	/* 原销售订单zby_sale_order_prize 的id */
	OrderPrizeId int32 `gorm:"column:order_prize_id;type:int(11);" json:"order_prize_id"`
	/* 订单商品类型1-新品 2 -旧料  3-赠品 */
	OrderGoodsType int32 `gorm:"column:order_goods_type;type:int(11);" json:"order_goods_type"`
	/* 货品编号 */
	StockCode string `gorm:"column:stock_code;type:varchar(128);" json:"stock_code"`
	/* 货品信息 */
	StockInfo string `gorm:"column:stock_info;type:json;" json:"stock_info"`
	/* 原销售单ID  */
	SaleOrderId int32 `gorm:"column:sale_order_id;type:int(11);" json:"sale_order_id"`
	/* 原销售单编号 */
	SaleOrderSn string `gorm:"column:sale_order_sn;type:varchar(50);" json:"sale_order_sn"`
	/* stock_status的id */
	StockId int32 `gorm:"column:stock_id;type:int(11);" json:"stock_id"`
	/* 商品ID */
	GoodsId int32 `gorm:"column:goods_id;type:int(11);" json:"goods_id"`
	/* 迁移前的商品id */
	OldGoodsId int32 `gorm:"column:old_goods_id;type:int(11);" json:"old_goods_id"`
	/* 商品分类id */
	GoodsTypeId int32 `gorm:"column:goods_type_id;type:int(11);" json:"goods_type_id"`
	/* 品牌名称 */
	GoodsTypeName string `gorm:"column:goods_type_name;type:varchar(100);" json:"goods_type_name"`
	/*  */
	GoodsTypeIdOld int32 `gorm:"column:goods_type_id_old;type:int(11);" json:"goods_type_id_old"`
	/* 商品名称 */
	GoodsName string `gorm:"column:goods_name;type:varchar(64);" json:"goods_name"`
	/* 首饰类型ID */
	ClassifyId int32 `gorm:"column:classify_id;type:int(11);" json:"classify_id"`
	/* 商品品牌ID */
	GoodsBrandId int32 `gorm:"column:goods_brand_id;type:int(11);" json:"goods_brand_id"`
	/* 品牌名称 */
	BrandName string `gorm:"column:brand_name;type:varchar(100);" json:"brand_name"`
	/* 柜台ID */
	LibraryId int32 `gorm:"column:library_id;type:int(11);" json:"library_id"`
	/* 商品价格 */
	GoodsPrice int32 `gorm:"column:goods_price;type:int(11);" json:"goods_price"`
	/* 商品数量 */
	GoodsNum int32 `gorm:"column:goods_num;type:int(11);" json:"goods_num"`
	/* 可退货数量 */
	TotalNum int32 `gorm:"column:total_num;type:int(11);" json:"total_num"`
	/* 总优惠金额 */
	TotalDiscountAmount int32 `gorm:"column:total_discount_amount;type:int(11);" json:"total_discount_amount"`
	/* 成色 */
	QualityId int32 `gorm:"column:quality_id;type:int(11);" json:"quality_id"`
	/*  */
	QualityIdOld int32 `gorm:"column:quality_id_old;type:int(11);" json:"quality_id_old"`
	/* 成色名称 */
	QualityName string `gorm:"column:quality_name;type:varchar(64);" json:"quality_name"`
	/* 销售计价方式 */
	SaleChargeType int32 `gorm:"column:sale_charge_type;type:tinyint(1);" json:"sale_charge_type"`
	/* 折旧总额 */
	Earning int32 `gorm:"column:earning;type:int(11);" json:"earning"`
	/* 原商品实收 */
	PayAmount int32 `gorm:"column:pay_amount;type:int(11);" json:"pay_amount"`
	/* 折旧 */
	DepreciationAmount int32 `gorm:"column:depreciation_amount;type:int(11);" json:"depreciation_amount"`
	/* 实退金额 */
	RefundAmount int32 `gorm:"column:refund_amount;type:int(11);" json:"refund_amount"`
	/* 应退 */
	TotalAmount int32 `gorm:"column:total_amount;type:int(11);" json:"total_amount"`
	/* 商品总价，包含工费 */
	GoodsAmount int32 `gorm:"column:goods_amount;type:int(11);" json:"goods_amount"`
	/* 成本金额 */
	CostAmount int32 `gorm:"column:cost_amount;type:int(11);" json:"cost_amount"`
	/* 原销售时间 */
	SaleTime int32 `gorm:"column:sale_time;type:int(11);" json:"sale_time"`
	/* 1-特价 0-非特价商品 */
	OnSale int32 `gorm:"column:on_sale;type:tinyint(4);" json:"on_sale"`
	/* 1-一码一货 2一码多货 */
	StockType int32 `gorm:"column:stock_type;type:int(11);" json:"stock_type"`
	/* 标签价 */
	LabelPrice int32 `gorm:"column:label_price;type:int(11);" json:"label_price"`
	/* 克重 */
	GoldWeight int32 `gorm:"column:gold_weight;type:int(11);" json:"gold_weight"`
	/* 抵扣金额，单位分 */
	DeductAmount int32 `gorm:"column:deduct_amount;type:int(11);" json:"deduct_amount"`
	/* 赠品参与活动名称 */
	ActName string `gorm:"column:act_name;type:varchar(100);" json:"act_name"`
	/* 是否本厂 默认 1-本厂 2-外厂 */
	IsOriginal int32 `gorm:"column:is_original;type:tinyint(1);" json:"is_original"`
	/* 退还使用积分 */
	UseBonus int32 `gorm:"column:use_bonus;type:int(11);" json:"use_bonus"`
	/* 退还使用积分金额 */
	UseBonusAmount int32 `gorm:"column:use_bonus_amount;type:int(11);" json:"use_bonus_amount"`
	/* 素非素 1素金2非素金 */
	GoldType int32 `gorm:"column:gold_type;type:tinyint(4);" json:"gold_type"`
	/* 素金旧料抵扣克重 */
	GoldDeductWeight int32 `gorm:"column:gold_deduct_weight;type:int(11);" json:"gold_deduct_weight"`
	/* 非素金旧料抵扣克重 */
	UngoldDeductWeight int32 `gorm:"column:ungold_deduct_weight;type:int(11);" json:"ungold_deduct_weight"`
	/* 素金旧料抵扣金额 */
	GoldDeductAmount int32 `gorm:"column:gold_deduct_amount;type:int(11);" json:"gold_deduct_amount"`
	/* 非素金旧料抵扣金额 */
	UngoldDeductAmount int32 `gorm:"column:ungold_deduct_amount;type:int(11);" json:"ungold_deduct_amount"`
	/* 创建时间 */
	CreateTime int32 `gorm:"column:create_time;type:int(11);" json:"create_time"`
	/* 修改时间 */
	UpdateTime int32 `gorm:"column:update_time;type:int(11);" json:"update_time"`
	/* 删除时间 */
	DeleteTime int32 `gorm:"column:delete_time;type:int(11);" json:"delete_time"`
	/* 旧料金额 */
	OldAmount int32 `gorm:"column:old_amount;type:int(11);" json:"old_amount"`
}

// 角色用户功能字段配置表
type RoleTemplateField struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/*  */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/*  */
	RoleId int32 `gorm:"column:role_id;type:int(11);" json:"role_id"`
	/*  */
	UserId int32 `gorm:"column:user_id;type:int(11);" json:"user_id"`
	/* 商户权限组 */
	AuthGroupId int32 `gorm:"column:auth_group_id;type:int(11);" json:"auth_group_id"`
	/* 标题名称 */
	ModuleKey string `gorm:"column:module_key;type:varchar(100);" json:"module_key"`
	/* 标题名称 */
	ModuleName string `gorm:"column:module_name;type:varchar(100);" json:"module_name"`
	/* 字段集 */
	FieldInfo string `gorm:"column:field_info;type:json;" json:"field_info"`
	/* 创建时间 */
	CreateTime string `gorm:"column:create_time;type:timestamp;" json:"create_time"`
	/*  */
	UpdateTime string `gorm:"column:update_time;type:timestamp;" json:"update_time"`
	/*  */
	DeleteTime string `gorm:"column:delete_time;type:timestamp;" json:"delete_time"`
}

// 销售单据表
type SaleOrder struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/* 销售单号 */
	OrderSn string `gorm:"column:order_sn;type:varchar(50);" json:"order_sn"`
	/* 订单号后5位 */
	SuffixOrderSn string `gorm:"column:suffix_order_sn;type:varchar(100);" json:"suffix_order_sn"`
	/* 商户ID */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 门店ID */
	MerchantShopId int32 `gorm:"column:merchant_shop_id;type:int(11);" json:"merchant_shop_id"`
	/* 门店名称 */
	MerchantShopName string `gorm:"column:merchant_shop_name;type:varchar(50);" json:"merchant_shop_name"`
	/* 客户ID */
	CustomerId int32 `gorm:"column:customer_id;type:int(11);" json:"customer_id"`
	/* 客户姓名 */
	CustomerName string `gorm:"column:customer_name;type:varchar(50);" json:"customer_name"`
	/* 客户电话 */
	CustomerTelephone string `gorm:"column:customer_telephone;type:varchar(25);" json:"customer_telephone"`
	/* 卡券 id */
	CouponId int32 `gorm:"column:coupon_id;type:int(11);" json:"coupon_id"`
	/* 优惠券主券id */
	MainCouponId int32 `gorm:"column:main_coupon_id;type:int(11);" json:"main_coupon_id"`
	/* 联客的id */
	UnionCustomerId int32 `gorm:"column:union_customer_id;type:int(11);" json:"union_customer_id"`
	/* 联客姓名 */
	UnionCustomerName string `gorm:"column:union_customer_name;type:varchar(50);" json:"union_customer_name"`
	/* 联客手机号 */
	UnionCustomerTelephone string `gorm:"column:union_customer_telephone;type:varchar(15);" json:"union_customer_telephone"`
	/*  应收金额（没减抹零） (错的 要自己算)--订单应收金额取 pay_amount + exchange_old_amount +exchange_new_amount */
	TotalAmount int32 `gorm:"column:total_amount;type:int(11);" json:"total_amount"`
	/* 订单总额（商品总价+工费） */
	GoodsAmount int32 `gorm:"column:goods_amount;type:int(11);" json:"goods_amount"`
	/* 订单总额（商品总价+工费） */
	OrderAmount int32 `gorm:"column:order_amount;type:int(11);" json:"order_amount"`
	/* 销售商品数量 */
	GoodsNum int32 `gorm:"column:goods_num;type:int(11);" json:"goods_num"`
	/* 截金抵扣金额 */
	CutAmount int32 `gorm:"column:cut_amount;type:int(11);" json:"cut_amount"`
	/* 线下优惠券抵扣金额 */
	CouponAmount int32 `gorm:"column:coupon_amount;type:int(11);" json:"coupon_amount"`
	/* 抖音优惠券金额 */
	TiktokCouponAmount int32 `gorm:"column:tiktok_coupon_amount;type:int(11);" json:"tiktok_coupon_amount"`
	/* 抖音json */
	TiktokCouponJson string `gorm:"column:tiktok_coupon_json;type:json;" json:"tiktok_coupon_json"`
	/* 金价+工费折扣优惠 */
	PriceDiscountAmount int32 `gorm:"column:price_discount_amount;type:int(11);" json:"price_discount_amount"`
	/* 收益金金额抵扣 */
	EarningAmount int32 `gorm:"column:earning_amount;type:int(11);" json:"earning_amount"`
	/* 活动折扣金额抵扣，单位：分 */
	ActDiscountAmount int32 `gorm:"column:act_discount_amount;type:int(11);" json:"act_discount_amount"`
	/* 店员折扣金额抵扣 */
	DiscountAmount int32 `gorm:"column:discount_amount;type:int(11);" json:"discount_amount"`
	/* 销售员折扣率 */
	DiscountRate int32 `gorm:"column:discount_rate;type:int(11);" json:"discount_rate"`
	/* 新品换购商品名称 */
	GoodsNewName string `gorm:"column:goods_new_name;type:varchar(1000);" json:"goods_new_name"`
	/* 新品换购抵扣金额 */
	ExchangeNewAmount int32 `gorm:"column:exchange_new_amount;type:int(11);" json:"exchange_new_amount"`
	/* 旧料换购克重 */
	ExchangeNewWeight int32 `gorm:"column:exchange_new_weight;type:int(11);" json:"exchange_new_weight"`
	/* 新品换购数量 */
	ExchangeNewNum int32 `gorm:"column:exchange_new_num;type:int(11);" json:"exchange_new_num"`
	/* 旧料换购名称 */
	GoodsOldName string `gorm:"column:goods_old_name;type:varchar(1000);" json:"goods_old_name"`
	/* 旧料换购抵扣金额 */
	ExchangeOldAmount int32 `gorm:"column:exchange_old_amount;type:int(11);" json:"exchange_old_amount"`
	/* 旧料换购数量 */
	ExchangeOldNum int32 `gorm:"column:exchange_old_num;type:int(11);" json:"exchange_old_num"`
	/* 旧料换购克重 */
	ExchangeOldWeight int32 `gorm:"column:exchange_old_weight;type:int(11);" json:"exchange_old_weight"`
	/* 旧料工费 */
	ExchangeOldCost int32 `gorm:"column:exchange_old_cost;type:int(11);" json:"exchange_old_cost"`
	/* 订金抵扣金额 */
	DepositAmount int32 `gorm:"column:deposit_amount;type:int(11);" json:"deposit_amount"`
	/* 退还定金金额 */
	ReturnDepositAmount int32 `gorm:"column:return_deposit_amount;type:int(11);" json:"return_deposit_amount"`
	/* 抹零金额抵扣 */
	EarseAmount int32 `gorm:"column:earse_amount;type:int(11);" json:"earse_amount"`
	/* 实收总额 */
	PayAmount int32 `gorm:"column:pay_amount;type:int(11);" json:"pay_amount"`
	/* 成本总价 */
	CostAmount int32 `gorm:"column:cost_amount;type:int(11);" json:"cost_amount"`
	/* 销售员ID */
	AdminUserId int32 `gorm:"column:admin_user_id;type:int(11);" json:"admin_user_id"`
	/* 销售员姓名 */
	AdminUserName string `gorm:"column:admin_user_name;type:varchar(50);" json:"admin_user_name"`
	/* 副销售员ID */
	SecondaryerUserId string `gorm:"column:secondaryer_user_id;type:varchar(128);" json:"secondaryer_user_id"`
	/* 副销售员姓名 */
	SecondaryerUserName string `gorm:"column:secondaryer_user_name;type:varchar(128);" json:"secondaryer_user_name"`
	/* 操作员id(收银员) */
	OperatorId int32 `gorm:"column:operator_id;type:int(11);" json:"operator_id"`
	/* 操作员姓名(收银员) */
	OperatorName string `gorm:"column:operator_name;type:varchar(50);" json:"operator_name"`
	/* 商品名称代表 */
	OrderGoodsName string `gorm:"column:order_goods_name;type:varchar(128);" json:"order_goods_name"`
	/* 1-正常销售单;2-新品换购抵消订单; */
	Type int32 `gorm:"column:type;type:tinyint(4);" json:"type"`
	/* 挂单来源 1-开单页; 2-结算页 */
	HangFrom int32 `gorm:"column:hang_from;type:tinyint(4);" json:"hang_from"`
	/* 订单总折扣率，80代表8折 */
	OrderDiscount int32 `gorm:"column:order_discount;type:int(11);" json:"order_discount"`
	/* 订单总折扣金额 */
	OrderDiscountAmount int32 `gorm:"column:order_discount_amount;type:int(11);" json:"order_discount_amount"`
	/* 订单毛利单位分 */
	ProfitAmount int32 `gorm:"column:profit_amount;type:int(11);" json:"profit_amount"`
	/* 是否该用户首购 默认0-非首购 */
	FirstBuy int32 `gorm:"column:first_buy;type:tinyint(4);" json:"first_buy"`
	/* 订单备注 */
	Remarks string `gorm:"column:remarks;type:varchar(500);" json:"remarks"`
	/*  */
	CustomerLevel int32 `gorm:"column:customer_level;type:int(11);" json:"customer_level"`
	/*  */
	CustomerLevelName string `gorm:"column:customer_level_name;type:varchar(25);" json:"customer_level_name"`
	/*  */
	CutWeight int32 `gorm:"column:cut_weight;type:int(11);" json:"cut_weight"`
	/* 截金数 */
	CutNum int32 `gorm:"column:cut_num;type:int(11);" json:"cut_num"`
	/* 单据编号 */
	OrderNumber string `gorm:"column:order_number;type:char(3);" json:"order_number"`
	/* 退货/撤单备注信息 */
	UnRemarks string `gorm:"column:un_remarks;type:varchar(255);" json:"un_remarks"`
	/* 退货/撤单审核人id */
	CheckUserId int32 `gorm:"column:check_user_id;type:int(11);" json:"check_user_id"`
	/* 退货/撤单审核人姓名 */
	CheckUserName string `gorm:"column:check_user_name;type:varchar(50);" json:"check_user_name"`
	/* 班次状态 0=无班次,1=A班,2=B班,3=C班 */
	ShiftStatus int32 `gorm:"column:shift_status;type:tinyint(4);" json:"shift_status"`
	/* 订单来源平台 1-收银端（默认） 2-在线橱窗 */
	Platform int32 `gorm:"column:platform;type:tinyint(4);" json:"platform"`
	/* 下单后获得的积分 */
	GetBonus float64 `gorm:"column:get_bonus;type:decimal(10,2);" json:"get_bonus"`
	/* 下单后赠送的优惠券，保存文本，逗号隔开 */
	GetCoupons string `gorm:"column:get_coupons;type:json;" json:"get_coupons"`
	/* 使用的积分数量 */
	UseBonus float64 `gorm:"column:use_bonus;type:decimal(8,2);" json:"use_bonus"`
	/* 积分抵现金额，分 */
	UseBonusAmount int32 `gorm:"column:use_bonus_amount;type:int(11);" json:"use_bonus_amount"`
	/* 创建时间 */
	CreateTime int32 `gorm:"column:create_time;type:int(11);" json:"create_time"`
	/* 修改时间 */
	UpdateTime int32 `gorm:"column:update_time;type:int(11);" json:"update_time"`
	/* 删除时间 */
	DeleteTime int32 `gorm:"column:delete_time;type:int(11);" json:"delete_time"`
	/* 赠品数量 */
	PrizeNum int32 `gorm:"column:prize_num;type:int(11);" json:"prize_num"`
	/* 提货码 */
	TakeCode string `gorm:"column:take_code;type:varchar(16);" json:"take_code"`
	/* 业绩分配比例 */
	PerformanceSetting string `gorm:"column:performance_setting;type:json;" json:"performance_setting"`
	/* 1-订单完成;2-挂单中;3-收订挂单完成;4-收订挂单取消;5-挂单删除;6-待处理;7-已完成订单删除(消单);8-整单退货 */
	Status int32 `gorm:"column:status;type:tinyint(1);" json:"status"`
	/* 是否是补单 */
	IsChangeDatetime int32 `gorm:"column:is_change_datetime;type:tinyint(4);" json:"is_change_datetime"`
	/*  */
	DepositTime int32 `gorm:"column:deposit_time;type:int(11);" json:"deposit_time"`
	/* 退订时间 */
	DelDepositTime int32 `gorm:"column:del_deposit_time;type:int(11);" json:"del_deposit_time"`
	/* 收定备注 */
	DepositRemarks string `gorm:"column:deposit_remarks;type:varchar(500);" json:"deposit_remarks"`
	/* 退订备注 */
	DelDepositRemarks string `gorm:"column:del_deposit_remarks;type:varchar(500);" json:"del_deposit_remarks"`
	/* 销售时间 */
	SaleDatetime string `gorm:"column:sale_datetime;type:timestamp;" json:"sale_datetime"`
	/* 操作人id（如删单操作) */
	OrderOperatorId int32 `gorm:"column:order_operator_id;type:int(11);" json:"order_operator_id"`
	/* 操作人姓名(如删单操作) */
	OrderOperatorName string `gorm:"column:order_operator_name;type:varchar(50);" json:"order_operator_name"`
	/* 删单时间 */
	DeleteOrderTime string `gorm:"column:delete_order_time;type:timestamp;" json:"delete_order_time"`
	/* 联客积分 */
	UnionBonus float64 `gorm:"column:union_bonus;type:decimal(10,2);" json:"union_bonus"`
	/* 该单是否属于 先抵后折 1 是 2 否 */
	IsOldDiscount int32 `gorm:"column:is_old_discount;type:tinyint(1);" json:"is_old_discount"`
	/* 所有支付id  */
	PayIds string `gorm:"column:pay_ids;type:varchar(100);" json:"pay_ids"`
}

// 销售单据商品明细表
type SaleOrderGoods struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/* 销售单号ID  */
	OrderId int32 `gorm:"column:order_id;type:int(11);" json:"order_id"`
	/* 销售单号 */
	OrderSn string `gorm:"column:order_sn;type:varchar(50);" json:"order_sn"`
	/* 商户ID */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 门店ID */
	MerchantShopId int32 `gorm:"column:merchant_shop_id;type:int(11);" json:"merchant_shop_id"`
	/* 门店名称 */
	MerchantShopName string `gorm:"column:merchant_shop_name;type:varchar(50);" json:"merchant_shop_name"`
	/* 销售员ID */
	AdminUserId int32 `gorm:"column:admin_user_id;type:int(11);" json:"admin_user_id"`
	/* 销售员名称 */
	AdminUserName string `gorm:"column:admin_user_name;type:varchar(255);" json:"admin_user_name"`
	/* 客户ID */
	CustomerId int32 `gorm:"column:customer_id;type:int(11);" json:"customer_id"`
	/* 客户姓名 */
	CustomerName string `gorm:"column:customer_name;type:varchar(50);" json:"customer_name"`
	/* 客户电话 */
	CustomerTelephone string `gorm:"column:customer_telephone;type:varchar(25);" json:"customer_telephone"`
	/* 优惠券id */
	CouponId int32 `gorm:"column:coupon_id;type:int(11);" json:"coupon_id"`
	/*  */
	CouponSn string `gorm:"column:coupon_sn;type:varchar(64);" json:"coupon_sn"`
	/* stock_status表的ID */
	StockId int32 `gorm:"column:stock_id;type:int(11);" json:"stock_id"`
	/* 商品ID */
	GoodsId int32 `gorm:"column:goods_id;type:int(11);" json:"goods_id"`
	/* 迁移前的商品id */
	OldGoodsId int32 `gorm:"column:old_goods_id;type:int(11);" json:"old_goods_id"`
	/* 商品条码 */
	StockCode string `gorm:"column:stock_code;type:varchar(128);" json:"stock_code"`
	/* 证书编号 */
	CerNumber string `gorm:"column:cer_number;type:varchar(254);" json:"cer_number"`
	/* 商品名称 */
	GoodsName string `gorm:"column:goods_name;type:varchar(50);" json:"goods_name"`
	/* 标记 */
	Mark string `gorm:"column:mark;type:char(5);" json:"mark"`
	/* 商品分类id */
	GoodsTypeId int32 `gorm:"column:goods_type_id;type:int(11);" json:"goods_type_id"`
	/*  */
	GoodsTypeIdOld int32 `gorm:"column:goods_type_id_old;type:int(11);" json:"goods_type_id_old"`
	/* 商品分类名称 */
	GoodsTypeName string `gorm:"column:goods_type_name;type:varchar(255);" json:"goods_type_name"`
	/* 货品位置id */
	LibraryId int32 `gorm:"column:library_id;type:int(11);" json:"library_id"`
	/*  */
	LibraryIdOld int32 `gorm:"column:library_id_old;type:int(11);" json:"library_id_old"`
	/* 货品位置名称 */
	LibraryName string `gorm:"column:library_name;type:varchar(100);" json:"library_name"`
	/* 品牌id */
	BrandId int32 `gorm:"column:brand_id;type:int(11);" json:"brand_id"`
	/*  */
	BrandIdOld int32 `gorm:"column:brand_id_old;type:int(11);" json:"brand_id_old"`
	/* 品牌名称 */
	BrandName string `gorm:"column:brand_name;type:varchar(100);" json:"brand_name"`
	/* 销售来源，0：全部，1：线上商城 2： 线下门店3：营销活动 */
	SourceType int32 `gorm:"column:source_type;type:tinyint(1);" json:"source_type"`
	/* 首饰分类id */
	ClassifyId int32 `gorm:"column:classify_id;type:int(11);" json:"classify_id"`
	/*  */
	ClassifyIdOld int32 `gorm:"column:classify_id_old;type:int(11);" json:"classify_id_old"`
	/* 首饰分类名称 */
	ClassifyName string `gorm:"column:classify_name;type:varchar(100);" json:"classify_name"`
	/* 销售金价类型 */
	GoldSalePriceType int32 `gorm:"column:gold_sale_price_type;type:int(11);" json:"gold_sale_price_type"`
	/* 使用金价 单位：分 */
	GoldPrice int32 `gorm:"column:gold_price;type:int(11);" json:"gold_price"`
	/* 修改后的金价 */
	NewGoldPrice int32 `gorm:"column:new_gold_price;type:int(11);" json:"new_gold_price"`
	/* 净金重 */
	GoldWeight int32 `gorm:"column:gold_weight;type:int(11);" json:"gold_weight"`
	/* 总件重 */
	Weight int32 `gorm:"column:weight;type:int(11);" json:"weight"`
	/* 成色id */
	QualityId int32 `gorm:"column:quality_id;type:int(11);" json:"quality_id"`
	/*  */
	QualityIdOld int32 `gorm:"column:quality_id_old;type:int(11);" json:"quality_id_old"`
	/* 成色名称 */
	QualityName string `gorm:"column:quality_name;type:varchar(64);" json:"quality_name"`
	/* 1-标签价不含工费(默认) 2-标签价含工费 */
	SaleLabourType int32 `gorm:"column:sale_labour_type;type:tinyint(1);" json:"sale_labour_type"`
	/* 销售计价类型 1-按克 2-按建 */
	ChargeType int32 `gorm:"column:charge_type;type:tinyint(1);" json:"charge_type"`
	/* 成本计价方式 */
	CostChargeType int32 `gorm:"column:cost_charge_type;type:tinyint(1);" json:"cost_charge_type"`
	/* 销售工费计价类型 1-按克 2-按件 */
	LabourChargeType int32 `gorm:"column:labour_charge_type;type:tinyint(1);" json:"labour_charge_type"`
	/* 销售工费单价 */
	SaleLabourUnitPrice int32 `gorm:"column:sale_labour_unit_price;type:int(11);" json:"sale_labour_unit_price"`
	/* 修改后工费单价 */
	NewSaleLabourUnitPrice int32 `gorm:"column:new_sale_labour_unit_price;type:int(11);" json:"new_sale_labour_unit_price"`
	/* 购买数量 */
	GoodsNum int32 `gorm:"column:goods_num;type:int(11);" json:"goods_num"`
	/* 被退货数量 */
	RefundNum int32 `gorm:"column:refund_num;type:int(11);" json:"refund_num"`
	/* 被换购数量 */
	ExchangeNum int32 `gorm:"column:exchange_num;type:int(11);" json:"exchange_num"`
	/* 标签价 单位：分  不包括工费 */
	LabelPrice int32 `gorm:"column:label_price;type:int(11);" json:"label_price"`
	/* 应收金额，抹零前金额(错的）pay_amount+old_amount+new_amount */
	TotalAmount int32 `gorm:"column:total_amount;type:int(11);" json:"total_amount"`
	/* 商品总价，单价*数量+工费 */
	GoodsAmount int32 `gorm:"column:goods_amount;type:int(11);" json:"goods_amount"`
	/* 旧料金额 */
	OldAmount int32 `gorm:"column:old_amount;type:int(11);" json:"old_amount"`
	/* 减旧差值 */
	OldDiff int32 `gorm:"column:old_diff;type:int(11);" json:"old_diff"`
	/* 新品换购金额 */
	NewAmount int32 `gorm:"column:new_amount;type:int(11);" json:"new_amount"`
	/* 截金克重 mg */
	CutWeight int32 `gorm:"column:cut_weight;type:int(11);" json:"cut_weight"`
	/* 截金折扣金额 */
	CutAmount int32 `gorm:"column:cut_amount;type:int(11);" json:"cut_amount"`
	/* 截金预估毛利 */
	CutOldProfitAmount int32 `gorm:"column:cut_old_profit_amount;type:int(11);" json:"cut_old_profit_amount"`
	/* 截金新品成本 */
	CutCostAmount int32 `gorm:"column:cut_cost_amount;type:int(11);" json:"cut_cost_amount"`
	/* 优惠券抵扣金额 单位：分 */
	CouponAmount int32 `gorm:"column:coupon_amount;type:int(11);" json:"coupon_amount"`
	/* 收益金抵扣金额 */
	EarningAmount int32 `gorm:"column:earning_amount;type:int(11);" json:"earning_amount"`
	/* 商品活动折扣率,单位% */
	ActDiscount float64 `gorm:"column:act_discount;type:decimal(8,2);" json:"act_discount"`
	/* act_discount 初始值 */
	ActDiscountSource string `gorm:"column:act_discount_source;type:varchar(128);" json:"act_discount_source"`
	/* 活动折扣抵扣金额 */
	ActDiscountAmount int32 `gorm:"column:act_discount_amount;type:int(11);" json:"act_discount_amount"`
	/* 参与的活动折扣ids 多个用,隔开 */
	ActDiscountIds string `gorm:"column:act_discount_ids;type:varchar(255);" json:"act_discount_ids"`
	/* 折扣活动工费优惠金额 */
	ActLabourDiscountAmount int32 `gorm:"column:act_labour_discount_amount;type:int(11);" json:"act_labour_discount_amount"`
	/* 折扣活动金价优惠金额 */
	ActPriceDiscountAmount int32 `gorm:"column:act_price_discount_amount;type:int(11);" json:"act_price_discount_amount"`
	/* 原始最低折扣 */
	HistoryDiscount float64 `gorm:"column:history_discount;type:decimal(8,4);" json:"history_discount"`
	/* 店员折扣率,单位% */
	Discount float64 `gorm:"column:discount;type:decimal(8,4);" json:"discount"`
	/* 店员折扣抵扣金额 */
	DiscountAmount int32 `gorm:"column:discount_amount;type:int(11);" json:"discount_amount"`
	/* 抹零抵扣金额 */
	EarseAmount int32 `gorm:"column:earse_amount;type:int(11);" json:"earse_amount"`
	/* 实收金额，这里如果需要算实收金额需要减去旧料和新品换购的金额 */
	PayAmount int32 `gorm:"column:pay_amount;type:int(11);" json:"pay_amount"`
	/* 成本单价，=cost_amount/goods_num */
	CostPrice int32 `gorm:"column:cost_price;type:int(11);" json:"cost_price"`
	/* 成本价 */
	CostAmount int32 `gorm:"column:cost_amount;type:int(11);" json:"cost_amount"`
	/* 平均成本金价 按克商品 单位分 */
	AvgCostPrice int32 `gorm:"column:avg_cost_price;type:int(11);" json:"avg_cost_price"`
	/* 金价及工费优惠 */
	PriceDiscountAmount int32 `gorm:"column:price_discount_amount;type:int(11);" json:"price_discount_amount"`
	/* 工费折扣优惠金额 */
	LabourDiscountAmount int32 `gorm:"column:labour_discount_amount;type:int(11);" json:"labour_discount_amount"`
	/* 商品图片地址 */
	PicUrl string `gorm:"column:pic_url;type:varchar(256);" json:"pic_url"`
	/* 是否特价 1-特价 0-非特价(默认) */
	OnSale int32 `gorm:"column:on_sale;type:tinyint(4);" json:"on_sale"`
	/* 1-开单页挂单 2-结算页挂单 */
	HangFrom int32 `gorm:"column:hang_from;type:tinyint(4);" json:"hang_from"`
	/*  */
	TotalNum int32 `gorm:"column:total_num;type:int(11);" json:"total_num"`
	/* 货品周转时间，单位天 */
	CycleTime int32 `gorm:"column:cycle_time;type:int(11);" json:"cycle_time"`
	/* 1-一码一货 2一码多货 */
	StockType int32 `gorm:"column:stock_type;type:tinyint(1);" json:"stock_type"`
	/* 销售工费总价 */
	SaleLabourPrice int32 `gorm:"column:sale_labour_price;type:int(11);" json:"sale_labour_price"`
	/* 修改后的销售工费总价 */
	NewSaleLabourPrice int32 `gorm:"column:new_sale_labour_price;type:int(11);" json:"new_sale_labour_price"`
	/* 主石重 */
	MainStoneWeight int32 `gorm:"column:main_stone_weight;type:int(11);" json:"main_stone_weight"`
	/*  */
	GoodsRemark string `gorm:"column:goods_remark;type:varchar(255);" json:"goods_remark"`
	/* 1 收银端 2 在线橱窗 */
	Platform int32 `gorm:"column:platform;type:tinyint(4);" json:"platform"`
	/* 下单后获得的积分 */
	GetBonus float64 `gorm:"column:get_bonus;type:decimal(8,2);" json:"get_bonus"`
	/* 使用的积分数量 */
	UseBonus float64 `gorm:"column:use_bonus;type:decimal(8,2);" json:"use_bonus"`
	/* 积分抵现金额，分 */
	UseBonusAmount int32 `gorm:"column:use_bonus_amount;type:int(11);" json:"use_bonus_amount"`
	/* 素非素 1素金2非素金 */
	GoldType int32 `gorm:"column:gold_type;type:tinyint(4);" json:"gold_type"`
	/* 商品销售状态,0-待处理,1-售出2-锁定3-退货,4-换购,5-已删除(消单),6-旧料换购,7-旧料回收 */
	Status int32 `gorm:"column:status;type:tinyint(1);" json:"status"`
	/* 销售时间 */
	SaleDatetime string `gorm:"column:sale_datetime;type:timestamp;" json:"sale_datetime"`
	/* 员工提成金额 */
	CommissionAmount int32 `gorm:"column:commission_amount;type:int(11);" json:"commission_amount"`
	/* 修改后员工提成金额 */
	CommissionAmountModify int32 `gorm:"column:commission_amount_modify;type:int(11);" json:"commission_amount_modify"`
	/* 入库天数 */
	StayDays int32 `gorm:"column:stay_days;type:int(11);" json:"stay_days"`
	/* 调拨到库天数 */
	AllocationDays int32 `gorm:"column:allocation_days;type:int(11);" json:"allocation_days"`
	/* 销售小备注 */
	Remarks string `gorm:"column:remarks;type:varchar(500);" json:"remarks"`
	/* 商品备注 1 */
	Remark1 string `gorm:"column:remark_1;type:varchar(255);" json:"remark_1"`
	/* 商品备注 2 */
	Remark2 string `gorm:"column:remark_2;type:varchar(255);" json:"remark_2"`
	/* 商品备注 3 */
	Remark3 string `gorm:"column:remark_3;type:varchar(255);" json:"remark_3"`
	/* 创建时间 */
	CreateTime int32 `gorm:"column:create_time;type:int(11);" json:"create_time"`
	/* 修改时间 */
	UpdateTime int32 `gorm:"column:update_time;type:int(11);" json:"update_time"`
	/* 删除时间 */
	DeleteTime int32 `gorm:"column:delete_time;type:int(11);" json:"delete_time"`
	/* 折扣超限审核（限制优惠） */
	AuditLimitData string `gorm:"column:audit_limit_data;type:varchar(128);" json:"audit_limit_data"`
	/* 折扣超限审核（限制优惠类型） */
	AuditChargeType string `gorm:"column:audit_charge_type;type:varchar(128);" json:"audit_charge_type"`
	/* 是否与旧料进行绑定 0：无  1：是 */
	IsBindOldStock int32 `gorm:"column:is_bind_old_stock;type:tinyint(4);" json:"is_bind_old_stock"`
	/* 主石净度 */
	MainStonePurityName string `gorm:"column:main_stone_purity_name;type:varchar(255);" json:"main_stone_purity_name"`
	/* 主石颜色 */
	MainStoneColorName string `gorm:"column:main_stone_color_name;type:varchar(255);" json:"main_stone_color_name"`
	/* 主石切工 */
	MainStoneSectionName string `gorm:"column:main_stone_section_name;type:varchar(255);" json:"main_stone_section_name"`
	/* 主石形状 */
	MainStoneShapeName string `gorm:"column:main_stone_shape_name;type:varchar(255);" json:"main_stone_shape_name"`
	/* 商品工艺 */
	TechniqueName string `gorm:"column:technique_name;type:varchar(255);" json:"technique_name"`
	/* 手寸 */
	HandSize string `gorm:"column:hand_size;type:varchar(255);" json:"hand_size"`
	/* 商品标记id(旧数据没有.) */
	GoodsTagId int32 `gorm:"column:goods_tag_id;type:int(11);" json:"goods_tag_id"`
	/* 商品标记id数组 */
	GoodsTagIds string `gorm:"column:goods_tag_ids;type:varchar(255);" json:"goods_tag_ids"`
	/* 商品标记名称 */
	GoodsTagName string `gorm:"column:goods_tag_name;type:varchar(255);" json:"goods_tag_name"`
	/* 素金旧料抵扣克重 */
	GoldDeductWeight int32 `gorm:"column:gold_deduct_weight;type:int(11);" json:"gold_deduct_weight"`
	/* 非素金旧料抵扣克重 */
	UngoldDeductWeight int32 `gorm:"column:ungold_deduct_weight;type:int(11);" json:"ungold_deduct_weight"`
	/* 素金旧料抵扣金额 */
	GoldDeductAmount int32 `gorm:"column:gold_deduct_amount;type:int(11);" json:"gold_deduct_amount"`
	/* 非素金旧料抵扣金额 */
	UngoldDeductAmount int32 `gorm:"column:ungold_deduct_amount;type:int(11);" json:"ungold_deduct_amount"`
	/* 非素金旧料抵扣金额的利润 */
	UngoldDeductProfitAmount int32 `gorm:"column:ungold_deduct_profit_amount;type:int(11);" json:"ungold_deduct_profit_amount"`
	/* 素金旧料抵扣金额的利润 */
	GoldDeductProfitAmount int32 `gorm:"column:gold_deduct_profit_amount;type:int(11);" json:"gold_deduct_profit_amount"`
	/* 电子优惠券抵扣金额 单位：分 */
	OnlineCouponAmount int32 `gorm:"column:online_coupon_amount;type:int(11);" json:"online_coupon_amount"`
	/* 纸质优惠券抵扣金额 单位：分 */
	OfflineCouponAmount int32 `gorm:"column:offline_coupon_amount;type:int(11);" json:"offline_coupon_amount"`
	/* 排序 */
	Sort int32 `gorm:"column:sort;type:tinyint(4);" json:"sort"`
	/* 折扣活动配置 */
	ActSetting []byte `gorm:"column:act_setting;type:json;" json:"act_setting"`
	/* 所有活动的积分金额 */
	TotalActBonusAmount int32 `gorm:"column:total_act_bonus_amount;type:int(11);" json:"total_act_bonus_amount"`
	/* 所有活动的业绩金额 */
	TotalActPerformanceAmount int32 `gorm:"column:total_act_performance_amount;type:int(11);" json:"total_act_performance_amount"`
	/* 抖音抵扣金额 */
	TiktokCouponAmount int32 `gorm:"column:tiktok_coupon_amount;type:int(11);" json:"tiktok_coupon_amount"`
	/* 一码多货克管理：2.否、1.是 */
	StockTypeVariedWeight int32 `gorm:"column:stock_type_varied_weight;type:tinyint(1);" json:"stock_type_varied_weight"`
	/* 财务成本单价 */
	FinancialCostPrice int32 `gorm:"column:financial_cost_price;type:int(11);" json:"financial_cost_price"`
	/* 市场成本单价 */
	MarketCostPrice int32 `gorm:"column:market_cost_price;type:int(11);" json:"market_cost_price"`
	/* 财务成本总价 */
	FinancialCostAmount int32 `gorm:"column:financial_cost_amount;type:int(11);" json:"financial_cost_amount"`
	/* 市场成本总价 */
	MarketCostAmount int32 `gorm:"column:market_cost_amount;type:int(11);" json:"market_cost_amount"`
	/* 财务截后成本总价 */
	FinancialCutCostAmount int32 `gorm:"column:financial_cut_cost_amount;type:int(11);" json:"financial_cut_cost_amount"`
	/* 市场截后成本总价 */
	MarketCutCostAmount int32 `gorm:"column:market_cut_cost_amount;type:int(11);" json:"market_cut_cost_amount"`
	/* 凭证图片字段 */
	CertificateList []byte `gorm:"column:certificate_list;type:json;" json:"certificate_list"`
}

// 历史销售单据表
type SaleOrderHistory struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/* 销售单号 */
	OrderSn string `gorm:"column:order_sn;type:varchar(50);" json:"order_sn"`
	/* 订单号后5位 */
	SuffixOrderSn string `gorm:"column:suffix_order_sn;type:varchar(100);" json:"suffix_order_sn"`
	/* 商户ID */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 门店ID */
	MerchantShopId int32 `gorm:"column:merchant_shop_id;type:int(11);" json:"merchant_shop_id"`
	/* 门店名称 */
	MerchantShopName string `gorm:"column:merchant_shop_name;type:varchar(50);" json:"merchant_shop_name"`
	/* 客户ID */
	CustomerId int32 `gorm:"column:customer_id;type:int(11);" json:"customer_id"`
	/* 客户姓名 */
	CustomerName string `gorm:"column:customer_name;type:varchar(50);" json:"customer_name"`
	/* 客户电话 */
	CustomerTelephone string `gorm:"column:customer_telephone;type:varchar(25);" json:"customer_telephone"`
	/* 商品条码 */
	StockCode string `gorm:"column:stock_code;type:varchar(50);" json:"stock_code"`
	/* 商品名称 */
	GoodsName string `gorm:"column:goods_name;type:varchar(100);" json:"goods_name"`
	/*  */
	GoodsTypeName string `gorm:"column:goods_type_name;type:varchar(100);" json:"goods_type_name"`
	/*  */
	BrandName string `gorm:"column:brand_name;type:varchar(100);" json:"brand_name"`
	/* 销售商品数量 */
	GoodsNum int32 `gorm:"column:goods_num;type:int(11);" json:"goods_num"`
	/* 净金重 */
	GoodsGoldWeight float64 `gorm:"column:goods_gold_weight;type:decimal(11,3);" json:"goods_gold_weight"`
	/* 销售计价方式 按克 / 按件 */
	SaleChargeType string `gorm:"column:sale_charge_type;type:varchar(10);" json:"sale_charge_type"`
	/* 新品/旧料商品名称 */
	GoodsDeductName string `gorm:"column:goods_deduct_name;type:varchar(100);" json:"goods_deduct_name"`
	/* 新品/旧料回收金价 */
	GoodsDeductRecycleGoldPrice float64 `gorm:"column:goods_deduct_recycle_gold_price;type:decimal(11,3);" json:"goods_deduct_recycle_gold_price"`
	/* 新品/旧料抵扣金重 */
	GoodsDeductGoldWeight float64 `gorm:"column:goods_deduct_gold_weight;type:decimal(11,3);" json:"goods_deduct_gold_weight"`
	/* 新品/旧料抵扣金额 */
	GoodsDeductAmount float64 `gorm:"column:goods_deduct_amount;type:decimal(11,2);" json:"goods_deduct_amount"`
	/* 标签价 */
	GoodsLabelPrice float64 `gorm:"column:goods_label_price;type:decimal(11,2);" json:"goods_label_price"`
	/* 成本价 */
	CostAmount float64 `gorm:"column:cost_amount;type:decimal(11,2);" json:"cost_amount"`
	/* 折后金额 */
	GoodsDiscountAmount float64 `gorm:"column:goods_discount_amount;type:decimal(11,2);" json:"goods_discount_amount"`
	/* 订单总额 */
	OrderGoodsAmount float64 `gorm:"column:order_goods_amount;type:decimal(11,2);" json:"order_goods_amount"`
	/* 订单优惠总额 */
	OrderCouponAmount float64 `gorm:"column:order_coupon_amount;type:decimal(11,2);" json:"order_coupon_amount"`
	/* 订单抵扣总额 */
	OrderDeductAmount float64 `gorm:"column:order_deduct_amount;type:decimal(11,2);" json:"order_deduct_amount"`
	/* 订单应收金额 */
	OrderAmount float64 `gorm:"column:order_amount;type:decimal(11,2);" json:"order_amount"`
	/* 订单实收金额 */
	OrderPayAmount float64 `gorm:"column:order_pay_amount;type:decimal(11,2);" json:"order_pay_amount"`
	/* 付款方式 */
	OrderPayInfo string `gorm:"column:order_pay_info;type:varchar(255);" json:"order_pay_info"`
	/* 销售员姓名 */
	SalerName string `gorm:"column:saler_name;type:varchar(100);" json:"saler_name"`
	/* 销售时间 */
	SaleDatetime string `gorm:"column:sale_datetime;type:timestamp;" json:"sale_datetime"`
	/* 备注 */
	Remark string `gorm:"column:remark;type:varchar(255);" json:"remark"`
	/* 备注1 */
	Remark1 string `gorm:"column:remark1;type:varchar(255);" json:"remark1"`
	/* 备注2 */
	Remark2 string `gorm:"column:remark2;type:varchar(255);" json:"remark2"`
	/* 备注3 */
	Remark3 string `gorm:"column:remark3;type:varchar(255);" json:"remark3"`
	/* 创建时间 */
	CreateTime int32 `gorm:"column:create_time;type:int(11);" json:"create_time"`
	/* 更新时间 */
	UpdateTime int32 `gorm:"column:update_time;type:int(11);" json:"update_time"`
	/* 导入员工id */
	AdminUserId int32 `gorm:"column:admin_user_id;type:int(11);" json:"admin_user_id"`
	/* 导入员工姓名 */
	AdminUserName string `gorm:"column:admin_user_name;type:varchar(100);" json:"admin_user_name"`
	/* 删除时间 */
	DeleteTime int32 `gorm:"column:delete_time;type:int(10) unsigned;" json:"delete_time"`
	/* 导入记录id */
	RecordId int32 `gorm:"column:record_id;type:int(11);" json:"record_id"`
}

// 新品换购订单信息
type SaleOrderNewExchange struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/* 原订单商品ID */
	OrderGoodsId int32 `gorm:"column:order_goods_id;type:int(11);" json:"order_goods_id"`
	/* 销售单号ID  */
	OrderId int32 `gorm:"column:order_id;type:int(11);" json:"order_id"`
	/* 销售单号 */
	OrderSn string `gorm:"column:order_sn;type:varchar(50);" json:"order_sn"`
	/* 商户ID */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 门店ID */
	MerchantShopId int32 `gorm:"column:merchant_shop_id;type:int(11);" json:"merchant_shop_id"`
	/* 原销售ID  */
	SaleOrderId int32 `gorm:"column:sale_order_id;type:int(11);" json:"sale_order_id"`
	/* 原销售单号 */
	SaleOrderSn string `gorm:"column:sale_order_sn;type:varchar(128);" json:"sale_order_sn"`
	/* 商品分类id */
	GoodsTypeId int32 `gorm:"column:goods_type_id;type:int(11);" json:"goods_type_id"`
	/*  */
	GoodsTypeIdOld int32 `gorm:"column:goods_type_id_old;type:int(11);" json:"goods_type_id_old"`
	/* 商品分类名称 */
	GoodsTypeName string `gorm:"column:goods_type_name;type:varchar(255);" json:"goods_type_name"`
	/* 商品标记id数组 */
	GoodsTagIds string `gorm:"column:goods_tag_ids;type:varchar(255);" json:"goods_tag_ids"`
	/* 商品标记名称数组 */
	GoodsTagName string `gorm:"column:goods_tag_name;type:varchar(255);" json:"goods_tag_name"`
	/* 素非素 1素金2非素金 */
	GoldType int32 `gorm:"column:gold_type;type:tinyint(4);" json:"gold_type"`
	/* 首饰分类id */
	ClassifyId int32 `gorm:"column:classify_id;type:int(11);" json:"classify_id"`
	/*  */
	ClassifyIdOld int32 `gorm:"column:classify_id_old;type:int(11);" json:"classify_id_old"`
	/* 首饰分类名称 */
	ClassifyName string `gorm:"column:classify_name;type:varchar(254);" json:"classify_name"`
	/* 柜台ID */
	LibraryId int32 `gorm:"column:library_id;type:int(11);" json:"library_id"`
	/* 商品品牌ID */
	GoodsBrandId int32 `gorm:"column:goods_brand_id;type:int(11);" json:"goods_brand_id"`
	/* 品牌名称 */
	BrandName string `gorm:"column:brand_name;type:varchar(100);" json:"brand_name"`
	/* 1-一码一货 2一码多货 */
	StockType int32 `gorm:"column:stock_type;type:tinyint(4);" json:"stock_type"`
	/* 商品编号 */
	StockCode string `gorm:"column:stock_code;type:varchar(128);" json:"stock_code"`
	/* 商品信息 */
	StockInfo string `gorm:"column:stock_info;type:json;" json:"stock_info"`
	/* 商品ID */
	GoodsId int32 `gorm:"column:goods_id;type:int(11);" json:"goods_id"`
	/* 迁移前的商品id */
	OldGoodsId int32 `gorm:"column:old_goods_id;type:int(11);" json:"old_goods_id"`
	/* 商品数量 */
	GoodsNum int32 `gorm:"column:goods_num;type:int(11);" json:"goods_num"`
	/* 商品名称 */
	GoodsName string `gorm:"column:goods_name;type:varchar(128);" json:"goods_name"`
	/* 销售计价方式  */
	SaleChargeType int32 `gorm:"column:sale_charge_type;type:tinyint(1);" json:"sale_charge_type"`
	/* 净金重 */
	GoldWeight int32 `gorm:"column:gold_weight;type:int(11);" json:"gold_weight"`
	/* 标签价 */
	LabelAmount int32 `gorm:"column:label_amount;type:int(11);" json:"label_amount"`
	/*  */
	TotalAmount int32 `gorm:"column:total_amount;type:int(11);" json:"total_amount"`
	/*  */
	GoodsAmount int32 `gorm:"column:goods_amount;type:int(11);" json:"goods_amount"`
	/* 实收金额 */
	PayAmount int32 `gorm:"column:pay_amount;type:int(11);" json:"pay_amount"`
	/* 成本价 */
	CostAmount int32 `gorm:"column:cost_amount;type:int(11);" json:"cost_amount"`
	/* 折旧费用 */
	DepreAmount int32 `gorm:"column:depre_amount;type:int(11);" json:"depre_amount"`
	/* 抵扣金额 */
	DeductAmount int32 `gorm:"column:deduct_amount;type:int(11);" json:"deduct_amount"`
	/* 0-待处理1-已换购2-已退货3-已删除4-逻辑删除 */
	Status int32 `gorm:"column:status;type:tinyint(4);" json:"status"`
	/* 创建时间 */
	CreateTime int32 `gorm:"column:create_time;type:int(11);" json:"create_time"`
	/* 修改时间 */
	UpdateTime int32 `gorm:"column:update_time;type:int(11);" json:"update_time"`
	/* 删除时间 */
	DeleteTime int32 `gorm:"column:delete_time;type:int(11);" json:"delete_time"`
	/* 主石净度 */
	MainStonePurityName string `gorm:"column:main_stone_purity_name;type:varchar(255);" json:"main_stone_purity_name"`
	/* 主石颜色 */
	MainStoneColorName string `gorm:"column:main_stone_color_name;type:varchar(255);" json:"main_stone_color_name"`
	/* 主石切工 */
	MainStoneSectionName string `gorm:"column:main_stone_section_name;type:varchar(255);" json:"main_stone_section_name"`
	/* 主石形状 */
	MainStoneShapeName string `gorm:"column:main_stone_shape_name;type:varchar(255);" json:"main_stone_shape_name"`
	/* 商品工艺 */
	TechniqueName string `gorm:"column:technique_name;type:varchar(255);" json:"technique_name"`
	/* 手寸 */
	HandSize string `gorm:"column:hand_size;type:varchar(255);" json:"hand_size"`
}

// 销售订单旧换新商品表
type SaleOrderOldExchange struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/*  */
	OrderId int32 `gorm:"column:order_id;type:int(11);" json:"order_id"`
	/*  */
	OrderSn string `gorm:"column:order_sn;type:varchar(254);" json:"order_sn"`
	/* 商品分类 */
	GoodsTypeId string `gorm:"column:goods_type_id;type:varchar(15);" json:"goods_type_id"`
	/*  */
	GoodsTypeIdOld string `gorm:"column:goods_type_id_old;type:varchar(15);" json:"goods_type_id_old"`
	/* 分类名称 */
	GoodsTypeName string `gorm:"column:goods_type_name;type:varchar(15);" json:"goods_type_name"`
	/* 商户ID */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 门店ID */
	MerchantShopId int32 `gorm:"column:merchant_shop_id;type:int(11);" json:"merchant_shop_id"`
	/*  */
	CustomerId int32 `gorm:"column:customer_id;type:int(11);" json:"customer_id"`
	/*  */
	CustomerName string `gorm:"column:customer_name;type:varchar(255);" json:"customer_name"`
	/*  */
	CustomerTelephone int32 `gorm:"column:customer_telephone;type:int(11);" json:"customer_telephone"`
	/* 销售员id */
	AdminUserId int32 `gorm:"column:admin_user_id;type:int(11);" json:"admin_user_id"`
	/* 销售员名称 */
	AdminUserName string `gorm:"column:admin_user_name;type:varchar(20);" json:"admin_user_name"`
	/*  */
	BrandId int32 `gorm:"column:brand_id;type:int(11);" json:"brand_id"`
	/*  */
	BrandIdOld int32 `gorm:"column:brand_id_old;type:int(11);" json:"brand_id_old"`
	/*  */
	BrandName string `gorm:"column:brand_name;type:varchar(255);" json:"brand_name"`
	/* 商品类别 */
	ClassifyId int32 `gorm:"column:classify_id;type:int(11);" json:"classify_id"`
	/*  */
	ClassifyIdOld int32 `gorm:"column:classify_id_old;type:int(11);" json:"classify_id_old"`
	/* 类别名称 */
	ClassifyName string `gorm:"column:classify_name;type:varchar(128);" json:"classify_name"`
	/* 商品成色 */
	QualityId int32 `gorm:"column:quality_id;type:int(11);" json:"quality_id"`
	/*  */
	QualityIdOld int32 `gorm:"column:quality_id_old;type:int(11);" json:"quality_id_old"`
	/* 成色名称 */
	QualityName string `gorm:"column:quality_name;type:varchar(128);" json:"quality_name"`
	/* 是否本厂 默认 1-本厂 2-外厂 */
	IsOriginal int32 `gorm:"column:is_original;type:tinyint(1);" json:"is_original"`
	/* 是否为在库本厂 1是 2 否 */
	IsOriginalType int32 `gorm:"column:is_original_type;type:tinyint(4);" json:"is_original_type"`
	/* 旧料名称 */
	GoodsName string `gorm:"column:goods_name;type:varchar(128);" json:"goods_name"`
	/* 计价方式，1-按克计费 2-按件计费 */
	CostChargeType int32 `gorm:"column:cost_charge_type;type:tinyint(4);" json:"cost_charge_type"`
	/* 主石名称 */
	MainStoneName string `gorm:"column:main_stone_name;type:varchar(128);" json:"main_stone_name"`
	/* 主石类型id */
	MainStoneTypeId int32 `gorm:"column:main_stone_type_id;type:int(11);" json:"main_stone_type_id"`
	/* 主石类型名称 */
	MainStoneTypeName string `gorm:"column:main_stone_type_name;type:varchar(100);" json:"main_stone_type_name"`
	/* 主石重 */
	MainStoneWeight int32 `gorm:"column:main_stone_weight;type:int(11);" json:"main_stone_weight"`
	/* 副石类型名称 */
	ViceStoneTypeName string `gorm:"column:vice_stone_type_name;type:varchar(100);" json:"vice_stone_type_name"`
	/* 副石重量 */
	ViceStoneWeight int32 `gorm:"column:vice_stone_weight;type:int(11);" json:"vice_stone_weight"`
	/* 证书号 */
	CerNumber string `gorm:"column:cer_number;type:varchar(128);" json:"cer_number"`
	/*  */
	GoldPrice int32 `gorm:"column:gold_price;type:int(11);" json:"gold_price"`
	/* 回收金价，单位分 */
	RecycleGoldPrice int32 `gorm:"column:recycle_gold_price;type:int(11);" json:"recycle_gold_price"`
	/* 总件重 */
	Weight int32 `gorm:"column:weight;type:int(11);" json:"weight"`
	/* 净金重 */
	GoldWeight int32 `gorm:"column:gold_weight;type:int(11);" json:"gold_weight"`
	/* 净金重(实际/回收) */
	GoldWeightRecycle int32 `gorm:"column:gold_weight_recycle;type:int(11);" json:"gold_weight_recycle"`
	/* 回收损耗，20代表20% */
	GoldWeightLossRate float64 `gorm:"column:gold_weight_loss_rate;type:decimal(11,2);" json:"gold_weight_loss_rate"`
	/* 回收工费单价 */
	RecycleLabourPrice int32 `gorm:"column:recycle_labour_price;type:int(11);" json:"recycle_labour_price"`
	/* 回收工费总计 */
	RecycleLabourAmount int32 `gorm:"column:recycle_labour_amount;type:int(11);" json:"recycle_labour_amount"`
	/* 回收数量 */
	GoodsNum int32 `gorm:"column:goods_num;type:int(11);" json:"goods_num"`
	/* 回收总金额，成本总价 单位分 计算得来 */
	RecycleAmount int32 `gorm:"column:recycle_amount;type:int(11);" json:"recycle_amount"`
	/* 抵扣金额，单位分 */
	DeductAmount int32 `gorm:"column:deduct_amount;type:int(11);" json:"deduct_amount"`
	/* 旧料换购利润预估值 */
	OldProfitAmount int32 `gorm:"column:old_profit_amount;type:int(11);" json:"old_profit_amount"`
	/* 旧料原价 */
	OriginalAmount int32 `gorm:"column:original_amount;type:int(11);" json:"original_amount"`
	/* 旧料原标签价（旧料超兑不显示） */
	OldOriginalLabelPrice int32 `gorm:"column:old_original_label_price;type:int(11);" json:"old_original_label_price"`
	/* 回收克差 */
	RecycleWeightDiff int32 `gorm:"column:recycle_weight_diff;type:int(11);" json:"recycle_weight_diff"`
	/* 旧料原标签克重 */
	OldOriginalLabelWeight int32 `gorm:"column:old_original_label_weight;type:int(11);" json:"old_original_label_weight"`
	/* 旧料原成本价 */
	OldOriginalCostPrice int32 `gorm:"column:old_original_cost_price;type:int(11);" json:"old_original_cost_price"`
	/* 新旧天数 */
	NewOldDiffDay int32 `gorm:"column:new_old_diff_day;type:int(11);" json:"new_old_diff_day"`
	/* 订单商品表中的id,已售商品当旧料换购 */
	OrderGoodsId int32 `gorm:"column:order_goods_id;type:int(11);" json:"order_goods_id"`
	/* 旧料条码 */
	StockCode string `gorm:"column:stock_code;type:varchar(128);" json:"stock_code"`
	/* 旧料小备注 */
	Remarks string `gorm:"column:remarks;type:varchar(255);" json:"remarks"`
	/* 0-待处理1-已换购2-已退货3-已删除4-逻辑删除 */
	Status int32 `gorm:"column:status;type:tinyint(4);" json:"status"`
	/* 创建时间 */
	CreateTime int32 `gorm:"column:create_time;type:int(11);" json:"create_time"`
	/* 修改时间 */
	UpdateTime int32 `gorm:"column:update_time;type:int(11);" json:"update_time"`
	/* 删除时间 */
	DeleteTime int32 `gorm:"column:delete_time;type:int(11);" json:"delete_time"`
	/* 同订单旧料序号 */
	SortNum int32 `gorm:"column:sort_num;type:int(11);" json:"sort_num"`
	/* 1素金2非素金 */
	GoldType int32 `gorm:"column:gold_type;type:int(11);" json:"gold_type"`
	/* 换购类型 1正常旧料  4：超兑旧料 */
	SourceType int32 `gorm:"column:source_type;type:tinyint(4);" json:"source_type"`
	/* 旧料与新品绑定的ids */
	GoodsStockIds string `gorm:"column:goods_stock_ids;type:varchar(128);" json:"goods_stock_ids"`
	/* 是否属于兑换 0：不是  1：是 */
	IsOverExchange int32 `gorm:"column:is_over_exchange;type:tinyint(4);" json:"is_over_exchange"`
	/* 超兑对应的数据（前端需要用到后端冗余） */
	NewProductId string `gorm:"column:new_product_id;type:varchar(256);" json:"new_product_id"`
	/* 旧料类型： 1：普通兑换 2：克换克正常兑换 3：克换克超兑数据 */
	Stats int32 `gorm:"column:stats;type:tinyint(4);" json:"stats"`
	/* 克换克旧料克重字段 */
	TotalOldWeight int32 `gorm:"column:total_old_weight;type:int(11);" json:"total_old_weight"`
	/* 主石重 */
	MainStoneNum string `gorm:"column:main_stone_num;type:varchar(255);" json:"main_stone_num"`
	/* 主石颜色 */
	MainStoneColorName string `gorm:"column:main_stone_color_name;type:varchar(255);" json:"main_stone_color_name"`
	/* 主石净度 */
	MainStonePurityName string `gorm:"column:main_stone_purity_name;type:varchar(255);" json:"main_stone_purity_name"`
	/* 副石数 */
	ViceStoneNum string `gorm:"column:vice_stone_num;type:varchar(255);" json:"vice_stone_num"`
	/* 旧料换购图片 */
	PicUrl string `gorm:"column:pic_url;type:varchar(255);" json:"pic_url"`
	/* 图片 */
	PicUrlList string `gorm:"column:pic_url_list;type:json;" json:"pic_url_list"`
	/* 原回收金价，单位分 */
	OriginalGoldPrice int32 `gorm:"column:original_gold_price;type:int(11);" json:"original_gold_price"`
	/* 旧料类型id */
	OldStockTypeId int32 `gorm:"column:old_stock_type_id;type:int(11);" json:"old_stock_type_id"`
	/* 旧料类型名称 */
	OldStockTypeName string `gorm:"column:old_stock_type_name;type:varchar(100);" json:"old_stock_type_name"`
}

// 销售单据赠品明细表
type SaleOrderPrize struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/* 销售单号ID  */
	OrderId int32 `gorm:"column:order_id;type:int(11);" json:"order_id"`
	/* 销售单号 */
	OrderSn string `gorm:"column:order_sn;type:varchar(50);" json:"order_sn"`
	/* 商户ID */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 门店ID */
	MerchantShopId int32 `gorm:"column:merchant_shop_id;type:int(11);" json:"merchant_shop_id"`
	/* 门店名称 */
	MerchantShopName string `gorm:"column:merchant_shop_name;type:varchar(50);" json:"merchant_shop_name"`
	/* 销售员ID */
	AdminUserId int32 `gorm:"column:admin_user_id;type:int(11);" json:"admin_user_id"`
	/* 销售员名称 */
	AdminUserName string `gorm:"column:admin_user_name;type:varchar(255);" json:"admin_user_name"`
	/* 客户ID */
	CustomerId int32 `gorm:"column:customer_id;type:int(11);" json:"customer_id"`
	/* 客户姓名 */
	CustomerName string `gorm:"column:customer_name;type:varchar(50);" json:"customer_name"`
	/* 客户电话 */
	CustomerTelephone string `gorm:"column:customer_telephone;type:varchar(25);" json:"customer_telephone"`
	/* 商品ID */
	GoodsId int32 `gorm:"column:goods_id;type:int(11);" json:"goods_id"`
	/* 迁移前的商品id */
	OldGoodsId int32 `gorm:"column:old_goods_id;type:int(11);" json:"old_goods_id"`
	/* 商品条码 */
	StockCode string `gorm:"column:stock_code;type:varchar(128);" json:"stock_code"`
	/* 证书编号 */
	CerNumber string `gorm:"column:cer_number;type:varchar(254);" json:"cer_number"`
	/* 商品名称 */
	GoodsName string `gorm:"column:goods_name;type:varchar(50);" json:"goods_name"`
	/* 商品分类id */
	GoodsTypeId int32 `gorm:"column:goods_type_id;type:int(11);" json:"goods_type_id"`
	/* 商品分类名称 */
	GoodsTypeName string `gorm:"column:goods_type_name;type:varchar(255);" json:"goods_type_name"`
	/* 商品标记id数组 */
	GoodsTagIds string `gorm:"column:goods_tag_ids;type:varchar(255);" json:"goods_tag_ids"`
	/* 商品标记名称数组 */
	GoodsTagName string `gorm:"column:goods_tag_name;type:varchar(255);" json:"goods_tag_name"`
	/* 货品位置id */
	LibraryId int32 `gorm:"column:library_id;type:int(11);" json:"library_id"`
	/* 货品位置名称 */
	LibraryName string `gorm:"column:library_name;type:varchar(100);" json:"library_name"`
	/* 品牌id */
	BrandId int32 `gorm:"column:brand_id;type:int(11);" json:"brand_id"`
	/* 品牌名称 */
	BrandName string `gorm:"column:brand_name;type:varchar(100);" json:"brand_name"`
	/* 首饰分类id */
	ClassifyId int32 `gorm:"column:classify_id;type:int(11);" json:"classify_id"`
	/* 首饰分类名称 */
	ClassifyName string `gorm:"column:classify_name;type:varchar(100);" json:"classify_name"`
	/* 成色id */
	QualityId int32 `gorm:"column:quality_id;type:int(11);" json:"quality_id"`
	/* 成色名称 */
	QualityName string `gorm:"column:quality_name;type:varchar(64);" json:"quality_name"`
	/* 销售金价类型 */
	GoldSalePriceType int32 `gorm:"column:gold_sale_price_type;type:int(11);" json:"gold_sale_price_type"`
	/* 使用金价 单位：分 */
	GoldPrice int32 `gorm:"column:gold_price;type:int(11);" json:"gold_price"`
	/* 净金重 */
	GoldWeight int32 `gorm:"column:gold_weight;type:int(11);" json:"gold_weight"`
	/* 总件重 */
	Weight int32 `gorm:"column:weight;type:int(11);" json:"weight"`
	/* 1-标签价不含工费(默认) 2-标签价含工费 */
	SaleLabourType int32 `gorm:"column:sale_labour_type;type:tinyint(1);" json:"sale_labour_type"`
	/* 销售计价类型 1-按克 2-按建 */
	ChargeType int32 `gorm:"column:charge_type;type:tinyint(1);" json:"charge_type"`
	/* 销售工费计价类型 1-按克 2-按件 */
	LabourChargeType int32 `gorm:"column:labour_charge_type;type:tinyint(1);" json:"labour_charge_type"`
	/* 销售工费单价 */
	SaleLabourUnitPrice int32 `gorm:"column:sale_labour_unit_price;type:int(11);" json:"sale_labour_unit_price"`
	/* 标签价 单位：分  不包括工费 */
	LabelPrice int32 `gorm:"column:label_price;type:int(11);" json:"label_price"`
	/* 应收金额，抹零前金额 */
	TotalAmount int32 `gorm:"column:total_amount;type:int(11);" json:"total_amount"`
	/* 商品总价，单价*数量+工费 */
	GoodsAmount int32 `gorm:"column:goods_amount;type:int(11);" json:"goods_amount"`
	/* 实收金额，这里如果需要算实收金额需要减去旧料和新品换购的金额 */
	PayAmount int32 `gorm:"column:pay_amount;type:int(11);" json:"pay_amount"`
	/* 成本价 */
	CostAmount int32 `gorm:"column:cost_amount;type:int(11);" json:"cost_amount"`
	/*  */
	TotalNum int32 `gorm:"column:total_num;type:int(11);" json:"total_num"`
	/* 货品周转时间，单位天 */
	CycleTime int32 `gorm:"column:cycle_time;type:int(11);" json:"cycle_time"`
	/* 1-一码一货 2一码多货 */
	StockType int32 `gorm:"column:stock_type;type:tinyint(1);" json:"stock_type"`
	/* 主石重 */
	MainStoneWeight int32 `gorm:"column:main_stone_weight;type:int(11);" json:"main_stone_weight"`
	/* 购买数量 */
	GoodsNum int32 `gorm:"column:goods_num;type:int(11);" json:"goods_num"`
	/* 当数量为多个时，保存goods_id */
	GoodsIdList string `gorm:"column:goods_id_list;type:varchar(255);" json:"goods_id_list"`
	/* 是否属于超送 1-超送 0-非超送 */
	IsOverSend int32 `gorm:"column:is_over_send;type:tinyint(1);" json:"is_over_send"`
	/* 限制赠品数量 */
	GoodsNumLimit int32 `gorm:"column:goods_num_limit;type:int(11);" json:"goods_num_limit"`
	/* 促销活动ID */
	ActId int32 `gorm:"column:act_id;type:int(11);" json:"act_id"`
	/* 促销活动名称 */
	ActName string `gorm:"column:act_name;type:varchar(255);" json:"act_name"`
	/* 商品销售状态,0-待处理,1-售出2-锁定3-退货,4-换购,5-已删除(消单),6-旧料换购,7-旧料回收 */
	Status int32 `gorm:"column:status;type:tinyint(1);" json:"status"`
	/* 该商品是否属于超送活动 */
	IsPromtion int32 `gorm:"column:is_promtion;type:tinyint(1);" json:"is_promtion"`
	/* 销售时间 */
	SaleDatetime string `gorm:"column:sale_datetime;type:timestamp;" json:"sale_datetime"`
	/* 创建时间 */
	CreateTime int32 `gorm:"column:create_time;type:int(11);" json:"create_time"`
	/* 修改时间 */
	UpdateTime int32 `gorm:"column:update_time;type:int(11);" json:"update_time"`
	/* 删除时间 */
	DeleteTime int32 `gorm:"column:delete_time;type:int(11);" json:"delete_time"`
}

// 商品扩展表
type GoodsStockExtend struct {
	/*  */
	Id int32 `gorm:"primary_key;column:id;type:int(11);" json:"id"`
	/* 商户ID */
	MerchantId int32 `gorm:"column:merchant_id;type:int(11);" json:"merchant_id"`
	/* 货品条码 */
	Code string `gorm:"column:code;type:varchar(50);" json:"code"`
	/* 商品介绍,富文本 */
	Description string `gorm:"column:description;type:text;" json:"description"`
	/* 创建时间 */
	CreateTime int32 `gorm:"column:create_time;type:int(11);" json:"create_time"`
	/* 操作用户 */
	AdminUserId int32 `gorm:"column:admin_user_id;type:int(11);" json:"admin_user_id"`
	/* 操作用户 */
	AdminUserName string `gorm:"column:admin_user_name;type:varchar(100);" json:"admin_user_name"`
	/*  */
	UpdateTime int32 `gorm:"column:update_time;type:int(11);" json:"update_time"`
	/* 图片列表 */
	PicUrlList string `gorm:"column:pic_url_list;type:json;" json:"pic_url_list"`
}

package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)

type User struct {
	Id int `json:"user_id"` //用户编号
	Name string `orm:"size(32);unique" json:"name"` //用户昵称
	Password_hash string `orm:"size(128)" json:"password"`
	Mobile string  `orm:"size(11)" json:"Mobile"`
	Real_name 	string `orm:"size(32)" json:"real_name"`
	Id_Card string `orm:"size(20)" json:"id_card"`
	Avatar_url		string			`orm:"size(256)" json:"avatar_url"`		//用户头像路径
	Houses			[] *House		`orm:"reverse(many)" json:"houses"`	//用户发布的房屋信息
	Orders			[] *OrderHouse	`orm:"reverse(many)" json:"orders"`		//用户下的订单
}

/* 户层信息 table_name = house */
type House struct {
	Id 				int				`json:"house_id"`				//房屋编号
	User 			*User			`orm:"rel(fk)" json:"user_id"`		//房屋主人的用户编号
	Area 			*Area			`orm:"rel(fk)" json:"area_id"`		//归属地的区域编号
	Title			string			`orm:"size(64)" json:"title"`		//房屋标题
	Price			int				`orm:"default(0)" json:"price"`	//单价，单位：分
	Address 		string			`orm:"size(512)" orm:"default("")" json:"address"` //地址
	Room_count		int				`orm:"default(1)" json:"room_count"`		//房间数目
	Acreage			int				`orm:"default(0)" json:"acreage"`			//房屋总面积
	Unit 			string			`orm:"size(32)" orm:"default("")" json:"unit"`	//房屋单元，如 几室几厅
	Capacity		int				`orm:"default(1)" json:"capacity"`		//房屋容纳的总人数
	Beds			string			`orm:"size(64)"	orm:"default("")" json:"beds"`	//房屋床铺的位置
	Deposit			int				`orm:"default(0)" json:"deposit"`		//押金
	Min_days		int				`orm:"default(1)" json:"min_days"`	//最少入住天数
	Max_days 		int				`orm:"default(0)" json:"max_days"`	//最多住天数 0表示不限制
	Order_count		int				`orm:"default(0)" json:"order_count"`	//预定完成的该房屋的订单数
	Index_image_url	string			`orm:"size(256)" orm:"default("")" json:"index_image_url"`	//房屋主图片路径
	Facilities		[] *Facility	`orm:"reverse(many)" json:"facilities"`		//房屋设施
	Images			[] *HouseImage	`orm:"reverse(many)" json:"img_urls"`		//房屋的图片
	Orders			[] *OrderHouse	`orm:"reverse(many)" json:"orders"`		//房屋的订单
	Ctime			time.Time		`orm:"auto_now_add;type(datetime)" json:"ctime"`
}

var HOME_PAGE_MAX_HOUSES int =5

var HOUSE_LIST_CAPACITY int =2

type Area struct{
	Id int `json:"aid"`
	Name string `orm:"size(32)" json:"aname"`
}
type Facility struct {
	Id int  `json:"fid"`
	Name string `orm:"size(32)"`
	House []*House `orm:"rel(m2m)"`
}
type HouseImage struct {
	Id 		int		`json:"house_image_id"`			//图片id
	Url 	string	`orm:"size(256)" json:"url"`	//图片url
	House 	*House	`orm:"rel(fk)" json:"house_id"` //图片所属房屋编号
}
// 定义 订单状态常量
const(
	ORDER_STATUS_WAIT_ACCEPT  	= "WAIT_ACCEPT"		//待接单
	ORDER_STATUS_WAIT_PAYMENT  	= "WAIT_PAYMENT"	//待支付
	ORDER_STATUS_PAID  			= "PAID"			//已支付
	ORDER_STATUS_WAIT_COMMENT  	= "COMMENT"			//待评价
	ORDER_STATUS_COMPLETE 	= "COMPLETE"			//已完成
	ORDER_STATUS_CANCELED  	= "CANCELED"			//已取消
	ORDER_STATUS_REJECTED  	= "REJECTED"			//已拒单
)
type OrderHouse struct {
	Id				int			`json:"order_id"`					//订单编号
	User			*User		`orm:"rel(fk)" json:"user_id"`		//下单的用户编号
	House			*House		`orm:"rel(fk)" json:"house_id"`	//预定的房间编号
	Begin_data		time.Time	`orm:"type(datetime)"`				//预定的起始时间
	End_data		time.Time	`orm:"type(datetime)"`				//预定的结束时间
	Days			int												//预定总天数
	House_price		int												//房屋的单价
	Amount			int												//订单总金额
	Status			string		`orm:"default(WAIT_ACCEPT)"`		//订单状态
	Comment			string		`orm:"size(512)"`					//订单评论
	Ctime			time.Time	`orm:"auto_now_add;type(datetime)" json:"ctime"`	//
}
func init(){
	//初始化 设置数据库
	orm.RegisterDataBase("default", "mysql", "root:1234@tcp(127.0.0.1:3306)/ihomebeego?charset=utf8", 30)
	//注册框架的 model orm映射表
	orm.RegisterModel(new(User),new(House),new(Area),new(Facility),new(HouseImage),new(OrderHouse))

	orm.RunSyncdb("default",false,true )

}
package orderdao

import (
	"3g-samples/database"
	"3g-samples/pkg/logging"
	"3g-samples/pkg/page"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Order struct {
	ID       int64  `gorm:"primary_key:yes"`
	UserName string `gorm:"column:user_name"`
	UserNo   string `gorm:"column:user_no"`

	OrderNo            string  `gorm:"column:order_no"`
	OrderTotalAmount   float64 `gorm:"column:order_total_amount"`
	OrderPayableAmount float64 `gorm:"column:order_payable_amount"`
	OrderPaidinAmount  float64 `gorm:"column:order_paidin_amount"`

	SkuNo              string  `gorm:"column:sku_no"`
	SkuName            string  `gorm:"column:sku_name"`
	SkuOriginalPrice   float64 `gorm:"column:sku_original_price"`
	SkuSalePrice       float64 `gorm:"column:sku_sale_price"`
	SkuDiscountRate    float64 `gorm:"column:sku_discount_rate"`
	SkuDiscountedPrice float64 `gorm:"column:sku_discounted_price"`

	CreateAt   time.Time `gorm:"column:create_at;NOT NULL;autoCreateTime"`
	UpdateAt   time.Time `gorm:"column:update_at;NOT NULL;autoUpdateTime"`
	CreateUser int64     `gorm:"column:create_user;NOT NULL;default:1"`
	UpdateUser int64     `gorm:"column:update_user;NOT NULL;default:1"`
	Yn         int       `gorm:"column:yn;NOT NULL;default:1"`
}

type Ordero struct {
	ID     int64  `gorm:"primary_key:yes"`
	UserNo string `gorm:"column:user_no"`

	OrderNo            string  `gorm:"column:order_no"`
	OrderTotalAmount   float64 `gorm:"column:order_total_amount"`
	OrderPayableAmount float64 `gorm:"column:order_payable_amount"`
	OrderPaidinAmount  float64 `gorm:"column:order_paidin_amount"`

	CreateAt   time.Time `gorm:"column:create_at;NOT NULL;autoCreateTime"`
	UpdateAt   time.Time `gorm:"column:update_at;NOT NULL;autoUpdateTime"`
	CreateUser int64     `gorm:"column:create_user;NOT NULL;default:1"`
	UpdateUser int64     `gorm:"column:update_user;NOT NULL;default:1"`
	Yn         int       `gorm:"column:yn;NOT NULL;default:1"`
}

// Ordero 结构体默认对应的表是orderos，如果想改变这个默认需要使用下面的方式进行重定义对应关系
func (order Ordero) TableName() string {
	return "orders"
}

// err := database.GetDB().Raw("SELECT u.user_name, u.login_name, o.order_no,o.order_total_amount,o.order_paidin_amount,oi.sku_name,oi.sku_original_price,oi.sku_discounted_price FROM users u,orders o,order_items oi WHERE u.user_no=o.user_no AND o.order_no=oi.order_no AND u.user_no = ? LIMIT ?,?", order.UserNo, pager.Page, pager.PageSize).Scan(&orders).Error
// 多表、分页、不定条件查询
func QueryOrdersByPage(order *Order, pager *page.Page) ([]Order, error) {
	var orders []Order
	queryConditions := "SELECT u.user_name, u.login_name, o.order_no,o.order_total_amount,o.order_paidin_amount,oi.sku_name,oi.sku_original_price,oi.sku_discounted_price FROM users u,orders o,order_items oi WHERE u.user_no=o.user_no AND o.order_no=oi.order_no AND 1 = 1"
	if order.UserName != "" {
		queryConditions += " AND u.user_name = '" + order.UserName + "'"
	}
	if order.UserNo != "" {
		queryConditions += " AND u.user_no = '" + order.UserNo + "'"
	}
	if order.OrderNo != "" {
		queryConditions += " AND o.order_no = '" + order.OrderNo + "'"
	}
	queryConditions += " LIMIT " + strconv.Itoa(pager.Page) + "," + strconv.Itoa(pager.PageSize)
	logging.Info("分页查询订单的SQL拼串:", queryConditions)

	err := database.GetDB().Raw(queryConditions).Scan(&orders).Error
	if err != nil {
		logging.Error("分页查询订单时，查询执行异常", err)
		return nil, err
	}
	return orders, err
}

func (order *Ordero) AddOrder(tx *gorm.DB) (int64, error) {
	ret := tx.Create(&order)
	return ret.RowsAffected, ret.Error
}

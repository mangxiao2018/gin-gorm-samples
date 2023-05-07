package orderdao

import (
	"gorm.io/gorm"
	"time"
)

type Orderx struct {
	UserNo             string      `json:"user_no"`
	OrderNo            string      `json:"order_no"`
	OrderTotalAmount   float64     `json:"order_total_amount"`
	OrderPayableAmount float64     `json:"order_payable_amount"`
	OrderPaidinAmount  float64     `json:"order_paidin_amount"`
	User               Userx       `json:"users"`       // order:user=1:1
	OrderItem          []OrderItem `json:"order_items"` // order:order_item=1:n
}

type Userx struct {
	ID       int64  `json:"id"`
	UserName string `json:"user_name"`
	UserNo   string `json:"user_no"`
}

type OrderItem struct {
	OrderNo            string  `gorm:"column:order_no" json:"order_no"`
	SkuNo              string  `gorm:"column:sku_no" json:"sku_no"`
	SkuName            string  `gorm:"column:sku_name" json:"sku_name"`
	SkuOriginalPrice   float64 `gorm:"column:sku_original_price" json:"sku_original_price"`
	SkuSalePrice       float64 `gorm:"column:sku_sale_price" json:"sku_sale_price"`
	SkuDiscountRate    float64 `gorm:"column:sku_discount_rate" json:"sku_discount_rate"`
	SkuDiscountedPrice float64 `gorm:"column:sku_discounted_price" json:"sku_discounted_price"`

	CreateAt   time.Time `gorm:"column:create_at;NOT NULL;autoCreateTime"`
	UpdateAt   time.Time `gorm:"column:update_at;NOT NULL;autoUpdateTime"`
	CreateUser int64     `gorm:"column:create_user;NOT NULL;default:1"`
	UpdateUser int64     `gorm:"column:update_user;NOT NULL;default:1"`
	Yn         int       `gorm:"column:yn;NOT NULL;default:1"`
}

// 批量把数据插入订单明细表
func AddOrderItems(orderItems []OrderItem, tx *gorm.DB) (int64, error) {
	ret := tx.CreateInBatches(orderItems, len(orderItems))
	return ret.RowsAffected, ret.Error
}

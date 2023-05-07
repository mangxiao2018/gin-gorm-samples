package orderservices

import (
	"3g-samples/dao/orderdao"
	"3g-samples/dao/userdao"
	"3g-samples/database"
	pager "3g-samples/pkg/page"
	"3g-samples/pkg/util"
)

type OrderService struct {
	UserNo             string      `json:"user_no"`
	OrderNo            string      `json:"order_no"`
	OrderTotalAmount   float64     `json:"order_total_amount"`
	OrderPayableAmount float64     `json:"order_payable_amount"`
	OrderPaidinAmount  float64     `json:"order_paidin_amount"`
	User               User        `json:"users"`       // order:user=1:1
	OrderItem          []OrderItem `json:"order_items"` // order:order_item=1:n
}

type User struct {
	ID       int64  `json:"id"`
	UserName string `json:"user_name"`
	UserNo   string `json:"user_no"`
}

type OrderItem struct {
	OrderNo            string  `json:"order_no"`
	SkuNo              string  `json:"sku_no"`
	SkuName            string  `json:"sku_name"`
	SkuOriginalPrice   float64 `json:"sku_original_price"`
	SkuSalePrice       float64 `json:"sku_sale_price"`
	SkuDiscountRate    float64 `json:"sku_discount_rate"`
	SkuDiscountedPrice float64 `json:"sku_discounted_price"`
}

func (orderService *OrderService) QueryOrdersByPage(page *pager.Page) ([]orderdao.Order, error) {
	order := orderdao.Order{
		UserNo:  orderService.UserNo,
		OrderNo: orderService.OrderNo,
	}
	return orderdao.QueryOrdersByPage(&order, page)
}

// 多表保存
// 使用了事务
// 从EndPoint拿到order:user=1:1 order:order_item=1:n
// 先根据用户编号user_no从数据库中拿到user的全量信息
// 接着用user的全量信息+EndPoint得到的order信息，包装完整的Order信息并保存到order表中
// 从Endpoint得到的order_item信息批量保存到order_item表中
func (orderService *OrderService) SaveOrder() (bool, error) {
	tx := database.GetDB().Begin()
	//1.根据用户编号user_no获取用户信息
	user := userdao.User{
		UserNo: orderService.User.UserNo,
	}

	Exist, err := user.QueryUserByNo(tx)
	if err != nil {
		tx.Rollback()
	}
	if !Exist {
		tx.Rollback()
	}
	//2.保存order信息1条
	orderDao := orderdao.Ordero{
		OrderNo:            orderService.OrderNo,
		UserNo:             orderService.UserNo,
		OrderTotalAmount:   orderService.OrderTotalAmount,
		OrderPayableAmount: orderService.OrderPayableAmount,
		OrderPaidinAmount:  orderService.OrderPaidinAmount,
	}
	rows_affected, err_order := orderDao.AddOrder(tx)
	if err_order != nil && rows_affected > 0 {
		tx.Rollback()
	}
	//3.保存ordr_item信息多条
	orderx := orderdao.Orderx{}
	util.StructUtils(&orderService, &orderx)
	rows_affected_item, err_order_item := orderdao.AddOrderItems(orderx.OrderItem, tx)
	//保存失败，则回滚事务
	if err_order_item != nil && rows_affected_item > 0 {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return true, nil
}

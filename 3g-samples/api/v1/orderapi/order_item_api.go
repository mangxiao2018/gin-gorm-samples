package orderapi

import (
	"3g-samples/pkg/app"
	"3g-samples/pkg/e"
	"3g-samples/pkg/util"
	"3g-samples/service/orderservices"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type OrderxForm struct {
	UserName string `json:"user_name"`
	UserNo   string `json:"user_no"`

	OrderNo            string           `json:"order_no"`
	OrderTotalAmount   float64          `json:"order_total_amount"`
	OrderPayableAmount float64          `json:"order_payable_amount"`
	OrderPaidinAmount  float64          `json:"order_paidin_amount"`
	UserxForm          UserxForm        `json:"users"`       // order:user=1:1
	OrderItemxForms    []OrderItemxForm `json:"order_items"` // order:order_item=1:n
}

type UserxForm struct {
	UserName   string `json:"user_name"`
	NickName   string `json:"nick_name"`
	UserAvatar string `json:"user_avatar"`
	LoginName  string `json:"login_name"`
	LoginPwd   string `json:"login_pwd"`
	UserNo     string `json:"user_no"`
	Gender     int8   `json:"gender"`
	Email      string `json:"email"`
	MobileNo   string `json:"mobile_no"`
}

type OrderItemxForm struct {
	OrderNo            string  `json:"order_no"`
	SkuNo              string  `json:"sku_no"`
	SkuName            string  `json:"sku_name"`
	SkuOriginalPrice   float64 `json:"sku_original_price"`
	SkuSalePrice       float64 `json:"sku_sale_price"`
	SkuDiscountRate    float64 `json:"sku_discount_rate"`
	SkuDiscountedPrice float64 `json:"sku_discounted_price"`

	CreateAt   time.Time `json:"create_at"`
	UpdateAt   time.Time `json:"update_at"`
	CreateUser int64     `json:"create_user"`
	UpdateUser int64     `json:"update_user"`
	Yn         int       `json:"yn"`
}

// 接收来自web的Json串
// 该Json口串包含1:1、1:n关系
// 多表保存数据到库
func AddOrder(c *gin.Context) {
	var (
		appG       = app.Gin{C: c}
		orderxForm = OrderxForm{}
	)
	err := c.ShouldBindJSON(&orderxForm) //绑定web来的Json
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ORDERS_FAIL, nil)
		return
	}
	orderService := orderservices.OrderService{}
	util.StructUtils(&orderxForm, &orderService)
	_, err = orderService.SaveOrder()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ORDERS_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

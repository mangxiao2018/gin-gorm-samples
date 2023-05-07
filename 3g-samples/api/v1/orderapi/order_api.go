package orderapi

import (
	"3g-samples/pkg/app"
	"3g-samples/pkg/e"
	pager "3g-samples/pkg/page"
	"3g-samples/service/orderservices"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderForm struct {
	ID       int64  `form:"id"`
	UserName string `form:"user_name"`
	UserNo   string `form:"user_no"`

	OrderNo            string  `form:"order_no"`
	OrderTotalAmount   float64 `from:"order_total_amount"`
	OrderPayableAmount float64 `from:"order_payable_amount"`
	OrderPaidinAmount  float64 `from:"order_paidin_amount"`

	SkuNo              string  `from:"sku_no"`
	SkuName            string  `from:"sku_name"`
	SkuOriginalPrice   float64 `from:"sku_original_price"`
	SkuSalePrice       float64 `from:"sku_sale_price"`
	SkuDiscountRate    float64 `from:"sku_discount_rate"`
	SkuDiscountedPrice float64 `from:"sku_discounted_price"`
}

// 分页跨表不定条件查询
func QueryOrdersByPage(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	page, _ := strconv.Atoi(c.Query("page"))

	pageSize, _ := strconv.Atoi(c.Query("page_size"))

	pager := pager.Page{
		Page:     page,
		PageSize: pageSize,
	}
	orderService := orderservices.OrderService{
		UserNo:  c.Query("user_no"),
		OrderNo: c.Query("order_no"),
	}
	orders, err := orderService.QueryOrdersByPage(pager.GetPage())
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ORDERS_FAIL, nil)
		return
	}
	appG.ResponseByPage(http.StatusOK, e.SUCCESS, orders, pager)
}

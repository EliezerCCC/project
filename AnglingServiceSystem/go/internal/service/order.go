package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"go/internal/dao"
	"go/internal/models"
	"time"
)

var (
	appID      = "2021000122670624"
	privateKey = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCxclf6fSK5nArpMma9LAZO4hMQ+vJchR4hdaVr3wY6vtu0UJz19gxS1r/Z2KdpRlXFrFYkSIlPkZHSxJVj8tKRUCdxi6V08DJNfLq2FN61NVS+brgkN+xoeMqmlupwhNzx5abhgLIDji99TNqjxOBHWC85JB5KoAJkLg+FVdHQ8DdCOIOr/sVjWOvg3e/L20NUZxUMJc3dNsrd1HFGPZA191uk9atwRBfR2UqIPWxU1I2rJRl0PYxH9IAkO4aCow9rozK8+tZK/jqA00V2JurqJLxKGEjAvJPW8z4z8GKdHThXiaeakWVMgnmpbt3MLO22TYFrwWx1UNO2DCfxXJApAgMBAAECggEAEyG67pzzU4PJgV8XyKmofHTPjXMhEmyf1Qe0MC85orfVRFm0sqg2p+/Te49qbWVR9iAgoYTYmSWezVS4rrANl/FGb22ZB0Smh8g88UjKwkrfujCB1hXZfWIYp2F6IWuRztrt5T5U8yEvaZkRsbqWW9rcLJMCWRVrbJWywrOADcDKnR+E+6AK00IhYclTPOq+TK9+WszWIc7FpwWov/c6lADOPH2Q0KqTM8Wam88ABEIzVsYfFSxr+fToy9C5669nt7wLGdPBYuOGUq28nZelc+InegboTA99CixUHbKUi5eukuyZk6wodUXFgxnvKdGtqUC63rJvtyLJ6yPLopK76QKBgQDYspY1gyKZgFq/HHU2Ko/R7UvwOfTnWv/WhvTTxxRGVR5eiPtXfw9mttvFSV2RrDqM/iVOSjyLklRK+veqqxoW9Ya1SgNR6+NC0zNh9hmOZGyB22/dymGQYBXwDlRjFghedogIvChET1l6ywKhIb+1YbsVwkO8PTIy+BM/4m1CEwKBgQDRoU8wsxNxHGs0FrKYzAvsmtmHt8XFcZaxd27NxLksfM2alvL2UN/rNJygzTMLTk8YC4GS9f7upz8oKibia71O3U1xR9KEyG2+t7G9RAvNSV/+3SUb8jurzOWgrXIuHOt4eJ8zw3cb3Wobv6XCpxAg0dcuKiHLaDbvpzr2LLPMUwKBgDmx6s5amNa31vF1OTvvvxUMkiBmve7vSpiqm/ra0WocdZaD8C+Ok+IsI312mCQV8AzkCbi7iT07+5dvFIvrPGoKexA/oXypSnvbUZ3ucUV5NW6ppzItFSbkflLI7ezny9iixHtW5X2yhsdrz8po+urnF3GmpHQ08u2hITbsKVzjAoGBAJylyvRcTQHwN5fkAi+2Ao8bI8SMKkih2fbhhbZT8//vfwCvHf28b6dmWEi9LhCbwvT33CrCkm8qIULVRNtboUomizFEd2bsEz46LKJ/D8amtwpPnORmPfrikRFhqvoUeE087dKTlRUIfrsfmjF7/8xWolKH/1rtlGe/cor02qKrAoGAHz3xT+r0YLBmRNXi5cYgbfGX3oMY4oK6QCpn3S8C+NGUuj6zQK/qr1UkVZmYUycgSXrwiGrnqrxl9tHKrEloMnKdcHPNA6tqvaIhsIOHXmgQC3p55D931xHOnuPT6BkYAqF+VgCi8+wYh7d4a7N3y/ql+HzoE8CzaBrbgJemHJU="
	client, _  = alipay.New(appID, privateKey, false)
)

type PayList struct {
	PayList []models.Commodity `json:"pay_list"`
}

// Pay 支付接口
func Pay(c *gin.Context) {

	client.LoadAliPayPublicKey("MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqqpXQTgAzQvh8OLvFzuf/9KoWFQRccNbFK9waM6flpleOJ0fxmrhqyReyWONViU7FQEbHgksRucTv5DVR9b72+/ypuLapStp00mqITHdpBH5/LyxjXnTSHSmgAiXDPThE/P7N9cbhgc0DRiAnNH8i4MRFVO6ykuHcQnjWONM6zXhW9ELeXMwYEUqo+kTWsf7b2BPUsW+iZnsN0nAQXlAy/VZUsdAkI0sVJKbwIL7ekklt8tjO1zpV0KRTsjEludXJjaZQGS4Br/lhTalmrqjE+Un75OU6IENRDYJCySVQurYahfn8cd1/XCY/eeGkoxmQ5aMvhv1z9JGey/j/zS6EwIDAQAB")
	var order models.Order
	c.ShouldBind(&order)

	userID, _ := c.Get("user_id")
	userIDValue, _ := userID.(string)
	token := c.MustGet("token")

	order.CreateTime = time.Now()
	order.Status = "已付款"
	order.UserID = userIDValue

	newOrder, err := dao.AddOrder(order)

	var p = alipay.TradePagePay{}
	p.NotifyURL = "https://www.baidu.com"
	p.ReturnURL = "http://localhost:8080/#/shop"         //订单付款后跳转的网址页面
	p.Subject = fmt.Sprintf("商城购买")                      //付款标题
	p.OutTradeNo = fmt.Sprintf("%d", newOrder.ID)        //商家订单号
	p.TotalAmount = fmt.Sprintf("%.2f", newOrder.Amount) //价格
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := client.TradePagePay(p)
	if err != nil {
		fmt.Println(err)
	}

	var payURL = url.String() //扫码支付的网页链接，返回前端后打开

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		fmt.Println("获取失败！")
	} else {
		fmt.Println("获取成功！")
		c.JSON(200, gin.H{
			"result":  "获取成功！",
			"token":   token,
			"pay_url": payURL,
		})
	}
}

//GetOrder 获取某用户订单
func GetOrder(c *gin.Context) {
	var user models.User
	c.ShouldBind(&user)

	orderList, err := dao.GetOrder(user.ID)
	token := c.MustGet("token")

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{
			"order_list": orderList,
			"token":      token,
		})
	}
}

// UpdateOrder 修改订单
func UpdateOrder(c *gin.Context) {
	var order models.Order
	c.ShouldBind(&order)
	token := c.MustGet("token")
	fmt.Println("修改订单:", order)

	err := dao.UpdateOrder(order)

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		fmt.Println("修改失败！")
	} else {
		fmt.Println("修改成功！")
		c.JSON(200, gin.H{
			"result": "修改成功！",
			"token":  token,
		})
	}

}

// GetAllOrder 获取所有订单
func GetAllOrder(c *gin.Context) {
	orderList, err := dao.GetAllOrder()
	token := c.MustGet("token")

	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{
			"order_list": orderList,
			"token":      token,
		})
	}
}

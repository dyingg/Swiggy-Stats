package requests

import (
	"fmt"
	"sgstory/orders"

	"github.com/go-resty/resty/v2"
)

type OrderResponse struct {
	StatusCode int `json:"statusCode"`
	Data       struct {
		TotalOrders int            `json:"total_orders"`
		Orders      []orders.Order `json:"orders"`
	} `json:"data"`
}

func FetchOrders(authToken string, orderID string) (OrderResponse, error) {
	client := resty.New()

	var orderResponseBody OrderResponse
	resp, err := client.R().
		SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:106.0) Gecko/20100101 Firefox/106.0").
		SetHeader("Cookie", "_session_tid="+authToken).
		SetHeader("Content-Type", "application/json").
		SetResult(&orderResponseBody).
		Get("https://www.swiggy.com/dapi/order/all?order_id=" + orderID)

	if err != nil || resp.StatusCode() != 200 || orderResponseBody.StatusCode != 0 {
		fmt.Println("Error:  ", err)
		return OrderResponse{}, fmt.Errorf("failed to fetch orders")
	}

	return orderResponseBody, nil

}
